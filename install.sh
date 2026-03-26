#!/usr/bin/env sh

set -eu

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
APP_NAME=${APP_NAME:-calc}
BUILD_DIR=${BUILD_DIR:-"$SCRIPT_DIR/bin"}
DEFAULT_PREFIX=/usr/local/bin
TARGET_DIR=${PREFIX:-$DEFAULT_PREFIX}
BINARY_PATH=$BUILD_DIR/$APP_NAME

can_write_to_dir() {
    target=$1

    if [ -d "$target" ]; then
        [ -w "$target" ]
        return
    fi

    parent=$(dirname "$target")
    [ -d "$parent" ] && [ -w "$parent" ]
}

if [ "${1:-}" = "--help" ]; then
    printf '%s\n' \
        "Usage: PREFIX=/custom/bin ./install.sh" \
        "Builds the calculator module and installs it as ${APP_NAME}."
    exit 0
fi

if [ ! -f "$SCRIPT_DIR/go.mod" ]; then
    printf 'go.mod topilmadi: %s/go.mod\n' "$SCRIPT_DIR" >&2
    exit 1
fi

if ! can_write_to_dir "$TARGET_DIR"; then
    if [ "$TARGET_DIR" = "$DEFAULT_PREFIX" ]; then
        TARGET_DIR="$HOME/.local/bin"
        printf 'write access yoq, fallback ishlatildi: %s\n' "$TARGET_DIR"
    else
        printf 'write access yoq: %s\n' "$TARGET_DIR" >&2
        exit 1
    fi
fi

mkdir -p "$BUILD_DIR" "$TARGET_DIR"

(
    cd "$SCRIPT_DIR"
    go build -o "$BINARY_PATH" .
)
install -m 0755 "$BINARY_PATH" "$TARGET_DIR/$APP_NAME"

printf 'installed: %s\n' "$TARGET_DIR/$APP_NAME"

case ":${PATH:-}:" in
    *":$TARGET_DIR:"*)
        ;;
    *)
        printf 'PATH ga qoshing: export PATH="%s:$PATH"\n' "$TARGET_DIR"
        ;;
esac
