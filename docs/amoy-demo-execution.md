# Diamond ERC-3643 — Amoy Testnet Demo Execution

**Network:** Polygon Amoy Testnet (chainId: 80002)
**Diamond:** [`0xAb07CEf1BEeDBb30F5795418c79879794b31C521`](https://amoy.polygonscan.com/address/0xAb07CEf1BEeDBb30F5795418c79879794b31C521)
**YieldDistributor:** [`0xa8c6A9AAfd18545bEe6BC734eba702C55beA95dF`](https://amoy.polygonscan.com/address/0xa8c6A9AAfd18545bEe6BC734eba702C55beA95dF)
**Date:** 2026-03-17
**Block explorer:** https://amoy.polygonscan.com

---

## Test Accounts

| # | Name | Address | Country |
|---|------|---------|---------|
| 0 | Alice | [`0x229cF220...C005AE83`](https://amoy.polygonscan.com/address/0x229cF2204fe396e45b24b0ACb61A8070C005AE83) | BR (76) |
| 1 | Bob | [`0x0bacE580...6b660E`](https://amoy.polygonscan.com/address/0x0bacE5808b4d3cC737BB5745dD2158F94d6b660E) | US (840) |
| 2 | Carlos | [`0x697464D2...442e81`](https://amoy.polygonscan.com/address/0x697464D235584Eb0946d5dc91cd8e2C2aF442e81) | PT (620) |
| 3 | Yuki | [`0xFBfE3e22...93c8c2`](https://amoy.polygonscan.com/address/0xFBfE3e221E97A88a90752DE9Dee428B5e493c8c2) | JP (392) |
| 4 | Hans | [`0x82172628...c75dd2`](https://amoy.polygonscan.com/address/0x821726280964bE2B94F06c1CdCc9A7E375c75dd2) | DE (276) |
| 5 | Marie | [`0xe3A776eA...3C42C`](https://amoy.polygonscan.com/address/0xe3A776eA46D2c9EF400e639503EF24c45Ac3C42C) | FR (250) |
| 6 | James | [`0x35C38F43...D36B4a`](https://amoy.polygonscan.com/address/0x35C38F430552793e05161024c385c4c107D36B4a) | GB (826) |
| 7 | Klaus | [`0x36E48F97...CB9b5`](https://amoy.polygonscan.com/address/0x36E48F9740F26cd34Ad51180047a60c78B3C2977) | CH (756) |
| 8 | Wei | [`0x769ce45e...9CB9b5`](https://amoy.polygonscan.com/address/0x769ce45e613a6b2E8e3848ef81d311d3029CB9b5) | SG (702) |
| 9 | Omar | [`0x4F581784...fb41C6`](https://amoy.polygonscan.com/address/0x4F581784D3079f27cdEB8F524D8D638f06fb41C6) | AE (784) |

---

## Registered Assets (20 tokens)

| tokenId | Name | Symbol | SupplyCap | Jurisdictions | Purpose |
|---------|------|--------|-----------|---------------|---------|
| 1 | RWA Token | RWA | 1,000,000 | All | Original test token |
| 2 | USDT Reward Pool | rUSDT | Unlimited | All | Reward token |
| 3 | USDC Reward Pool | rUSDC | Unlimited | All | Reward token |
| 4 | DAI Reward Pool | rDAI | Unlimited | All | Reward token |
| 5 | Staking Reward Points | SRP | 10,000,000 | All | Reward token |
| 6 | Governance Reward Token | GRT | 5,000,000 | All | Reward token |
| 7 | Brazil Real Estate Fund | BREF | 100,000 | BR only | RWA — real estate |
| 8 | US Treasury Bond 2026 | USTB26 | 500,000 | US only | RWA — government bond |
| 9 | European Carbon Credits | ECC | 1,000,000 | DE,FR,GB,PT | RWA — ESG |
| 10 | Gold Backed Token | GBT | 50,000 | All | RWA — commodity |
| 11 | Silver Commodity Token | SCT | 200,000 | All | RWA — commodity |
| 12 | Tokyo Office REIT | TKREIT | 300,000 | JP only | RWA — real estate |
| 13 | Art Collection Fund | ARTF | 10,000 | All | RWA — alternative |
| 14 | Wine Vintage Portfolio | WINE | 25,000 | FR,PT | RWA — alternative |
| 15 | Singapore Infrastructure Bond | SGIB | 750,000 | SG,JP | RWA — infra |
| 16 | Dubai Real Estate Trust | DRET | 500,000 | AE only | Project token |
| 17 | Swiss Private Equity | SPEF | 200,000 | CH only | Project token |
| 18 | Sao Paulo Debentures | SPDB | 400,000 | BR only | Project token |
| 19 | London Green Bond | LGRB | 300,000 | GB only | Project token |
| 20 | Abu Dhabi Sukuk | ADSK | 600,000 | AE only | Project token |

---

## Execution Log

### Phase 1: Gas Funding

1 POL sent from owner to each test wallet.

| # | Recipient | Tx |
|---|-----------|-----|
| 0 | Alice | [0x6c594f5e...fc95a508](https://amoy.polygonscan.com/tx/0x6c594f5e42b6518eecbb6f0e6057a635edbc7d1b44145bc4c59895a3fc95a508) |
| 1 | Bob | [0x50c631bd...b1956a56](https://amoy.polygonscan.com/tx/0x50c631bd49e6ff46dbd321345ff84398e6d7d14c91d915499f550aafb1956a56) |
| 2 | Carlos | [0x03478b47...fa4ba5c](https://amoy.polygonscan.com/tx/0x03478b470d754224ec033e8108ad0f7aa5451fa0f4a9233697dc1c980fa4ba5c) |
| 3 | Yuki | [0x04d9f652...10bc5145](https://amoy.polygonscan.com/tx/0x04d9f6521538cc3706435042897f6589db817cc145793f5c1110cac210bc5145) |
| 4 | Hans | [0xa0eac100...7198daea](https://amoy.polygonscan.com/tx/0xa0eac100395b9678bd6fdf622a030042c421880a6b8275de054704ea7198daea) |
| 5 | Marie | [0xa1184f82...3eb2e8](https://amoy.polygonscan.com/tx/0xa1184f82ecf1927989c083e95378d600d89ba66612d351b392a31ed8bc3eb2e8) |
| 6 | James | [0x5651ed96...e789c0e0](https://amoy.polygonscan.com/tx/0x5651ed961615d0dd4e71c43f427d4df5a61e2a08700089e15fd41c7ee789c0e0) |
| 7 | Klaus | [0x3c8deb4e...ba83f93c](https://amoy.polygonscan.com/tx/0x3c8deb4e2d19f128edc2ce9d555bf126dbe0ee71ad9ff90f439a8500ba83f93c) |
| 8 | Wei | [0x470c7f29...cbab3d0b](https://amoy.polygonscan.com/tx/0x470c7f2913aeb00fe8aa1c98ed17e992cd85f4db345e98eac35ccd92cbab3d0b) |
| 9 | Omar | [0x843ec50d...f9b3d6](https://amoy.polygonscan.com/tx/0x843ec50dfa1abb6f6628b3cc9eec59ee8675bf07626917c9605987e707f9b3d6) |

---

### Phase 2: Identity Registration

Batch registration of 10 wallets with ONCHAINID (mock) and ISO 3166-1 country codes.

**Call:** `batchRegisterIdentity(wallets[], identities[], countries[])`
**Tx:** [0x6219bf42...43c18c23](https://amoy.polygonscan.com/tx/0x6219bf42f3ffc7513079e6c6d3bc4427e8cac832d0dba04660f0555d43c18c23)

**Events emitted (10x):**
```
IdentityBound(wallet, identity, country)
```

---

### Phase 3: Asset Registration (19 new tokens)

**Call:** `registerAsset((name, symbol, uri, supplyCap, identityProfileId, complianceModules[], issuer, allowedCountries[]))`

| tokenId | Symbol | Tx |
|---------|--------|-----|
| 2 | rUSDT | [0xb42ecc82...f74a52e3](https://amoy.polygonscan.com/tx/0xb42ecc82bf21c97352614208c1c732b30f51bdc4c4a62ce0a9d93ce2f74a52e3) |
| 3 | rUSDC | [0x9d66ddb0...fa23a390](https://amoy.polygonscan.com/tx/0x9d66ddb0d2673db67ca32e94c43361ea1ec308f871073aaa9c7e09cefa23a390) |
| 4 | rDAI | [0x15469264...4601c2f8](https://amoy.polygonscan.com/tx/0x1546926480e86a0a442b8da9dcb7a94ddc4c7f07b1a5fbc5ec0dd8804601c2f8) |
| 5 | SRP | [0xe5ac3283...3320a825](https://amoy.polygonscan.com/tx/0xe5ac328356e880e29f35a5dbe22de6303a1c3a1e8d511ad6365d78703320a825) |
| 6 | GRT | [0xefb29a9a...7647b0bc](https://amoy.polygonscan.com/tx/0xefb29a9a368ff2178cfdacedf3a0c1ba9332b7d8b79f399af16c49117647b0bc) |
| 7 | BREF | [0x5bb6b86d...d7ba21](https://amoy.polygonscan.com/tx/0x5bb6b86d2e7a592a8087a66a5e764bd777457c426a6b5a6d5eeaab0915d7ba21) |
| 8 | USTB26 | [0x8505d504...dc9d29](https://amoy.polygonscan.com/tx/0x8505d50443304eccc96795c18d183dda2a9ced50e070904b823f4c9440dc9d29) |
| 9 | ECC | [0x4c9a3a7a...5629c57](https://amoy.polygonscan.com/tx/0x4c9a3a7ab5b013bcd1ffe20ab3e826de507a91ecb7f2b5de5fed1b2c95629c57) |
| 10 | GBT | [0x53ad9dfa...5611a1](https://amoy.polygonscan.com/tx/0x53ad9dfa94c8e900c8fdd0a265b1b79b62ee2629ba6d994bddd6d12b645611a1) |
| 11 | SCT | [0x0a2707ec...013f50a](https://amoy.polygonscan.com/tx/0x0a2707ec43aeb3b72e24358c4d2d6b45b58479135ed8f62c043d10022013f50a) |
| 12 | TKREIT | [0xb0215c33...8657e6f](https://amoy.polygonscan.com/tx/0xb0215c33b64bc82e552d163a72cc9488b4094ccb1c25902c6006697298657e6f) |
| 13 | ARTF | [0xc64f1162...ad62c8c](https://amoy.polygonscan.com/tx/0xc64f1162c4b3ad3c4fec6b8034c27cb240fc965c5a2408200c7a1b274ad62c8c) |
| 14 | WINE | [0x479462f6...f1f30cf2](https://amoy.polygonscan.com/tx/0x479462f6ec88690cd47986c96507d728ee5d7605277884e1227ec1a5f1f30cf2) |
| 15 | SGIB | [0x522511ca...00827d9](https://amoy.polygonscan.com/tx/0x522511ca1287c7e7267a7258b4ac223509544934827dfe77c6a2c39f900827d9) |
| 16 | DRET | [0x99e52294...92ec1b5b](https://amoy.polygonscan.com/tx/0x99e522941b899068aeed2a1aa735c164c0ca33b055c2aa01594fe3b692ec1b5b) |
| 17 | SPEF | [0x7f47ad47...219b9261](https://amoy.polygonscan.com/tx/0x7f47ad4715081c8f91ca80d576852fe20ab3e46698fb87f11d11bab5219b9261) |
| 18 | SPDB | [0x4d642290...c94af449](https://amoy.polygonscan.com/tx/0x4d642290bb86c3093fe0d193e9fdc61188fe1ee297e10dbb75fcf085c94af449) |
| 19 | LGRB | [0x11d6500a...9071ca0](https://amoy.polygonscan.com/tx/0x11d6500a7d9f567cbb65611b6080c1c00e4b41e7cf9f1e05e4d20f10e9071ca0) |
| 20 | ADSK | [0x670aaf20...fb9a465](https://amoy.polygonscan.com/tx/0x670aaf204f3cdf340964e20b0e432b2ed55929ed81f891efec5243fb0fb9a465) |

**Event per registration:** `AssetRegistered(tokenId, issuer, profileId)`

---

### Phase 4: Minting

**Call:** `mint(tokenId, to, amount)`

Each mint emits:
```
TransferSingle(operator, address(0), to, tokenId, amount)
Minted(tokenId, to, amount)
```

#### Reward Tokens (2-6) — Minted to Owner

| Token | Amount | Tx |
|-------|--------|-----|
| 2 (rUSDT) | 1,000,000 | [0x254b5f99...3224640e](https://amoy.polygonscan.com/tx/0x254b5f994524b3db9b07fa4aa1b4b474912b1714836933ea47d7a0a93224640e) |
| 3 (rUSDC) | 1,000,000 | [0xc67d2bdb...8035347](https://amoy.polygonscan.com/tx/0xc67d2bdb57ab5d2ab2e87ac6da31a76efd27ef852b06a246b25ab78fa8035347) |
| 4 (rDAI) | 1,000,000 | [0x29c9bd3c...6898d8](https://amoy.polygonscan.com/tx/0x29c9bd3c7d1ac038698d718e661aaaf8d9c1b8ae66a01e91780d5553286898d8) |
| 5 (SRP) | 500,000 | [0x1f76ada5...ee127ee0](https://amoy.polygonscan.com/tx/0x1f76ada5fd93a9e2b603597e227cfaca42c347a78809c96bb4ee127ee0dcf767) |
| 6 (GRT) | 250,000 | [0xaac9f4d2...1bfbc6a](https://amoy.polygonscan.com/tx/0xaac9f4d21c8f705ea8ec0e20c53791b63ead2934189ac7fc6bb17931d1bfbc6a) |

#### RWA & Project Tokens (7-20) — Minted to Investors

| Token | Investor | Amount | Tx |
|-------|----------|--------|-----|
| 7 | Alice (BR) | 5,000 | [0x9ccd2950...a6632b60](https://amoy.polygonscan.com/tx/0x9ccd29509df0bbd1ea613f93b3231916651f00108d72667c5caefe85a6632b60) |
| 7 | Bob (US) | 3,000 | [0xeda947ca...b9df964c](https://amoy.polygonscan.com/tx/0xeda947ca3ef0c4fc824efc48d468d3c478c4b2ea93915efcd3412853b9df964c) |
| 7 | Carlos (PT) | 2,000 | [0x91fe1903...e8d3ca](https://amoy.polygonscan.com/tx/0x91fe1903fad491bbc41a884a39867fa06cf6dac24e8b0f60201a495454e8d3ca) |
| 8 | Bob (US) | 10,000 | [0xf5a13891...c4ed8d3](https://amoy.polygonscan.com/tx/0xf5a138912c4b18acda73c52b399ab7bf16acacf335b8b016e49292e9cc4ed8d3) |
| 8 | Hans (DE) | 5,000 | [0x2a86e0d0...bf361cb](https://amoy.polygonscan.com/tx/0x2a86e0d0fce3cef915221e03f1c22a61e8c3a476da1a4b60bdf7aa864bf361cb) |
| 8 | Klaus (CH) | 5,000 | [0x5b91460c...83f2f9](https://amoy.polygonscan.com/tx/0x5b91460c393fc3552c9bb76d9787000bb995837e105bbe5452b7309d2183f2f9) |
| 9 | Hans (DE) | 8,000 | [0x6139d69c...cda4c2](https://amoy.polygonscan.com/tx/0x6139d69c6fefb0f1e9bb5267e065283c3b54ca6fe4230c177eb7544e90cda4c2) |
| 9 | Marie (FR) | 6,000 | [0x28619ac8...77a7f44](https://amoy.polygonscan.com/tx/0x28619ac84b86c06d00b8742344a43e7555ebaa3ef61a6507ae5ec5eb377a7f44) |
| 9 | James (GB) | 4,000 | [0x6ee8f3d9...9f15d3](https://amoy.polygonscan.com/tx/0x6ee8f3d90444ea22bf21e105716f736d4a853ca781c6ca154387a00dbd9f15d3) |
| 9 | Carlos (PT) | 2,000 | [0xc413bf14...cc99cb5](https://amoy.polygonscan.com/tx/0xc413bf14191e34166b10ce24e6e9e36072df817198f52ef933354df12cc99cb5) |
| 10 | Alice (BR) | 1,000 | [0x73d44fb3...9e0c3093](https://amoy.polygonscan.com/tx/0x73d44fb31b191bdcf3adf54816c4909a858678b74539d8b74d3000ca9e0c3093) |
| 10 | Yuki (JP) | 2,000 | [0x65f7d4e0...bfc786](https://amoy.polygonscan.com/tx/0x65f7d4e0d3a701f29ef70dddd2cd020ad9dbefd6bba51bce0ff5efb7a1bfc786) |
| 10 | Omar (AE) | 3,000 | [0xbb8b0509...a24682e](https://amoy.polygonscan.com/tx/0xbb8b0509249dced00f7b5e30e15a382fcf862926deeb1e802ac2fd823a24682e) |
| 10 | Wei (SG) | 4,000 | [0x79039e16...2f7013](https://amoy.polygonscan.com/tx/0x79039e162a4dce7f06a400ccfbdf9e9108fcb058501130be9b75586f8d2f7013) |
| 11 | Bob (US) | 5,000 | [0x2f5b48c2...9c57f9](https://amoy.polygonscan.com/tx/0x2f5b48c2fcf0578ea4ea00304b5b7eef722c07ebd3c5601422e0faaa8d9c57f9) |
| 11 | Marie (FR) | 3,000 | [0xe28be7c7...e61645](https://amoy.polygonscan.com/tx/0xe28be7c7b77498c70e76b10136b2d3162c7a3b5b58c7ff16cf5e4b3b29e61645) |
| 11 | James (GB) | 2,000 | [0x9ed9b8ae...c8538b](https://amoy.polygonscan.com/tx/0x9ed9b8aea0f958340f1e2557ef67dcd03a40f48cb5fe522752f5cd2e25c8538b) |
| 12 | Yuki (JP) | 8,000 | [0x005f957b...98cca1](https://amoy.polygonscan.com/tx/0x005f957b50c2fd21ffc47d126019a49b77e269cbbee5f03816cfb0506098cca1) |
| 12 | Wei (SG) | 5,000 | [0x712e6319...9e24a1](https://amoy.polygonscan.com/tx/0x712e631958a4bd6297bfe9efc05effb3c97f29f76a85e273bef65f73999e24a1) |
| 12 | Omar (AE) | 2,000 | [0x71653849...de285f](https://amoy.polygonscan.com/tx/0x716538499f95c73b2198db9421437240f3788045d74b10fa17b6e961fede285f) |
| 13 | Klaus (CH) | 500 | [0x4e4c3156...b7645a](https://amoy.polygonscan.com/tx/0x4e4c3156d2ee38ec9182bbd2fe5193938fe302fa6273d37c6beec9b35cb7645a) |
| 13 | Alice (BR) | 300 | [0x3e92026c...982972](https://amoy.polygonscan.com/tx/0x3e92026c66510d65a30ddca713a2adb2cb24d6719f59d7d8dfd53286e4982972) |
| 13 | Marie (FR) | 200 | [0x43518d7e...84ea2f](https://amoy.polygonscan.com/tx/0x43518d7e1f3ebcd2c1073598de402e26bd8e94e2aed1ba707801c2422084ea2f) |
| 14 | Marie (FR) | 500 | [0x0b05e449...5a43fe](https://amoy.polygonscan.com/tx/0x0b05e4490dce74b93efa0925d457d5899aeab4886b59ed3b2209bc881c5a43fe) |
| 14 | Carlos (PT) | 400 | [0xb62b6f74...b88c2c](https://amoy.polygonscan.com/tx/0xb62b6f748188d9724d3a70b771001c8539aa4350b6883677ebe63516b8b88c2c) |
| 14 | Klaus (CH) | 100 | [0x28608f62...c0406](https://amoy.polygonscan.com/tx/0x28608f62d2af8f4cf0ae9652ddd0e2e316cf0a6195ed59d3a64e266416ec0406) |
| 15 | Wei (SG) | 10,000 | [0xc7d98cf4...6a016a](https://amoy.polygonscan.com/tx/0xc7d98cf48b1b15bd1d46c7c4fe34724260b2e3ec031708914c6a8151636a016a) |
| 15 | Yuki (JP) | 8,000 | [0x2f35f901...14ea2a](https://amoy.polygonscan.com/tx/0x2f35f901a1688aeb7bb372bb127d6e7f15af5c6081cf1c2aff5489bff914ea2a) |
| 15 | Omar (AE) | 5,000 | [0x1e80477e...06de98](https://amoy.polygonscan.com/tx/0x1e80477e21b84bec25fb45e4ec22b657314522e9c37acfaa14599aa4c706de98) |
| 16 | Omar (AE) | 15,000 | [0xba2bd3ad...9c2d72](https://amoy.polygonscan.com/tx/0xba2bd3ad7d5da5b5c677f26009ca48e96f8b72fb8aeaef66d9a22ee7639c2d72) |
| 16 | Wei (SG) | 10,000 | [0x157450ca...de7283](https://amoy.polygonscan.com/tx/0x157450cab690b085063ebd59e70820ae2fd05a9d2fb36db88125207209de7283) |
| 16 | Klaus (CH) | 5,000 | [0x5d97cc49...ade084](https://amoy.polygonscan.com/tx/0x5d97cc49809a8089a8af68737ab4a941ccef3d6ce9b2a78063c9466f78ade084) |
| 17 | Klaus (CH) | 8,000 | [0x25ea06bc...860696](https://amoy.polygonscan.com/tx/0x25ea06bc98c16e4a967ee490291920ad21ad342c0f06c85271e58060d2860696) |
| 17 | Hans (DE) | 6,000 | [0xbe9d82fc...9d26f3](https://amoy.polygonscan.com/tx/0xbe9d82fc6ff4ae6305462a5b8d05523c7f7733679c526aa6ec1539963b9d26f3) |
| 17 | James (GB) | 4,000 | [0x0433c2ac...5312b](https://amoy.polygonscan.com/tx/0x0433c2acc9e8e9228e7ed81ca089c27dfc35a8518b4143b415a99bc06055312b) |
| 17 | Bob (US) | 2,000 | [0x8c575b76...efccf](https://amoy.polygonscan.com/tx/0x8c575b7697ae643e34a019564365e0d41d7e5e1335163bcf56e2cc5a458efccf) |
| 18 | Alice (BR) | 12,000 | [0x2cab454d...dcdb933](https://amoy.polygonscan.com/tx/0x2cab454daa4d11972f2db885f824df4417eb3f2476b74ec88e5c5174bdcdb933) |
| 18 | Carlos (PT) | 8,000 | [0x54919483...514bfd](https://amoy.polygonscan.com/tx/0x549194830e5ffccb29b1485ae9650ef41f2e815e16354011a94035024d514bfd) |
| 18 | Bob (US) | 5,000 | [0x36ee5484...2da7d62](https://amoy.polygonscan.com/tx/0x36ee5484d45c26a5d37f1f248c00e2ac43c8d04b303575d70659ea08a2da7d62) |
| 19 | James (GB) | 10,000 | [0xeaf405ce...ae281a8](https://amoy.polygonscan.com/tx/0xeaf405ce5a41d8186a1d76573ffc35550a1040ca445f557cb1f2283bcae281a8) |
| 19 | Hans (DE) | 7,000 | [0x3ed10dde...7242a1](https://amoy.polygonscan.com/tx/0x3ed10ddeddaf89726a47f9d978a024b16a4a6a83638d82a94539a0de677242a1) |
| 19 | Marie (FR) | 3,000 | [0x6b0263af...1eefaf](https://amoy.polygonscan.com/tx/0x6b0263af6847b3c475d8fe3269cbf3e5502c165967df3e91454b4820f01eefaf) |
| 20 | Omar (AE) | 20,000 | [0xe1e9c0a3...cca0da](https://amoy.polygonscan.com/tx/0xe1e9c0a3de7ee921ff1d2728a67a4a3aa08cca0da0d6bf32daed9611e3fc0b6f) |
| 20 | Wei (SG) | 15,000 | [0xdfa09cb1...fbc1be2](https://amoy.polygonscan.com/tx/0xdfa09cb1490c0bace00910fd9a426cb65ae09f2b75fb1864bb0802edefbc1be2) |
| 20 | Yuki (JP) | 10,000 | [0x1d0e7ba0...d61400](https://amoy.polygonscan.com/tx/0x1d0e7ba06cee24ce2523b1cdf55860e97c4538cd5af7f303e4e53d7d78d61400) |
| 20 | Alice (BR) | 5,000 | [0x7d816cd0...f4c106](https://amoy.polygonscan.com/tx/0x7d816cd051e2f10a69e8db8b15372b555dcc835ba3a2d01e963036fc52f4c106) |

---

### Phase 5: Compliance Module Configuration

#### External Modules

| Module | Address |
|--------|---------|
| CountryRestrictModule | [`0x0c15e06c...bfFa50D2`](https://amoy.polygonscan.com/address/0x0c15e06c36b07E44aEe6D49a75554bc7bfFa50D2) |
| MaxBalanceModule | [`0x4B3cCd1F...cD89B44`](https://amoy.polygonscan.com/address/0x4B3cCd1F7BB1aF5F41b73e7fE3010023FcD89B44) |
| MaxHoldersModule | [`0xC40Bf7bb...D9999a`](https://amoy.polygonscan.com/address/0xC40Bf7bb339DD4485b1F5c3c0C5FE78DACD9999a) |

#### 5a. Configure Restrictions

| Module | Token | Parameters | Tx |
|--------|-------|------------|-----|
| CountryRestrict | 7 | Block US(840), JP(392) | [0x6e15e9f3...38ae13](https://amoy.polygonscan.com/tx/0x6e15e9f3b434c88b9048f4079c57fa5c66ab3e43f6a29d28b8d3a7bdbe38ae13) |
| CountryRestrict | 8 | Block BR(76), AE(784) | [0x5b4278ab...03a022c](https://amoy.polygonscan.com/tx/0x5b4278abb611b5c66e89fcd4abbcf4674fa85a5d541d244235b86152603a022c) |
| CountryRestrict | 16 | Block US(840), GB(826) | [0xa5c28a7d...e8b9f1](https://amoy.polygonscan.com/tx/0xa5c28a7dd7607a89bc8248093dba13f774078857c4a790b6b315143308e8b9f1) |
| MaxBalance | 8 | Max 12,000/wallet | [0x05936ff9...3155de6](https://amoy.polygonscan.com/tx/0x05936ff9195769cc63a8c9068f7c0a63105a9ae63a41391ee629f0b033155de6) |
| MaxBalance | 10 | Max 5,000/wallet | [0x7b5e1176...2f8b665](https://amoy.polygonscan.com/tx/0x7b5e1176f31c9f7013c0c627d192219da768ea1f24227e666875de8ad2f8b665) |
| MaxBalance | 16 | Max 20,000/wallet | [0x472b5b09...418dff](https://amoy.polygonscan.com/tx/0x472b5b094aea5ae419779485cd2e0367c04539f5059dc4e04a292a5434418dff) |
| MaxHolders | 9 | Max 5 holders | [0xa4952fcf...99fc0351](https://amoy.polygonscan.com/tx/0xa4952fcf5c669caf9957b2d2693e701aced99fc0351cb64e6287d02e7b30c4cd) |

#### 5b. Plug Modules into Diamond

**Call:** `addComplianceModule(tokenId, module)` → emits `ComplianceModuleAdded(tokenId, module)`

| Token | Module | Tx |
|-------|--------|-----|
| 7 | CountryRestrict | [0x60942722...4192e7](https://amoy.polygonscan.com/tx/0x60942722a33870132ac5a2374f58125a169d914649699291231b1cc69e4192e7) |
| 8 | CountryRestrict | [0x417215a3...57d78fc](https://amoy.polygonscan.com/tx/0x417215a3f398356f99f332a6fa72683c7156440ee0d3fee481c43679757d78fc) |
| 8 | MaxBalance | [0x50d060c9...a80e6](https://amoy.polygonscan.com/tx/0x50d060c9b558e0ef2b413b3ae2f2614ceaa5c0b575e24263143684cc0f0a80e6) |
| 9 | MaxHolders | [0x305d9f2c...dc981](https://amoy.polygonscan.com/tx/0x305d9f2c67953c51ed56865bac4da94a373977939f5ab035b1a5b0b62b6dc981) |
| 10 | MaxBalance | [0xae45f200...62824](https://amoy.polygonscan.com/tx/0xae45f2007c1ecf847e9434b510db061b2e64187cf01a8aaeab61b30740662824) |
| 16 | CountryRestrict | [0x763c43b4...7834312](https://amoy.polygonscan.com/tx/0x763c43b47ab34914c53a3f86e179696f17834312a9cf5afe3caf27d185db7845) |
| 16 | MaxBalance | [0x0266d1ed...73f53307](https://amoy.polygonscan.com/tx/0x0266d1ed5dbbaed700631710e9de4c8706401bcda735c73654a9a34673f53307) |

#### Final Compliance Configuration

| Token | Modules | Rules |
|-------|---------|-------|
| 7 (BREF) | CountryRestrict | US, JP blocked |
| 8 (USTB26) | CountryRestrict + MaxBalance | BR, AE blocked; max 12k/wallet |
| 9 (ECC) | MaxHolders | Max 5 unique holders |
| 10 (GBT) | MaxBalance | Max 5k/wallet |
| 16 (DRET) | CountryRestrict + MaxBalance | US, GB blocked; max 20k/wallet |

---

### Phase 6: Compliance Transfer Tests

#### 6a. Expected Reverts (Compliance Rejection)

All 4 tests **reverted on-chain** with `ERC1155Facet__ComplianceRejected(tokenId, reason)` — status `0x0` (failed).

| Test | Token | From → To | Amount | Reason | Tx (reverted) |
|------|-------|-----------|--------|--------|----------------|
| A | 7 (BREF) | Alice(BR) → Bob(US) | 500 | `REASON_COUNTRY_RESTRICTED` — US blocked | [0x893f1370...c6824149](https://amoy.polygonscan.com/tx/0x893f13709a016c35433adfa987bda03e2eb56868b6c82c8c8da0481bc6824149) |
| B | 8 (USTB26) | Bob(US) → Alice(BR) | 100 | `REASON_COUNTRY_RESTRICTED` — BR blocked | [0x52eed900...ec767e94](https://amoy.polygonscan.com/tx/0x52eed90072502cd3feb82ac8f20f39b7be5638d251f75d985db10a4cec767e94) |
| C | 10 (GBT) | Alice → Wei | 2,000 | `REASON_HOLDING_LIMIT` — Wei 4k+2k > 5k max | [0x54dd6da0...df824e9d](https://amoy.polygonscan.com/tx/0x54dd6da0ebd886d5bb6cf43de3d379bc5cbbd59f3cec356a3a58b5d0df824e9d) |
| D | 9 (ECC) | Hans → Alice | 100 | `MAX_HOLDERS_EXCEEDED` — 4 holders + Alice > 5 max | [0x18f75214...8650d82f](https://amoy.polygonscan.com/tx/0x18f752146fa84ab0771afeec64c03448f90542d6a6dd4782be80bd478650d82f) |

**Nota:** O Polygonscan exibe apenas "execution reverted" genérico. O motivo real está no **revert data** da tx — decodificável via `cast`:

```
Error selector: 0xd8c2ffe6 = ERC1155Facet__ComplianceRejected(uint256 tokenId, bytes32 reason)
```

**Reason codes (bytes32):**
```
COUNTRY_RESTRICTED = 0xca38fc82...fe49d6e  → keccak256("REASON_COUNTRY_RESTRICTED")
HOLDING_LIMIT      = 0x01e653fe...d90dc3c0 → keccak256("REASON_HOLDING_LIMIT")
MAX_HOLDERS        = 0x34801528...5643baa0 → keccak256("MAX_HOLDERS_EXCEEDED")
```

#### 6b. Successful Transfers

| Token | From → To | Amount | Why OK | Tx |
|-------|-----------|--------|--------|-----|
| 7 | Alice(BR) → Carlos(PT) | 500 | PT not blocked | [0x466768cf...01ab5fe](https://amoy.polygonscan.com/tx/0x466768cf93bddd71fbf395d56ceb9526d0fe80ce5f4154b9f7f2b191601ab5fe) |
| 8 | Bob(US) → Klaus(CH) | 2,000 | CH ok; 7k < 12k max | [0x3629255f...ec1125](https://amoy.polygonscan.com/tx/0x3629255fff543989a239b2a2ebc2ae4ec17cf58764fcfa167e158cc6beec1125) |
| 10 | Alice → Yuki | 500 | 2.5k < 5k max | [0xe148097a...1989be](https://amoy.polygonscan.com/tx/0xe148097ac1b65bfc896965f3dc1ddaaf1aebf3b74ec880e9d5d1a0ad4b1989be) |
| 12 | Yuki(JP) → Wei(SG) | 1,000 | No modules | [0x5ec5c891...db96c83](https://amoy.polygonscan.com/tx/0x5ec5c89144f895fe56b67562377ed1817267eac61965b1ccf3795bbccdb96c83) |
| 20 | Omar(AE) → Alice(BR) | 3,000 | No modules | [0xa6470116...959a4e](https://amoy.polygonscan.com/tx/0xa647011681f89624ff54644345a4ef23f8e82d22ff03b3838d1b362538959a4e) |

**Events per transfer:**
```
TransferSingle(operator, from, to, tokenId, amount)
RegulatoryTransfer(tokenId, from, to, amount, REASON_OK)
```

---

### Phase 7: YieldDistributor — Token 7 (BREF) + Token 10 (GBT)

#### 7a. Add Plugin to Tokens

| Token | Tx |
|-------|-----|
| 7 (BREF) | [0xa3666623...50030830](https://amoy.polygonscan.com/tx/0xa3666623e98a56d1f576b72d6a3028f63f17b46437c7be04e2e2e37150030830) |
| 10 (GBT) | [0x13b2400f...35b25e7b](https://amoy.polygonscan.com/tx/0x13b2400f4eb99e6bf427bc77cd08fc517a26978fbadbd35b29131d0235b25e7b) |

#### 7b. Register Reward Assets

| Staked Token | Reward | Type | Tx |
|-------------|--------|------|-----|
| 7 (BREF) | tokenId 2 (rUSDT) | ERC-1155 | [0x23b9a97a...ea95b28d](https://amoy.polygonscan.com/tx/0x23b9a97a25485a620c0f71f67c1280331aaef39830d938a553b5b90aea95b28d) |
| 7 (BREF) | tokenId 5 (SRP) | ERC-1155 | [0x2904d96b...4a2a6f5](https://amoy.polygonscan.com/tx/0x2904d96b3de9d7068e360508e93a10b1e4bff85db73837501b1512e3d4a2a6f5) |
| 10 (GBT) | tokenId 3 (rUSDC) | ERC-1155 | [0xe2c5f4f9...40f77bc](https://amoy.polygonscan.com/tx/0xe2c5f4f96e1cafbef58b42551a5b87e297cab5c3d3d791de040f77bc885ad070) |

#### 7c. Deposit Yield

**Approval:** [0xb4ffeb44...6b71315f](https://amoy.polygonscan.com/tx/0xb4ffeb44dcf4f85d2e0da797c4f552ef098e5d2251a2a9171d6001ca6b71315f) — `setApprovalForAll(YieldDistributor, true)`

| Staked Token | Reward | Amount | Tx |
|-------------|--------|--------|-----|
| 7 (BREF) | rUSDT | 10,000 | [0x95674772...1581ba08](https://amoy.polygonscan.com/tx/0x9567477288960171e94c151fa9a5238afeb132c6309b535cc0635eaa1581ba08) |
| 7 (BREF) | SRP | 5,000 | [0xee4028e9...50c18c2a](https://amoy.polygonscan.com/tx/0xee4028e9d474517f0a0a557ad2db1eb3638522df688a843f91149ab750c18c2a) |
| 10 (GBT) | rUSDC | 8,000 | [0x730079c9...0936569f](https://amoy.polygonscan.com/tx/0x730079c9c59d518b3eed639e33717609c9346152f00ad5b1cc57b5870936569f) |

#### 7d. Claims (Alice & Bob)

| Holder | Reward | Claimed | Tx |
|--------|--------|---------|-----|
| Alice | rUSDT | 3,166 | [0x4bcdcd7b...e2fc4dc5](https://amoy.polygonscan.com/tx/0x4bcdcd7b70c080b1bd0bfb2406b7935aa6607554f8b63ea9a2d65f05e2fc4dc5) |
| Alice | SRP | 1,583 | [0x815c5bf7...14933df5](https://amoy.polygonscan.com/tx/0x815c5bf7647e6d7d48b1381d8669e4a95f6c0f5082b0eaa4bc6614933df5427f) |
| Bob | rUSDT+SRP | 1,999+999 | [0x4e3e3c64...2643b947](https://amoy.polygonscan.com/tx/0x4e3e3c64d876cdff83bd528a3c0ecf73aca80816470b238fc6959ffc2643b947) |

---

### Phase 8: Full 10-Holder Reward Distribution (Token 11 — Silver)

This phase demonstrates yield distribution to **all 10 investors** from a single deposit.

#### 8a. Mint SCT (token 11) to All 10 Holders

7 additional mints (Bob, Marie, James already held SCT):

| Investor | Amount | Tx |
|----------|--------|-----|
| Alice | 2,000 | [0x2fe07bb0...f7f29ec0](https://amoy.polygonscan.com/tx/0x2fe07bb0c6f9057e6ca42f9d86f059ba3858274a901381bfa0cc28ccf7f29ec0) |
| Carlos | 1,500 | [0xeeacb024...7e30caad](https://amoy.polygonscan.com/tx/0xeeacb02421b6ab24598e8b9f4f1a4f91d1fd20c5ef628ae08f86664c7e30caad) |
| Yuki | 3,000 | [0xd469f8e9...bf6b998](https://amoy.polygonscan.com/tx/0xd469f8e979acb968e98ef69571a238697ec6d9a04a003567288bdd795bf6b998) |
| Hans | 2,500 | [0x42dd74bc...e2a5144](https://amoy.polygonscan.com/tx/0x42dd74bcc6f5552d27c717c57dae83d55f76045cb5f48d188a684a0abe2a5144) |
| Klaus | 1,000 | [0x6fb643cb...a184467](https://amoy.polygonscan.com/tx/0x6fb643cb0ea2b06030462d63128b8ce7ceb627af42ee1cd4cd93427eba184467) |
| Wei | 4,000 | [0x9b9ff5c1...4209ef4](https://amoy.polygonscan.com/tx/0x9b9ff5c1d084a16d19a42caccced600ea7cf5753c2db78530d41272774209ef4) |
| Omar | 3,000 | [0x66883321...49b0523](https://amoy.polygonscan.com/tx/0x66883321e62d703520c41b2633fcb1781762345c9e45e76fe33469b4049b0523) |

**Post-mint state:** 13 holders (10 active + 3 legacy), totalSupply = 37,000

#### 8b. Setup YieldDistributor on Token 11

| Step | Tx |
|------|-----|
| Add plugin | [0x1b27738e...05302b67](https://amoy.polygonscan.com/tx/0x1b27738ea88845debd79776e29390ac81455b57b1e965d685aead75a05302b67) |
| Register rDAI (tokenId 4) as reward | [0x4726f795...4303ef9](https://amoy.polygonscan.com/tx/0x4726f7957683f06797b1dcd18967dabf0d4dd24f3fccd123e753ed3554303ef9) |
| Deposit 37,000 rDAI | [0x7289f507...826ba638](https://amoy.polygonscan.com/tx/0x7289f5072a8649901f455e94c9c47cda5e60660f4b8d8587c1877639826ba638) |

**Accumulator math:**
```
totalSupply(11) = 37,000
deposit = 37,000 rDAI
accRewardPerShare = 37,000 * 1e36 / 37,000 = 1e36 (exactly 1:1)
→ each holder receives exactly as many rDAI as they hold SCT
```

#### 8c. Claimable Yields (pre-claim)

| Holder | SCT Balance | Share | Claimable rDAI |
|--------|-------------|-------|----------------|
| Alice | 2,000 | 5.4% | 2,000 |
| Bob | 5,000 | 13.5% | 5,000 |
| Carlos | 1,500 | 4.1% | 1,500 |
| Yuki | 3,000 | 8.1% | 3,000 |
| Hans | 2,500 | 6.8% | 2,500 |
| Marie | 3,000 | 8.1% | 3,000 |
| James | 2,000 | 5.4% | 2,000 |
| Klaus | 1,000 | 2.7% | 1,000 |
| Wei | 4,000 | 10.8% | 4,000 |
| Omar | 3,000 | 8.1% | 3,000 |
| **Total** | **27,000** | **73.0%** | **27,000** |

*Remaining 10,000 rDAI allocated to legacy Anvil wallets (unclaimed).*

#### 8d. All 10 Holders Claim rDAI

Each holder calls `claimAllYield(11)` — emits `YieldClaimed(11, rewardKey, holder, amount)` + `TransferSingle(yieldModule, yieldModule, holder, 4, amount)`.

| Holder | rDAI Before | rDAI After | Claimed | Tx |
|--------|-------------|------------|---------|-----|
| Alice | 0 | 2,000 | 2,000 | [0xa8846840...e1bc1b0](https://amoy.polygonscan.com/tx/0xa8846840b2caf9ba1a5e5c22d9f71c75c40c00f19e4616186dd3d58fae1bc1b0) |
| Bob | 0 | 5,000 | 5,000 | [0x0e8e2bee...b52357c](https://amoy.polygonscan.com/tx/0x0e8e2bee4e70555ae0670d13a5d36eeb4c68ea1d5a4aabce89c462ad0b52357c) |
| Carlos | 0 | 1,500 | 1,500 | [0x65efff98...174bff0](https://amoy.polygonscan.com/tx/0x65efff98c70e72189b7532f252563fff9f0cd47fe27397100405804bd174bff0) |
| Yuki | 0 | 3,000 | 3,000 | [0x797b6c2f...7235018](https://amoy.polygonscan.com/tx/0x797b6c2ffba74927ac11cddfca4baabbec2439046dbed1d27e66337e37235018) |
| Hans | 0 | 2,500 | 2,500 | [0xefcabd3a...938b0c8](https://amoy.polygonscan.com/tx/0xefcabd3a13e58fd4aaa6112a9c61bb7159511b52f3f4d08b17031bc5d938b0c8) |
| Marie | 0 | 3,000 | 3,000 | [0xc0d8d303...cacaf38](https://amoy.polygonscan.com/tx/0xc0d8d303d835dd3c9d5366195c592b92818b5058f607529a8f8a960eacacaf38) |
| James | 0 | 2,000 | 2,000 | [0xf6b76d4d...d542f64](https://amoy.polygonscan.com/tx/0xf6b76d4d63456e33d8064907d60441d1151ad4d01faee39225835677dd542f64) |
| Klaus | 0 | 1,000 | 1,000 | [0x33542cbb...072cfb5a](https://amoy.polygonscan.com/tx/0x33542cbbf995bdc03b58df0504509803c39669232b229dc97b507c50072cfb5a) |
| Wei | 0 | 4,000 | 4,000 | [0xf13c9281...aafd80f](https://amoy.polygonscan.com/tx/0xf13c9281bfaeba7f9b9e652a984286be725fd2905af1040f972a48807aafd80f) |
| Omar | 0 | 3,000 | 3,000 | [0x2f76fda8...2a75ad](https://amoy.polygonscan.com/tx/0x2f76fda8ff2c5c1b6ee42dd8cf5f04336e5a70b1bcd04742298beb041d2a75ad) |

**Total claimed: 27,000 rDAI** across 10 independent transactions.

---

## Summary of Demonstrated Features

| Feature | Status | Tokens |
|---------|--------|--------|
| Asset registration (20 tokens) | ✅ | 1-20 |
| Batch identity registration (10 wallets) | ✅ | — |
| Minting with TransferSingle events | ✅ | 2-20 |
| CountryRestrict compliance (per-tokenId) | ✅ | 7, 8, 16 |
| MaxBalance compliance (per-tokenId) | ✅ | 8, 10, 16 |
| MaxHolders compliance (per-tokenId) | ✅ | 9 |
| Transfer revert on country block | ✅ | 7, 8 |
| Transfer revert on balance limit | ✅ | 10 |
| Transfer revert on holder limit | ✅ | 9 |
| Successful compliant transfers | ✅ | 7, 8, 10, 12, 20 |
| Multi-module compliance (2 modules) | ✅ | 8, 16 |
| YieldDistributor plugin setup | ✅ | 7, 10, 11 |
| Multi-asset rewards (rUSDT + SRP) | ✅ | 7 |
| Cross-tokenId ERC-1155 rewards | ✅ | 7, 10, 11 |
| Yield deposit with accumulator | ✅ | 7, 10, 11 |
| Yield claim (single + claimAll) | ✅ | 7, 11 |
| **Full 10-holder reward distribution** | ✅ | **11** |

---

## Contracts

| Contract | Address |
|----------|---------|
| Diamond | [`0xAb07CEf1BEeDBb30F5795418c79879794b31C521`](https://amoy.polygonscan.com/address/0xAb07CEf1BEeDBb30F5795418c79879794b31C521) |
| CountryRestrictModule | [`0x0c15e06c36b07E44aEe6D49a75554bc7bfFa50D2`](https://amoy.polygonscan.com/address/0x0c15e06c36b07E44aEe6D49a75554bc7bfFa50D2) |
| MaxBalanceModule | [`0x4B3cCd1F7BB1aF5F41b73e7fE3010023FcD89B44`](https://amoy.polygonscan.com/address/0x4B3cCd1F7BB1aF5F41b73e7fE3010023FcD89B44) |
| MaxHoldersModule | [`0xC40Bf7bb339DD4485b1F5c3c0C5FE78DACD9999a`](https://amoy.polygonscan.com/address/0xC40Bf7bb339DD4485b1F5c3c0C5FE78DACD9999a) |
| YieldDistributorModule | [`0xa8c6A9AAfd18545bEe6BC734eba702C55beA95dF`](https://amoy.polygonscan.com/address/0xa8c6A9AAfd18545bEe6BC734eba702C55beA95dF) |
