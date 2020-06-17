#!/bin/bash

NAMESPACE="${NAMESPACE:-atlas}"

kubectl --namespace="${NAMESPACE}" run "atlas-run-$(date +%F)-${RANDOM}" -it \
    --image=binarly/atlas:dev --rm=true --restart=Never --image-pull-policy=Always

