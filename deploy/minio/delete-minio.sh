#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

source "minio-instance-params.sh"

kubectl delete --namespace="${NAMESPACE}" -f <( \
    cat minioinstance-template.yaml |          \
    NAMESPACE="${NAMESPACE}"                   \
    ACCESS_KEY="${ACCESS_KEY}"                 \
    SECRET_KEY="${SECRET_KEY}"                 \
    ACCESS_KEY_BASE64="${ACCESS_KEY_BASE64}"   \
    SECRET_KEY_BASE64="${SECRET_KEY_BASE64}"   \
    SERVERS_NUM="${SERVERS_NUM}"               \
    VOLUMES_PER_SERVER="${VOLUMES_PER_SERVER}" \
    VOLUME_SIZE="${VOLUME_SIZE}"               \
    STORAGE_CLASS_NAME="${STORAGE_CLASS_NAME}" \
    envsubst
)
