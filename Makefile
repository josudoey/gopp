PATH := ${CURDIR}/bin:$(PATH)

protoc_version = 3.19.4
protoc_arch = x86_64

ifeq ($(shell uname -s),Darwin)
	protoc_os = osx
else
	protoc_os = linux

	ifeq ($(shell uname -m),aarch64)
		protoc_arch = aarch_64
	endif
endif

.PHONY: clean
clean:
	rm bin/protoc-gen-grpcadapter

.PHONY: gorpc
gorpc: bin/protoc bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-grpcadapter
	./bin/protoc \
		-I=. \
		--go_out=paths=source_relative:. \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. \
		./gorpc/*.proto

.PHONY: wire
wire: bin/wire
	bin/wire  ./...

bin/wire:
	GOBIN=$(abspath bin) go install github.com/google/wire/cmd/wire@v0.5.0

# see https://github.com/protocolbuffers/protobuf-go
bin/protoc-gen-go:
	GOBIN=$(abspath bin) go install google.golang.org/protobuf/cmd/protoc-gen-go

# see https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc
bin/protoc-gen-go-grpc:
	GOBIN=$(abspath bin) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

bin/protoc-$(protoc_version).zip:
	mkdir -p $(dir $@)
	curl -L -o $@ https://github.com/protocolbuffers/protobuf/releases/download/v$(protoc_version)/protoc-$(protoc_version)-$(protoc_os)-$(protoc_arch).zip

bin/protoc-$(protoc_version): bin/protoc-$(protoc_version).zip
	mkdir -p $@
	unzip -d $@ -o $<

bin/protoc: bin/protoc-$(protoc_version)
	rm -rf $@
	ln -s ./protoc-$(protoc_version)/bin/protoc $@