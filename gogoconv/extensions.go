package gogoconv

import (
	"reflect"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var (
	E_GoprotoEnumPrefix       = GetExtensionInfo(gogoproto.E_GoprotoEnumPrefix)
	E_GoprotoEnumStringer     = GetExtensionInfo(gogoproto.E_GoprotoEnumStringer)
	E_EnumStringer            = GetExtensionInfo(gogoproto.E_EnumStringer)
	E_EnumCustomname          = GetExtensionInfo(gogoproto.E_EnumCustomname)
	E_Enumdecl                = GetExtensionInfo(gogoproto.E_Enumdecl)
	E_EnumvalueCustomname     = GetExtensionInfo(gogoproto.E_EnumvalueCustomname)
	E_GoprotoGettersAll       = GetExtensionInfo(gogoproto.E_GoprotoGettersAll)
	E_GoprotoEnumPrefixAll    = GetExtensionInfo(gogoproto.E_GoprotoEnumPrefixAll)
	E_GoprotoStringerAll      = GetExtensionInfo(gogoproto.E_GoprotoStringerAll)
	E_VerboseEqualAll         = GetExtensionInfo(gogoproto.E_VerboseEqualAll)
	E_FaceAll                 = GetExtensionInfo(gogoproto.E_FaceAll)
	E_GostringAll             = GetExtensionInfo(gogoproto.E_GostringAll)
	E_PopulateAll             = GetExtensionInfo(gogoproto.E_PopulateAll)
	E_StringerAll             = GetExtensionInfo(gogoproto.E_StringerAll)
	E_OnlyoneAll              = GetExtensionInfo(gogoproto.E_OnlyoneAll)
	E_EqualAll                = GetExtensionInfo(gogoproto.E_EqualAll)
	E_DescriptionAll          = GetExtensionInfo(gogoproto.E_DescriptionAll)
	E_TestgenAll              = GetExtensionInfo(gogoproto.E_TestgenAll)
	E_BenchgenAll             = GetExtensionInfo(gogoproto.E_BenchgenAll)
	E_MarshalerAll            = GetExtensionInfo(gogoproto.E_MarshalerAll)
	E_UnmarshalerAll          = GetExtensionInfo(gogoproto.E_UnmarshalerAll)
	E_StableMarshalerAll      = GetExtensionInfo(gogoproto.E_StableMarshalerAll)
	E_SizerAll                = GetExtensionInfo(gogoproto.E_SizerAll)
	E_GoprotoEnumStringerAll  = GetExtensionInfo(gogoproto.E_GoprotoEnumStringerAll)
	E_EnumStringerAll         = GetExtensionInfo(gogoproto.E_EnumStringerAll)
	E_UnsafeMarshalerAll      = GetExtensionInfo(gogoproto.E_UnsafeMarshalerAll)
	E_UnsafeUnmarshalerAll    = GetExtensionInfo(gogoproto.E_UnsafeUnmarshalerAll)
	E_GoprotoExtensionsMapAll = GetExtensionInfo(gogoproto.E_GoprotoExtensionsMapAll)
	E_GoprotoUnrecognizedAll  = GetExtensionInfo(gogoproto.E_GoprotoUnrecognizedAll)
	E_GogoprotoImport         = GetExtensionInfo(gogoproto.E_GogoprotoImport)
	E_ProtosizerAll           = GetExtensionInfo(gogoproto.E_ProtosizerAll)
	E_CompareAll              = GetExtensionInfo(gogoproto.E_CompareAll)
	E_TypedeclAll             = GetExtensionInfo(gogoproto.E_TypedeclAll)
	E_EnumdeclAll             = GetExtensionInfo(gogoproto.E_EnumdeclAll)
	E_GoprotoRegistration     = GetExtensionInfo(gogoproto.E_GoprotoRegistration)
	E_MessagenameAll          = GetExtensionInfo(gogoproto.E_MessagenameAll)
	E_GoprotoSizecacheAll     = GetExtensionInfo(gogoproto.E_GoprotoSizecacheAll)
	E_GoprotoUnkeyedAll       = GetExtensionInfo(gogoproto.E_GoprotoUnkeyedAll)
	E_GoprotoGetters          = GetExtensionInfo(gogoproto.E_GoprotoGetters)
	E_GoprotoStringer         = GetExtensionInfo(gogoproto.E_GoprotoStringer)
	E_VerboseEqual            = GetExtensionInfo(gogoproto.E_VerboseEqual)
	E_Face                    = GetExtensionInfo(gogoproto.E_Face)
	E_Gostring                = GetExtensionInfo(gogoproto.E_Gostring)
	E_Populate                = GetExtensionInfo(gogoproto.E_Populate)
	E_Stringer                = GetExtensionInfo(gogoproto.E_Stringer)
	E_Onlyone                 = GetExtensionInfo(gogoproto.E_Onlyone)
	E_Equal                   = GetExtensionInfo(gogoproto.E_Equal)
	E_Description             = GetExtensionInfo(gogoproto.E_Description)
	E_Testgen                 = GetExtensionInfo(gogoproto.E_Testgen)
	E_Benchgen                = GetExtensionInfo(gogoproto.E_Benchgen)
	E_Marshaler               = GetExtensionInfo(gogoproto.E_Marshaler)
	E_Unmarshaler             = GetExtensionInfo(gogoproto.E_Unmarshaler)
	E_StableMarshaler         = GetExtensionInfo(gogoproto.E_StableMarshaler)
	E_Sizer                   = GetExtensionInfo(gogoproto.E_Sizer)
	E_UnsafeMarshaler         = GetExtensionInfo(gogoproto.E_UnsafeMarshaler)
	E_UnsafeUnmarshaler       = GetExtensionInfo(gogoproto.E_UnsafeUnmarshaler)
	E_GoprotoExtensionsMap    = GetExtensionInfo(gogoproto.E_GoprotoExtensionsMap)
	E_GoprotoUnrecognized     = GetExtensionInfo(gogoproto.E_GoprotoUnrecognized)
	E_Protosizer              = GetExtensionInfo(gogoproto.E_Protosizer)
	E_Compare                 = GetExtensionInfo(gogoproto.E_Compare)
	E_Typedecl                = GetExtensionInfo(gogoproto.E_Typedecl)
	E_Messagename             = GetExtensionInfo(gogoproto.E_Messagename)
	E_GoprotoSizecache        = GetExtensionInfo(gogoproto.E_GoprotoSizecache)
	E_GoprotoUnkeyed          = GetExtensionInfo(gogoproto.E_GoprotoUnkeyed)
	E_Nullable                = GetExtensionInfo(gogoproto.E_Nullable)
	E_Embed                   = GetExtensionInfo(gogoproto.E_Embed)
	E_Customtype              = GetExtensionInfo(gogoproto.E_Customtype)
	E_Customname              = GetExtensionInfo(gogoproto.E_Customname)
	E_Jsontag                 = GetExtensionInfo(gogoproto.E_Jsontag)
	E_Moretags                = GetExtensionInfo(gogoproto.E_Moretags)
	E_Casttype                = GetExtensionInfo(gogoproto.E_Casttype)
	E_Castkey                 = GetExtensionInfo(gogoproto.E_Castkey)
	E_Castvalue               = GetExtensionInfo(gogoproto.E_Castvalue)
	E_Stdtime                 = GetExtensionInfo(gogoproto.E_Stdtime)
	E_Stdduration             = GetExtensionInfo(gogoproto.E_Stdduration)
	E_Wktpointer              = GetExtensionInfo(gogoproto.E_Wktpointer)
)

var GetExtension = proto.GetExtension

// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/proto/extensions_gogo.go#L65
func GetBoolExtension(pb proto.Message, extension *protoimpl.ExtensionInfo, ifnotset bool) bool {
	if reflect.ValueOf(pb).IsNil() {
		return ifnotset
	}
	value, err := GetExtension(pb, extension)
	if err != nil {
		return ifnotset
	}
	if value == nil {
		return ifnotset
	}
	if value.(*bool) == nil {
		return ifnotset
	}
	return *(value.(*bool))
}


