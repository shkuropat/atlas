#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

kubectl create namespace "${NAMESPACE}" || true

kubectl apply --namespace="${NAMESPACE}" -f <(\
    curl https://raw.githubusercontent.com/minio/minio-operator/master/minio-operator.yaml | \
    NAMESPACE="${NAMESPACE}" sed -e "s/namespace: default/namespace: ${NAMESPACE}/g" \
)


