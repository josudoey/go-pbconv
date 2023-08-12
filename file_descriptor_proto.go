package pbconv

import (
	"bytes"
	"compress/gzip"
	"io"

	goproto "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetFileDescriptorProtoByRaw(rawDesc []byte) (*descriptorpb.FileDescriptorProto, error) {
	var file descriptorpb.FileDescriptorProto
	if err := proto.Unmarshal(rawDesc, &file); err != nil {
		return nil, err
	}

	return &file, nil
}

func GetFileDescriptorProtoByGzippedRaw(gzippedRawDesc []byte) (*descriptorpb.FileDescriptorProto, error) {
	reader, err := gzip.NewReader(bytes.NewReader(gzippedRawDesc))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	rawDesc, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return GetFileDescriptorProtoByRaw(rawDesc)
}

func GetFileDescriptorProtoByFilename(filename string) (*descriptorpb.FileDescriptorProto, error) {
	// ref https://github.com/golang/protobuf/blob/5d5e8c018a13017f9d5b8bf4fad64aaa42a87308/protoc-gen-go/generator/generator.go#L2597
	// data from gzipped FileDescriptorProto
	return GetFileDescriptorProtoByGzippedRaw(goproto.FileDescriptor(filename))
}

func GetFileDescriptorProtoByMessage(message goproto.Message) (*descriptorpb.FileDescriptorProto, error) {
	messageReflect := goproto.MessageReflect(message)
	return GetFileDescriptorProtoByFilename(messageReflect.Descriptor().ParentFile().Path())
}
