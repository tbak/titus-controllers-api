export

# Customizable stuff
GO_VERSION := 1.13.7

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GOBIN_TOOL := $(shell which gobin || echo $(GOBIN)/gobin)

.PHONY: all
all: generate test

# Run tests
test: generate vet
	go test ./... -coverprofile cover.out

# Run go fmt against code
.PHONY: fmt
fmt: $(GOBIN_TOOL)
	$(GOBIN_TOOL) -m -run golang.org/x/tools/cmd/goimports -w $(shell go list -f '{{.Dir}}' ./...)
	jsonnetfmt -i nfpm.jsonnet

# Run go vet against code
.PHONY: lint
lint: $(GOBIN_TOOL)
	$(GOBIN_TOOL) -m -run github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose

.PHONY: vet
vet:
	go vet ./...

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

.PHONY: clean
clean:
	rm -rf build

$(GOBIN_TOOL):
	go get github.com/myitcv/gobin
	go install github.com/myitcv/gobin
