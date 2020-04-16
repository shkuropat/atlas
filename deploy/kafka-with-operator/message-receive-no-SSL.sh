#!/bin/bash

kubectl --namespace=kafka run kafka-consumer -it \
    --image=banzaicloud/kafka:2.13-2.4.0 --rm=true --restart=Never -- \
    /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka-headless:29092 --topic my-topic --from-beginning
