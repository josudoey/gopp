package main

import (
	"flag"

	"github.com/josudoey/gopp/protoc-gen-grpcadapter/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	// see https://github.com/protocolbuffers/protobuf-go/blob/master/cmd/protoc-gen-go/main.go
	// see https://github.com/grpc/grpc-go/blob/master/cmd/protoc-gen-go-grpc/main.go
	var flags flag.FlagSet
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			gen.GenerateFile(plugin, file)
		}

		return nil
	})
}
