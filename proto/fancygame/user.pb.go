// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: user.proto

package fancygame

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

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xda, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12,
	0x13, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6c, 0x6f,
	0x62, 0x62, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e,
	0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0d,
	0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e,
	0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x67, 0x61, 0x6d, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_proto_goTypes = []interface{}{
	(*RegisterInfo)(nil),      // 0: backend.RegisterInfo
	(*LoginInfo)(nil),         // 1: backend.LoginInfo
	(*LogoutInfo)(nil),        // 2: backend.LogoutInfo
	(*EmptyReq)(nil),          // 3: shared.EmptyReq
	(*SetMember)(nil),         // 4: lobby.SetMember
	(*KickOutMemberInfo)(nil), // 5: lobby.KickOutMemberInfo
	(*RegisterRes)(nil),       // 6: backend.RegisterRes
	(*LoginRes)(nil),          // 7: backend.LoginRes
	(*LogoutRes)(nil),         // 8: backend.LogoutRes
	(*GetMemberListRes)(nil),  // 9: lobby.GetMemberListRes
	(*SetMemberRes)(nil),      // 10: lobby.SetMemberRes
	(*KickOutMemberRes)(nil),  // 11: lobby.KickOutMemberRes
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: backenduser.User.Register:input_type -> backend.RegisterInfo
	1,  // 1: backenduser.User.Login:input_type -> backend.LoginInfo
	2,  // 2: backenduser.User.Logout:input_type -> backend.LogoutInfo
	3,  // 3: backenduser.User.GetMemberList:input_type -> shared.EmptyReq
	4,  // 4: backenduser.User.SetMemberData:input_type -> lobby.SetMember
	5,  // 5: backenduser.User.KickOutMember:input_type -> lobby.KickOutMemberInfo
	6,  // 6: backenduser.User.Register:output_type -> backend.RegisterRes
	7,  // 7: backenduser.User.Login:output_type -> backend.LoginRes
	8,  // 8: backenduser.User.Logout:output_type -> backend.LogoutRes
	9,  // 9: backenduser.User.GetMemberList:output_type -> lobby.GetMemberListRes
	10, // 10: backenduser.User.SetMemberData:output_type -> lobby.SetMemberRes
	11, // 11: backenduser.User.KickOutMember:output_type -> lobby.KickOutMemberRes
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_backend_proto_init()
	file_lobby_proto_init()
	file_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}