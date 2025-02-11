// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: program.proto

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

type Program struct {
	state                          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ProgramActivity     *ProgramActivity       `protobuf:"bytes,1,opt,name=program_activity,json=programActivity"`
	xxx_hidden_SessionCloseCommand *SessionCloseCommand   `protobuf:"bytes,2,opt,name=session_close_command,json=sessionCloseCommand"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *Program) Reset() {
	*x = Program{}
	mi := &file_program_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_program_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Program) GetProgramActivity() *ProgramActivity {
	if x != nil {
		return x.xxx_hidden_ProgramActivity
	}
	return nil
}

func (x *Program) GetSessionCloseCommand() *SessionCloseCommand {
	if x != nil {
		return x.xxx_hidden_SessionCloseCommand
	}
	return nil
}

func (x *Program) SetProgramActivity(v *ProgramActivity) {
	x.xxx_hidden_ProgramActivity = v
}

func (x *Program) SetSessionCloseCommand(v *SessionCloseCommand) {
	x.xxx_hidden_SessionCloseCommand = v
}

func (x *Program) HasProgramActivity() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ProgramActivity != nil
}

func (x *Program) HasSessionCloseCommand() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_SessionCloseCommand != nil
}

func (x *Program) ClearProgramActivity() {
	x.xxx_hidden_ProgramActivity = nil
}

func (x *Program) ClearSessionCloseCommand() {
	x.xxx_hidden_SessionCloseCommand = nil
}

type Program_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	ProgramActivity     *ProgramActivity
	SessionCloseCommand *SessionCloseCommand
}

func (b0 Program_builder) Build() *Program {
	m0 := &Program{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_ProgramActivity = b.ProgramActivity
	x.xxx_hidden_SessionCloseCommand = b.SessionCloseCommand
	return m0
}

var File_program_proto protoreflect.FileDescriptor

var file_program_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x73, 0x74, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x12, 0x3f, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x15, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x13, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x42, 0x0e, 0x5a, 0x04, 0x2f, 0x61, 0x73, 0x74, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02,
	0x10, 0x03, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var file_program_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_program_proto_goTypes = []any{
	(*Program)(nil),             // 0: ast.Program
	(*ProgramActivity)(nil),     // 1: ast.ProgramActivity
	(*SessionCloseCommand)(nil), // 2: ast.SessionCloseCommand
}
var file_program_proto_depIdxs = []int32{
	1, // 0: ast.Program.program_activity:type_name -> ast.ProgramActivity
	2, // 1: ast.Program.session_close_command:type_name -> ast.SessionCloseCommand
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_program_proto_init() }
func file_program_proto_init() {
	if File_program_proto != nil {
		return
	}
	file_program_activity_proto_init()
	file_session_commands_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_program_proto_rawDesc), len(file_program_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_program_proto_goTypes,
		DependencyIndexes: file_program_proto_depIdxs,
		MessageInfos:      file_program_proto_msgTypes,
	}.Build()
	File_program_proto = out.File
	file_program_proto_goTypes = nil
	file_program_proto_depIdxs = nil
}
