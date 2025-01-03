package gogoconv

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

func IsEmbed(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Embed, false)
}

func IsNullable(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Nullable, true)
}

func IsStdTime(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Stdtime, false)
}

func IsStdDuration(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Stdduration, false)
}

func IsStdDouble(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.DoubleValue"
}

func IsStdFloat(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.FloatValue"
}

func IsStdInt64(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.Int64Value"
}

func IsStdUInt64(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.UInt64Value"
}

func IsStdInt32(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.Int32Value"
}

func IsStdUInt32(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.UInt32Value"
}

func IsStdBool(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.BoolValue"
}

func IsStdString(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.StringValue"
}

func IsStdBytes(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.BytesValue"
}

func IsStdType(field *descriptorpb.FieldDescriptorProto) bool {
	return (IsStdTime(field) || IsStdDuration(field) ||
		IsStdDouble(field) || IsStdFloat(field) ||
		IsStdInt64(field) || IsStdUInt64(field) ||
		IsStdInt32(field) || IsStdUInt32(field) ||
		IsStdBool(field) ||
		IsStdString(field) || IsStdBytes(field))
}

func IsWktPtr(field *descriptorpb.FieldDescriptorProto) bool {
	return GetBoolExtension(field.Options, E_Wktpointer, false)
}

func IsCustomType(field *descriptorpb.FieldDescriptorProto) bool {
	return len(GetCustomType(field)) > 0
}

func IsCastType(field *descriptorpb.FieldDescriptorProto) bool {
	return len(GetCastType(field)) > 0
}

func IsCastKey(field *descriptorpb.FieldDescriptorProto) bool {
	return len(GetCastKey(field)) > 0
}

func IsCastValue(field *descriptorpb.FieldDescriptorProto) bool {
	return len(GetCastValue(field)) > 0
}

func HasEnumDecl(file *descriptorpb.FileDescriptorProto, enum *descriptorpb.EnumDescriptorProto) bool {
	return GetBoolExtension(enum.Options, E_Enumdecl, GetBoolExtension(file.Options, E_EnumdeclAll, true))
}

func HasTypeDecl(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Typedecl, GetBoolExtension(file.Options, E_TypedeclAll, true))
}

func GetCustomType(field *descriptorpb.FieldDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Customtype)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetCastType(field *descriptorpb.FieldDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Casttype)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetCastKey(field *descriptorpb.FieldDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Castkey)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetCastValue(field *descriptorpb.FieldDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Castvalue)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func IsCustomName(field *descriptorpb.FieldDescriptorProto) bool {
	return len(GetCustomName(field)) > 0
}

func IsEnumCustomName(field *descriptorpb.EnumDescriptorProto) bool {
	return len(GetEnumCustomName(field)) > 0
}

func IsEnumValueCustomName(field *descriptorpb.EnumValueDescriptorProto) bool {
	return len(GetEnumValueCustomName(field)) > 0
}

func GetCustomName(field *descriptorpb.FieldDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Customname)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetEnumCustomName(field *descriptorpb.EnumDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_EnumCustomname)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetEnumValueCustomName(field *descriptorpb.EnumValueDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_EnumvalueCustomname)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetJsonTag(field *descriptorpb.FieldDescriptorProto) *string {
	if field == nil {
		return nil
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Jsontag)
		if err == nil && v.(*string) != nil {
			return (v.(*string))
		}
	}
	return nil
}

func GetMoreTags(field *descriptorpb.FieldDescriptorProto) *string {
	if field == nil {
		return nil
	}
	if field.Options != nil {
		v, err := GetExtension(field.Options, E_Moretags)
		if err == nil && v.(*string) != nil {
			return (v.(*string))
		}
	}
	return nil
}

type EnableFunc func(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool

func EnabledGoEnumPrefix(file *descriptorpb.FileDescriptorProto, enum *descriptorpb.EnumDescriptorProto) bool {
	return GetBoolExtension(enum.Options, E_GoprotoEnumPrefix, GetBoolExtension(file.Options, E_GoprotoEnumPrefixAll, true))
}

func EnabledGoStringer(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_GoprotoStringer, GetBoolExtension(file.Options, E_GoprotoStringerAll, true))
}

