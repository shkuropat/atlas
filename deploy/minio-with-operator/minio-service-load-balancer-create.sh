#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

kubectl apply --namespace="${NAMESPACE}" -f minio-service-nlb.yaml

