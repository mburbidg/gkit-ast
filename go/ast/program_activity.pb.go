// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: program_activity.proto

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

type ProgramActivity struct {
	state           protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Type isProgramActivity_Type `protobuf_oneof:"type"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ProgramActivity) Reset() {
	*x = ProgramActivity{}
	mi := &file_program_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProgramActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramActivity) ProtoMessage() {}

func (x *ProgramActivity) ProtoReflect() protoreflect.Message {
	mi := &file_program_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ProgramActivity) GetSessionActivity() *SessionActivity {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*programActivity_SessionActivity); ok {
			return x.SessionActivity
		}
	}
	return nil
}

func (x *ProgramActivity) GetTransactionActivity() *TransactionActivity {
	if x != nil {
		if x, ok := x.xxx_hidden_Type.(*programActivity_TransactionActivity); ok {
			return x.TransactionActivity
		}
	}
	return nil
}

func (x *ProgramActivity) SetSessionActivity(v *SessionActivity) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &programActivity_SessionActivity{v}
}

func (x *ProgramActivity) SetTransactionActivity(v *TransactionActivity) {
	if v == nil {
		x.xxx_hidden_Type = nil
		return
	}
	x.xxx_hidden_Type = &programActivity_TransactionActivity{v}
}

func (x *ProgramActivity) HasType() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Type != nil
}

func (x *ProgramActivity) HasSessionActivity() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*programActivity_SessionActivity)
	return ok
}

func (x *ProgramActivity) HasTransactionActivity() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Type.(*programActivity_TransactionActivity)
	return ok
}

func (x *ProgramActivity) ClearType() {
	x.xxx_hidden_Type = nil
}

func (x *ProgramActivity) ClearSessionActivity() {
	if _, ok := x.xxx_hidden_Type.(*programActivity_SessionActivity); ok {
		x.xxx_hidden_Type = nil
	}
}

func (x *ProgramActivity) ClearTransactionActivity() {
	if _, ok := x.xxx_hidden_Type.(*programActivity_TransactionActivity); ok {
		x.xxx_hidden_Type = nil
	}
}

const ProgramActivity_Type_not_set_case case_ProgramActivity_Type = 0
const ProgramActivity_SessionActivity_case case_ProgramActivity_Type = 1
const ProgramActivity_TransactionActivity_case case_ProgramActivity_Type = 2

func (x *ProgramActivity) WhichType() case_ProgramActivity_Type {
	if x == nil {
		return ProgramActivity_Type_not_set_case
	}
	switch x.xxx_hidden_Type.(type) {
	case *programActivity_SessionActivity:
		return ProgramActivity_SessionActivity_case
	case *programActivity_TransactionActivity:
		return ProgramActivity_TransactionActivity_case
	default:
		return ProgramActivity_Type_not_set_case
	}
}

type ProgramActivity_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_Type:
	SessionActivity     *SessionActivity
	TransactionActivity *TransactionActivity
	// -- end of xxx_hidden_Type
}

func (b0 ProgramActivity_builder) Build() *ProgramActivity {
	m0 := &ProgramActivity{}
	b, x := &b0, m0
	_, _ = b, x
	if b.SessionActivity != nil {
		x.xxx_hidden_Type = &programActivity_SessionActivity{b.SessionActivity}
	}
	if b.TransactionActivity != nil {
		x.xxx_hidden_Type = &programActivity_TransactionActivity{b.TransactionActivity}
	}
	return m0
}

type case_ProgramActivity_Type protoreflect.FieldNumber

func (x case_ProgramActivity_Type) String() string {
	md := file_program_activity_proto_msgTypes[0].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isProgramActivity_Type interface {
	isProgramActivity_Type()
}

type programActivity_SessionActivity struct {
	SessionActivity *SessionActivity `protobuf:"bytes,1,opt,name=session_activity,json=sessionActivity,oneof"`
}

type programActivity_TransactionActivity struct {
	TransactionActivity *TransactionActivity `protobuf:"bytes,2,opt,name=transaction_activity,json=transactionActivity,oneof"`
}

func (*programActivity_SessionActivity) isProgramActivity_Type() {}

func (*programActivity_TransactionActivity) isProgramActivity_Type() {}

type SessionActivity struct {
	state                           protoimpl.MessageState  `protogen:"opaque.v1"`
	xxx_hidden_SessionSetCommands   *[]*SessionSetCommand   `protobuf:"bytes,1,rep,name=session_set_commands,json=sessionSetCommands"`
	xxx_hidden_SessionResetCommands *[]*SessionResetCommand `protobuf:"bytes,2,rep,name=session_reset_commands,json=sessionResetCommands"`
	unknownFields                   protoimpl.UnknownFields
	sizeCache                       protoimpl.SizeCache
}

func (x *SessionActivity) Reset() {
	*x = SessionActivity{}
	mi := &file_program_activity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionActivity) ProtoMessage() {}

func (x *SessionActivity) ProtoReflect() protoreflect.Message {
	mi := &file_program_activity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SessionActivity) GetSessionSetCommands() []*SessionSetCommand {
	if x != nil {
		if x.xxx_hidden_SessionSetCommands != nil {
			return *x.xxx_hidden_SessionSetCommands
		}
	}
	return nil
}

func (x *SessionActivity) GetSessionResetCommands() []*SessionResetCommand {
	if x != nil {
		if x.xxx_hidden_SessionResetCommands != nil {
			return *x.xxx_hidden_SessionResetCommands
		}
	}
	return nil
}

func (x *SessionActivity) SetSessionSetCommands(v []*SessionSetCommand) {
	x.xxx_hidden_SessionSetCommands = &v
}

func (x *SessionActivity) SetSessionResetCommands(v []*SessionResetCommand) {
	x.xxx_hidden_SessionResetCommands = &v
}

type SessionActivity_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	SessionSetCommands   []*SessionSetCommand
	SessionResetCommands []*SessionResetCommand
}

func (b0 SessionActivity_builder) Build() *SessionActivity {
	m0 := &SessionActivity{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_SessionSetCommands = &b.SessionSetCommands
	x.xxx_hidden_SessionResetCommands = &b.SessionResetCommands
	return m0
}

type TransactionActivity struct {
	state                              protoimpl.MessageState   `protogen:"opaque.v1"`
	xxx_hidden_StartTransactionCommand *StartTransactionCommand `protobuf:"bytes,1,opt,name=startTransactionCommand"`
	xxx_hidden_ProcedureSpecification  *ProcedureSpecification  `protobuf:"bytes,2,opt,name=procedureSpecification"`
	xxx_hidden_EndTransactionCommand   *EndTransactionCommand   `protobuf:"bytes,3,opt,name=endTransactionCommand"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *TransactionActivity) Reset() {
	*x = TransactionActivity{}
	mi := &file_program_activity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionActivity) ProtoMessage() {}

