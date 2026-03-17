# ─────────────────────────────────────────────────────────────────────────────
# diamond-erc3643 — contracts Makefile
# ─────────────────────────────────────────────────────────────────────────────

.DEFAULT_GOAL := help
.PHONY: help build clean fmt fmt-check lint lint-fix \
        test test-v test-vv test-vvv test-vvvv \
        test-unit test-unit-v test-unit-vv test-unit-vvv \
        test-fuzz test-invariant test-fork \
        test-contract \
        coverage slither \
        deploy-local deploy-dry deploy-amoy deploy-compliance \
        set-implementation verify-all configure-owner snapshot

# ─────────────────────────────────────────────────────────────────────────────
# Config — override via env or CLI: make deploy-amoy RPC_AMOY=...
# ─────────────────────────────────────────────────────────────────────────────

RPC_LOCAL   ?= http://localhost:8545
RPC_AMOY    ?= $(AMOY_RPC_URL)
CHAIN_ID    ?= 80002
FUZZ_RUNS   ?= 10000
DEPTH       ?= 500
CONTRACT    ?=   # e.g. make test-contract CONTRACT=OwnershipFacet

# ─────────────────────────────────────────────────────────────────────────────
# Help
# ─────────────────────────────────────────────────────────────────────────────

help:
	@echo ""
	@echo "  BUILD"
	@echo "    make build             forge build --sizes"
	@echo "    make clean             forge clean"
	@echo "    make snapshot          forge snapshot"
	@echo ""
	@echo "  FORMAT / LINT"
	@echo "    make fmt               forge fmt (write)"
	@echo "    make fmt-check         forge fmt --check"
	@echo "    make lint              solhint src/**/*.sol"
	@echo "    make lint-fix          solhint src/**/*.sol --fix"
	@echo ""
	@echo "  TEST — verbosity"
	@echo "    make test              all tests, -v"
	@echo "    make test-v            all tests, -v"
	@echo "    make test-vv           all tests, -vv"
	@echo "    make test-vvv          all tests, -vvv"
	@echo "    make test-vvvv         all tests, -vvvv"
	@echo ""
	@echo "  TEST — by type"
	@echo "    make test-unit         test/unit/**  -vvv"
	@echo "    make test-unit-v       test/unit/**  -v"
	@echo "    make test-unit-vv      test/unit/**  -vv"
	@echo "    make test-unit-vvv     test/unit/**  -vvv"
	@echo "    make test-fuzz         test/fuzz/**  --fuzz-runs \$$FUZZ_RUNS"
	@echo "    make test-invariant    test/invariant/** --fuzz-runs \$$FUZZ_RUNS"
	@echo "    make test-fork         test/fork/**"
	@echo ""
	@echo "  TEST — by contract"
	@echo "    make test-contract CONTRACT=OwnershipFacet"
	@echo "    make test-contract CONTRACT=Diamond"
	@echo ""
	@echo "  COVERAGE / ANALYSIS"
	@echo "    make coverage          forge coverage --report lcov"
	@echo "    make slither           slither . --config-file slither.config.json"
	@echo ""
	@echo "  DEPLOY — local (Anvil)"
	@echo "    make deploy-local      Deploy.s.sol against local Anvil"
	@echo "    make deploy-dry        Deploy.s.sol simulation (no broadcast)"
	@echo ""
	@echo "  DEPLOY — Polygon Amoy"
	@echo "    make deploy-amoy       Full Diamond deploy to Amoy"
	@echo "    make deploy-compliance Deploy compliance modules to Amoy"
	@echo "    make set-implementation Deploy DiamondABI stub + set EIP-1967 slot"
	@echo "    make verify-all        Verify all contracts on Polygonscan"
	@echo "    make configure-owner   Grant all roles to OWNER_ADDRESS"
	@echo ""

# ─────────────────────────────────────────────────────────────────────────────
# Build
# ─────────────────────────────────────────────────────────────────────────────

build:
	forge build --sizes

clean:
	forge clean

snapshot:
	forge snapshot

# ─────────────────────────────────────────────────────────────────────────────
# Format / Lint
# ─────────────────────────────────────────────────────────────────────────────

fmt:
	forge fmt

fmt-check:
	forge fmt --check

lint:
	npx solhint 'src/**/*.sol'

lint-fix:
	npx solhint 'src/**/*.sol' --fix

# ─────────────────────────────────────────────────────────────────────────────
# Test — all tests, varying verbosity
# ─────────────────────────────────────────────────────────────────────────────

test: test-v

test-v:
	forge test -v

test-vv:
	forge test -vv

test-vvv:
	forge test -vvv

test-vvvv:
	forge test -vvvv

# ─────────────────────────────────────────────────────────────────────────────
# Test — by type
# ─────────────────────────────────────────────────────────────────────────────

test-unit:
	forge test --match-path 'test/unit/**' -vvv

test-unit-v:
	forge test --match-path 'test/unit/**' -v

test-unit-vv:
	forge test --match-path 'test/unit/**' -vv

test-unit-vvv:
	forge test --match-path 'test/unit/**' -vvv

test-fuzz:
	forge test --match-path 'test/fuzz/**' --fuzz-runs $(FUZZ_RUNS) -vvv

test-invariant:
	forge test --match-path 'test/invariant/**' \
		--fuzz-runs $(FUZZ_RUNS) \
		--depth $(DEPTH) \
		-vvv

test-fork:
	forge test --match-path 'test/fork/**' -vvv

# ─────────────────────────────────────────────────────────────────────────────
# Test — by contract name
# usage: make test-contract CONTRACT=OwnershipFacet
# ─────────────────────────────────────────────────────────────────────────────

test-contract:
ifndef CONTRACT
	$(error CONTRACT is required. Usage: make test-contract CONTRACT=OwnershipFacet)
