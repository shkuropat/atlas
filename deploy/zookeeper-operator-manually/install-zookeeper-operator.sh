#!/bin/bash

CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/zookeeper-operator-common.sh"

# OPERATOR_NAMESPACE is declared in common/configuration file. Check in there for default value

echo "OPTIONS"
echo "Install operator into ${OPERATOR_NAMESPACE} namespace"
echo ""
echo "!!! IMPORTANT !!!"
echo "If you do not agree with specified options, press ctrl-c now"
sleep 30
echo "Apply options now..."

# CRD
kubectl apply -f https://raw.githubusercontent.com/pravega/zookeeper-operator/master/deploy/crds/zookeeper_v1beta1_zookeepercluster_crd.yaml

# RBAC
kubectl apply --namespace="${OPERATOR_NAMESPACE}" -f <( \
    get_file https://raw.githubusercontent.com/pravega/zookeeper-operator/master/deploy/default_ns/rbac.yaml | \
        OPERATOR_NAMESPACE=${OPERATOR_NAMESPACE} sed "s/namespace: default/namespace: ${OPERATOR_NAMESPACE}/" \
)

# Operator
kubectl apply --namespace="${OPERATOR_NAMESPACE}" -f https://raw.githubusercontent.com/pravega/zookeeper-operator/master/deploy/default_ns/operator.yaml