func (x *TransactionActivity) ProtoReflect() protoreflect.Message {
	mi := &file_program_activity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TransactionActivity) GetStartTransactionCommand() *StartTransactionCommand {
	if x != nil {
		return x.xxx_hidden_StartTransactionCommand
	}
	return nil
}

func (x *TransactionActivity) GetProcedureSpecification() *ProcedureSpecification {
	if x != nil {
		return x.xxx_hidden_ProcedureSpecification
	}
	return nil
}

func (x *TransactionActivity) GetEndTransactionCommand() *EndTransactionCommand {
	if x != nil {
		return x.xxx_hidden_EndTransactionCommand
	}
	return nil
}

func (x *TransactionActivity) SetStartTransactionCommand(v *StartTransactionCommand) {
	x.xxx_hidden_StartTransactionCommand = v
}

func (x *TransactionActivity) SetProcedureSpecification(v *ProcedureSpecification) {
	x.xxx_hidden_ProcedureSpecification = v
}

func (x *TransactionActivity) SetEndTransactionCommand(v *EndTransactionCommand) {
	x.xxx_hidden_EndTransactionCommand = v
}

func (x *TransactionActivity) HasStartTransactionCommand() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_StartTransactionCommand != nil
}

func (x *TransactionActivity) HasProcedureSpecification() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ProcedureSpecification != nil
}