func HasGoGetters(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_GoprotoGetters, GetBoolExtension(file.Options, E_GoprotoGettersAll, true))
}

func IsUnion(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Onlyone, GetBoolExtension(file.Options, E_OnlyoneAll, false))
}

func HasGoString(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Gostring, GetBoolExtension(file.Options, E_GostringAll, false))
}

func HasEqual(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Equal, GetBoolExtension(file.Options, E_EqualAll, false))
}

func HasVerboseEqual(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_VerboseEqual, GetBoolExtension(file.Options, E_VerboseEqualAll, false))
}

func IsStringer(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Stringer, GetBoolExtension(file.Options, E_StringerAll, false))
}

func IsFace(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Face, GetBoolExtension(file.Options, E_FaceAll, false))
}

func HasDescription(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Description, GetBoolExtension(file.Options, E_DescriptionAll, false))
}

func HasPopulate(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Populate, GetBoolExtension(file.Options, E_PopulateAll, false))
}

func HasTestGen(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Testgen, GetBoolExtension(file.Options, E_TestgenAll, false))
}

func HasBenchGen(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Benchgen, GetBoolExtension(file.Options, E_BenchgenAll, false))
}

func IsMarshaler(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Marshaler, GetBoolExtension(file.Options, E_MarshalerAll, false))
}

func IsUnmarshaler(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Unmarshaler, GetBoolExtension(file.Options, E_UnmarshalerAll, false))
}

func IsStableMarshaler(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_StableMarshaler, GetBoolExtension(file.Options, E_StableMarshalerAll, false))
}

func IsSizer(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Sizer, GetBoolExtension(file.Options, E_SizerAll, false))
}

func IsProtoSizer(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Protosizer, GetBoolExtension(file.Options, E_ProtosizerAll, false))
}

func IsGoEnumStringer(file *descriptorpb.FileDescriptorProto, enum *descriptorpb.EnumDescriptorProto) bool {
	return GetBoolExtension(enum.Options, E_GoprotoEnumStringer, GetBoolExtension(file.Options, E_GoprotoEnumStringerAll, true))
}

func IsEnumStringer(file *descriptorpb.FileDescriptorProto, enum *descriptorpb.EnumDescriptorProto) bool {
	return GetBoolExtension(enum.Options, E_EnumStringer, GetBoolExtension(file.Options, E_EnumStringerAll, false))
}

func IsUnsafeMarshaler(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_UnsafeMarshaler, GetBoolExtension(file.Options, E_UnsafeMarshalerAll, false))
}

func IsUnsafeUnmarshaler(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_UnsafeUnmarshaler, GetBoolExtension(file.Options, E_UnsafeUnmarshalerAll, false))
}

func HasExtensionsMap(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_GoprotoExtensionsMap, GetBoolExtension(file.Options, E_GoprotoExtensionsMapAll, true))
}

func HasUnrecognized(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_GoprotoUnrecognized, GetBoolExtension(file.Options, E_GoprotoUnrecognizedAll, true))
}

func IsProto3(file *descriptorpb.FileDescriptorProto) bool {
	return file.GetSyntax() == "proto3"
}

func ImportsGoGoProto(file *descriptorpb.FileDescriptorProto) bool {
	return GetBoolExtension(file.Options, E_GogoprotoImport, true)
}

func HasCompare(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Compare, GetBoolExtension(file.Options, E_CompareAll, false))
}

func RegistersGolangProto(file *descriptorpb.FileDescriptorProto) bool {
	return GetBoolExtension(file.Options, E_GoprotoRegistration, false)
}

func HasMessageName(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_Messagename, GetBoolExtension(file.Options, E_MessagenameAll, false))
}

func HasSizecache(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_GoprotoSizecache, GetBoolExtension(file.Options, E_GoprotoSizecacheAll, true))
}

func HasUnkeyed(file *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto) bool {
	return GetBoolExtension(message.Options, E_GoprotoUnkeyed, GetBoolExtension(file.Options, E_GoprotoUnkeyedAll, true))
}
