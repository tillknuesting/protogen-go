// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/model/v1alpha/task_text_to_image.proto

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

// TextToImageInput represents the input of a text-to-image task.
type TextToImageInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt text.
	Prompt string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Prompt image, only for multimodal input.
	//
	// Types that are assignable to Type:
	//
	//	*TextToImageInput_PromptImageUrl
	//	*TextToImageInput_PromptImageBase64
	Type isTextToImageInput_Type `protobuf_oneof:"type"`
	// Steps, defaults to 5.
	Steps *int32 `protobuf:"varint,4,opt,name=steps,proto3,oneof" json:"steps,omitempty"`
	// Guidance scale, defaults to 7.5.
	CfgScale *float32 `protobuf:"fixed32,5,opt,name=cfg_scale,json=cfgScale,proto3,oneof" json:"cfg_scale,omitempty"`
	// Seed, defaults to 0.
	Seed *int32 `protobuf:"varint,6,opt,name=seed,proto3,oneof" json:"seed,omitempty"`
	// Number of generated samples, default is 1.
	Samples *int32 `protobuf:"varint,7,opt,name=samples,proto3,oneof" json:"samples,omitempty"`
	// Extra parameters.
	ExtraParams *structpb.Struct `protobuf:"bytes,9,opt,name=extra_params,json=extraParams,proto3" json:"extra_params,omitempty"`
}

func (x *TextToImageInput) Reset() {
	*x = TextToImageInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_text_to_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextToImageInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextToImageInput) ProtoMessage() {}

func (x *TextToImageInput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_text_to_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextToImageInput.ProtoReflect.Descriptor instead.
func (*TextToImageInput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_text_to_image_proto_rawDescGZIP(), []int{0}
}

func (x *TextToImageInput) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (m *TextToImageInput) GetType() isTextToImageInput_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *TextToImageInput) GetPromptImageUrl() string {
	if x, ok := x.GetType().(*TextToImageInput_PromptImageUrl); ok {
		return x.PromptImageUrl
	}
	return ""
}

func (x *TextToImageInput) GetPromptImageBase64() string {
	if x, ok := x.GetType().(*TextToImageInput_PromptImageBase64); ok {
		return x.PromptImageBase64
	}
	return ""
}

func (x *TextToImageInput) GetSteps() int32 {
	if x != nil && x.Steps != nil {
		return *x.Steps
	}
	return 0
}

func (x *TextToImageInput) GetCfgScale() float32 {
	if x != nil && x.CfgScale != nil {
		return *x.CfgScale
	}
	return 0
}

func (x *TextToImageInput) GetSeed() int32 {
	if x != nil && x.Seed != nil {
		return *x.Seed
	}
	return 0
}

func (x *TextToImageInput) GetSamples() int32 {
	if x != nil && x.Samples != nil {
		return *x.Samples
	}
	return 0
}

func (x *TextToImageInput) GetExtraParams() *structpb.Struct {
	if x != nil {
		return x.ExtraParams
	}
	return nil
}

type isTextToImageInput_Type interface {
	isTextToImageInput_Type()
}

type TextToImageInput_PromptImageUrl struct {
	// Image URL.
	PromptImageUrl string `protobuf:"bytes,2,opt,name=prompt_image_url,json=promptImageUrl,proto3,oneof"`
}

type TextToImageInput_PromptImageBase64 struct {
	// Base64-encoded image.
	PromptImageBase64 string `protobuf:"bytes,3,opt,name=prompt_image_base64,json=promptImageBase64,proto3,oneof"`
}

func (*TextToImageInput_PromptImageUrl) isTextToImageInput_Type() {}

func (*TextToImageInput_PromptImageBase64) isTextToImageInput_Type() {}

// TextToImageOutput contains the result of a text-to-image task.
type TextToImageOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of generated images, encoded in base64.
	Images []string `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *TextToImageOutput) Reset() {
	*x = TextToImageOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_text_to_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextToImageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextToImageOutput) ProtoMessage() {}

func (x *TextToImageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_text_to_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextToImageOutput.ProtoReflect.Descriptor instead.
func (*TextToImageOutput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_text_to_image_proto_rawDescGZIP(), []int{1}
}

func (x *TextToImageOutput) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

var File_model_model_v1alpha_task_text_to_image_proto protoreflect.FileDescriptor

var file_model_model_v1alpha_task_text_to_image_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x74, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x10, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x11, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x36, 0x34, 0x12, 0x1e, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x01, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x09, 0x63, 0x66, 0x67, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x02, 0x52, 0x08, 0x63, 0x66,
	0x67, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x73, 0x65, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x03, 0x52, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x04, 0x52,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0c, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x06, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x63, 0x66, 0x67, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x73, 0x65, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x22, 0x30, 0x0a, 0x11, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x42, 0xe1, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x14, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x4d, 0x4d,
	0x58, 0xaa, 0x02, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1f,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_model_v1alpha_task_text_to_image_proto_rawDescOnce sync.Once
	file_model_model_v1alpha_task_text_to_image_proto_rawDescData = file_model_model_v1alpha_task_text_to_image_proto_rawDesc
)

func file_model_model_v1alpha_task_text_to_image_proto_rawDescGZIP() []byte {
	file_model_model_v1alpha_task_text_to_image_proto_rawDescOnce.Do(func() {
		file_model_model_v1alpha_task_text_to_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_model_v1alpha_task_text_to_image_proto_rawDescData)
	})
	return file_model_model_v1alpha_task_text_to_image_proto_rawDescData
}

var file_model_model_v1alpha_task_text_to_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_model_v1alpha_task_text_to_image_proto_goTypes = []interface{}{
	(*TextToImageInput)(nil),  // 0: model.model.v1alpha.TextToImageInput
	(*TextToImageOutput)(nil), // 1: model.model.v1alpha.TextToImageOutput
	(*structpb.Struct)(nil),   // 2: google.protobuf.Struct
}
var file_model_model_v1alpha_task_text_to_image_proto_depIdxs = []int32{
	2, // 0: model.model.v1alpha.TextToImageInput.extra_params:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_model_v1alpha_task_text_to_image_proto_init() }
func file_model_model_v1alpha_task_text_to_image_proto_init() {
	if File_model_model_v1alpha_task_text_to_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_model_v1alpha_task_text_to_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextToImageInput); i {
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
		file_model_model_v1alpha_task_text_to_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextToImageOutput); i {
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
	file_model_model_v1alpha_task_text_to_image_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TextToImageInput_PromptImageUrl)(nil),
		(*TextToImageInput_PromptImageBase64)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_model_v1alpha_task_text_to_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_v1alpha_task_text_to_image_proto_goTypes,
		DependencyIndexes: file_model_model_v1alpha_task_text_to_image_proto_depIdxs,
		MessageInfos:      file_model_model_v1alpha_task_text_to_image_proto_msgTypes,
	}.Build()
	File_model_model_v1alpha_task_text_to_image_proto = out.File
	file_model_model_v1alpha_task_text_to_image_proto_rawDesc = nil
	file_model_model_v1alpha_task_text_to_image_proto_goTypes = nil
	file_model_model_v1alpha_task_text_to_image_proto_depIdxs = nil
}
