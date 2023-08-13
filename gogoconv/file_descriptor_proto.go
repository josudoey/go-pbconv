package gogoconv

import (
	gogoproto "github.com/gogo/protobuf/proto"
	pbconv "github.com/josudoey/pbconv"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetFileDescriptorProtoByFilename(filename string) (*descriptorpb.FileDescriptorProto, error) {
	// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/protoc-gen-gogo/generator/generator.go#L3256
	// data from gzipped FileDescriptorProto
	return pbconv.GetFileDescriptorProtoByGzippedRaw(gogoproto.FileDescriptor(filename))
}
