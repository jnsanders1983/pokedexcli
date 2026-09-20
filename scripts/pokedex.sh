#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"
BINARY="$ROOT_DIR/bin/pokedexcli"

usage() {
    cat <<'EOF'
Usage: ./scripts/pokedex.sh <command>

Commands:
  build       Build the CLI into bin/pokedexcli
  run         Run the CLI with go run
  test        Run all Go tests
  debug       Start the CLI in Delve's interactive debugger
  help        Show this help message
EOF
}

build() {
    mkdir -p "$ROOT_DIR/bin"
    go build -o "$BINARY" ./cmd/pokedexcli
    printf 'Built %s\n' "$BINARY"
}

run() {
    go run ./cmd/pokedexcli "$@"
}

test() {
    go test ./...
}

debug() {
    if ! command -v dlv >/dev/null 2>&1; then
        printf 'Delve is required for debug. Install it with:\n' >&2
        printf '  go install github.com/go-delve/delve/cmd/dlv@latest\n' >&2
        exit 1
    fi

    dlv debug ./cmd/pokedexcli -- "$@"
}

cd "$ROOT_DIR"

command="${1:-help}"
if [[ $# -gt 0 ]]; then
    shift
fi

case "$command" in
    build)
        build
        ;;
    run)
        run "$@"
        ;;
    test)
        test
        ;;
    debug)
        debug "$@"
        ;;
    help|-h|--help)
        usage
        ;;
    *)
        printf 'Unknown command: %s\n\n' "$command" >&2
        usage >&2
        exit 1
        ;;
esac
