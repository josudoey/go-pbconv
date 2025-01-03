package gogoconv

import (
	"testing"

	_ "github.com/gogo/protobuf/test/example"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestHelper(t *testing.T) {
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

	// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/test/example/example.proto#L47
	A := file.GetMessageType()[0]
	if want := true; IsFace(file, A) != want {
		t.Errorf("message is face got %v\nwant %v", IsFace(file, A), want)
	}

	if want := false; HasGoGetters(file, A) != want {
		t.Errorf("message has go getter %v\nwant %v", HasGoGetters(file, A), want)
	}

	Description := A.GetField()[0]
	if want := false; IsNullable(Description) != want {
		t.Errorf("field is nullable %v\nwant %v", IsNullable(Description), want)
	}

	Id := A.GetField()[2]
	if want := "github.com/gogo/protobuf/test.Uuid"; GetCustomType(Id) != want {
		t.Errorf("field custom type %v\nwant %v", GetCustomType(Id), want)
	}

	B := file.GetMessageType()[1]
	if want := true; HasDescription(file, B) != want {
		t.Errorf("message has description %v\nwant %v", HasDescription(file, B), want)
	}

	B_A := B.Field[0]
	if want := true; IsEmbed(B_A) != want {
		t.Errorf("field is embed %v\nwant %v", IsEmbed(B_A), want)
	}

	C := file.GetMessageType()[2]
	size := C.GetField()[0]
	if want := "MySize"; GetCustomName(size) != want {
		t.Errorf("field custom name %v\nwant %v", GetCustomName(size), want)
	}

	U := file.GetMessageType()[3]
	if want := true; IsUnion(file, U) != want {
		t.Errorf("message is union %v\nwant %v", IsUnion(file, U), want)
	}

	E := file.GetMessageType()[4]
	if want := false; HasExtensionsMap(file, E) != want {
		t.Errorf("message has extension map %v\nwant %v", HasExtensionsMap(file, E), want)
	}

	R := file.GetMessageType()[5]
	if want := false; HasUnrecognized(file, R) != want {
		t.Errorf("message has unrecognized %v\nwant %v", HasUnrecognized(file, R), want)
	}

	CastType := file.GetMessageType()[6]
	CastType_Int32 := CastType.GetField()[0]
	if want := "int32"; GetCastType(CastType_Int32) != want {
		t.Errorf("message cast type %v\nwant %v", GetCastType(CastType_Int32), want)
	}
}
