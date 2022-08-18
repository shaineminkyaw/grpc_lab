// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: create_login.proto

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

type RequestID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestID) Reset() {
	*x = RequestID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestID) ProtoMessage() {}

func (x *RequestID) ProtoReflect() protoreflect.Message {
	mi := &file_create_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestID.ProtoReflect.Descriptor instead.
func (*RequestID) Descriptor() ([]byte, []int) {
	return file_create_login_proto_rawDescGZIP(), []int{0}
}

func (x *RequestID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResponseToken) Reset() {
	*x = ResponseToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseToken) ProtoMessage() {}

func (x *ResponseToken) ProtoReflect() protoreflect.Message {
	mi := &file_create_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseToken.ProtoReflect.Descriptor instead.
func (*ResponseToken) Descriptor() ([]byte, []int) {
	return file_create_login_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId []*RequestID `protobuf:"bytes,1,rep,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_create_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_create_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetRequestId() []*RequestID {
	if x != nil {
		return x.RequestId
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllToken []*ResponseToken `protobuf:"bytes,1,rep,name=all_token,json=allToken,proto3" json:"all_token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_create_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_create_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetAllToken() []*ResponseToken {
	if x != nil {
		return x.AllToken
	}
	return nil
}

var File_create_login_proto protoreflect.FileDescriptor

var file_create_login_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x1b, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x3f, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x42, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x6d, 0x69, 0x6e, 0x6b, 0x79, 0x61, 0x77,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x61, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_create_login_proto_rawDescOnce sync.Once
	file_create_login_proto_rawDescData = file_create_login_proto_rawDesc
)

func file_create_login_proto_rawDescGZIP() []byte {
	file_create_login_proto_rawDescOnce.Do(func() {
		file_create_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_create_login_proto_rawDescData)
	})
	return file_create_login_proto_rawDescData
}

var file_create_login_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_create_login_proto_goTypes = []interface{}{
	(*RequestID)(nil),     // 0: login.RequestID
	(*ResponseToken)(nil), // 1: login.ResponseToken
	(*LoginRequest)(nil),  // 2: login.LoginRequest
	(*LoginResponse)(nil), // 3: login.LoginResponse
}
var file_create_login_proto_depIdxs = []int32{
	0, // 0: login.LoginRequest.request_id:type_name -> login.RequestID
	1, // 1: login.LoginResponse.all_token:type_name -> login.ResponseToken
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_create_login_proto_init() }
func file_create_login_proto_init() {
	if File_create_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_create_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestID); i {
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
		file_create_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseToken); i {
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
		file_create_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_create_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_create_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_create_login_proto_goTypes,
		DependencyIndexes: file_create_login_proto_depIdxs,
		MessageInfos:      file_create_login_proto_msgTypes,
	}.Build()
	File_create_login_proto = out.File
	file_create_login_proto_rawDesc = nil
	file_create_login_proto_goTypes = nil
	file_create_login_proto_depIdxs = nil
}
