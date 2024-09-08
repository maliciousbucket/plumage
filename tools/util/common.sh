#!/usr/bin/env bash

RED='\033[1;31m'
GRN='\033[1;32m'
YEL='\033[1;33m'
BLU='\033[1;34m'
WHT='\033[1;37m'
MGT='\033[1;95m'
CYA='\033[1;96m'
END='\033[0m'
BLOCK='\033[1;37m'


# highlight a new section
section() {
  echo ""
  log "***** Section: ${MGT}$1${END} *****";
  echo ""
}

# show command and execute it
exe() {
  echo "\$ $1"
  eval "$1"
}

warn() { log "Warning: ${YEL}$1${END}"; }

success() { log "Success: ${CYA}$1${END}"; }

# output a "log" line with bold leading >>>
log() { >&2 printf "${BLOCK}#${END} %s\n" "$1"; }