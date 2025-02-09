// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: catalog.proto

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

type CatalogParentAndNameType int32

const (
	CatalogParentAndNameType_GRAPH      CatalogParentAndNameType = 0
	CatalogParentAndNameType_GRAPH_TYPE CatalogParentAndNameType = 1
)

// Enum value maps for CatalogParentAndNameType.
var (
	CatalogParentAndNameType_name = map[int32]string{
		0: "GRAPH",
		1: "GRAPH_TYPE",
	}
	CatalogParentAndNameType_value = map[string]int32{
		"GRAPH":      0,
		"GRAPH_TYPE": 1,
	}
)

func (x CatalogParentAndNameType) Enum() *CatalogParentAndNameType {
	p := new(CatalogParentAndNameType)
	*p = x
	return p
}

func (x CatalogParentAndNameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CatalogParentAndNameType) Descriptor() protoreflect.EnumDescriptor {
	return file_catalog_proto_enumTypes[0].Descriptor()
}

func (CatalogParentAndNameType) Type() protoreflect.EnumType {
	return &file_catalog_proto_enumTypes[0]
}

func (x CatalogParentAndNameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CatalogParentAndNameType.Descriptor instead.
func (CatalogParentAndNameType) EnumDescriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{0}
}

type CatalogParentAndName struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Type          CatalogParentAndNameType      `protobuf:"varint,1,opt,name=type,proto3,enum=ast.CatalogParentAndNameType" json:"type,omitempty"`
	Parent        *CatalogObjectParentReference `protobuf:"bytes,2,opt,name=parent,proto3,oneof" json:"parent,omitempty"`
	Name          *Identifier                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CatalogParentAndName) Reset() {
	*x = CatalogParentAndName{}
	mi := &file_catalog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CatalogParentAndName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogParentAndName) ProtoMessage() {}

func (x *CatalogParentAndName) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogParentAndName.ProtoReflect.Descriptor instead.
func (*CatalogParentAndName) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *CatalogParentAndName) GetType() CatalogParentAndNameType {
	if x != nil {
		return x.Type
	}
	return CatalogParentAndNameType_GRAPH
}

func (x *CatalogParentAndName) GetParent() *CatalogObjectParentReference {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *CatalogParentAndName) GetName() *Identifier {
	if x != nil {
		return x.Name
	}
	return nil
}

