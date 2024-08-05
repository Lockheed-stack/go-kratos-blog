// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: api/category/category_err.proto

package category

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CategoryErrorReason int32

const (
	CategoryErrorReason_ERR_CATEGORY_NOT_EXIST              CategoryErrorReason = 0
	CategoryErrorReason_ERR_CATEGORY_INVALID_TITLE          CategoryErrorReason = 1
	CategoryErrorReason_ERR_CATEGORY_INVALID_ID             CategoryErrorReason = 2
	CategoryErrorReason_ERR_CATEGORY_PRE_EXISTING           CategoryErrorReason = 3
	CategoryErrorReason_ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT CategoryErrorReason = 4
)

// Enum value maps for CategoryErrorReason.
var (
	CategoryErrorReason_name = map[int32]string{
		0: "ERR_CATEGORY_NOT_EXIST",
		1: "ERR_CATEGORY_INVALID_TITLE",
		2: "ERR_CATEGORY_INVALID_ID",
		3: "ERR_CATEGORY_PRE_EXISTING",
		4: "ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT",
	}
	CategoryErrorReason_value = map[string]int32{
		"ERR_CATEGORY_NOT_EXIST":              0,
		"ERR_CATEGORY_INVALID_TITLE":          1,
		"ERR_CATEGORY_INVALID_ID":             2,
		"ERR_CATEGORY_PRE_EXISTING":           3,
		"ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT": 4,
	}
)

func (x CategoryErrorReason) Enum() *CategoryErrorReason {
	p := new(CategoryErrorReason)
	*p = x
	return p
}

func (x CategoryErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CategoryErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_category_category_err_proto_enumTypes[0].Descriptor()
}

func (CategoryErrorReason) Type() protoreflect.EnumType {
	return &file_api_category_category_err_proto_enumTypes[0]
}

func (x CategoryErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CategoryErrorReason.Descriptor instead.
func (CategoryErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_category_category_err_proto_rawDescGZIP(), []int{0}
}

var File_api_category_category_err_proto protoreflect.FileDescriptor

var file_api_category_category_err_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x65, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a,
	0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd4, 0x01, 0x0a, 0x13, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x16,
	0x45, 0x52, 0x52, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x24,
	0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x21, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x2d, 0x0a, 0x23,
	0x45, 0x52, 0x52, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x46, 0x4f, 0x52,
	0x45, 0x49, 0x47, 0x4e, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41,
	0x49, 0x4e, 0x54, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x42, 0x20, 0x5a, 0x1e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x3b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_category_category_err_proto_rawDescOnce sync.Once
	file_api_category_category_err_proto_rawDescData = file_api_category_category_err_proto_rawDesc
)

func file_api_category_category_err_proto_rawDescGZIP() []byte {
	file_api_category_category_err_proto_rawDescOnce.Do(func() {
		file_api_category_category_err_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_category_category_err_proto_rawDescData)
	})
	return file_api_category_category_err_proto_rawDescData
}

var file_api_category_category_err_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_category_category_err_proto_goTypes = []any{
	(CategoryErrorReason)(0), // 0: api.category.CategoryErrorReason
}
var file_api_category_category_err_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_category_category_err_proto_init() }
func file_api_category_category_err_proto_init() {
	if File_api_category_category_err_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_category_category_err_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_category_category_err_proto_goTypes,
		DependencyIndexes: file_api_category_category_err_proto_depIdxs,
		EnumInfos:         file_api_category_category_err_proto_enumTypes,
	}.Build()
	File_api_category_category_err_proto = out.File
	file_api_category_category_err_proto_rawDesc = nil
	file_api_category_category_err_proto_goTypes = nil
	file_api_category_category_err_proto_depIdxs = nil
}
