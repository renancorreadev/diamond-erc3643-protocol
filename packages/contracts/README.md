# Diamond ERC-3643 — Contracts

Solidity smart contracts for the Diamond ERC-3643 protocol. See the [root README](../../README.md) for full architecture documentation.

## Quick Start

```bash
forge install
make build       # forge build --sizes
make test        # all 476 tests
```

## Structure

```
src/
├── Diamond.sol                    # EIP-2535 proxy
├── facets/                        # 21 facets by domain
│   ├── core/                      # DiamondCut, Loupe, Ownership
│   ├── token/                     # ERC1155, Supply, Metadata, AssetManager
│   ├── security/                  # AccessControl, Pause, Emergency
│   ├── identity/                  # IdentityRegistry, ClaimTopics, TrustedIssuer
│   ├── compliance/                # ComplianceRouterFacet
│   ├── plugins/                   # GlobalPluginFacet
│   ├── routers/                   # PluginRouterFacet
│   └── rwa/                       # Freeze, Recovery, Snapshot, Dividend, AssetGroup
├── compliance/modules/            # Pluggable compliance modules (gating)
│   ├── CountryRestrictModule.sol
│   ├── MaxBalanceModule.sol
│   └── MaxHoldersModule.sol
├── plugins/modules/               # Pluggable plugin modules (reactive)
│   └── YieldDistributorModule.sol
├── interfaces/
│   ├── plugins/                   # IPluginModule, IHookablePlugin
│   └── ...
├── storage/                       # 12 namespaced storage libraries
├── libraries/                     # LibDiamond, LibAppStorage
└── initializers/                  # DiamondInit

test/
├── unit/                          # 1:1 with facets + modules
│   ├── GlobalPluginFacet.t.sol    # 30 tests
│   └── modules/plugins/           # YieldDistributorModule tests (50 tests)
├── fuzz/                          # Property-based fuzz tests
├── invariant/                     # FREI-PI invariant tests
├── echidna/                       # Echidna fuzzing
└── helpers/                       # DiamondHelper, MockERC20, MockPluginModule

script/
├── Deploy.s.sol                   # Full deploy (21 facets + Diamond + EIP-1967)
└── ...
```

## Commands

```bash
make build                    # forge build --sizes
make test                     # all 476 tests
make test-vvv                 # verbose
make test-fuzz                # fuzz tests (10k runs)
make test-contract CONTRACT=YieldDistributorModuleTest
make slither                  # static analysis
make coverage                 # forge coverage
make deploy-local             # deploy to local Anvil
make deploy-amoy              # deploy to Polygon Amoy
```
