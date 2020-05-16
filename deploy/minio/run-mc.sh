#!/bin/bash

NAMESPACE="${NAMESPACE:-minio-operator-ns}"

kubectl --namespace="${NAMESPACE}" run mc -it \
    --image=minio/mc:latest --rm=true --restart=Never -- \
    mc ls
