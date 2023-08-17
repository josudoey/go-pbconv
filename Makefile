PATH := ${CURDIR}/bin:$(PATH)

PROTOC_VERSION = 24.0
PROTOC_ARCH = x86_64

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	PROTOC_OS = linux
endif
ifeq ($(UNAME_S),Darwin)
	PROTOC_OS = osx
endif


.PHONY: build
build: bin/protoc bin/protoc-gen-go bin/protoc-gen-iface
	./bin/protoc \
		-I=. \
		--go_out=paths=source_relative:. \
		--iface_out=paths=source_relative:. \
		./internal/fixture/*.proto

# see https://github.com/protocolbuffers/protobuf-go
bin/protoc-gen-go:
	GOBIN=$(abspath bin) go install google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-$(PROTOC_VERSION).zip:
	mkdir -p $(dir $@)
	curl -L -o $@ https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip

bin/protoc-$(PROTOC_VERSION): bin/protoc-$(PROTOC_VERSION).zip
	mkdir -p $@
	unzip -d $@ -o $<

bin/protoc: bin/protoc-$(PROTOC_VERSION)
	rm -rf $@
	ln -s ./protoc-$(PROTOC_VERSION)/bin/protoc $@

bin/protoc-gen-iface:
	go build -o $@ ./protoc-gen-iface/main.go
