#!/bin/bash

# Namespace to install Zookeeper
NAMESPACE="${NAMESPACE:-zookeeper}"

CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo "OPTIONS"
echo "Install Zookeeper into ${NAMESPACE} namespace"
echo ""
echo "!!! IMPORTANT !!!"
echo "If you do not agree with specified options, press ctrl-c now"
sleep 30
echo "Apply options now..."

kubectl apply --namespace="${NAMESPACE}" -f "${CUR_DIR}/zookeeper.yaml"