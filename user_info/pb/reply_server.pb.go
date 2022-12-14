// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: reply_server.proto

package pb

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

type LoginServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoginServerRequest) Reset() {
	*x = LoginServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginServerRequest) ProtoMessage() {}

func (x *LoginServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reply_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginServerRequest.ProtoReflect.Descriptor instead.
func (*LoginServerRequest) Descriptor() ([]byte, []int) {
	return file_reply_server_proto_rawDescGZIP(), []int{0}
}

func (x *LoginServerRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseToLoginServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResponseToLoginServer) Reset() {
	*x = ResponseToLoginServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseToLoginServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseToLoginServer) ProtoMessage() {}

func (x *ResponseToLoginServer) ProtoReflect() protoreflect.Message {
	mi := &file_reply_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseToLoginServer.ProtoReflect.Descriptor instead.
func (*ResponseToLoginServer) Descriptor() ([]byte, []int) {
	return file_reply_server_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseToLoginServer) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type AllIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllId []*LoginServerRequest `protobuf:"bytes,1,rep,name=all_id,json=allId,proto3" json:"all_id,omitempty"`
}

func (x *AllIDRequest) Reset() {
	*x = AllIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllIDRequest) ProtoMessage() {}

func (x *AllIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reply_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllIDRequest.ProtoReflect.Descriptor instead.
func (*AllIDRequest) Descriptor() ([]byte, []int) {
	return file_reply_server_proto_rawDescGZIP(), []int{2}
}

func (x *AllIDRequest) GetAllId() []*LoginServerRequest {
	if x != nil {
		return x.AllId
	}
	return nil
}

var File_reply_server_proto protoreflect.FileDescriptor

var file_reply_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x24, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x44, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x69, 0x6e,
	0x65, 0x6d, 0x69, 0x6e, 0x6b, 0x79, 0x61, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x61,
	0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reply_server_proto_rawDescOnce sync.Once
	file_reply_server_proto_rawDescData = file_reply_server_proto_rawDesc
)

func file_reply_server_proto_rawDescGZIP() []byte {
	file_reply_server_proto_rawDescOnce.Do(func() {
		file_reply_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_reply_server_proto_rawDescData)
	})
	return file_reply_server_proto_rawDescData
}

var file_reply_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reply_server_proto_goTypes = []interface{}{
	(*LoginServerRequest)(nil),    // 0: user_info.LoginServerRequest
	(*ResponseToLoginServer)(nil), // 1: user_info.ResponseToLoginServer
	(*AllIDRequest)(nil),          // 2: user_info.AllIDRequest
}
var file_reply_server_proto_depIdxs = []int32{
	0, // 0: user_info.AllIDRequest.all_id:type_name -> user_info.LoginServerRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reply_server_proto_init() }
func file_reply_server_proto_init() {
	if File_reply_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reply_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginServerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reply_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseToLoginServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reply_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reply_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reply_server_proto_goTypes,
		DependencyIndexes: file_reply_server_proto_depIdxs,
		MessageInfos:      file_reply_server_proto_msgTypes,
	}.Build()
	File_reply_server_proto = out.File
	file_reply_server_proto_rawDesc = nil
	file_reply_server_proto_goTypes = nil
	file_reply_server_proto_depIdxs = nil
}
