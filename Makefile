#This makefile is used by ci-operator

CGO_ENABLED=0
GOOS=linux

install:
	go install ./cmd/controller/ ./cmd/webhook/ ./pkg/provisioners/kafka ./cmd/fanoutsidecar
	go build -o $(GOPATH)/bin/in-memory-channel-controller ./pkg/controller/eventing/inmemory/controller
.PHONY: install

test-install:
	go install ./test/test_images/k8sevents
.PHONY: test-install

test-e2e:
	sh openshift/e2e-tests-openshift.sh
.PHONY: test-e2e
