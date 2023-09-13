// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.2
// source: protofiles/center.proto

package protofiles

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

type AvailableKeysReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys int32 `protobuf:"varint,1,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (x *AvailableKeysReq) Reset() {
	*x = AvailableKeysReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableKeysReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableKeysReq) ProtoMessage() {}

func (x *AvailableKeysReq) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableKeysReq.ProtoReflect.Descriptor instead.
func (*AvailableKeysReq) Descriptor() ([]byte, []int) {
	return file_protofiles_center_proto_rawDescGZIP(), []int{0}
}

func (x *AvailableKeysReq) GetKeys() int32 {
	if x != nil {
		return x.Keys
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protofiles_center_proto_rawDescGZIP(), []int{1}
}

type UsersNotAdmittedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users int32 `protobuf:"varint,1,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersNotAdmittedReq) Reset() {
	*x = UsersNotAdmittedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersNotAdmittedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersNotAdmittedReq) ProtoMessage() {}

func (x *UsersNotAdmittedReq) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersNotAdmittedReq.ProtoReflect.Descriptor instead.
func (*UsersNotAdmittedReq) Descriptor() ([]byte, []int) {
	return file_protofiles_center_proto_rawDescGZIP(), []int{2}
}

func (x *UsersNotAdmittedReq) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

type ContinueServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Continue bool `protobuf:"varint,1,opt,name=continue,proto3" json:"continue,omitempty"`
}

func (x *ContinueServiceReq) Reset() {
	*x = ContinueServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinueServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinueServiceReq) ProtoMessage() {}

func (x *ContinueServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinueServiceReq.ProtoReflect.Descriptor instead.
func (*ContinueServiceReq) Descriptor() ([]byte, []int) {
	return file_protofiles_center_proto_rawDescGZIP(), []int{3}
}

func (x *ContinueServiceReq) GetContinue() bool {
	if x != nil {
		return x.Continue
	}
	return false
}

type FinalNotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfUsersFailed int32 `protobuf:"varint,1,opt,name=NumberOfUsersFailed,proto3" json:"NumberOfUsersFailed,omitempty"`
}

func (x *FinalNotifyRequest) Reset() {
	*x = FinalNotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalNotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalNotifyRequest) ProtoMessage() {}

func (x *FinalNotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalNotifyRequest.ProtoReflect.Descriptor instead.
func (*FinalNotifyRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_center_proto_rawDescGZIP(), []int{4}
}

func (x *FinalNotifyRequest) GetNumberOfUsersFailed() int32 {
	if x != nil {
		return x.NumberOfUsersFailed
	}
	return 0
}

type FinalNotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FinalNotifyResponse) Reset() {
	*x = FinalNotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_center_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalNotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalNotifyResponse) ProtoMessage() {}

func (x *FinalNotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_center_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalNotifyResponse.ProtoReflect.Descriptor instead.
func (*FinalNotifyResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_center_proto_rawDescGZIP(), []int{5}
}

func (x *FinalNotifyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protofiles_center_proto protoreflect.FileDescriptor

var file_protofiles_center_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4e,
	0x6f, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x69, 0x6e, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x69, 0x6e, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x2f, 0x0a,
	0x13, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe9,
	0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x3b, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x0e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x12, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x4c, 0x0a, 0x16,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4e, 0x6f, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4e, 0x6f, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x66, 0x0a, 0x11, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x51, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x54, 0x61, 0x72, 0x65, 0x61, 0x2d, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_center_proto_rawDescOnce sync.Once
	file_protofiles_center_proto_rawDescData = file_protofiles_center_proto_rawDesc
)

func file_protofiles_center_proto_rawDescGZIP() []byte {
	file_protofiles_center_proto_rawDescOnce.Do(func() {
		file_protofiles_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_center_proto_rawDescData)
	})
	return file_protofiles_center_proto_rawDescData
}

var file_protofiles_center_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protofiles_center_proto_goTypes = []interface{}{
	(*AvailableKeysReq)(nil),    // 0: protofiles.AvailableKeysReq
	(*Empty)(nil),               // 1: protofiles.Empty
	(*UsersNotAdmittedReq)(nil), // 2: protofiles.UsersNotAdmittedReq
	(*ContinueServiceReq)(nil),  // 3: protofiles.ContinueServiceReq
	(*FinalNotifyRequest)(nil),  // 4: protofiles.FinalNotifyRequest
	(*FinalNotifyResponse)(nil), // 5: protofiles.FinalNotifyResponse
}
var file_protofiles_center_proto_depIdxs = []int32{
	0, // 0: protofiles.NotifyKeys.SendKeys:input_type -> protofiles.AvailableKeysReq
	3, // 1: protofiles.NotifyKeys.NotifyContinue:input_type -> protofiles.ContinueServiceReq
	2, // 2: protofiles.NotifyKeys.UsersNotAdmittedNotify:input_type -> protofiles.UsersNotAdmittedReq
	4, // 3: protofiles.FinalNotification.NotifyRegional:input_type -> protofiles.FinalNotifyRequest
	1, // 4: protofiles.NotifyKeys.SendKeys:output_type -> protofiles.Empty
	3, // 5: protofiles.NotifyKeys.NotifyContinue:output_type -> protofiles.ContinueServiceReq
	1, // 6: protofiles.NotifyKeys.UsersNotAdmittedNotify:output_type -> protofiles.Empty
	5, // 7: protofiles.FinalNotification.NotifyRegional:output_type -> protofiles.FinalNotifyResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofiles_center_proto_init() }
func file_protofiles_center_proto_init() {
	if File_protofiles_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_center_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableKeysReq); i {
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
		file_protofiles_center_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protofiles_center_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersNotAdmittedReq); i {
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
		file_protofiles_center_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinueServiceReq); i {
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
		file_protofiles_center_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalNotifyRequest); i {
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
		file_protofiles_center_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalNotifyResponse); i {
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
			RawDescriptor: file_protofiles_center_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protofiles_center_proto_goTypes,
		DependencyIndexes: file_protofiles_center_proto_depIdxs,
		MessageInfos:      file_protofiles_center_proto_msgTypes,
	}.Build()
	File_protofiles_center_proto = out.File
	file_protofiles_center_proto_rawDesc = nil
	file_protofiles_center_proto_goTypes = nil
	file_protofiles_center_proto_depIdxs = nil
}