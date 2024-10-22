// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: program/program.proto

package program

import (
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

type SendMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SendMessage) Reset() {
	*x = SendMessage{}
	mi := &file_program_program_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessage) ProtoMessage() {}

func (x *SendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessage.ProtoReflect.Descriptor instead.
func (*SendMessage) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessage) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SendMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type MessagePublished struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Timestamp  int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MessagePublished) Reset() {
	*x = MessagePublished{}
	mi := &file_program_program_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessagePublished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePublished) ProtoMessage() {}

func (x *MessagePublished) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePublished.ProtoReflect.Descriptor instead.
func (*MessagePublished) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{1}
}

func (x *MessagePublished) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *MessagePublished) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ClientName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ClientName) Reset() {
	*x = ClientName{}
	mi := &file_program_program_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientName) ProtoMessage() {}

func (x *ClientName) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientName.ProtoReflect.Descriptor instead.
func (*ClientName) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{2}
}

func (x *ClientName) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type SendToStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageStream string `protobuf:"bytes,1,opt,name=messageStream,proto3" json:"messageStream,omitempty"`
	Timestamp     int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SendToStream) Reset() {
	*x = SendToStream{}
	mi := &file_program_program_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendToStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToStream) ProtoMessage() {}

func (x *SendToStream) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToStream.ProtoReflect.Descriptor instead.
func (*SendToStream) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{3}
}

func (x *SendToStream) GetMessageStream() string {
	if x != nil {
		return x.MessageStream
	}
	return ""
}

func (x *SendToStream) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_program_program_proto protoreflect.FileDescriptor

var file_program_program_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x22, 0x59, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x50, 0x0a, 0x10, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x20, 0x0a,
	0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x52, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x32, 0x84, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3b, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_program_program_proto_rawDescOnce sync.Once
	file_program_program_proto_rawDescData = file_program_program_proto_rawDesc
)

func file_program_program_proto_rawDescGZIP() []byte {
	file_program_program_proto_rawDescOnce.Do(func() {
		file_program_program_proto_rawDescData = protoimpl.X.CompressGZIP(file_program_program_proto_rawDescData)
	})
	return file_program_program_proto_rawDescData
}

var file_program_program_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_program_program_proto_goTypes = []any{
	(*SendMessage)(nil),      // 0: program.SendMessage
	(*MessagePublished)(nil), // 1: program.MessagePublished
	(*ClientName)(nil),       // 2: program.ClientName
	(*SendToStream)(nil),     // 3: program.SendToStream
}
var file_program_program_proto_depIdxs = []int32{
	2, // 0: program.ChittyChat.ConnectUser:input_type -> program.ClientName
	0, // 1: program.ChittyChat.SendChat:input_type -> program.SendMessage
	3, // 2: program.ChittyChat.ConnectUser:output_type -> program.SendToStream
	1, // 3: program.ChittyChat.SendChat:output_type -> program.MessagePublished
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_program_program_proto_init() }
func file_program_program_proto_init() {
	if File_program_program_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_program_program_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_program_program_proto_goTypes,
		DependencyIndexes: file_program_program_proto_depIdxs,
		MessageInfos:      file_program_program_proto_msgTypes,
	}.Build()
	File_program_program_proto = out.File
	file_program_program_proto_rawDesc = nil
	file_program_program_proto_goTypes = nil
	file_program_program_proto_depIdxs = nil
}
