#!/bin/bash

# Build configuration options

CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
SRC_ROOT="$(realpath "${CUR_DIR}/..")"

MANIFESTS_ROOT="${SRC_ROOT}/deploy"
PKG_ROOT="${SRC_ROOT}/pkg"
CONFIG_DIR="${SRC_ROOT}/config"

REPO="github.com/binarly-io/atlas"

# 1.2.3
VERSION=$(cd "${SRC_ROOT}"; cat release)
# 885c3f7
GIT_SHA=$(cd "${CUR_DIR}"; git rev-parse --short HEAD)
# 2020-01-02 12:34:56
NOW=$(date "+%FT%T")

# Service binary name can be specified externally
SERVICE_BIN="${SERVICE_BIN:-${SRC_ROOT}/dev/bin/service}"
SERVICE_BUILDER_SCRIPT="go_build_service.sh"

# Client binary name can be specified externally
CLIENT_BIN="${CLIENT_BIN:-${SRC_ROOT}/dev/bin/client}"
CLIENT_BUILDER_SCRIPT="go_build_client.sh"

# Consumer binary name can be specified externally
CONSUMER_BIN="${CONSUMER_BIN:-${SRC_ROOT}/dev/bin/consumer}"
CONSUMER_BUILDER_SCRIPT="go_build_consumer.sh"

# Where modules kept
MODULES_DIR=vendor
