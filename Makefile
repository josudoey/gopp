PATH := ${CURDIR}/bin:$(PATH)

.PHONY: wire
wire: bin/wire
	bin/wire  ./...

bin/wire: go.sum
	GOBIN=$(abspath bin) go install github.com/google/wire/cmd/wire@v0.5.0

