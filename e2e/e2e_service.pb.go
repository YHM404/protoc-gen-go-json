// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: e2e_service.proto

package e2e

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_e2e_service_proto protoreflect.FileDescriptor

var file_e2e_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x32, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x32, 0x65, 0x1a, 0x09, 0x65, 0x32, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x28, 0x0a, 0x03, 0x45, 0x32, 0x45, 0x12, 0x21, 0x0a, 0x05, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0a, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x1a,
	0x0a, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x65, 0x32, 0x65, 0x3b, 0x65, 0x32, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_e2e_service_proto_goTypes = []any{
	(*Basic)(nil), // 0: e2e.Basic
}
var file_e2e_service_proto_depIdxs = []int32{
	0, // 0: e2e.E2E.Hello:input_type -> e2e.Basic
	0, // 1: e2e.E2E.Hello:output_type -> e2e.Basic
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_e2e_service_proto_init() }
func file_e2e_service_proto_init() {
	if File_e2e_service_proto != nil {
		return
	}
	file_e2e_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_e2e_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_e2e_service_proto_goTypes,
		DependencyIndexes: file_e2e_service_proto_depIdxs,
	}.Build()
	File_e2e_service_proto = out.File
	file_e2e_service_proto_rawDesc = nil
	file_e2e_service_proto_goTypes = nil
	file_e2e_service_proto_depIdxs = nil
}
