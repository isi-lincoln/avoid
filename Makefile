prefix ?= /usr

#VERSION = $(shell git describe --always --long --dirty --tags)
#LDFLAGS = "-X github.com/isi-lincoln/pkg/common.Version=$(VERSION)"

all: dns

protobuf: protobuf-avoid

#code: gateway management dns path

# TODO: build mock/test infrastructure for integration
#mock: build/dmock

dns: build/avoid-dns build/avoid-dns-cli
#gateway: build/avoid-gw build/avoid-gw-cli
#management: build/avoid-mgmt build/avoid-mgmt-cli
#path: build/avoid-path build/avoid-path-cli

test:
	go test -v ./...

#build/avoid-gw: gw/main.go | build protobuf-avoid
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#
#build/avoid-gw-cli: gw/cli/main.go | build protobuf-avoid
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#
#build/avoid-mgmt: mgmt/main.go | build protobuf-avoid
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#
#build/avoid-mgmt-cli: mgmt/cli/main.go | build protobuf-avoid
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#

build/avoid-coredns-binary:
	$(MAKE) -C coredns

build/avoid-dns: dns/main.go | build protobuf-avoid
	go build -ldflags=$(LDFLAGS) -o $@ $<

build/avoid-dns-cli: dns/cli/main.go | build protobuf-avoid
	go build -o $@ $<
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#
#build/avoid-path: path/main.go | build protobuf-avoid
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#
#build/avoid-path-cli: path/cli/main.go | build protobuf-avoid
#	go build -ldflags=$(LDFLAGS) -o $@ $<
#
build:
	mkdir -p build

clean:
	rm -rf build

protobuf-avoid:
	protoc -I=./protocol --go_out=./protocol --go_opt=paths=source_relative \
		--go-grpc_out=./protocol --go-grpc_opt=paths=source_relative  \
		 ./protocol/*.proto

REGISTRY ?= docker.io
REPO ?= isilincoln
TAG ?= latest
BUILD_ARGS ?= --no-cache

docker: $(REGISTRY)/$(REPO)/avoid-gateway-service \
	    $(REGISTRY)/$(REPO)/avoid-management-service \
		$(REGISTRY)/$(REPO)/avoid-dns-service \
		$(REGISTRY)/$(REPO)/avoid-path-service

$(REGISTRY)/$(REPO)/avoid-gateway-service:
	@docker build ${BUILD_ARGS} $(DOCKER_QUIET) -f Dockerfile -t $(@):$(TAG) .
	$(if ${PUSH},$(call docker-push))

$(REGISTRY)/$(REPO)/avoid-management-service:
	@docker build ${BUILD_ARGS} $(DOCKER_QUIET) -f Dockerfile -t $(@):$(TAG) .
	$(if ${PUSH},$(call docker-push))

$(REGISTRY)/$(REPO)/avoid-dns-service:
	@docker build ${BUILD_ARGS} $(DOCKER_QUIET) -f Dockerfile -t $(@):$(TAG) .
	$(if ${PUSH},$(call docker-push))

$(REGISTRY)/$(REPO)/avoid-path-service:
	@docker build ${BUILD_ARGS} $(DOCKER_QUIET) -f Dockerfile -t $(@):$(TAG) .
	$(if ${PUSH},$(call docker-push))

define docker-push
	@docker push $(@):$(TAG)
endef
