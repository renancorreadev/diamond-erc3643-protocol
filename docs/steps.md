# Diamond ERC-3643 — Steps to First Transfer

Guia passo a passo: do Diamond recém-deployado até uma transferência entre investidores.

> **Diamond (Amoy):** `0xc9f624Bc1B3e9514b9d7C112408cf05AdC886377`
> **Owner:** `0xB40061C7bf8394eb130Fcb5EA06868064593BFAa`

---

## Visão Geral do Fluxo

```
1. Grant Roles          ← owner configura o time
2. Create Profile       ← define regras de KYC
3. Register Asset       ← cria a classe de ativo (tokenId)
4. Register Identities  ← vincula wallets dos investidores
5. Mint                 ← emite tokens para investidores
6. Transfer             ← investidor transfere para outro
```

---

## Step 1 — Grant Roles

> **Quem executa:** Diamond Owner
> **Página no frontend:** Admin → Security → Role Management
> **Função:** `grantRole(bytes32 role, address account)`

O owner precisa dar permissões ao time. Sem isso, ninguém além do owner consegue operar.

| Role | Hash (keccak256) | Permite |
|------|-----------------|---------|
| `COMPLIANCE_ADMIN` | `keccak256("COMPLIANCE_ADMIN")` | Registrar assets, gerenciar módulos de compliance |
| `ISSUER_ROLE` | `keccak256("ISSUER_ROLE")` | Mint, burn de tokens |
| `TRANSFER_AGENT` | `keccak256("TRANSFER_AGENT")` | Registrar identidades, forced transfers, recovery |
| `CLAIM_ISSUER_ROLE` | `keccak256("CLAIM_ISSUER_ROLE")` | Criar perfis de identidade |
| `PAUSER_ROLE` | `keccak256("PAUSER_ROLE")` | Pausar protocolo ou assets individuais |
| `RECOVERY_AGENT` | `keccak256("RECOVERY_AGENT")` | Recuperar wallets comprometidas |

**Erro se pular:** Cada facet retorna `*Facet__Unauthorized()` se o caller não tem a role necessária.

**Dica:** O owner pode executar tudo sem roles. Se você é o owner, pode pular este step e ir direto para o Step 2.

---

## Step 2 — Create Identity Profile

> **Quem executa:** Owner ou `CLAIM_ISSUER_ROLE`
> **Página no frontend:** Admin → Compliance → Create Claim Topic Profile
> **Função:** `createProfile(uint256[] claimTopics)`

Define quais claim topics (tipos de KYC) são obrigatórios para investidores deste perfil.

**Parâmetros:**
- `claimTopics` — array de topic IDs ERC-735 (ex: `[1, 2, 3]`)
  - `1` = KYC básico
  - `2` = AML check
  - `3` = Accredited investor
  - (os números são convenção do seu protocolo)

**Retorna:** `profileId` (auto-incrementado, começa em 1)

**Exemplo:**
```
claimTopics: [1, 2]  →  profileId: 1 (KYC + AML)
```

**Erro se pular:** O `registerAsset` no Step 3 precisa de um `identityProfileId` válido. Se passar `0` ou um ID que não existe, transfers vão falhar na verificação de compliance.

---

## Step 3 — Register Asset

> **Quem executa:** Owner ou `COMPLIANCE_ADMIN`
> **Página no frontend:** Admin → Assets → Register New Asset
> **Função:** `registerAsset(RegisterAssetParams)`

Cria uma nova classe de ativo (tokenId). Cada tokenId = um tipo de asset diferente.

**Parâmetros:**

