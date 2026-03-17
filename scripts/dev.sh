#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

# Load .env
set -a
source "$ROOT_DIR/.env"
set +a

# Colors
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
GRAY='\033[0;90m'
NC='\033[0m'

cleanup() {
  echo ""
  echo -e "${GRAY}Shutting down...${NC}"
  kill $INDEXER_PID $APP_PID 2>/dev/null || true
  wait $INDEXER_PID $APP_PID 2>/dev/null || true
  echo -e "${GRAY}Done.${NC}"
}
trap cleanup EXIT INT TERM

# ── Build indexer ──────────────────────────────────────────
echo -e "${MAGENTA}[indexer]${NC} Building..."
cd "$ROOT_DIR/packages/indexer"
export CGO_CFLAGS="-I/opt/homebrew/include"
export CGO_LDFLAGS="-L/opt/homebrew/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd"
export CGO_ENABLED=1
go build -o bin/indexer ./cmd/indexer/ 2>&1
echo -e "${MAGENTA}[indexer]${NC} Built."

# ── Start indexer ──────────────────────────────────────────
mkdir -p "$DATA_DIR"
echo -e "${MAGENTA}[indexer]${NC} Starting on $HTTP_LISTEN (Diamond: $DIAMOND_ADDRESS)"
./bin/indexer 2>&1 | sed "s/^/$(printf "${MAGENTA}[indexer]${NC} ")/" &
INDEXER_PID=$!

# ── Start frontend ────────────────────────────────────────
cd "$ROOT_DIR/packages/app"
echo -e "${CYAN}[app]${NC}     Starting on http://localhost:3000"
npx vite --port 3000 2>&1 | sed "s/^/$(printf "${CYAN}[app]${NC}     ")/" &
APP_PID=$!

echo ""
echo -e "${GRAY}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "  ${CYAN}Frontend${NC}  → http://localhost:3000"
echo -e "  ${MAGENTA}Indexer${NC}   → http://localhost:8080"
echo -e "  ${MAGENTA}GraphQL${NC}  → http://localhost:8080/graphql"
echo -e "  ${GRAY}Diamond${NC}   → $DIAMOND_ADDRESS"
echo -e "${GRAY}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""

# Wait for both (Ctrl+C triggers cleanup)
wait $INDEXER_PID $APP_PID
