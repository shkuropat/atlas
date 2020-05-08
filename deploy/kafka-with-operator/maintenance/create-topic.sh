#!/bin/bash

CLUSTER="${CLUSTER:-kafka}"
TOPIC="${TOPIC:-}"
PARTITIONS="${PARTITIONS:-1}"
REPLICATION_FACTOR="${REPLICATION_FACTOR:-1}"
NAMESPACE="${NAMESPACE:-kafka}"

if [[ -z "${TOPIC}" ]]; then
    echo "Please specify TOPIC. Abort."
    exit 1
fi

cat << EOF | \
    CLUSTER="${CLUSTER}" \
    TOPIC="${TOPIC}" \
    PARTITIONS="${PARTITIONS}" \
    REPLICATION_FACTOR="${REPLICATION_FACTOR}" \
    envsubst | kubectl apply --namespace "${NAMESPACE}" -f -
apiVersion: kafka.banzaicloud.io/v1alpha1
kind: KafkaTopic
metadata:
  name: ${TOPIC}
spec:
  clusterRef:
    name: ${CLUSTER}
  name: ${TOPIC}
  partitions: ${PARTITIONS}
  replicationFactor: ${REPLICATION_FACTOR}
EOF
