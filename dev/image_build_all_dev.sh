#!/bin/bash

# Build clickhouse-operator
# Do not forget to update version

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}"

echo "Build Service..."
DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}" "${CUR_DIR}"/image_build_service_dev.sh

echo "Build Client..."
DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}" "${CUR_DIR}"/image_build_client_dev.sh

echo "Build Consumer..."
DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}" "${CUR_DIR}"/image_build_consumer_dev.sh

echo "Build Full Distro..."
DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}" "${CUR_DIR}"/image_build_atlas_dev.sh
