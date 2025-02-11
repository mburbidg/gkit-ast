// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: graph_type.proto

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

type GraphType struct {
	state           protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Type isGraphType_Type       `protobuf_oneof:"type"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GraphType) Reset() {
	*x = GraphType{}
	mi := &file_graph_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GraphType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphType) ProtoMessage() {}

func (x *GraphType) ProtoReflect() protoreflect.Message {
	mi := &file_graph_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GraphType) GetGraphTypeLikeGraph() *GraphTypeLikeGraph {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*graphType_GraphTypeLikeGraph); ok {
			return x.GraphTypeLikeGraph
		}
	}
	return nil
}

func (x *GraphType) GetGraphTypeReference() *GraphTypeReference {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*graphType_GraphTypeReference); ok {
			return x.GraphTypeReference
		}
	}
	return nil
}

func (x *GraphType) GetNestedGraphTypeSpecification() *NestedGraphTypeSpecification {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*graphType_NestedGraphTypeSpecification); ok {
			return x.NestedGraphTypeSpecification
		}
	}
	return nil
}

func (x *GraphType) SetGraphTypeLikeGraph(v *GraphTypeLikeGraph) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &graphType_GraphTypeLikeGraph{v}
}

func (x *GraphType) SetGraphTypeReference(v *GraphTypeReference) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &graphType_GraphTypeReference{v}
}

func (x *GraphType) SetNestedGraphTypeSpecification(v *NestedGraphTypeSpecification) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &graphType_NestedGraphTypeSpecification{v}
}

func (x *GraphType) HasType() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Type != nil
}

func (x *GraphType) HasGraphTypeLikeGraph() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*graphType_GraphTypeLikeGraph)
	return ok
}

func (x *GraphType) HasGraphTypeReference() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*graphType_GraphTypeReference)
	return ok
}

func (x *GraphType) HasNestedGraphTypeSpecification() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*graphType_NestedGraphTypeSpecification)
	return ok
}

func (x *GraphType) ClearType() {
	x.xxx_hidden_Type = nil
}

func (x *GraphType) ClearGraphTypeLikeGraph() {
	if _, ok := x.xxx_hidden_Type.(*graphType_GraphTypeLikeGraph); ok {
		x.xxx_hidden_Type = nil
	}
}

func (x *GraphType) ClearGraphTypeReference() {
	if _, ok := x.xxx_hidden_Type.(*graphType_GraphTypeReference); ok {
		x.xxx_hidden_Type = nil
	}
}

func (x *GraphType) ClearNestedGraphTypeSpecification() {
	if _, ok := x.xxx_hidden_Type.(*graphType_NestedGraphTypeSpecification); ok {
		x.xxx_hidden_Type = nil
	}
}

const GraphType_Type_not_set_case case_GraphType_Type = 0
const GraphType_GraphTypeLikeGraph_case case_GraphType_Type = 1
const GraphType_GraphTypeReference_case case_GraphType_Type = 2
const GraphType_NestedGraphTypeSpecification_case case_GraphType_Type = 3

func (x *GraphType) WhichType() case_GraphType_Type {
	if x == nil {
		return GraphType_Type_not_set_case
	}
	switch x.xxx_hidden_Type.(type) {
	case *graphType_GraphTypeLikeGraph:
		return GraphType_GraphTypeLikeGraph_case
	case *graphType_GraphTypeReference:
		return GraphType_GraphTypeReference_case
	case *graphType_NestedGraphTypeSpecification:
		return GraphType_NestedGraphTypeSpecification_case
	default:
		return GraphType_Type_not_set_case
	}
}

type GraphType_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_Type:
	GraphTypeLikeGraph           *GraphTypeLikeGraph
	GraphTypeReference           *GraphTypeReference
	NestedGraphTypeSpecification *NestedGraphTypeSpecification
	// -- end of xxx_hidden_Type
}

func (b0 GraphType_builder) Build() *GraphType {
	m0 := &GraphType{}
	b, x := &b0, m0
	_, _ = b, x
	if b.GraphTypeLikeGraph != nil {
		x.xxx_hidden_Type = &graphType_GraphTypeLikeGraph{b.GraphTypeLikeGraph}
	}
	if b.GraphTypeReference != nil {
		x.xxx_hidden_Type = &graphType_GraphTypeReference{b.GraphTypeReference}
	}
	if b.NestedGraphTypeSpecification != nil {
		x.xxx_hidden_Type = &graphType_NestedGraphTypeSpecification{b.NestedGraphTypeSpecification}
	}
	return m0
}

type case_GraphType_Type protoreflect.FieldNumber

func (x case_GraphType_Type) String() string {
	md := file_graph_type_proto_msgTypes[0].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isGraphType_Type interface {
	isGraphType_Type()
}

type graphType_GraphTypeLikeGraph struct {
	GraphTypeLikeGraph *GraphTypeLikeGraph `protobuf:"bytes,1,opt,name=graph_type_like_graph,json=graphTypeLikeGraph,oneof"`
}

type graphType_GraphTypeReference struct {
	GraphTypeReference *GraphTypeReference `protobuf:"bytes,2,opt,name=graph_type_reference,json=graphTypeReference,oneof"`
}

type graphType_NestedGraphTypeSpecification struct {
	NestedGraphTypeSpecification *NestedGraphTypeSpecification `protobuf:"bytes,3,opt,name=nested_graph_type_specification,json=nestedGraphTypeSpecification,oneof"`
}

func (*graphType_GraphTypeLikeGraph) isGraphType_Type() {}

func (*graphType_GraphTypeReference) isGraphType_Type() {}

func (*graphType_NestedGraphTypeSpecification) isGraphType_Type() {}

type GraphTypeLikeGraph struct {
	state                      protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_GraphExpression *GraphExpression       `protobuf:"bytes,1,opt,name=graph_expression,json=graphExpression"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *GraphTypeLikeGraph) Reset() {
	*x = GraphTypeLikeGraph{}
	mi := &file_graph_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GraphTypeLikeGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTypeLikeGraph) ProtoMessage() {}

