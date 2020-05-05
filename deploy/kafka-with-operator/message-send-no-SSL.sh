#!/bin/bash

TOPIC="${TOPIC:-}"
NAMESPACE="${NAMESPACE:-kafka}"

if [[ -z "${TOPIC}" ]]; then
    echo "Please specify TOPIC. Abort."
    exit 1
fi

kubectl --namespace="${NAMESPACE}" run kafka-producer -it \
    --image=banzaicloud/kafka:2.13-2.4.0 --rm=true --restart=Never -- \
    /opt/kafka/bin/kafka-console-producer.sh --broker-list kafka-headless:29092 --topic "${TOPIC}"
