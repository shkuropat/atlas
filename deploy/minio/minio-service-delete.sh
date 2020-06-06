#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

kubectl delete --namespace="${NAMESPACE}" -f minioservice.yaml

