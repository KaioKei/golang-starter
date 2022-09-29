#!/usr/bin/env bash

SCRIPT_PATH="$(realpath "$0")"
PROJECT_DIR="$(dirname "${SCRIPT_PATH}")"
CMD_DIR="${PROJECT_DIR}/cmd"
BIN_DIR="${PROJECT_DIR}/bin"
commands=()

# list dirs in commands
# i.e dirs with a main.go file inside
# ARG1: path of the root commands dir
function set_commands () {
    for pathname in "$1"/*; do
        if [ -d "${pathname}" ]; then
            set_commands "${pathname}"
        elif [ -e "${pathname}" ]; then
            case "${pathname}" in *main.go)
                commands+=("$(dirname "${pathname}")")
            esac
        fi
    done
}

function build_commands() {
  for cmd in ${commands[*]}; do
      (cd "${cmd}" || exit; go build)
      mv "${cmd}/$(basename "${cmd}")" "${BIN_DIR}"
      echo "${BIN_DIR}/$(basename "${cmd}")"
  done
}

mkdir -p "${BIN_DIR}"
# get all commands in cmd dir
set_commands "${CMD_DIR}"
# build each
build_commands

exit 0