endif
	forge test --match-contract $(CONTRACT) -vvv

# ─────────────────────────────────────────────────────────────────────────────
# Coverage / Analysis
# ─────────────────────────────────────────────────────────────────────────────

coverage:
	forge coverage --report lcov

slither:
	slither . --config-file slither.config.json

# ─────────────────────────────────────────────────────────────────────────────
# Deploy — local Anvil
# ─────────────────────────────────────────────────────────────────────────────

deploy-local:
	ETHERSCAN_API_KEY=dummy \
	forge script script/Deploy.s.sol \
		--rpc-url $(RPC_LOCAL) \
		--broadcast \
		--sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
		--unlocked \
		-vvvv

deploy-dry:
	ETHERSCAN_API_KEY=dummy \
	forge script script/Deploy.s.sol \
		--rpc-url $(RPC_LOCAL) \
		-vvvv

# ─────────────────────────────────────────────────────────────────────────────
# Deploy — Polygon Amoy Testnet
# Requires: PRIVATE_KEY, AMOY_RPC_URL, POLYGONSCAN_API_KEY in .env or env
# ─────────────────────────────────────────────────────────────────────────────

deploy-amoy:
	@echo "=== Step 1: Deploy Diamond + all facets ==="
	forge script script/Deploy.s.sol \
		--tc Deploy \
		--rpc-url $(RPC_AMOY) \
		--private-key $(PRIVATE_KEY) \
		--broadcast \
		--verify \
		--etherscan-api-key $(POLYGONSCAN_API_KEY) \
		--chain-id $(CHAIN_ID) \
		-vvvv

deploy-compliance:
	@echo "=== Deploy compliance modules ==="
	forge script script/DeployComplianceModules.s.sol \
		--rpc-url $(RPC_AMOY) \
		--private-key $(PRIVATE_KEY) \
		--broadcast \
		--verify \
		--etherscan-api-key $(POLYGONSCAN_API_KEY) \
		--chain-id $(CHAIN_ID) \
		-vvvv

set-implementation:
	@echo "=== Deploy DiamondABI stub + set EIP-1967 implementation slot ==="
	forge script script/SetDiamondImplementation.s.sol \
		--tc SetDiamondImplementation \
		--rpc-url $(RPC_AMOY) \
		--private-key $(PRIVATE_KEY) \
		--broadcast \
		--verify \
		--etherscan-api-key $(POLYGONSCAN_API_KEY) \
		--chain-id $(CHAIN_ID) \
		-vvvv

# ─────────────────────────────────────────────────────────────────────────────
# Verify — manual verification for contracts that failed auto-verify
# Usage: make verify-contract ADDR=0x... NAME=DiamondABI
# ─────────────────────────────────────────────────────────────────────────────

verify-contract:
ifndef ADDR
	$(error ADDR is required. Usage: make verify-contract ADDR=0x... NAME=DiamondABI)
endif
ifndef NAME
	$(error NAME is required. Usage: make verify-contract ADDR=0x... NAME=DiamondABI)
endif
	forge verify-contract $(ADDR) $(NAME) \
		--chain-id $(CHAIN_ID) \
		--etherscan-api-key $(POLYGONSCAN_API_KEY) \
		--watch

# ─────────────────────────────────────────────────────────────────────────────
# Configure — grant all roles to OWNER_ADDRESS
# Requires: DIAMOND_ADDRESS, OWNER_ADDRESS, PRIVATE_KEY, AMOY_RPC_URL
# ─────────────────────────────────────────────────────────────────────────────

configure-owner:
ifndef OWNER_ADDRESS
	$(error OWNER_ADDRESS is required)
endif
ifndef DIAMOND_ADDRESS
	$(error DIAMOND_ADDRESS is required)
endif
	@echo "=== Granting all roles to $(OWNER_ADDRESS) ==="
	@for ROLE in \
		"GOVERNANCE_ROLE" \
		"UPGRADER_ROLE" \
		"PAUSER_ROLE" \
		"ISSUER_ROLE" \
		"COMPLIANCE_ADMIN" \
		"TRANSFER_AGENT" \
		"RECOVERY_AGENT" \
		"CLAIM_ISSUER_ROLE"; do \
		HASH=$$(cast keccak "$$ROLE"); \
		echo "Granting $$ROLE ($$HASH)..."; \
		cast send $(DIAMOND_ADDRESS) \
			"grantRole(bytes32,address)" \
			"$$HASH" $(OWNER_ADDRESS) \
			--rpc-url $(RPC_AMOY) \
			--private-key $(PRIVATE_KEY); \
	done
	@echo "=== All roles granted ==="

# ─────────────────────────────────────────────────────────────────────────────
# Full deploy pipeline: build → test → deploy → compliance → eip1967 → roles
# ─────────────────────────────────────────────────────────────────────────────

deploy-full: build test deploy-amoy deploy-compliance set-implementation
	@echo ""
	@echo "=== Full deploy complete ==="
	@echo "Next steps:"
	@echo "  1. Update DIAMOND_ADDRESS in .env / packages/app/.env / deployments/amoy.json"
	@echo "  2. Run: make configure-owner OWNER_ADDRESS=0x... DIAMOND_ADDRESS=0x..."
	@echo "  3. Register identity: cast send \$$DIAMOND 'registerIdentity(address,address,uint16)' ..."
	@echo "  4. Create profile:    cast send \$$DIAMOND 'createProfile(uint256[])' ..."
	@echo "  5. Register asset:    cast send \$$DIAMOND 'registerAsset(...)' ..."
	@echo "  6. Mint tokens:       cast send \$$DIAMOND 'mint(uint256,address,uint256)' ..."
	@echo ""
