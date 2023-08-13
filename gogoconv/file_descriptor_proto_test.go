package gogoconv

import (
	"testing"

	_ "github.com/gogo/protobuf/gogoproto"
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
