#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"

if ! command -v dlv >/dev/null 2>&1; then
    printf 'Delve is required for debug. Install it with:\n' >&2
    printf '  go install github.com/go-delve/delve/cmd/dlv@latest\n' >&2
    exit 1
fi

cd "$ROOT_DIR"
dlv debug . -- "$@"