func (x *GraphTypeLikeGraph) ProtoReflect() protoreflect.Message {
	mi := &file_graph_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GraphTypeLikeGraph) GetGraphExpression() *GraphExpression {
	if x != nil {
		return x.xxx_hidden_GraphExpression
	}
	return nil
}

func (x *GraphTypeLikeGraph) SetGraphExpression(v *GraphExpression) {
	x.xxx_hidden_GraphExpression = v
}

func (x *GraphTypeLikeGraph) HasGraphExpression() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_GraphExpression != nil
}

func (x *GraphTypeLikeGraph) ClearGraphExpression() {
	x.xxx_hidden_GraphExpression = nil
}

type GraphTypeLikeGraph_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	GraphExpression *GraphExpression
}

func (b0 GraphTypeLikeGraph_builder) Build() *GraphTypeLikeGraph {
	m0 := &GraphTypeLikeGraph{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_GraphExpression = b.GraphExpression
	return m0
}

type NestedGraphTypeSpecification struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NestedGraphTypeSpecification) Reset() {
	*x = NestedGraphTypeSpecification{}
	mi := &file_graph_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NestedGraphTypeSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedGraphTypeSpecification) ProtoMessage() {}

func (x *NestedGraphTypeSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_graph_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type NestedGraphTypeSpecification_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 NestedGraphTypeSpecification_builder) Build() *NestedGraphTypeSpecification {
	m0 := &NestedGraphTypeSpecification{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type GraphTypeReference struct {
	state           protoimpl.MessageState    `protogen:"opaque.v1"`
	xxx_hidden_Type isGraphTypeReference_Type `protobuf_oneof:"type"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GraphTypeReference) Reset() {
	*x = GraphTypeReference{}
	mi := &file_graph_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GraphTypeReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTypeReference) ProtoMessage() {}

func (x *GraphTypeReference) ProtoReflect() protoreflect.Message {
	mi := &file_graph_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GraphTypeReference) GetParentAndName() *CatalogParentAndName {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*graphTypeReference_ParentAndName); ok {
			return x.ParentAndName
		}
	}
	return nil
}

func (x *GraphTypeReference) GetReferenceParameterSpecification() *ReferenceParameterSpecification {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*graphTypeReference_ReferenceParameterSpecification); ok {
			return x.ReferenceParameterSpecification
		}
	}
	return nil
}

func (x *GraphTypeReference) SetParentAndName(v *CatalogParentAndName) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &graphTypeReference_ParentAndName{v}
}

func (x *GraphTypeReference) SetReferenceParameterSpecification(v *ReferenceParameterSpecification) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &graphTypeReference_ReferenceParameterSpecification{v}
}

func (x *GraphTypeReference) HasType() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Type != nil
}

func (x *GraphTypeReference) HasParentAndName() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*graphTypeReference_ParentAndName)
	return ok
}

func (x *GraphTypeReference) HasReferenceParameterSpecification() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*graphTypeReference_ReferenceParameterSpecification)
	return ok
}

func (x *GraphTypeReference) ClearType() {
	x.xxx_hidden_Type = nil
}

func (x *GraphTypeReference) ClearParentAndName() {
	if _, ok := x.xxx_hidden_Type.(*graphTypeReference_ParentAndName); ok {
		x.xxx_hidden_Type = nil
	}
}

func (x *GraphTypeReference) ClearReferenceParameterSpecification() {
	if _, ok := x.xxx_hidden_Type.(*graphTypeReference_ReferenceParameterSpecification); ok {
		x.xxx_hidden_Type = nil
	}
}

const GraphTypeReference_Type_not_set_case case_GraphTypeReference_Type = 0
const GraphTypeReference_ParentAndName_case case_GraphTypeReference_Type = 1
const GraphTypeReference_ReferenceParameterSpecification_case case_GraphTypeReference_Type = 2

func (x *GraphTypeReference) WhichType() case_GraphTypeReference_Type {
	if x == nil {
		return GraphTypeReference_Type_not_set_case
	}
	switch x.xxx_hidden_Type.(type) {
	case *graphTypeReference_ParentAndName:
		return GraphTypeReference_ParentAndName_case
	case *graphTypeReference_ReferenceParameterSpecification:
		return GraphTypeReference_ReferenceParameterSpecification_case
	default:
		return GraphTypeReference_Type_not_set_case
	}
}

type GraphTypeReference_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_Type:
	ParentAndName                   *CatalogParentAndName
	ReferenceParameterSpecification *ReferenceParameterSpecification
	// -- end of xxx_hidden_Type
}

func (b0 GraphTypeReference_builder) Build() *GraphTypeReference {
	m0 := &GraphTypeReference{}
	b, x := &b0, m0
	_, _ = b, x
	if b.ParentAndName != nil {
		x.xxx_hidden_Type = &graphTypeReference_ParentAndName{b.ParentAndName}
	}
	if b.ReferenceParameterSpecification != nil {
		x.xxx_hidden_Type = &graphTypeReference_ReferenceParameterSpecification{b.ReferenceParameterSpecification}
	}
	return m0
}

type case_GraphTypeReference_Type protoreflect.FieldNumber

func (x case_GraphTypeReference_Type) String() string {
	md := file_graph_type_proto_msgTypes[3].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isGraphTypeReference_Type interface {
	isGraphTypeReference_Type()
}

type graphTypeReference_ParentAndName struct {
	ParentAndName *CatalogParentAndName `protobuf:"bytes,1,opt,name=parent_and_name,json=parentAndName,oneof"`
}

type graphTypeReference_ReferenceParameterSpecification struct {
	ReferenceParameterSpecification *ReferenceParameterSpecification `protobuf:"bytes,2,opt,name=reference_parameter_specification,json=referenceParameterSpecification,oneof"`
}

func (*graphTypeReference_ParentAndName) isGraphTypeReference_Type() {}

func (*graphTypeReference_ReferenceParameterSpecification) isGraphTypeReference_Type() {}

var File_graph_type_proto protoreflect.FileDescriptor

var file_graph_type_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x73, 0x74, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x4c, 0x0a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6c, 0x69,
	0x6b, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x6b, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x48, 0x00, 0x52, 0x12, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x4b, 0x0a,
	0x14, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x73,
	0x74, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x12, 0x67, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x1f, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x55,
	0x0a, 0x12, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x12, 0x3f, 0x0a, 0x10, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x61, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x1c, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x54, 0x79, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd5, 0x01, 0x0a, 0x12, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x48, 0x00, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x72, 0x0a, 0x21, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x5a,
	0x04, 0x2f, 0x61, 0x73, 0x74, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x03, 0x62, 0x08, 0x65,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var file_graph_type_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_graph_type_proto_goTypes = []any{
	(*GraphType)(nil),                       // 0: ast.GraphType
	(*GraphTypeLikeGraph)(nil),              // 1: ast.GraphTypeLikeGraph
	(*NestedGraphTypeSpecification)(nil),    // 2: ast.NestedGraphTypeSpecification
	(*GraphTypeReference)(nil),              // 3: ast.GraphTypeReference
	(*GraphExpression)(nil),                 // 4: ast.GraphExpression
	(*CatalogParentAndName)(nil),            // 5: ast.CatalogParentAndName
	(*ReferenceParameterSpecification)(nil), // 6: ast.ReferenceParameterSpecification
}
var file_graph_type_proto_depIdxs = []int32{
	1, // 0: ast.GraphType.graph_type_like_graph:type_name -> ast.GraphTypeLikeGraph
	3, // 1: ast.GraphType.graph_type_reference:type_name -> ast.GraphTypeReference
	2, // 2: ast.GraphType.nested_graph_type_specification:type_name -> ast.NestedGraphTypeSpecification
	4, // 3: ast.GraphTypeLikeGraph.graph_expression:type_name -> ast.GraphExpression
	5, // 4: ast.GraphTypeReference.parent_and_name:type_name -> ast.CatalogParentAndName
	6, // 5: ast.GraphTypeReference.reference_parameter_specification:type_name -> ast.ReferenceParameterSpecification
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_graph_type_proto_init() }
func file_graph_type_proto_init() {
	if File_graph_type_proto != nil {
		return
	}
	file_catalog_proto_init()
	file_literals_proto_init()
	file_graph_expression_proto_init()
	file_graph_type_proto_msgTypes[0].OneofWrappers = []any{
		(*graphType_GraphTypeLikeGraph)(nil),
		(*graphType_GraphTypeReference)(nil),
		(*graphType_NestedGraphTypeSpecification)(nil),
	}
	file_graph_type_proto_msgTypes[3].OneofWrappers = []any{
		(*graphTypeReference_ParentAndName)(nil),
		(*graphTypeReference_ReferenceParameterSpecification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_graph_type_proto_rawDesc), len(file_graph_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_graph_type_proto_goTypes,
		DependencyIndexes: file_graph_type_proto_depIdxs,
		MessageInfos:      file_graph_type_proto_msgTypes,
	}.Build()
	File_graph_type_proto = out.File
	file_graph_type_proto_goTypes = nil
	file_graph_type_proto_depIdxs = nil
}
