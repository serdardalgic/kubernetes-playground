SHELL = /bin/bash -o pipefail

## Frontend App Name (default: dummyproxy)
FE_APPNAME ?= dummyproxy
## Frontend App Version (default: v1)
FE_VERSION ?= v1

## Backend App Name (default: go-app)
BE_APPNAME ?= go-app
## Backend App Version (default: v3)
BE_VERSION ?= v3

## Programs that should be installed (default: kubectl and minikube)
PREREQUISITES = kubectl minikube
$(foreach prg, $(PREREQUISITES), \
	$(if $(shell type $(prg) 2> /dev/null), , $(error $(prg) must be installed)))

## Default parameter, prints the help message
all: deploy

## Prints info about the project
info:
	@echo $(PWD)
	@printf '%s: %s\n' "FE_APPNAME" $(FE_APPNAME)
	@printf '%s: %s\n' "FE_VERSION" $(FE_VERSION)
	@printf '%s: %s\n' "BE_APPNAME" $(BE_APPNAME)
	@printf '%s: %s\n' "BE_VERSION" $(BE_VERSION)


## Builds the docker images
build: build-fe build-be

## Builds frontend image
build-fe:
	$(eval FE_IMAGE := $(shell eval $$(minikube docker-env) ;\
		docker images -q $(FE_APPNAME):$(FE_VERSION) 2> /dev/null))
	@if [ -z $(FE_IMAGE) ]; then \
		echo ;\
		echo Building Frontend Image: $(FE_APPNAME):$(FE_VERSION) ;\
		echo ;\
		eval $$(minikube docker-env) ;\
		cd frontend ;\
		docker build --no-cache -t $(FE_APPNAME):$(FE_VERSION) . ;\
	else \
		echo "Frontend image $(FE_APPNAME):$(FE_VERSION) is present in the local docker registry." ;\
	fi


## Builds backend image
build-be:
	$(eval BE_IMAGE := $(shell eval $$(minikube docker-env) ;\
		docker images -q $(BE_APPNAME):$(BE_VERSION) 2> /dev/null))
	@if [ -z $(BE_IMAGE) ]; then \
		echo ;\
		echo Building Backend Image: $(BE_APPNAME):$(BE_VERSION) ;\
		echo ;\
		eval $$(minikube docker-env) ;\
		cd frontend ;\
		docker build --no-cache -t $(BE_APPNAME):$(BE_VERSION) . ;\
	else \
		echo "Backend image $(BE_APPNAME):$(BE_VERSION) is present in the local docker registry." ;\
	fi

## Deploys on kubernetes
deploy: build
	@echo
	@echo Deploying ...
	@echo
	kubectl apply -f kubernetes/redis.yaml
	kubectl apply -f kubernetes/backend.yaml
	kubectl apply -f kubernetes/frontend.yaml

## Tests the system
test: ; # TODO: Test the system

## Deletes all the project resources
clean:
	@echo
	@echo Deleting ...
	@echo
	kubectl delete -f kubernetes/frontend.yaml
	kubectl delete -f kubernetes/backend.yaml
	kubectl delete -f kubernetes/redis.yaml

# Taken from https://github.com/ianstormtaylor/makefile-help
## Show this help prompt.
help:
	@ echo
	@ echo '  Usage:'
	@ echo ''
	@ echo '    make <target> [flags...]'
	@ echo ''
	@ echo '  Targets:'
	@ echo ''
	@ awk '/^##/{ comment = substr($$0,3) } comment && /^[a-zA-Z][a-zA-Z0-9_-]+ ?:/{ print "   ", $$1, comment }' $(MAKEFILE_LIST) | column -t -s ':' | sort
	@ echo ''
	@ echo '  Flags:'
	@ echo ''
	@ awk '/^##/{ comment = substr($$0,3) } comment && /^[a-zA-Z][a-zA-Z0-9_-]+ ?(\?)*=/{ print "   ", $$1, $$2, comment }' $(MAKEFILE_LIST) | column -t -s '?=' | sort
	@ echo ''

.PHONY: clean test info help
