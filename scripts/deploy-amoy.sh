#!/usr/bin/env bash
set -euo pipefail

# ── Deploy Diamond ERC-3643 to Polygon Amoy + update all configs ──
#
# Usage:
#   # Option 1: private key in env
#   PRIVATE_KEY=0x... ./scripts/deploy-amoy.sh
#
#   # Option 2: cast wallet (import first: cast wallet import deployer --interactive)
#   ACCOUNT=deployer ./scripts/deploy-amoy.sh

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
CONTRACTS_DIR="$ROOT_DIR/packages/contracts"

# Load .env
set -a
source "$ROOT_DIR/.env"
set +a

# Colors
GREEN='\033[0;32m'
CYAN='\033[0;36m'
RED='\033[0;31m'
NC='\033[0m'

# ── Auth flags ────────────────────────────────────────────
AUTH_FLAGS=""
if [[ -n "${PRIVATE_KEY:-}" ]]; then
  AUTH_FLAGS="--private-key $PRIVATE_KEY"
elif [[ -n "${ACCOUNT:-}" ]]; then
  AUTH_FLAGS="--account $ACCOUNT"
else
  echo -e "${RED}Error: set PRIVATE_KEY or ACCOUNT env var${NC}"
  echo "  PRIVATE_KEY=0x... ./scripts/deploy-amoy.sh"
  echo "  ACCOUNT=deployer  ./scripts/deploy-amoy.sh"
  exit 1
fi

OWNER_ADDRESS="${OWNER_ADDRESS:-0xB40061C7bf8394eb130Fcb5EA06868064593BFAa}"

echo -e "${CYAN}═══════════════════════════════════════════════════${NC}"
echo -e "${CYAN}  Diamond ERC-3643 — Fresh Deploy to Amoy${NC}"
echo -e "${CYAN}═══════════════════════════════════════════════════${NC}"
echo ""

# ── Step 1: Build ─────────────────────────────────────────
echo -e "${GREEN}[1/4]${NC} Building contracts..."
cd "$CONTRACTS_DIR"
forge build --sizes 2>&1 | tail -3

# ── Step 2: Deploy Diamond + all facets ───────────────────
echo ""
echo -e "${GREEN}[2/4]${NC} Deploying Diamond..."
DEPLOY_OUTPUT=$(OWNER_ADDRESS="$OWNER_ADDRESS" \
  forge script script/Deploy.s.sol --tc Deploy \
    --rpc-url "$AMOY_RPC_URL" \
    $AUTH_FLAGS \
    --broadcast -vvv 2>&1)

echo "$DEPLOY_OUTPUT" | grep -E "(Diamond |Facet |Init |===|---| OK|Facets|Error)" | head -30

# Extract Diamond address
DIAMOND=$(echo "$DEPLOY_OUTPUT" | grep "Diamond              :" | head -1 | awk '{print $NF}')
if [[ -z "$DIAMOND" ]]; then
  echo -e "${RED}Failed to extract Diamond address from deploy output${NC}"
  echo "$DEPLOY_OUTPUT" | tail -20
  exit 1
fi

echo ""
echo -e "${GREEN}Diamond deployed:${NC} $DIAMOND"

# ── Step 3: Deploy compliance modules ─────────────────────
echo ""
echo -e "${GREEN}[3/4]${NC} Deploying compliance modules..."
MODULES_OUTPUT=$(DIAMOND_ADDRESS="$DIAMOND" OWNER_ADDRESS="$OWNER_ADDRESS" \
  forge script script/DeployComplianceModules.s.sol \
    --rpc-url "$AMOY_RPC_URL" \
    $AUTH_FLAGS \
    --broadcast -vvv 2>&1)

echo "$MODULES_OUTPUT" | grep -E "(Module|Diamond|Owner|===)" | head -10

COUNTRY_MODULE=$(echo "$MODULES_OUTPUT" | grep "CountryRestrictModule:" | awk '{print $NF}')
MAX_BALANCE_MODULE=$(echo "$MODULES_OUTPUT" | grep "MaxBalanceModule" | awk '{print $NF}')
MAX_HOLDERS_MODULE=$(echo "$MODULES_OUTPUT" | grep "MaxHoldersModule" | awk '{print $NF}')

# ── Step 4: Extract all addresses and update configs ──────
echo ""
echo -e "${GREEN}[4/4]${NC} Updating configuration files..."