| Campo | Tipo | Exemplo | Descrição |
|-------|------|---------|-----------|
| `name` | string | `"Edifício Aurora Apt"` | Nome do ativo |
| `symbol` | string | `"AURORA"` | Símbolo |
| `uri` | string | `"ipfs://Qm..."` | Metadata URI |
| `supplyCap` | uint256 | `1000000` | Máximo de tokens (0 = ilimitado) |
| `identityProfileId` | uint32 | `1` | ID do perfil criado no Step 2 |
| `complianceModules` | address[] | `[]` | Módulos de compliance (pode ser vazio inicialmente) |
| `issuer` | address | `0xB400...` | Quem pode mintar este asset |
| `allowedCountries` | uint16[] | `[76, 840]` | Países permitidos (ISO 3166: 76=Brasil, 840=EUA) |

**Retorna:** `tokenId` (auto-incrementado, começa em 1)

**Erros:**
- `AssetManagerFacet__ZeroAddress()` — issuer não pode ser 0x0
- `AssetManagerFacet__EmptyString()` — name e symbol obrigatórios
- `AssetManagerFacet__TooManyModules()` — máximo 10 módulos

---

## Step 4 — Register Investor Identities

> **Quem executa:** Owner ou `TRANSFER_AGENT`
> **Página no frontend:** Admin → Identity → Register Identity
> **Função:** `registerIdentity(address wallet, address identity, uint16 country)`

Vincula a wallet do investidor ao seu contrato ONCHAINID e país.

**Parâmetros:**

| Campo | Tipo | Exemplo | Descrição |
|-------|------|---------|-----------|
| `wallet` | address | `0xAlice...` | EOA do investidor |
| `identity` | address | `0xIdentity...` | Contrato ONCHAINID do investidor |
| `country` | uint16 | `76` | País (ISO 3166-1 numérico) |

**Importante:** Precisa registrar **ambos** — sender E receiver — antes de qualquer transferência.

**Códigos de país comuns:**

| País | Código |
|------|--------|
| Brasil | 76 |
| EUA | 840 |
| Japão | 392 |
| Alemanha | 276 |
| UK | 826 |
| Portugal | 620 |

**Batch:** Use `batchRegisterIdentity(wallets[], identities[], countries[])` para múltiplos.

**Erros:**
- `IdentityRegistryFacet__ZeroAddress()` — wallet ou identity não pode ser 0x0
- `IdentityRegistryFacet__AlreadyRegistered(wallet)` — wallet já registrada

**Nota sobre ONCHAINID:** No estado atual, a verificação simplificada aceita qualquer endereço como identity (não valida claims on-chain). Para testes, pode usar qualquer endereço não-zero como identity.

---

## Step 5 — Mint Tokens

> **Quem executa:** Owner ou `ISSUER_ROLE`
> **Página no frontend:** Admin → Supply → Mint
> **Função:** `mint(uint256 tokenId, address to, uint256 amount)`

Emite tokens para um investidor registrado.

**Parâmetros:**

| Campo | Tipo | Exemplo | Descrição |
|-------|------|---------|-----------|
| `tokenId` | uint256 | `1` | ID do asset (Step 3) |
| `to` | address | `0xAlice...` | Investidor (deve estar registrado no Step 4) |
| `amount` | uint256 | `1000` | Quantidade de tokens |

**Checks executados (em ordem):**
1. ✅ Asset existe e não está pausado
2. ✅ Protocolo não está pausado globalmente
3. ✅ Wallet do receiver não está frozen (global ou por asset)
4. ✅ Supply cap não será excedido

**Erros:**
- `SupplyFacet__AssetNotRegistered(tokenId)` — tokenId não existe
- `SupplyFacet__AssetPaused(tokenId)` — asset está pausado
- `SupplyFacet__ProtocolPaused()` — protocolo pausado
- `SupplyFacet__SupplyCapExceeded(tokenId, current, cap)` — excede supply cap
- `SupplyFacet__WalletFrozenGlobal(wallet)` — wallet frozen global
- `SupplyFacet__WalletFrozenAsset(tokenId, wallet)` — wallet frozen no asset

