
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

##@ Tools

KUBECTL ?= kubectl
K3D ?= $(LOCALBIN)/k3d

## Tool Versions

K3D_VERSION ?= v5.7.3

.PHONY: k3d ## Install K3D to the project's /bin/ directory
k3d: $(K3D)
$(K3D): $(LOCALBIN)
	$(call go-install-tool,$(K3D),github.com/k3d-io/k3d/v5,$(K3D_VERSION))

.PHONY: cluster
cluster: ## Create the default K3D cluster
	@echo "Using K3D: $(K3D)"
	$(K3D) cluster create argocd --config kubernetes/k3d-config.yaml

clean: ## Delete the galah-monitoring cluster and remove generated charts
	$(K3D) cluster delete argocd

.PHONY: argo
install-argo: ## Install Argo with the default configuration on the cluster
	@kubectl create namespace argocd
	@kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	@kubectl config set-context --current --namespace=argocd

argo-ingress: ## Add an ingress pointing to the argocd server
	@kubectl apply -f kubernetes/ingress.yaml

argo-node: ## Open a nodepoprt to the argocd server
	$(SHELL ./tools/pf-argp.sh)



define go-install-tool
@[ -f "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f $(1) || true ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv $(1) $(1)-$(3) ;\
} ;\
ln -sf $(1)-$(3) $(1)
endef

.PHONY: help
help:
ifeq ($(OS),Windows_NT)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  %-40s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
else
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
endif