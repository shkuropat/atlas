#!/bin/bash


NAMESPACE="${NAMESPACE:-kafka}"

kubectl create namespace "${NAMESPACE}" || true

helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com/
helm install kafka-operator --namespace="${NAMESPACE}" banzaicloud-stable/kafka-operator

kubectl --namespace "${NAMESPACE}" get all
