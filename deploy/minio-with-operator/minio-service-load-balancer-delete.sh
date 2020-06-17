#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

kubectl delete --namespace="${NAMESPACE}" -f minio-service-nlb.yaml

