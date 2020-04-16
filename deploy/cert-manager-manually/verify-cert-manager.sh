#!/bin/bash

# Verify installation

#
#
#
function press_enter() {
    read -p "Press enter to continue"
}

#
#
#
function press_any_key() {
    # -n defines the required character count to stop reading
    # -s hides the user's input
    # -r causes the string to be interpreted "raw" (without considering backslash escapes)
    read -n 1 -s -r -p "Press any key to continue"
}

#
#
#
function cat_test_manifests() {
    cat <<EOF
apiVersion: v1
kind: Namespace
metadata:
  name: cert-manager-test
---
apiVersion: cert-manager.io/v1alpha2
kind: Issuer
metadata:
  name: test-selfsigned
  namespace: cert-manager-test
spec:
  selfSigned: {}
---
apiVersion: cert-manager.io/v1alpha2
kind: Certificate
metadata:
  name: selfsigned-cert
  namespace: cert-manager-test
spec:
  dnsNames:
    - example.com
  secretName: selfsigned-cert-tls
  issuerRef:
    name: test-selfsigned
EOF
}

echo "Check cert-manager pods are in place"
echo "This is an interactive process with 'wait' command, which can be terminated by Ctrl-C"
echo "Press any key to continue"
press_any_key
watch -n1 "kubectl --namespace cert-manager get pods"


cat_test_manifests | kubectl apply -f -

echo "Check issuer,certificate are in place"
echo "This is an interactive process with 'wait' command, which can be terminated by Ctrl-C"
echo "Press any key to continue"
press_any_key
watch -n1 "kubectl --namespace cert-manager-test get issuer,certificate"

kubectl --namespace cert-manager-test describe issuer,certificate | less -S

cat_test_manifests | kubectl delete -f -

clear
echo "Now read on how to setup certificates"
echo "https://cert-manager.io/docs/configuration/"
