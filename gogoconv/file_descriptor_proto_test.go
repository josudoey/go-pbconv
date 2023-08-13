package gogoconv

import (
	"testing"

	"github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/test/example"
	goproto "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestGetFileDescriptorProtoByFilename(t *testing.T) {
	var (
		filename string

		file *descriptorpb.FileDescriptorProto
		err  error
	)

	// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/gogoproto/gogo.pb.go#L787C1-L787C1
	filename = "gogo.proto"

	file, err = GetFileDescriptorProtoByFilename(filename)
	if err != nil {
		t.Fatalf("GetFileDescriptorProtoByFilename(%q): got %+v", filename, err)
	}

	// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/gogoproto/gogo.proto
	if want := "gogo.proto"; file.GetName() != want {
		t.Errorf("file name got %q\nwant %q", file.GetName(), want)
	}

	if want := "gogoproto"; file.GetPackage() != want {
		t.Errorf("file package got %q\nwant %q", file.GetPackage(), want)
	}

	if want := "github.com/gogo/protobuf/gogoproto"; file.GetOptions().GetGoPackage() != want {
		t.Errorf("go package option got %q\nwant %q", file.GetOptions().GetGoPackage(), want)
	}

	{
		extension := file.GetExtension()[0]
		if want := ".google.protobuf.EnumOptions"; extension.GetExtendee() != want {
			t.Errorf("extension got %q\nwant %q", extension.GetExtendee(), want)
		}

		if want := "goproto_enum_prefix"; extension.GetName() != want {
			t.Errorf("extension name got %q\nwant %q", extension.GetName(), want)
		}

		if want := int32(62001); extension.GetNumber() != want {
			t.Errorf("extension name got %q\nwant %q", extension.GetNumber(), want)
		}
	}
}

func TestGetFileDescriptorProtoByFilename_Example_Proto(t *testing.T) {
	var (
		filename string

		file *descriptorpb.FileDescriptorProto
		err  error
	)

	// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/test/example/example.pb.go#L371
	filename = "example.proto"

	file, err = GetFileDescriptorProtoByFilename(filename)
	if err != nil {
		t.Fatalf("GetFileDescriptorProtoByFilename(%q): got %+v", filename, err)
	}

	// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/gogoproto/gogo.proto
	if want := "example.proto"; file.GetName() != want {
		t.Errorf("file name got %q\nwant %q", file.GetName(), want)
	}

	if want := "test"; file.GetPackage() != want {
		t.Errorf("file package got %q\nwant %q", file.GetPackage(), want)
	}

	if want := ""; file.GetOptions().GetGoPackage() != want {
		t.Errorf("go package option got %q\nwant %q", file.GetOptions().GetGoPackage(), want)
	}

	{
		// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/test/example/example.proto#L47
		message := file.GetMessageType()[0]
		if want := "A"; message.GetName() != want {
			t.Errorf("message name got %q\nwant %q", message.GetName(), want)
		}

		v, err := goproto.GetExtension(message.GetOptions(), GetExtensionInfo(gogoproto.E_Face))
		if err != nil {
			t.Fatalf("get message option gogoproto.face got err: %+v", err)
		}

		switch face := v.(type) {
		case *bool:
			if want := true; *face != want {
				t.Errorf("get message option gogoproto.face got %v\nwant %v", *face, want)
			}
		default:
			t.Fatalf("get message option gogoproto.face got type: %t", face)
		}

		if want := "Description"; message.GetField()[0].GetName() != want {
			t.Errorf("message field name got %q\nwant %q", message.GetField()[0].GetName(), want)
		}

		if want := "Number"; message.GetField()[1].GetName() != want {
			t.Errorf("message field name got %q\nwant %q", message.GetField()[1].GetName(), want)
		}

		if want := "Id"; message.GetField()[2].GetName() != want {
			t.Errorf("message field name got %q\nwant %q", message.GetField()[2].GetName(), want)
		}

		v, err = goproto.GetExtension(message.GetField()[2].GetOptions(), GetExtensionInfo(gogoproto.E_Customtype))
		if err != nil {
			t.Fatalf("get message field option gogoproto.customtype got err: %+v", err)
		}

		switch customtype := v.(type) {
		case *string:
			if want := "github.com/gogo/protobuf/test.Uuid"; *customtype != want {
				t.Errorf("get message option gogoproto.customtype got %v\nwant %v", *customtype, want)
			}
		default:
			t.Fatalf("get message option gogoproto.customtype got type: %t", customtype)
		}

	}
}
