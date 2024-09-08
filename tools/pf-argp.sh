#!/usr/bin/env bash

CURR_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
[ -d "$CURR_DIR" ] || { echo "FATAL: Couldn't locate current directory"; exit 1; }

source "$CURR_DIR/util/common.sh"

exe "kubectl patch svc argocd-server -n argocd -p '{\"spec\": {\"type\": \"LoadBalancer\"}}'"

sleep 5

exe "kubectl port-forward svc/argocd-server -n argocd 8081:443"