type CatalogObjectParentReference struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SchemaReference *SchemaReference       `protobuf:"bytes,1,opt,name=schemaReference,proto3" json:"schemaReference,omitempty"`
	ObjectName      []*Identifier          `protobuf:"bytes,2,rep,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CatalogObjectParentReference) Reset() {
	*x = CatalogObjectParentReference{}
	mi := &file_catalog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CatalogObjectParentReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogObjectParentReference) ProtoMessage() {}

func (x *CatalogObjectParentReference) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogObjectParentReference.ProtoReflect.Descriptor instead.
func (*CatalogObjectParentReference) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CatalogObjectParentReference) GetSchemaReference() *SchemaReference {
	if x != nil {
		return x.SchemaReference
	}
	return nil
}

func (x *CatalogObjectParentReference) GetObjectName() []*Identifier {
	if x != nil {
		return x.ObjectName
	}
	return nil
}

type SchemaReference struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*SchemaReference_AbsoluteCatalogPath
	//	*SchemaReference_RelativeCatalogPath
	//	*SchemaReference_ReferenceParameterSpecification
	Type          isSchemaReference_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SchemaReference) Reset() {
	*x = SchemaReference{}
	mi := &file_catalog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchemaReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaReference) ProtoMessage() {}

func (x *SchemaReference) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaReference.ProtoReflect.Descriptor instead.
func (*SchemaReference) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *SchemaReference) GetType() isSchemaReference_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *SchemaReference) GetAbsoluteCatalogPath() *AbsoluteCatalogPath {
	if x != nil {
		if x, ok := x.Type.(*SchemaReference_AbsoluteCatalogPath); ok {
			return x.AbsoluteCatalogPath
		}
	}
	return nil
}

func (x *SchemaReference) GetRelativeCatalogPath() *RelativeCatalogPath {
	if x != nil {
		if x, ok := x.Type.(*SchemaReference_RelativeCatalogPath); ok {
			return x.RelativeCatalogPath
		}
	}
	return nil
}

func (x *SchemaReference) GetReferenceParameterSpecification() *ReferenceParameterSpecification {
	if x != nil {
		if x, ok := x.Type.(*SchemaReference_ReferenceParameterSpecification); ok {
			return x.ReferenceParameterSpecification
		}
	}
	return nil
}

type isSchemaReference_Type interface {
	isSchemaReference_Type()
}

type SchemaReference_AbsoluteCatalogPath struct {
	AbsoluteCatalogPath *AbsoluteCatalogPath `protobuf:"bytes,1,opt,name=absolute_catalog_path,json=absoluteCatalogPath,proto3,oneof"`
}

type SchemaReference_RelativeCatalogPath struct {
	RelativeCatalogPath *RelativeCatalogPath `protobuf:"bytes,2,opt,name=relative_catalog_path,json=relativeCatalogPath,proto3,oneof"`
}

type SchemaReference_ReferenceParameterSpecification struct {
	ReferenceParameterSpecification *ReferenceParameterSpecification `protobuf:"bytes,3,opt,name=reference_parameter_specification,json=referenceParameterSpecification,proto3,oneof"`
}

func (*SchemaReference_AbsoluteCatalogPath) isSchemaReference_Type() {}

func (*SchemaReference_RelativeCatalogPath) isSchemaReference_Type() {}

func (*SchemaReference_ReferenceParameterSpecification) isSchemaReference_Type() {}

type AbsoluteCatalogPath struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AbsoluteCatalogPath) Reset() {
	*x = AbsoluteCatalogPath{}
	mi := &file_catalog_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AbsoluteCatalogPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsoluteCatalogPath) ProtoMessage() {}

func (x *AbsoluteCatalogPath) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsoluteCatalogPath.ProtoReflect.Descriptor instead.
func (*AbsoluteCatalogPath) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *AbsoluteCatalogPath) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RelativeCatalogPath struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelativeCatalogPath) Reset() {
	*x = RelativeCatalogPath{}
	mi := &file_catalog_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelativeCatalogPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelativeCatalogPath) ProtoMessage() {}

func (x *RelativeCatalogPath) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelativeCatalogPath.ProtoReflect.Descriptor instead.
func (*RelativeCatalogPath) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *RelativeCatalogPath) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_catalog_proto protoreflect.FileDescriptor

var file_catalog_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x73, 0x74, 0x1a, 0x0e, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x73,
	0x74, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x3e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x22, 0x90, 0x01, 0x0a, 0x1c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x73, 0x74,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x62, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x41, 0x62, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68,
	0x48, 0x00, 0x52, 0x13, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x4e, 0x0a, 0x15, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68,
	0x48, 0x00, 0x52, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x72, 0x0a, 0x21, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x2b, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x35, 0x0a,
	0x18, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x41,
	0x50, 0x48, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x62, 0x75, 0x72, 0x62, 0x69, 0x64, 0x67, 0x2f, 0x67, 0x6b, 0x69, 0x74,
	0x2d, 0x61, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_catalog_proto_rawDescOnce sync.Once
	file_catalog_proto_rawDescData []byte
)

func file_catalog_proto_rawDescGZIP() []byte {
	file_catalog_proto_rawDescOnce.Do(func() {
		file_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_catalog_proto_rawDesc), len(file_catalog_proto_rawDesc)))
	})
	return file_catalog_proto_rawDescData
}

var file_catalog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_catalog_proto_goTypes = []any{
	(CatalogParentAndNameType)(0),           // 0: ast.CatalogParentAndNameType
	(*CatalogParentAndName)(nil),            // 1: ast.CatalogParentAndName
	(*CatalogObjectParentReference)(nil),    // 2: ast.CatalogObjectParentReference
	(*SchemaReference)(nil),                 // 3: ast.SchemaReference
	(*AbsoluteCatalogPath)(nil),             // 4: ast.AbsoluteCatalogPath
	(*RelativeCatalogPath)(nil),             // 5: ast.RelativeCatalogPath
	(*Identifier)(nil),                      // 6: ast.Identifier
	(*ReferenceParameterSpecification)(nil), // 7: ast.ReferenceParameterSpecification
}
var file_catalog_proto_depIdxs = []int32{
	0, // 0: ast.CatalogParentAndName.type:type_name -> ast.CatalogParentAndNameType
	2, // 1: ast.CatalogParentAndName.parent:type_name -> ast.CatalogObjectParentReference
	6, // 2: ast.CatalogParentAndName.name:type_name -> ast.Identifier
	3, // 3: ast.CatalogObjectParentReference.schemaReference:type_name -> ast.SchemaReference
	6, // 4: ast.CatalogObjectParentReference.object_name:type_name -> ast.Identifier
	4, // 5: ast.SchemaReference.absolute_catalog_path:type_name -> ast.AbsoluteCatalogPath
	5, // 6: ast.SchemaReference.relative_catalog_path:type_name -> ast.RelativeCatalogPath
	7, // 7: ast.SchemaReference.reference_parameter_specification:type_name -> ast.ReferenceParameterSpecification
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_catalog_proto_init() }
func file_catalog_proto_init() {
	if File_catalog_proto != nil {
		return
	}
	file_literals_proto_init()
	file_catalog_proto_msgTypes[0].OneofWrappers = []any{}
	file_catalog_proto_msgTypes[2].OneofWrappers = []any{
		(*SchemaReference_AbsoluteCatalogPath)(nil),
		(*SchemaReference_RelativeCatalogPath)(nil),
		(*SchemaReference_ReferenceParameterSpecification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_catalog_proto_rawDesc), len(file_catalog_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_catalog_proto_goTypes,
		DependencyIndexes: file_catalog_proto_depIdxs,
		EnumInfos:         file_catalog_proto_enumTypes,
		MessageInfos:      file_catalog_proto_msgTypes,
	}.Build()
	File_catalog_proto = out.File
	file_catalog_proto_goTypes = nil
	file_catalog_proto_depIdxs = nil
}
