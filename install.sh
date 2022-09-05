#!/usr/bin/env bash

SCRIPT_PATH="$(realpath "$0")"
PROJECT_DIR="$(dirname "${SCRIPT_PATH}")"
CMD_DIR="${PROJECT_DIR}/cmd"
BIN_DIR="${PROJECT_DIR}/bin"

mkdir -p "${BIN_DIR}"

# add your commands here
commands=( "hello" "shellcli/cobra" )
for cmd in ${commands[*]}; do
    (cd "${CMD_DIR}/${cmd}" || exit; go build)
    mv "${CMD_DIR}/${cmd}/$(basename "${cmd}")" "${BIN_DIR}"
    echo "${BIN_DIR}/$(basename "${cmd}")"
done
