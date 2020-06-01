#!/bin/bash

# Install CustomResourceDefinitions first
kubectl apply --validate=false -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.11/deploy/manifests/00-crds.yaml

# Add the jetstack helm repo
helm repo add jetstack https://charts.jetstack.io
helm repo update

kubectl create ns cert-manager || true

# Install cert-manager into the cluster
# Using helm3
# You have to create the namespace before executing following command
helm install cert-manager --namespace cert-manager --version v0.11.0 jetstack/cert-manager
