#!/bin/bash

NAMESPACE="${NAMESPACE:-kafka}"

kubectl --namespace="${NAMESPACE}" run kafka-topics -it \
    --image=banzaicloud/kafka:2.13-2.4.0 --rm=true --restart=Never -- \
    /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka-headless:29092 --list
