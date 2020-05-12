#!/bin/bash

NAMESPACE=keycloak

REPO="https://raw.githubusercontent.com/keycloak/keycloak-operator/master"
DEPLOY="${REPO}/deploy"
CRDS="${DEPLOY}/crds"

kubectl apply -f ${DEPLOY}/examples/keycloak/keycloak.yaml -n ${NAMESPACE}

echo " HOW TO FETCH admin username/password
kubectl -n keycloak describe pod/keycloak-0
look for     \
Environment:
      KEYCLOAK_USER:                     <set to the key 'ADMIN_USERNAME' in secret 'credential-example-keycloak'>  Optional: false
      KEYCLOAK_PASSWORD:                 <set to the key 'ADMIN_PASSWORD' in secret 'credential-example-keycloak'>  Optional: false

kubectl -n keycloak get -o yaml secret/credential-example-keycloak

look for

data:
  ADMIN_PASSWORD: TzJnWGx3d3FuSVdpNFE9PQ==
  ADMIN_USERNAME: YWRtaW4=

echo YWRtaW4= | base64 -d
admin
echo TzJnWGx3d3FuSVdpNFE9PQ== | base64 -d
O2gXlwwqnIWi4Q==
"
