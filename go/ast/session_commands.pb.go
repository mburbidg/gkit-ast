// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: session_commands.proto

package ast

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SessionCloseCommand struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SessionCloseCommand) Reset() {
	*x = SessionCloseCommand{}
	mi := &file_session_commands_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionCloseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCloseCommand) ProtoMessage() {}

func (x *SessionCloseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_session_commands_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type SessionCloseCommand_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 SessionCloseCommand_builder) Build() *SessionCloseCommand {
	m0 := &SessionCloseCommand{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type SessionResetCommand struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SessionResetCommand) Reset() {
	*x = SessionResetCommand{}
	mi := &file_session_commands_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionResetCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResetCommand) ProtoMessage() {}

func (x *SessionResetCommand) ProtoReflect() protoreflect.Message {
	mi := &file_session_commands_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type SessionResetCommand_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 SessionResetCommand_builder) Build() *SessionResetCommand {
	m0 := &SessionResetCommand{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type SessionSetCommand struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SessionSetCommand) Reset() {
	*x = SessionSetCommand{}
	mi := &file_session_commands_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionSetCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionSetCommand) ProtoMessage() {}

func (x *SessionSetCommand) ProtoReflect() protoreflect.Message {
	mi := &file_session_commands_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type SessionSetCommand_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 SessionSetCommand_builder) Build() *SessionSetCommand {
	m0 := &SessionSetCommand{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_session_commands_proto protoreflect.FileDescriptor

var file_session_commands_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x73, 0x74, 0x1a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x13,
	0x0a, 0x11, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x42, 0x10, 0x5a, 0x04, 0x2f, 0x61, 0x73, 0x74, 0x92, 0x03, 0x07, 0xd2, 0x3e,
	0x02, 0x10, 0x03, 0x08, 0x02, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70,
	0xe8, 0x07,
})

var file_session_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_session_commands_proto_goTypes = []any{
	(*SessionCloseCommand)(nil), // 0: ast.SessionCloseCommand
	(*SessionResetCommand)(nil), // 1: ast.SessionResetCommand
	(*SessionSetCommand)(nil),   // 2: ast.SessionSetCommand
}
var file_session_commands_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_session_commands_proto_init() }
func file_session_commands_proto_init() {
	if File_session_commands_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_session_commands_proto_rawDesc), len(file_session_commands_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_commands_proto_goTypes,
		DependencyIndexes: file_session_commands_proto_depIdxs,
		MessageInfos:      file_session_commands_proto_msgTypes,
	}.Build()
	File_session_commands_proto = out.File
	file_session_commands_proto_goTypes = nil
	file_session_commands_proto_depIdxs = nil
}
