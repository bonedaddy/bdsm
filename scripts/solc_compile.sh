#! /bin/bash

CONTRACT="$1"
OUTPUT="$2"
COMPILER="$3"

if [[ "$COMPILER" == "" ]]; then
    COMPILER=$(which solc)
fi

if [[ "$COMPILER" == "" ]]; then
    echo "[ERROR] no solc compiler found"
    exit 1
fi

"$COMPILER" \
    --bin \
    --abi \
    --optimize \
    --optimize-runs 200 \
    --hashes \
    --devdoc \
    --userdoc \
    --pretty-json \
    --output-dir="$OUTPUT" \
    --overwrite \
    "$CONTRACT"
	# solc --bin --abi --optimize --optimize-runs 200 --hashes --devdoc --userdoc --pretty-json --output-dir="$OUJTPUT" "$CONTRACT"