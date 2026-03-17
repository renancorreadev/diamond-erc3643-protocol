# Diamond ERC-3643 — Demo Amoy (Resumo)

Contrato Diamond: https://amoy.polygonscan.com/address/0xAb07CEf1BEeDBb30F5795418c79879794b31C521

---

## 1. Identidade — 10 wallets registradas (KYC)

Batch com 10 investidores de 10 países diferentes:
https://amoy.polygonscan.com/tx/0x6219bf42f3ffc7513079e6c6d3bc4427e8cac832d0dba04660f0555d43c18c23

---

## 2. Registro de 20 ativos (tokenIds 1-20)

Cada tokenId é uma classe de ativo independente com regras próprias de compliance.

- 5 reward tokens (rUSDT, rUSDC, rDAI, SRP, GRT)
- 10 RWA tokens (imóveis BR, bonds US, carbono EU, ouro, prata, REIT JP...)
- 5 project tokens (Dubai, Suíça, SP, Londres, Abu Dhabi)

Exemplo — registro do Brazil Real Estate Fund (BREF):
https://amoy.polygonscan.com/tx/0x5bb6b86d2e7a592a8087a66a5e764bd777457c426a6b5a6d5eeaab0915d7ba21

---

## 3. Minting — 50+ mints para investidores

Cada mint emite `TransferSingle(operator, address(0), to, tokenId, amount)`.

Exemplo — 5,000 BREF para Alice:
https://amoy.polygonscan.com/tx/0x9ccd29509df0bbd1ea613f93b3231916651f00108d72667c5caefe85a6632b60

---

## 4. Compliance — módulos plugados por tokenId

Cada ativo tem suas próprias regras. Módulos plugados via `addComplianceModule`:

| Token | Regra |
|-------|-------|
| BREF (7) | Bloqueia US e JP |
| USTB26 (8) | Bloqueia BR e AE + max 12k/wallet |
| ECC (9) | Máx 5 holders |
| GBT (10) | Máx 5k/wallet |

---

## 5. Transfers que FALHARAM (compliance revert)

Todas reverteram on-chain. O Polygonscan mostra "execution reverted" genérico, mas o revert data contém o erro decodificado: `ERC1155Facet__ComplianceRejected(tokenId, reason)`.

**A) Alice(BR) → Bob(US) — token 7 (BREF) — País bloqueado (US)**
https://amoy.polygonscan.com/tx/0x893f13709a016c35433adfa987bda03e2eb56868b6c82c8c8da0481bc6824149

**B) Bob(US) → Alice(BR) — token 8 (USTB26) — País bloqueado (BR)**
https://amoy.polygonscan.com/tx/0x52eed90072502cd3feb82ac8f20f39b7be5638d251f75d985db10a4cec767e94

**C) Alice → Wei — token 10 (GBT) — Limite de saldo excedido (4k+2k > 5k)**
https://amoy.polygonscan.com/tx/0x54dd6da0ebd886d5bb6cf43de3d379bc5cbbd59f3cec356a3a58b5d0df824e9d

**D) Hans → Alice — token 9 (ECC) — Limite de holders excedido (>5)**
https://amoy.polygonscan.com/tx/0x18f752146fa84ab0771afeec64c03448f90542d6a6dd4782be80bd478650d82f

---

## 6. Transfers que PASSARAM (compliance ok)

**Alice(BR) → Carlos(PT) — token 7 — PT não está bloqueado**
https://amoy.polygonscan.com/tx/0x466768cf93bddd71fbf395d56ceb9526d0fe80ce5f4154b9f7f2b191601ab5fe

**Bob(US) → Klaus(CH) — token 8 — CH ok + 7k < 12k max**
https://amoy.polygonscan.com/tx/0x3629255fff543989a239b2a2ebc2ae4ec17cf58764fcfa167e158cc6beec1125

**Alice → Yuki — token 10 — 2.5k < 5k max**
https://amoy.polygonscan.com/tx/0xe148097ac1b65bfc896965f3dc1ddaaf1aebf3b74ec880e9d5d1a0ad4b1989be

---

## 7. Yield Distribution — recompensas automáticas

O YieldDistributor distribui yield proporcionalmente em O(1) sem iterar holders.

**Setup:** plugin plugado no token 11 (Silver) + rDAI como reward
https://amoy.polygonscan.com/tx/0x1b27738ea88845debd79776e29390ac81455b57b1e965d685aead75a05302b67

**Depósito de 37,000 rDAI (1:1 com supply):**
https://amoy.polygonscan.com/tx/0x7289f5072a8649901f455e94c9c47cda5e60660f4b8d8587c1877639826ba638

**10 holders claimaram proporcionalmente:**

| Holder | Claimed | Tx |
|--------|---------|-----|
| Alice | 2,000 | https://amoy.polygonscan.com/tx/0xa8846840b2caf9ba1a5e5c22d9f71c75c40c00f19e4616186dd3d58fae1bc1b0 |
| Bob | 5,000 | https://amoy.polygonscan.com/tx/0x0e8e2bee4e70555ae0670d13a5d36eeb4c68ea1d5a4aabce89c462ad0b52357c |
| Carlos | 1,500 | https://amoy.polygonscan.com/tx/0x65efff98c70e72189b7532f252563fff9f0cd47fe27397100405804bd174bff0 |
| Yuki | 3,000 | https://amoy.polygonscan.com/tx/0x797b6c2ffba74927ac11cddfca4baabbec2439046dbed1d27e66337e37235018 |
| Hans | 2,500 | https://amoy.polygonscan.com/tx/0xefcabd3a13e58fd4aaa6112a9c61bb7159511b52f3f4d08b17031bc5d938b0c8 |
| Marie | 3,000 | https://amoy.polygonscan.com/tx/0xc0d8d303d835dd3c9d5366195c592b92818b5058f607529a8f8a960eacacaf38 |
| James | 2,000 | https://amoy.polygonscan.com/tx/0xf6b76d4d63456e33d8064907d60441d1151ad4d01faee39225835677dd542f64 |
| Klaus | 1,000 | https://amoy.polygonscan.com/tx/0x33542cbbf995bdc03b58df0504509803c39669232b229dc97b507c50072cfb5a |
| Wei | 4,000 | https://amoy.polygonscan.com/tx/0xf13c9281bfaeba7f9b9e652a984286be725fd2905af1040f972a48807aafd80f |
| Omar | 3,000 | https://amoy.polygonscan.com/tx/0x2f76fda8ff2c5c1b6ee42dd8cf5f04336e5a70b1bcd04742298beb041d2a75ad |

---

## Stack

- EIP-2535 Diamond Proxy (1 endereço, 18 facets upgradeáveis)
- ERC-1155 Multi-token (20 ativos num único contrato)
- ERC-3643 Compliance (KYC + restrições por jurisdição/balance/holders)
- Plugin system (YieldDistributor com padrão Synthetix accumulator O(1))
- Solidity 0.8.28 / Foundry / 396 testes
