#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

watch "kubectl -n ${NAMESPACE} get all"

