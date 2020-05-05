#!/bin/bash

NAMESPACE="${NAMESPACE:-kafka}"

# kubectl --namespace=kafka apply -f https://raw.githubusercontent.com/banzaicloud/kafka-operator/master/config/samples/simplekafkacluster.yaml
kubectl --namespace="${NAMESPACE}" apply -f simplekafkacluster.yaml

#kubectl --namespace=kafka apply -f https://raw.githubusercontent.com/banzaicloud/kafka-operator/master/config/samples/kafkacluster-prometheus.yaml
kubectl --namespace="${NAMESPACE}" apply -f kafkacluster-prometheus.yaml
