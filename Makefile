PATH := ${CURDIR}/bin:$(PATH)

protoc_version = 24.0
protoc_arch = x86_64

ifeq ($(shell uname -s),Darwin)
	protoc_os = osx
else
	protoc_os = linux

	ifeq ($(shell uname -m),aarch64)
		protoc_arch = aarch_64
	endif
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

bin/protoc-$(protoc_version).zip:
	mkdir -p $(dir $@)
	curl -L -o $@ https://github.com/protocolbuffers/protobuf/releases/download/v$(protoc_version)/protoc-$(protoc_version)-$(protoc_os)-$(protoc_arch).zip

bin/protoc-$(protoc_version): bin/protoc-$(protoc_version).zip
	mkdir -p $@
	unzip -d $@ -o $<

bin/protoc: bin/protoc-$(protoc_version)
	rm -rf $@
	ln -s ./protoc-$(protoc_version)/bin/protoc $@

bin/protoc-gen-iface:
	go build -o $@ ./protoc-gen-iface/main.go
