#!/usr/bin/env bash

CSS_FILE="css.go"
CSS_TEST_FILE="css.test"

./resources/scripts/generate.sh "$CSS_TEST_FILE" >/dev/null

if [[ ! -f "$CSS_TEST_FILE" ]]; then
    echo >&2 "could not generate $CSS_TEST_FILE"
    exit 1
fi

trap "rm -f $CSS_TEST_FILE" EXIT

if diff -u <(cat "$CSS_FILE") <(cat "$CSS_TEST_FILE") >/dev/null 2>&1; then
    echo "passed"
else
    echo >&2 "css content is outdated, please run \`make generate\`"
    exit 1
fi
