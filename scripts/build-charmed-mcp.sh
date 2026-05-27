#!/usr/bin/env bash
# Build the charmed-mcp binary from the sibling charmed-mcp Go project.
#
# Usage:
#   bash scripts/build-charmed-mcp.sh          # Build to ~/.local/bin/charmed-mcp
#   bash scripts/build-charmed-mcp.sh --local  # Build to <plugin-root>/bin/charmed-mcp
#
# After building, ensure the binary is on PATH or update .mcp.json to point to it.

set -euo pipefail

PLUGIN_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

# Find the Go source: prefer sibling directory, then GOROOT-style paths
GO_SRC_CANDIDATES=(
  "${PLUGIN_ROOT}/../charmed-mcp"
  "${PLUGIN_ROOT}/../charmed/charmed-mcp"
  "${HOME}/Charmed/charmed-mcp"
)

GO_SRC=""
for c in "${GO_SRC_CANDIDATES[@]}"; do
  if [[ -f "${c}/main.go" ]]; then
    GO_SRC="$c"
    break
  fi
done

if [[ -z "$GO_SRC" ]]; then
  echo "ERROR: Cannot find charmed-mcp Go source. Expected at one of:"
  printf '  %s\n' "${GO_SRC_CANDIDATES[@]}"
  echo ""
  echo "Clone it: git clone https://github.com/B-A-M-N/charmed-mcp.git ../charmed-mcp"
  exit 1
fi

# Determine output path
if [[ "${1:-}" == "--local" ]]; then
  OUT="${PLUGIN_ROOT}/bin/charmed-mcp"
  mkdir -p "$(dirname "$OUT")"
else
  OUT="${HOME}/.local/bin/charmed-mcp"
  mkdir -p "$(dirname "$OUT")"
fi

echo "Building charmed-mcp from ${GO_SRC} -> ${OUT}"
cd "$GO_SRC"
go build -ldflags="-s -w" -o "$OUT" .

echo ""
echo "✓ Built: ${OUT}"
echo "  Ensure this is on your PATH, or update .mcp.json 'command' to:"
echo "    \"command\": \"${OUT}\""
