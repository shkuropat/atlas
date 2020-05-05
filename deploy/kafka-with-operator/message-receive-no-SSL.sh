#!/bin/bash

TOPIC="${TOPIC:-}"
NAMESPACE="${NAMESPACE:-kafka}"

if [[ -z "${TOPIC}" ]]; then
    echo "Please specify TOPIC. Abort."
    exit 1
fi

kubectl --namespace="${NAMESPACE}" run kafka-consumer -it \
    --image=banzaicloud/kafka:2.13-2.4.0 --rm=true --restart=Never -- \
    /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka-headless:29092 --topic "${TOPIC}" --from-beginning