# Extract all facet addresses from deploy output
extract_addr() {
  echo "$DEPLOY_OUTPUT" | grep "$1" | head -1 | awk '{print $NF}'
}

DIAMOND_CUT=$(extract_addr "DiamondCutFacet")
DIAMOND_LOUPE=$(extract_addr "DiamondLoupeFacet")
OWNERSHIP=$(extract_addr "OwnershipFacet")
ACCESS_CONTROL=$(extract_addr "AccessControlFacet")
PAUSE=$(extract_addr "PauseFacet")
EMERGENCY=$(extract_addr "EmergencyFacet")
FREEZE=$(extract_addr "FreezeFacet")
RECOVERY=$(extract_addr "RecoveryFacet")
SNAPSHOT=$(extract_addr "SnapshotFacet")
DIVIDEND=$(extract_addr "DividendFacet")
ASSET_GROUP=$(extract_addr "AssetGroupFacet")
ASSET_MANAGER=$(extract_addr "AssetManagerFacet")
ERC1155=$(extract_addr "ERC1155Facet")
SUPPLY=$(extract_addr "SupplyFacet")
METADATA=$(extract_addr "MetadataFacet")
CLAIM_TOPICS=$(extract_addr "ClaimTopicsFacet")
TRUSTED_ISSUER=$(extract_addr "TrustedIssuerFacet")
IDENTITY_REGISTRY=$(extract_addr "IdentityRegistryFacet")
COMPLIANCE_ROUTER=$(extract_addr "ComplianceRouterFacet")
DIAMOND_INIT=$(extract_addr "DiamondInit")
EIP1967_INIT=$(extract_addr "EIP1967Init")
EXPLORER="https://amoy.polygonscan.com"
TODAY=$(date +%Y-%m-%d)

# Update .env (root — indexer)
sed -i '' "s|^DIAMOND_ADDRESS=.*|DIAMOND_ADDRESS=$DIAMOND|" "$ROOT_DIR/.env"
echo -e "  ✓ .env (DIAMOND_ADDRESS)"

# Update packages/app/.env (frontend)
sed -i '' "s|^VITE_DIAMOND_ADDRESS=.*|VITE_DIAMOND_ADDRESS=$DIAMOND|" "$ROOT_DIR/packages/app/.env"
if [[ -n "$COUNTRY_MODULE" ]]; then
  sed -i '' "s|^VITE_COUNTRY_RESTRICT_MODULE=.*|VITE_COUNTRY_RESTRICT_MODULE=$COUNTRY_MODULE|" "$ROOT_DIR/packages/app/.env"
  sed -i '' "s|^VITE_MAX_BALANCE_MODULE=.*|VITE_MAX_BALANCE_MODULE=$MAX_BALANCE_MODULE|" "$ROOT_DIR/packages/app/.env"
  sed -i '' "s|^VITE_MAX_HOLDERS_MODULE=.*|VITE_MAX_HOLDERS_MODULE=$MAX_HOLDERS_MODULE|" "$ROOT_DIR/packages/app/.env"
fi
echo -e "  ✓ packages/app/.env (VITE_DIAMOND_ADDRESS + modules)"

