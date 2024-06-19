// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: yesorno/v1/yesorno.proto

package yesornov1

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

type AskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *AskRequest) Reset() {
	*x = AskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yesorno_v1_yesorno_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskRequest) ProtoMessage() {}

func (x *AskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yesorno_v1_yesorno_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskRequest.ProtoReflect.Descriptor instead.
func (*AskRequest) Descriptor() ([]byte, []int) {
	return file_yesorno_v1_yesorno_proto_rawDescGZIP(), []int{0}
}

func (x *AskRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

type AskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *AskResponse) Reset() {
	*x = AskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yesorno_v1_yesorno_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskResponse) ProtoMessage() {}

func (x *AskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yesorno_v1_yesorno_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskResponse.ProtoReflect.Descriptor instead.
func (*AskResponse) Descriptor() ([]byte, []int) {
	return file_yesorno_v1_yesorno_proto_rawDescGZIP(), []int{1}
}

func (x *AskResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_yesorno_v1_yesorno_proto protoreflect.FileDescriptor

var file_yesorno_v1_yesorno_proto_rawDesc = []byte{
	0x0a, 0x18, 0x79, 0x65, 0x73, 0x6f, 0x72, 0x6e, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x79, 0x65, 0x73,
	0x6f, 0x72, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x79, 0x65, 0x73, 0x6f,
	0x72, 0x6e, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x25, 0x0a, 0x0b, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0x4a, 0x0a, 0x0e, 0x59, 0x65, 0x73, 0x4f, 0x72,
	0x4e, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x03, 0x41, 0x73, 0x6b,
	0x12, 0x16, 0x2e, 0x79, 0x65, 0x73, 0x6f, 0x72, 0x6e, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x79, 0x65, 0x73, 0x6f, 0x72,
	0x6e, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x6f, 0x6d, 0x69, 0x6e, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x79, 0x65, 0x73, 0x6f, 0x72, 0x6e, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x79, 0x65, 0x73,
	0x6f, 0x72, 0x6e, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yesorno_v1_yesorno_proto_rawDescOnce sync.Once
	file_yesorno_v1_yesorno_proto_rawDescData = file_yesorno_v1_yesorno_proto_rawDesc
)

func file_yesorno_v1_yesorno_proto_rawDescGZIP() []byte {
	file_yesorno_v1_yesorno_proto_rawDescOnce.Do(func() {
		file_yesorno_v1_yesorno_proto_rawDescData = protoimpl.X.CompressGZIP(file_yesorno_v1_yesorno_proto_rawDescData)
	})
	return file_yesorno_v1_yesorno_proto_rawDescData
}

var file_yesorno_v1_yesorno_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yesorno_v1_yesorno_proto_goTypes = []any{
	(*AskRequest)(nil),  // 0: yesorno.v1.AskRequest
	(*AskResponse)(nil), // 1: yesorno.v1.AskResponse
}
var file_yesorno_v1_yesorno_proto_depIdxs = []int32{
	0, // 0: yesorno.v1.YesOrNoService.Ask:input_type -> yesorno.v1.AskRequest
	1, // 1: yesorno.v1.YesOrNoService.Ask:output_type -> yesorno.v1.AskResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yesorno_v1_yesorno_proto_init() }
func file_yesorno_v1_yesorno_proto_init() {
	if File_yesorno_v1_yesorno_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yesorno_v1_yesorno_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AskRequest); i {
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
		file_yesorno_v1_yesorno_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AskResponse); i {
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
			RawDescriptor: file_yesorno_v1_yesorno_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yesorno_v1_yesorno_proto_goTypes,
		DependencyIndexes: file_yesorno_v1_yesorno_proto_depIdxs,
		MessageInfos:      file_yesorno_v1_yesorno_proto_msgTypes,
	}.Build()
	File_yesorno_v1_yesorno_proto = out.File
	file_yesorno_v1_yesorno_proto_rawDesc = nil
	file_yesorno_v1_yesorno_proto_goTypes = nil
	file_yesorno_v1_yesorno_proto_depIdxs = nil
}
