#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
BINARY="$ROOT_DIR/bin/pokedexcli"

cd "$ROOT_DIR"
mkdir -p "$ROOT_DIR/bin"
go build -o "$BINARY" ./cmd/pokedexcli
printf 'Built %s\n' "$BINARY"
