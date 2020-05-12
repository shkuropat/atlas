#!/bin/bash

NAMESPACE=keycloak

REPO="https://raw.githubusercontent.com/keycloak/keycloak-operator/master"
DEPLOY="${REPO}/deploy"
CRDS="${DEPLOY}/crds"

# kubectl apply -f deploy/crds/
# expands into

kubectl apply -f ${CRDS}/keycloak.org_keycloakbackups_crd.yaml
kubectl apply -f ${CRDS}/keycloak.org_keycloakclients_crd.yaml
kubectl apply -f ${CRDS}/keycloak.org_keycloakrealms_crd.yaml
kubectl apply -f ${CRDS}/keycloak.org_keycloaks_crd.yaml
kubectl apply -f ${CRDS}/keycloak.org_keycloakusers_crd.yaml

kubectl create namespace ${NAMESPACE}

kubectl apply -f ${DEPLOY}/role.yaml -n ${NAMESPACE}
kubectl apply -f ${DEPLOY}/role_binding.yaml -n ${NAMESPACE}
kubectl apply -f ${DEPLOY}/service_account.yaml -n ${NAMESPACE}

kubectl apply -f ${DEPLOY}/operator.yaml -n ${NAMESPACE}
