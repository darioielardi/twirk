#!/usr/bin/env bash

set -euo pipefail

# Pull into variable since both protoc-gen-go and protoc-gen-twirk need the map.
Y_IMPORT_MAPPING="y/y.proto=github.com/darioielardi/twirk/internal/twirktest/importmapping/y"

PROTOC_GEN_GO_PARAMS="M${Y_IMPORT_MAPPING}" \
PROTOC_GEN_twirk_PARAMS="go_import_mapping@${Y_IMPORT_MAPPING}" \
../../protoc_gen.sh x/x.proto
