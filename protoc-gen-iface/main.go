package main

import (
	"flag"

	"github.com/josudoey/go-pbconv/protoc-gen-iface/internal_geniface"
	"google.golang.org/protobuf/compiler/protogen"
)

// ref https://github.com/protocolbuffers/protobuf-go/blob/6d0a5dbd95005b70501b4cc2c5124dab07a1f4a0/cmd/protoc-gen-go/main.go#L27
func main() {
	var (
		flags flag.FlagSet
	)

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				internal_geniface.GenerateFile(gen, f)
			}
		}
		return nil
	})
}
