#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

kubectl apply --namespace="${NAMESPACE}" -f minioservice.yaml

