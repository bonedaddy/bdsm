#! /bin/bash
BIN="$1"
ABI="$2"
PKG="$3"
OUT="$4"
OPTS="$5"
abigen \
    --bin "$BIN" \
    --abi "$ABI" \
    --pkg "$PKG" \
    --out "$OUT" \
    "$OPTS"