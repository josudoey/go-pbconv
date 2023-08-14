package gogoconv

import (
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

func NormalizeExtendedType(extendedType protoiface.MessageV1) protoiface.MessageV1 {
	switch extendedType.(type) {
	case *descriptor.EnumOptions:
		return (*descriptorpb.EnumOptions)(nil)
	case *descriptor.EnumValueOptions:
		return (*descriptorpb.EnumValueOptions)(nil)
	case *descriptor.FieldOptions:
		return (*descriptorpb.FieldOptions)(nil)
	case *descriptor.FileOptions:
		return (*descriptorpb.FileOptions)(nil)
	case *descriptor.MessageOptions:
		return (*descriptorpb.MessageOptions)(nil)
	case *descriptor.MethodOptions:
		return (*descriptorpb.MethodOptions)(nil)
	case *descriptor.OneofOptions:
		return (*descriptorpb.OneofOptions)(nil)
	case *descriptor.ServiceOptions:
		return (*descriptorpb.ServiceOptions)(nil)
	}

	return extendedType
}

func GetExtensionInfo(extensionDesc *gogoproto.ExtensionDesc) *protoimpl.ExtensionInfo {
	return &protoimpl.ExtensionInfo{
		ExtendedType:  NormalizeExtendedType(extensionDesc.ExtendedType),
		ExtensionType: extensionDesc.ExtensionType,
		Field:         extensionDesc.Field,
		Name:          extensionDesc.Name,
		Tag:           extensionDesc.Tag,
		Filename:      extensionDesc.Filename,
	}
}