# Update deployments/amoy.json
cat > "$CONTRACTS_DIR/deployments/amoy.json" << EOJSON
{
  "network": "Polygon Amoy Testnet",
  "chainId": 80002,
  "deployer": "$OWNER_ADDRESS",
  "deployedAt": "$TODAY",
  "blockExplorer": "$EXPLORER",
  "contracts": {
    "Diamond": {
      "address": "$DIAMOND",
      "verified": false,
      "url": "$EXPLORER/address/$DIAMOND"
    },
    "DiamondCutFacet": {
      "address": "$DIAMOND_CUT",
      "verified": false,
      "url": "$EXPLORER/address/$DIAMOND_CUT"
    },
    "DiamondLoupeFacet": {
      "address": "$DIAMOND_LOUPE",
      "verified": false,
      "url": "$EXPLORER/address/$DIAMOND_LOUPE"
    },
    "OwnershipFacet": {
      "address": "$OWNERSHIP",
      "verified": false,
      "url": "$EXPLORER/address/$OWNERSHIP"
    },
    "AccessControlFacet": {
      "address": "$ACCESS_CONTROL",
      "verified": false,
      "url": "$EXPLORER/address/$ACCESS_CONTROL"
    },
    "PauseFacet": {
      "address": "$PAUSE",
      "verified": false,
      "url": "$EXPLORER/address/$PAUSE"
    },
    "EmergencyFacet": {
      "address": "$EMERGENCY",
      "verified": false,
      "url": "$EXPLORER/address/$EMERGENCY"
    },
    "FreezeFacet": {
      "address": "$FREEZE",
      "verified": false,
      "url": "$EXPLORER/address/$FREEZE"
    },
    "RecoveryFacet": {
      "address": "$RECOVERY",
      "verified": false,
      "url": "$EXPLORER/address/$RECOVERY"
    },
    "AssetManagerFacet": {
      "address": "$ASSET_MANAGER",
      "verified": false,
      "url": "$EXPLORER/address/$ASSET_MANAGER"
    },
    "ClaimTopicsFacet": {
      "address": "$CLAIM_TOPICS",
      "verified": false,
      "url": "$EXPLORER/address/$CLAIM_TOPICS"
    },
    "TrustedIssuerFacet": {
      "address": "$TRUSTED_ISSUER",
      "verified": false,
      "url": "$EXPLORER/address/$TRUSTED_ISSUER"
    },
    "IdentityRegistryFacet": {
      "address": "$IDENTITY_REGISTRY",
      "verified": false,
      "url": "$EXPLORER/address/$IDENTITY_REGISTRY"
    },
    "ComplianceRouterFacet": {
      "address": "$COMPLIANCE_ROUTER",
      "verified": false,
      "url": "$EXPLORER/address/$COMPLIANCE_ROUTER"
    },
    "ERC1155Facet": {
      "address": "$ERC1155",
      "verified": false,
      "url": "$EXPLORER/address/$ERC1155"
    },
    "SupplyFacet": {
      "address": "$SUPPLY",
      "verified": false,
      "url": "$EXPLORER/address/$SUPPLY",
      "note": "v2 — emits ERC-1155 TransferSingle on mint/burn/forcedTransfer"
    },
    "MetadataFacet": {
      "address": "$METADATA",
      "verified": false,
      "url": "$EXPLORER/address/$METADATA"
    },
    "SnapshotFacet": {
      "address": "$SNAPSHOT",
      "verified": false,
      "url": "$EXPLORER/address/$SNAPSHOT"
    },
    "DividendFacet": {
      "address": "$DIVIDEND",
      "verified": false,
      "url": "$EXPLORER/address/$DIVIDEND"
    },
    "AssetGroupFacet": {
      "address": "$ASSET_GROUP",
      "verified": false,
      "url": "$EXPLORER/address/$ASSET_GROUP"
    },
    "DiamondInit": {
      "address": "$DIAMOND_INIT",
      "verified": false,
      "url": "$EXPLORER/address/$DIAMOND_INIT"
    },
    "EIP1967Init": {
      "address": "$EIP1967_INIT",
      "verified": false,
      "url": "$EXPLORER/address/$EIP1967_INIT"
    },
    "CountryRestrictModule": {
      "address": "${COUNTRY_MODULE:-}",
      "verified": false,
      "url": "$EXPLORER/address/${COUNTRY_MODULE:-}"
    },
    "MaxBalanceModule": {
      "address": "${MAX_BALANCE_MODULE:-}",
      "verified": false,
      "url": "$EXPLORER/address/${MAX_BALANCE_MODULE:-}"
    },
    "MaxHoldersModule": {
      "address": "${MAX_HOLDERS_MODULE:-}",
      "verified": false,
      "url": "$EXPLORER/address/${MAX_HOLDERS_MODULE:-}"
    }
  }
}
EOJSON
echo -e "  ✓ packages/contracts/deployments/amoy.json"

# ── Done ──────────────────────────────────────────────────
echo ""
echo -e "${CYAN}═══════════════════════════════════════════════════${NC}"
echo -e "${GREEN}  Deploy complete!${NC}"
echo -e "${CYAN}═══════════════════════════════════════════════════${NC}"
echo ""
echo -e "  Diamond: ${GREEN}$DIAMOND${NC}"
echo -e "  Explorer: $EXPLORER/address/$DIAMOND"
echo ""
echo -e "  ${CYAN}Next steps:${NC}"
echo -e "  1. Verify contracts on Polygonscan"
echo -e "  2. Grant roles (ISSUER_ROLE, COMPLIANCE_ADMIN, etc.)"
echo -e "  3. Register assets"
echo -e "  4. Clear indexer data: rm -rf packages/indexer/data-amoy"
echo -e "  5. Restart dev: ./scripts/dev.sh"
