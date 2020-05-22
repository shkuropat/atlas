#!/bin/bash

NAMESPACE="${NAMESPACE:-kafka}"

kubectl --namespace="${NAMESPACE}" run "atlas-run-$(date +%F)-${RANDOM}" -it \
    --image=binarly/atlas:dev --rm=true --restart=Never --image-pull-policy=Always
