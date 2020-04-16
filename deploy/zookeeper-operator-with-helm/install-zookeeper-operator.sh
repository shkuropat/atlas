#!/bin/bash

helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com

kubectl create ns zookeeper

helm install zookeeper-operator --namespace=zookeeper banzaicloud-stable/zookeeper-operator