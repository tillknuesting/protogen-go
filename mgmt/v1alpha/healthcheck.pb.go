// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: mgmt/v1alpha/healthcheck.proto

package mgmtv1alpha

import (
	v1alpha "github.com/instill-ai/protogen-go/healthcheck/v1alpha"
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

// LivenessRequest represents a request to check a service liveness status
type LivenessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HealthCheckRequest message
	HealthCheckRequest *v1alpha.HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof" json:"health_check_request,omitempty"`
}

func (x *LivenessRequest) Reset() {
	*x = LivenessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessRequest) ProtoMessage() {}

func (x *LivenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessRequest.ProtoReflect.Descriptor instead.
func (*LivenessRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha_healthcheck_proto_rawDescGZIP(), []int{0}
}

func (x *LivenessRequest) GetHealthCheckRequest() *v1alpha.HealthCheckRequest {
	if x != nil {
		return x.HealthCheckRequest
	}
	return nil
}

// LivenessResponse represents a response for a service liveness status
type LivenessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HealthCheckResponse message
	HealthCheckResponse *v1alpha.HealthCheckResponse `protobuf:"bytes,1,opt,name=health_check_response,json=healthCheckResponse,proto3" json:"health_check_response,omitempty"`
}

func (x *LivenessResponse) Reset() {
	*x = LivenessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessResponse) ProtoMessage() {}

func (x *LivenessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessResponse.ProtoReflect.Descriptor instead.
func (*LivenessResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha_healthcheck_proto_rawDescGZIP(), []int{1}
}

func (x *LivenessResponse) GetHealthCheckResponse() *v1alpha.HealthCheckResponse {
	if x != nil {
		return x.HealthCheckResponse
	}
	return nil
}

// ReadinessRequest represents a request to check a service readiness status
type ReadinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HealthCheckRequest message
	HealthCheckRequest *v1alpha.HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof" json:"health_check_request,omitempty"`
}

func (x *ReadinessRequest) Reset() {
	*x = ReadinessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessRequest) ProtoMessage() {}

func (x *ReadinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadinessRequest.ProtoReflect.Descriptor instead.
func (*ReadinessRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha_healthcheck_proto_rawDescGZIP(), []int{2}
}

func (x *ReadinessRequest) GetHealthCheckRequest() *v1alpha.HealthCheckRequest {
	if x != nil {
		return x.HealthCheckRequest
	}
	return nil
}

// ReadinessResponse represents a response for a service readiness status
type ReadinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HealthCheckResponse message
	HealthCheckResponse *v1alpha.HealthCheckResponse `protobuf:"bytes,1,opt,name=health_check_response,json=healthCheckResponse,proto3" json:"health_check_response,omitempty"`
}

func (x *ReadinessResponse) Reset() {
	*x = ReadinessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessResponse) ProtoMessage() {}

func (x *ReadinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha_healthcheck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadinessResponse.ProtoReflect.Descriptor instead.
func (*ReadinessResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha_healthcheck_proto_rawDescGZIP(), []int{3}
}

func (x *ReadinessResponse) GetHealthCheckResponse() *v1alpha.HealthCheckResponse {
	if x != nil {
		return x.HealthCheckResponse
	}
	return nil
}

var File_mgmt_v1alpha_healthcheck_proto protoreflect.FileDescriptor

var file_mgmt_v1alpha_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x14, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x12, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x17, 0x0a, 0x15, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x10, 0x4c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x52,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x63, 0x0a, 0x14, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x12, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x71, 0x0a,
	0x11, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x13, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xb1, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x67, 0x6d,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x4d, 0x67,
	0x6d, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x0c, 0x4d, 0x67, 0x6d,
	0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x18, 0x4d, 0x67, 0x6d, 0x74,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mgmt_v1alpha_healthcheck_proto_rawDescOnce sync.Once
	file_mgmt_v1alpha_healthcheck_proto_rawDescData = file_mgmt_v1alpha_healthcheck_proto_rawDesc
)

func file_mgmt_v1alpha_healthcheck_proto_rawDescGZIP() []byte {
	file_mgmt_v1alpha_healthcheck_proto_rawDescOnce.Do(func() {
		file_mgmt_v1alpha_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_mgmt_v1alpha_healthcheck_proto_rawDescData)
	})
	return file_mgmt_v1alpha_healthcheck_proto_rawDescData
}

var file_mgmt_v1alpha_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mgmt_v1alpha_healthcheck_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),             // 0: mgmt.v1alpha.LivenessRequest
	(*LivenessResponse)(nil),            // 1: mgmt.v1alpha.LivenessResponse
	(*ReadinessRequest)(nil),            // 2: mgmt.v1alpha.ReadinessRequest
	(*ReadinessResponse)(nil),           // 3: mgmt.v1alpha.ReadinessResponse
	(*v1alpha.HealthCheckRequest)(nil),  // 4: healthcheck.v1alpha.HealthCheckRequest
	(*v1alpha.HealthCheckResponse)(nil), // 5: healthcheck.v1alpha.HealthCheckResponse
}
var file_mgmt_v1alpha_healthcheck_proto_depIdxs = []int32{
	4, // 0: mgmt.v1alpha.LivenessRequest.health_check_request:type_name -> healthcheck.v1alpha.HealthCheckRequest
	5, // 1: mgmt.v1alpha.LivenessResponse.health_check_response:type_name -> healthcheck.v1alpha.HealthCheckResponse
	4, // 2: mgmt.v1alpha.ReadinessRequest.health_check_request:type_name -> healthcheck.v1alpha.HealthCheckRequest
	5, // 3: mgmt.v1alpha.ReadinessResponse.health_check_response:type_name -> healthcheck.v1alpha.HealthCheckResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mgmt_v1alpha_healthcheck_proto_init() }
func file_mgmt_v1alpha_healthcheck_proto_init() {
	if File_mgmt_v1alpha_healthcheck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mgmt_v1alpha_healthcheck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LivenessRequest); i {
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
		file_mgmt_v1alpha_healthcheck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LivenessResponse); i {
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
		file_mgmt_v1alpha_healthcheck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadinessRequest); i {
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
		file_mgmt_v1alpha_healthcheck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadinessResponse); i {
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
	file_mgmt_v1alpha_healthcheck_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_mgmt_v1alpha_healthcheck_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mgmt_v1alpha_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mgmt_v1alpha_healthcheck_proto_goTypes,
		DependencyIndexes: file_mgmt_v1alpha_healthcheck_proto_depIdxs,
		MessageInfos:      file_mgmt_v1alpha_healthcheck_proto_msgTypes,
	}.Build()
	File_mgmt_v1alpha_healthcheck_proto = out.File
	file_mgmt_v1alpha_healthcheck_proto_rawDesc = nil
	file_mgmt_v1alpha_healthcheck_proto_goTypes = nil
	file_mgmt_v1alpha_healthcheck_proto_depIdxs = nil
}
