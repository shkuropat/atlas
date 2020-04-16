#!/bin/bash

# kubectl --namespace=kafka apply -f https://raw.githubusercontent.com/banzaicloud/kafka-operator/master/config/samples/simplekafkacluster.yaml
kubectl --namespace=kafka apply -f simplekafkacluster.yaml

#kubectl --namespace=kafka apply -f https://raw.githubusercontent.com/banzaicloud/kafka-operator/master/config/samples/kafkacluster-prometheus.yaml
kubectl --namespace=kafka apply -f kafkacluster-prometheus.yaml