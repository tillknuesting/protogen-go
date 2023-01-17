// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vdp/model/v1alpha/task_detection.proto

package modelv1alpha

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// DetectionObject represents a predicted object
type DetectionObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Detection object category
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// Detection object score
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Detection bounding box
	BoundingBox *BoundingBox `protobuf:"bytes,3,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
}

func (x *DetectionObject) Reset() {
	*x = DetectionObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_task_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectionObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectionObject) ProtoMessage() {}

func (x *DetectionObject) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_task_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectionObject.ProtoReflect.Descriptor instead.
func (*DetectionObject) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_task_detection_proto_rawDescGZIP(), []int{0}
}

func (x *DetectionObject) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *DetectionObject) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *DetectionObject) GetBoundingBox() *BoundingBox {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

// DetectionInput represents the input of detection task
type DetectionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input type
	//
	// Types that are assignable to Type:
	//	*DetectionInput_ImageUrl
	//	*DetectionInput_ImageBase64
	Type isDetectionInput_Type `protobuf_oneof:"type"`
}

func (x *DetectionInput) Reset() {
	*x = DetectionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_task_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectionInput) ProtoMessage() {}

func (x *DetectionInput) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_task_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectionInput.ProtoReflect.Descriptor instead.
func (*DetectionInput) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_task_detection_proto_rawDescGZIP(), []int{1}
}

func (m *DetectionInput) GetType() isDetectionInput_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *DetectionInput) GetImageUrl() string {
	if x, ok := x.GetType().(*DetectionInput_ImageUrl); ok {
		return x.ImageUrl
	}
	return ""
}

func (x *DetectionInput) GetImageBase64() string {
	if x, ok := x.GetType().(*DetectionInput_ImageBase64); ok {
		return x.ImageBase64
	}
	return ""
}

type isDetectionInput_Type interface {
	isDetectionInput_Type()
}

type DetectionInput_ImageUrl struct {
	// Image type URL
	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3,oneof"`
}

type DetectionInput_ImageBase64 struct {
	// Image type base64
	ImageBase64 string `protobuf:"bytes,2,opt,name=image_base64,json=imageBase64,proto3,oneof"`
}

func (*DetectionInput_ImageUrl) isDetectionInput_Type() {}

func (*DetectionInput_ImageBase64) isDetectionInput_Type() {}

// DetectionOutput represents the output of detection task
type DetectionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of detection objects
	Objects []*DetectionObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *DetectionOutput) Reset() {
	*x = DetectionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_task_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectionOutput) ProtoMessage() {}

func (x *DetectionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_task_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectionOutput.ProtoReflect.Descriptor instead.
func (*DetectionOutput) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_task_detection_proto_rawDescGZIP(), []int{2}
}

func (x *DetectionOutput) GetObjects() []*DetectionObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_vdp_model_v1alpha_task_detection_proto protoreflect.FileDescriptor

var file_vdp_model_v1alpha_task_detection_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x76, 0x64,
	0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a,
	0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x46, 0x0a, 0x0c,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42,
	0x6f, 0x78, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x42, 0x6f, 0x78, 0x22, 0x5c, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x54, 0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0xd3, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x64, 0x70, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x56, 0x4d, 0x58,
	0xaa, 0x02, 0x11, 0x56, 0x64, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x11, 0x56, 0x64, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1d, 0x56, 0x64, 0x70, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x56, 0x64, 0x70, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vdp_model_v1alpha_task_detection_proto_rawDescOnce sync.Once
	file_vdp_model_v1alpha_task_detection_proto_rawDescData = file_vdp_model_v1alpha_task_detection_proto_rawDesc
)

func file_vdp_model_v1alpha_task_detection_proto_rawDescGZIP() []byte {
	file_vdp_model_v1alpha_task_detection_proto_rawDescOnce.Do(func() {
		file_vdp_model_v1alpha_task_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_vdp_model_v1alpha_task_detection_proto_rawDescData)
	})
	return file_vdp_model_v1alpha_task_detection_proto_rawDescData
}

var file_vdp_model_v1alpha_task_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vdp_model_v1alpha_task_detection_proto_goTypes = []interface{}{
	(*DetectionObject)(nil), // 0: vdp.model.v1alpha.DetectionObject
	(*DetectionInput)(nil),  // 1: vdp.model.v1alpha.DetectionInput
	(*DetectionOutput)(nil), // 2: vdp.model.v1alpha.DetectionOutput
	(*BoundingBox)(nil),     // 3: vdp.model.v1alpha.BoundingBox
}
var file_vdp_model_v1alpha_task_detection_proto_depIdxs = []int32{
	3, // 0: vdp.model.v1alpha.DetectionObject.bounding_box:type_name -> vdp.model.v1alpha.BoundingBox
	0, // 1: vdp.model.v1alpha.DetectionOutput.objects:type_name -> vdp.model.v1alpha.DetectionObject
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vdp_model_v1alpha_task_detection_proto_init() }
func file_vdp_model_v1alpha_task_detection_proto_init() {
	if File_vdp_model_v1alpha_task_detection_proto != nil {
		return
	}
	file_vdp_model_v1alpha_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vdp_model_v1alpha_task_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectionObject); i {
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
		file_vdp_model_v1alpha_task_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectionInput); i {
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
		file_vdp_model_v1alpha_task_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectionOutput); i {
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
	file_vdp_model_v1alpha_task_detection_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DetectionInput_ImageUrl)(nil),
		(*DetectionInput_ImageBase64)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_model_v1alpha_task_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vdp_model_v1alpha_task_detection_proto_goTypes,
		DependencyIndexes: file_vdp_model_v1alpha_task_detection_proto_depIdxs,
		MessageInfos:      file_vdp_model_v1alpha_task_detection_proto_msgTypes,
	}.Build()
	File_vdp_model_v1alpha_task_detection_proto = out.File
	file_vdp_model_v1alpha_task_detection_proto_rawDesc = nil
	file_vdp_model_v1alpha_task_detection_proto_goTypes = nil
	file_vdp_model_v1alpha_task_detection_proto_depIdxs = nil
}
