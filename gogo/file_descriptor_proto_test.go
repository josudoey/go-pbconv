package gogoconv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/gogo/protobuf/gogoproto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var _ = Describe("GetFileDescriptorProtoByFilename", func() {
	var (
		filename string

		file *descriptorpb.FileDescriptorProto
		err  error
	)

	BeforeEach(func() {
		// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/gogoproto/gogo.pb.go#L787C1-L787C1
		filename = "gogo.proto"
	})

	JustBeforeEach(func() {
		file, err = GetFileDescriptorProtoByFilename(filename)
	})

	It("success", func() {
		Expect(err).To(Succeed())
		Expect(file).ToNot(BeNil())

		// ref https://github.com/gogo/protobuf/blob/f67b8970b736e53dbd7d0a27146c8f1ac52f74e5/gogoproto/gogo.proto
		Expect(file.GetPackage()).To(Equal("gogoproto"))
		Expect(file.GetOptions().GetGoPackage()).To(Equal("github.com/gogo/protobuf/gogoproto"))
		Expect(file.GetExtension()[0].GetExtendee()).To(Equal(".google.protobuf.EnumOptions"))
		Expect(file.GetExtension()[0].GetName()).To(Equal("goproto_enum_prefix"))
		Expect(file.GetExtension()[0].GetNumber()).To(BeEquivalentTo(62001))
	})
})