**Side effects:**
- Tokens vão para a **partição free** do investidor
- `totalSupply` incrementado
- Holder tracking atualizado
- Evento: `Minted(tokenId, to, amount)`

---

## Step 6 — Transfer

> **Quem executa:** O investidor (holder) ou operador aprovado
> **Página no frontend:** Portfolio → Transfer
> **Função:** `safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data)`

Transfere tokens entre investidores. Passa por **6 estágios de validação**.

**Parâmetros:**

| Campo | Tipo | Exemplo | Descrição |
|-------|------|---------|-----------|
| `from` | address | `0xAlice...` | Sender |
| `to` | address | `0xBob...` | Receiver |
| `id` | uint256 | `1` | tokenId |
| `amount` | uint256 | `500` | Quantidade |
| `data` | bytes | `0x` | Dados extras (normalmente vazio) |

**Pipeline de validação:**

```
safeTransferFrom(Alice, Bob, 1, 500, 0x)
  │
  ├─ 1. Operator check     → msg.sender == from OU isApprovedForAll?
  ├─ 2. Protocol paused?   → revert ProtocolPaused
  ├─ 3. Wallet frozen?     → revert WalletFrozenGlobal (sender OU receiver)
  ├─ 4. Asset exists?      → revert AssetNotRegistered
  ├─ 5. Asset paused?      → revert AssetPaused
  ├─ 6. Asset freeze?      → revert WalletFrozenAsset (sender OU receiver)
  ├─ 7. Lockup ativo?      → revert LockupActive
  ├─ 8. Free balance?      → revert InsufficientFreeBalance (balance - frozen >= amount)
  ├─ 9. Compliance check   → cada módulo retorna canTransfer()
  │
  ├─ Execute: atualiza partições free
  └─ Post-hook: module.transferred()
```

**Preflight (recomendado):** Use `canTransfer()` antes para verificar sem gastar gas:
```
ComplianceRouterFacet.canTransfer(tokenId, from, to, amount, data) → (bool, bytes32)
```

---

## Checklist Rápido

```
□ 1. [Owner]           grantRole(COMPLIANCE_ADMIN, myAddress)
□ 2. [Owner/Claim]     createProfile([1, 2])                    → profileId: 1
□ 3. [Owner/Compliance] registerAsset({
                           name: "Token A",
                           symbol: "TKA",
                           uri: "https://...",
                           supplyCap: 1000000,
                           identityProfileId: 1,
                           complianceModules: [],
                           issuer: owner,
                           allowedCountries: [76, 840]
                         })                                       → tokenId: 1
□ 4a. [Owner/Agent]    registerIdentity(Alice, identityAlice, 76)
□ 4b. [Owner/Agent]    registerIdentity(Bob, identityBob, 840)
□ 5.  [Owner/Issuer]   mint(1, Alice, 1000)
□ 6.  [Alice]          safeTransferFrom(Alice, Bob, 1, 500, 0x)  ✅
```

---

## Troubleshooting

| Erro | Causa | Solução |
|------|-------|---------|
| `*__Unauthorized()` | Caller não tem a role necessária | Grant role no Step 1 |
| `AssetNotRegistered` | tokenId não existe | Execute Step 3 primeiro |
| `AlreadyRegistered` | Wallet já tem identidade | Use `updateIdentity` ao invés |
| `SupplyCapExceeded` | Mint excede o cap | Aumente com `setSupplyCap` |
| `WalletFrozenGlobal` | Wallet congelada | Owner: `setWalletFrozen(wallet, false)` |
| `InsufficientFreeBalance` | Saldo insuficiente (descontando frozen) | Verifique `partitionBalanceOf` |
| `ComplianceRejected` | Módulo de compliance rejeitou | Verifique `canTransfer` para ver reason code |
| `LockupActive` | Tokens em lockup | Aguarde expiração ou remova lockup |
| `ProtocolPaused` | Protocolo pausado globalmente | Owner: `unpauseProtocol()` |
