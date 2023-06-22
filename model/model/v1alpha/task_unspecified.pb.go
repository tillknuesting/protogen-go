// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/model/v1alpha/task_unspecified.proto

package modelv1alpha

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// UnspecifiedInput represents the input of unspecified task
type UnspecifiedInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of unspecified task inputs
	RawInputs []*structpb.Struct `protobuf:"bytes,1,rep,name=raw_inputs,json=rawInputs,proto3" json:"raw_inputs,omitempty"`
}

func (x *UnspecifiedInput) Reset() {
	*x = UnspecifiedInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_unspecified_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnspecifiedInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnspecifiedInput) ProtoMessage() {}

func (x *UnspecifiedInput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_unspecified_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnspecifiedInput.ProtoReflect.Descriptor instead.
func (*UnspecifiedInput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_unspecified_proto_rawDescGZIP(), []int{0}
}

func (x *UnspecifiedInput) GetRawInputs() []*structpb.Struct {
	if x != nil {
		return x.RawInputs
	}
	return nil
}

// UnspecifiedOutput represents the output of unspecified task
type UnspecifiedOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of unspecified task outputs
	RawOutputs []*structpb.Struct `protobuf:"bytes,1,rep,name=raw_outputs,json=rawOutputs,proto3" json:"raw_outputs,omitempty"`
}

func (x *UnspecifiedOutput) Reset() {
	*x = UnspecifiedOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_unspecified_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnspecifiedOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnspecifiedOutput) ProtoMessage() {}

func (x *UnspecifiedOutput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_unspecified_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnspecifiedOutput.ProtoReflect.Descriptor instead.
func (*UnspecifiedOutput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_unspecified_proto_rawDescGZIP(), []int{1}
}

func (x *UnspecifiedOutput) GetRawOutputs() []*structpb.Struct {
	if x != nil {
		return x.RawOutputs
	}
	return nil
}

var File_model_model_v1alpha_task_unspecified_proto protoreflect.FileDescriptor

var file_model_model_v1alpha_task_unspecified_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x09, 0x72, 0x61, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x11,
	0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x42, 0xe1, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x14, 0x54, 0x61,
	0x73, 0x6b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x4d, 0x4d, 0x58, 0xaa, 0x02,
	0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1f, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_model_v1alpha_task_unspecified_proto_rawDescOnce sync.Once
	file_model_model_v1alpha_task_unspecified_proto_rawDescData = file_model_model_v1alpha_task_unspecified_proto_rawDesc
)

func file_model_model_v1alpha_task_unspecified_proto_rawDescGZIP() []byte {
	file_model_model_v1alpha_task_unspecified_proto_rawDescOnce.Do(func() {
		file_model_model_v1alpha_task_unspecified_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_model_v1alpha_task_unspecified_proto_rawDescData)
	})
	return file_model_model_v1alpha_task_unspecified_proto_rawDescData
}

var file_model_model_v1alpha_task_unspecified_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_model_v1alpha_task_unspecified_proto_goTypes = []interface{}{
	(*UnspecifiedInput)(nil),  // 0: model.model.v1alpha.UnspecifiedInput
	(*UnspecifiedOutput)(nil), // 1: model.model.v1alpha.UnspecifiedOutput
	(*structpb.Struct)(nil),   // 2: google.protobuf.Struct
}
var file_model_model_v1alpha_task_unspecified_proto_depIdxs = []int32{
	2, // 0: model.model.v1alpha.UnspecifiedInput.raw_inputs:type_name -> google.protobuf.Struct
	2, // 1: model.model.v1alpha.UnspecifiedOutput.raw_outputs:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_model_v1alpha_task_unspecified_proto_init() }
func file_model_model_v1alpha_task_unspecified_proto_init() {
	if File_model_model_v1alpha_task_unspecified_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_model_v1alpha_task_unspecified_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnspecifiedInput); i {
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
		file_model_model_v1alpha_task_unspecified_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnspecifiedOutput); i {
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
			RawDescriptor: file_model_model_v1alpha_task_unspecified_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_v1alpha_task_unspecified_proto_goTypes,
		DependencyIndexes: file_model_model_v1alpha_task_unspecified_proto_depIdxs,
		MessageInfos:      file_model_model_v1alpha_task_unspecified_proto_msgTypes,
	}.Build()
	File_model_model_v1alpha_task_unspecified_proto = out.File
	file_model_model_v1alpha_task_unspecified_proto_rawDesc = nil
	file_model_model_v1alpha_task_unspecified_proto_goTypes = nil
	file_model_model_v1alpha_task_unspecified_proto_depIdxs = nil
}
