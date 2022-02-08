#!/usr/bin/env bash

SCRIPT_PATH="$(realpath "$0")"
PROJECT_DIR="$(dirname "${BIN_DIR}")"
CMD_DIR="${PROJECT_DIR}/cmd"
BIN_DIR="${PROJECT_DIR}/bin"

mkdir -p "${BIN_DIR}"

# add your commands here
commands=( "hello" )
for cmd in ${commands[*]}; do
    (cd "${CMD_DIR}/${cmd}" || exit; go build)
    mv "${CMD_DIR}/${cmd}/${cmd}" "${BIN_DIR}"
    echo "${BIN_DIR}/${cmd}"
done
