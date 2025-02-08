// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: binding_variable_definition_block.proto

package ast

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BindingVariableDefinitionBlock struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BindingVariableDefinitionBlock) Reset() {
	*x = BindingVariableDefinitionBlock{}
	mi := &file_binding_variable_definition_block_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindingVariableDefinitionBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindingVariableDefinitionBlock) ProtoMessage() {}

func (x *BindingVariableDefinitionBlock) ProtoReflect() protoreflect.Message {
	mi := &file_binding_variable_definition_block_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindingVariableDefinitionBlock.ProtoReflect.Descriptor instead.
func (*BindingVariableDefinitionBlock) Descriptor() ([]byte, []int) {
	return file_binding_variable_definition_block_proto_rawDescGZIP(), []int{0}
}

var File_binding_variable_definition_block_proto protoreflect.FileDescriptor

var file_binding_variable_definition_block_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x73, 0x74, 0x22, 0x20,
	0x0a, 0x1e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x62, 0x75, 0x72, 0x62, 0x69, 0x64, 0x67, 0x2f, 0x67, 0x6b, 0x69, 0x74, 0x2d, 0x61, 0x73, 0x74,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_binding_variable_definition_block_proto_rawDescOnce sync.Once
	file_binding_variable_definition_block_proto_rawDescData []byte
)

func file_binding_variable_definition_block_proto_rawDescGZIP() []byte {
	file_binding_variable_definition_block_proto_rawDescOnce.Do(func() {
		file_binding_variable_definition_block_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_binding_variable_definition_block_proto_rawDesc), len(file_binding_variable_definition_block_proto_rawDesc)))
	})
	return file_binding_variable_definition_block_proto_rawDescData
}

var file_binding_variable_definition_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_binding_variable_definition_block_proto_goTypes = []any{
	(*BindingVariableDefinitionBlock)(nil), // 0: ast.BindingVariableDefinitionBlock
}
var file_binding_variable_definition_block_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_binding_variable_definition_block_proto_init() }
func file_binding_variable_definition_block_proto_init() {
	if File_binding_variable_definition_block_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_binding_variable_definition_block_proto_rawDesc), len(file_binding_variable_definition_block_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_binding_variable_definition_block_proto_goTypes,
		DependencyIndexes: file_binding_variable_definition_block_proto_depIdxs,
		MessageInfos:      file_binding_variable_definition_block_proto_msgTypes,
	}.Build()
	File_binding_variable_definition_block_proto = out.File
	file_binding_variable_definition_block_proto_goTypes = nil
	file_binding_variable_definition_block_proto_depIdxs = nil
}
