#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

kubectl -n "${NAMESPACE}" port-forward service/minio-service 10000:9000
