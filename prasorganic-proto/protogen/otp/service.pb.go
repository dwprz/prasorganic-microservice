// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/otp/service.proto

package otp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_otp_service_proto protoreflect.FileDescriptor

var file_proto_otp_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x74, 0x70, 0x1a, 0x18,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6f,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x77, 0x0a, 0x0a, 0x4f, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x6f, 0x74,
	0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x12, 0x13, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x77, 0x70,
	0x72, 0x7a, 0x2f, 0x70, 0x72, 0x61, 0x73, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x74,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_otp_service_proto_goTypes = []any{
	(*SendReq)(nil),       // 0: otp.type.SendReq
	(*VerifyReq)(nil),     // 1: otp.type.VerifyReq
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
	(*VerifyRes)(nil),     // 3: otp.type.VerifyRes
}
var file_proto_otp_service_proto_depIdxs = []int32{
	0, // 0: otp.OtpService.Send:input_type -> otp.type.SendReq
	1, // 1: otp.OtpService.Verify:input_type -> otp.type.VerifyReq
	2, // 2: otp.OtpService.Send:output_type -> google.protobuf.Empty
	3, // 3: otp.OtpService.Verify:output_type -> otp.type.VerifyRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_otp_service_proto_init() }
func file_proto_otp_service_proto_init() {
	if File_proto_otp_service_proto != nil {
		return
	}
	file_proto_otp_type_otp_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_otp_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_otp_service_proto_goTypes,
		DependencyIndexes: file_proto_otp_service_proto_depIdxs,
	}.Build()
	File_proto_otp_service_proto = out.File
	file_proto_otp_service_proto_rawDesc = nil
	file_proto_otp_service_proto_goTypes = nil
	file_proto_otp_service_proto_depIdxs = nil
}
