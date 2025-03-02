package plugin

import (
	"log"
	"strings"
	"testing"

	"github.com/josudoey/pbconv"
	"github.com/josudoey/pbconv/internal/fixture"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestGenGo(t *testing.T) {
	fileProtoPath := fixture.File_internal_fixture_file_proto.Path()
	fileProto, _ := pbconv.GetFileDescriptorProtoByFilename(fileProtoPath)
	req := &pluginpb.CodeGeneratorRequest{
		FileToGenerate: []string{fileProtoPath},
		ProtoFile: []*descriptorpb.FileDescriptorProto{
			fileProto,
		},
	}

	res, err := Run(req, []string{}, GenGo)
	if err != nil {
		t.Fatalf("os pipe got err: %+v\n", err)
	}

	if want := 1; len(res.File) != want {
		log.Fatalf("response file length got: %q want: %q", len(res.File), want)
	}

	if want := "github.com/josudoey/pbconv/internal/fixture/file.pb.go"; want != res.File[0].GetName() {
		t.Errorf("response file name got: %q want: %q", res.File[0].GetName(), want)
	}

	if want := `// Code generated by protoc-gen-go. DO NOT EDIT.`; want != strings.SplitN(res.File[0].GetContent(), "\n", -1)[0] {
		t.Errorf("response file content got: %q want: %q", res.File[0].GetContent(), want)
	}
}
