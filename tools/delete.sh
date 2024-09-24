#!/usr/bin/env bash


target="${1}"

demo_server() {
  kubectl delete deployment demo-server-deployment -n galah-testbed
}

demo_client() {
    kubectl delete deployment demo-client-deployment -n galah-testbed
}

argo() {
      kubectl delete deployment argocd -n argocd
}


all() {
  demo_server
  demo_client
  argo
}

default() {
  echo "unknown resurce $target"
}


case "${target}" in
  "argo" )
  argo
  ;;
  "server" )
  demo_server
  ;;
  "client" )
  demo_client
  ;;
  *)
  default
  ;;
esac