// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: admin/v1/index.proto

package admin

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

type ActivityDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ActivityDetailRequest) Reset() {
	*x = ActivityDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityDetailRequest) ProtoMessage() {}

func (x *ActivityDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityDetailRequest.ProtoReflect.Descriptor instead.
func (*ActivityDetailRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_index_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityDetailRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ActivityDetailRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ActivityDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ActivityDetailResponse) Reset() {
	*x = ActivityDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityDetailResponse) ProtoMessage() {}

func (x *ActivityDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityDetailResponse.ProtoReflect.Descriptor instead.
func (*ActivityDetailResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_index_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityDetailResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ActivityDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_admin_v1_index_proto protoreflect.FileDescriptor

var file_admin_v1_index_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x45, 0x0a,
	0x15, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x50, 0x0a, 0x08,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_index_proto_rawDescOnce sync.Once
	file_admin_v1_index_proto_rawDescData = file_admin_v1_index_proto_rawDesc
)

func file_admin_v1_index_proto_rawDescGZIP() []byte {
	file_admin_v1_index_proto_rawDescOnce.Do(func() {
		file_admin_v1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_index_proto_rawDescData)
	})
	return file_admin_v1_index_proto_rawDescData
}

var file_admin_v1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_v1_index_proto_goTypes = []interface{}{
	(*ActivityDetailRequest)(nil),  // 0: admin.ActivityDetailRequest
	(*ActivityDetailResponse)(nil), // 1: admin.ActivityDetailResponse
}
var file_admin_v1_index_proto_depIdxs = []int32{
	0, // 0: admin.Activity.Get:input_type -> admin.ActivityDetailRequest
	1, // 1: admin.Activity.Get:output_type -> admin.ActivityDetailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_v1_index_proto_init() }
func file_admin_v1_index_proto_init() {
	if File_admin_v1_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityDetailRequest); i {
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
		file_admin_v1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityDetailResponse); i {
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
			RawDescriptor: file_admin_v1_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_index_proto_goTypes,
		DependencyIndexes: file_admin_v1_index_proto_depIdxs,
		MessageInfos:      file_admin_v1_index_proto_msgTypes,
	}.Build()
	File_admin_v1_index_proto = out.File
	file_admin_v1_index_proto_rawDesc = nil
	file_admin_v1_index_proto_goTypes = nil
	file_admin_v1_index_proto_depIdxs = nil
}
