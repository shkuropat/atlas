#!/bin/bash


helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com/
kubectl create ns kafka
helm install kafka-operator --namespace=kafka banzaicloud-stable/kafka-operator
kubectl -n kafka get all