func (x *TransactionActivity) HasEndTransactionCommand() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_EndTransactionCommand != nil
}

func (x *TransactionActivity) ClearStartTransactionCommand() {
	x.xxx_hidden_StartTransactionCommand = nil
}

func (x *TransactionActivity) ClearProcedureSpecification() {
	x.xxx_hidden_ProcedureSpecification = nil
}

func (x *TransactionActivity) ClearEndTransactionCommand() {
	x.xxx_hidden_EndTransactionCommand = nil
}

type TransactionActivity_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	StartTransactionCommand *StartTransactionCommand
	ProcedureSpecification  *ProcedureSpecification
	EndTransactionCommand   *EndTransactionCommand
}

func (b0 TransactionActivity_builder) Build() *TransactionActivity {
	m0 := &TransactionActivity{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_StartTransactionCommand = b.StartTransactionCommand
	x.xxx_hidden_ProcedureSpecification = b.ProcedureSpecification
	x.xxx_hidden_EndTransactionCommand = b.EndTransactionCommand
	return m0
}

var File_program_activity_proto protoreflect.FileDescriptor

var file_program_activity_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x73, 0x74, 0x1a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x14, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x14, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x12, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12,
	0x4e, 0x0a, 0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x14, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22,
	0xa2, 0x02, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x17, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x01, 0x52, 0x17, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x53, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x15, 0x65,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x73, 0x74,
	0x2e, 0x45, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x01, 0x52, 0x15, 0x65,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x42, 0x10, 0x5a, 0x04, 0x2f, 0x61, 0x73, 0x74, 0x92, 0x03, 0x07, 0xd2,
	0x3e, 0x02, 0x10, 0x03, 0x08, 0x02, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x70, 0xe8, 0x07,
})

var file_program_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_program_activity_proto_goTypes = []any{
	(*ProgramActivity)(nil),         // 0: ast.ProgramActivity
	(*SessionActivity)(nil),         // 1: ast.SessionActivity
	(*TransactionActivity)(nil),     // 2: ast.TransactionActivity
	(*SessionSetCommand)(nil),       // 3: ast.SessionSetCommand
	(*SessionResetCommand)(nil),     // 4: ast.SessionResetCommand
	(*StartTransactionCommand)(nil), // 5: ast.StartTransactionCommand
	(*ProcedureSpecification)(nil),  // 6: ast.ProcedureSpecification
	(*EndTransactionCommand)(nil),   // 7: ast.EndTransactionCommand
}
var file_program_activity_proto_depIdxs = []int32{
	1, // 0: ast.ProgramActivity.session_activity:type_name -> ast.SessionActivity
	2, // 1: ast.ProgramActivity.transaction_activity:type_name -> ast.TransactionActivity
	3, // 2: ast.SessionActivity.session_set_commands:type_name -> ast.SessionSetCommand
	4, // 3: ast.SessionActivity.session_reset_commands:type_name -> ast.SessionResetCommand
	5, // 4: ast.TransactionActivity.startTransactionCommand:type_name -> ast.StartTransactionCommand
	6, // 5: ast.TransactionActivity.procedureSpecification:type_name -> ast.ProcedureSpecification
	7, // 6: ast.TransactionActivity.endTransactionCommand:type_name -> ast.EndTransactionCommand
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_program_activity_proto_init() }
func file_program_activity_proto_init() {
	if File_program_activity_proto != nil {
		return
	}
	file_session_commands_proto_init()
	file_transaction_commands_proto_init()
	file_procedure_specification_proto_init()
	file_program_activity_proto_msgTypes[0].OneofWrappers = []any{
		(*programActivity_SessionActivity)(nil),
		(*programActivity_TransactionActivity)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_program_activity_proto_rawDesc), len(file_program_activity_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_program_activity_proto_goTypes,
		DependencyIndexes: file_program_activity_proto_depIdxs,
		MessageInfos:      file_program_activity_proto_msgTypes,
	}.Build()
	File_program_activity_proto = out.File
	file_program_activity_proto_goTypes = nil
	file_program_activity_proto_depIdxs = nil
}
