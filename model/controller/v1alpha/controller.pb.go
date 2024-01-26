// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/controller/v1alpha/controller.proto

package controllerv1alpha

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1beta "github.com/instill-ai/protogen-go/common/healthcheck/v1beta"
	v1alpha "github.com/instill-ai/protogen-go/model/model/v1alpha"
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
	HealthCheckRequest *v1beta.HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof" json:"health_check_request,omitempty"`
}

func (x *LivenessRequest) Reset() {
	*x = LivenessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessRequest) ProtoMessage() {}

func (x *LivenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[0]
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
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{0}
}

func (x *LivenessRequest) GetHealthCheckRequest() *v1beta.HealthCheckRequest {
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
	HealthCheckResponse *v1beta.HealthCheckResponse `protobuf:"bytes,1,opt,name=health_check_response,json=healthCheckResponse,proto3" json:"health_check_response,omitempty"`
}

func (x *LivenessResponse) Reset() {
	*x = LivenessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessResponse) ProtoMessage() {}

func (x *LivenessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[1]
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
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{1}
}

func (x *LivenessResponse) GetHealthCheckResponse() *v1beta.HealthCheckResponse {
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
	HealthCheckRequest *v1beta.HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof" json:"health_check_request,omitempty"`
}

func (x *ReadinessRequest) Reset() {
	*x = ReadinessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessRequest) ProtoMessage() {}

func (x *ReadinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[2]
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
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{2}
}

func (x *ReadinessRequest) GetHealthCheckRequest() *v1beta.HealthCheckRequest {
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
	HealthCheckResponse *v1beta.HealthCheckResponse `protobuf:"bytes,1,opt,name=health_check_response,json=healthCheckResponse,proto3" json:"health_check_response,omitempty"`
}

func (x *ReadinessResponse) Reset() {
	*x = ReadinessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessResponse) ProtoMessage() {}

func (x *ReadinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[3]
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
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{3}
}

func (x *ReadinessResponse) GetHealthCheckResponse() *v1beta.HealthCheckResponse {
	if x != nil {
		return x.HealthCheckResponse
	}
	return nil
}

// Resource represents the current information of a resource
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permalink of a resource. For example:
	// "resources/{resource_uuid}/types/{type}"
	ResourcePermalink string `protobuf:"bytes,1,opt,name=resource_permalink,json=resourcePermalink,proto3" json:"resource_permalink,omitempty"`
	// Resource state
	//
	// Types that are assignable to State:
	//
	//	*Resource_ModelState
	//	*Resource_BackendState
	State isResource_State `protobuf_oneof:"state"`
	// Resource longrunning progress
	Progress *int32 `protobuf:"varint,6,opt,name=progress,proto3,oneof" json:"progress,omitempty"`
	// Resource longrunning workflow id
	WorkflowId *string `protobuf:"bytes,7,opt,name=workflow_id,json=workflowId,proto3,oneof" json:"workflow_id,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{4}
}

func (x *Resource) GetResourcePermalink() string {
	if x != nil {
		return x.ResourcePermalink
	}
	return ""
}

func (m *Resource) GetState() isResource_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (x *Resource) GetModelState() v1alpha.Model_State {
	if x, ok := x.GetState().(*Resource_ModelState); ok {
		return x.ModelState
	}
	return v1alpha.Model_State(0)
}

func (x *Resource) GetBackendState() v1beta.HealthCheckResponse_ServingStatus {
	if x, ok := x.GetState().(*Resource_BackendState); ok {
		return x.BackendState
	}
	return v1beta.HealthCheckResponse_ServingStatus(0)
}

func (x *Resource) GetProgress() int32 {
	if x != nil && x.Progress != nil {
		return *x.Progress
	}
	return 0
}

func (x *Resource) GetWorkflowId() string {
	if x != nil && x.WorkflowId != nil {
		return *x.WorkflowId
	}
	return ""
}

type isResource_State interface {
	isResource_State()
}

type Resource_ModelState struct {
	// Model state
	ModelState v1alpha.Model_State `protobuf:"varint,2,opt,name=model_state,json=modelState,proto3,enum=model.model.v1alpha.Model_State,oneof"`
}

type Resource_BackendState struct {
	// Backend service state
	BackendState v1beta.HealthCheckResponse_ServingStatus `protobuf:"varint,5,opt,name=backend_state,json=backendState,proto3,enum=common.healthcheck.v1beta.HealthCheckResponse_ServingStatus,oneof"`
}

func (*Resource_ModelState) isResource_State() {}

func (*Resource_BackendState) isResource_State() {}

// GetResourceRequest represents a request to query a resource's state
type GetResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permalink of a resource. For example:
	// "resources/{resource_uuid}/types/{type}"
	ResourcePermalink string `protobuf:"bytes,1,opt,name=resource_permalink,json=resourcePermalink,proto3" json:"resource_permalink,omitempty"`
}

func (x *GetResourceRequest) Reset() {
	*x = GetResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRequest) ProtoMessage() {}

func (x *GetResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{5}
}

func (x *GetResourceRequest) GetResourcePermalink() string {
	if x != nil {
		return x.ResourcePermalink
	}
	return ""
}

// GetResourceResponse represents a response to fetch a resource's state
type GetResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Retrieved resource state
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetResourceResponse) Reset() {
	*x = GetResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceResponse) ProtoMessage() {}

func (x *GetResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceResponse.ProtoReflect.Descriptor instead.
func (*GetResourceResponse) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{6}
}

func (x *GetResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// UpdateResourceRequest represents a request to update a resource's state
type UpdateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource state
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateResourceRequest) Reset() {
	*x = UpdateResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRequest) ProtoMessage() {}

func (x *UpdateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceRequest) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateResourceRequest) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// UpdateResourceResponse represents a response to update a resource's state
type UpdateResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Updated resource state
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateResourceResponse) Reset() {
	*x = UpdateResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceResponse) ProtoMessage() {}

func (x *UpdateResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceResponse.ProtoReflect.Descriptor instead.
func (*UpdateResourceResponse) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// DeleteResourceRequest represents a request to delete a resource's state
type DeleteResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permalink of a resource. For example:
	// "resources/{resource_uuid}/types/{type}"
	ResourcePermalink string `protobuf:"bytes,1,opt,name=resource_permalink,json=resourcePermalink,proto3" json:"resource_permalink,omitempty"`
}

func (x *DeleteResourceRequest) Reset() {
	*x = DeleteResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRequest) ProtoMessage() {}

func (x *DeleteResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceRequest) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteResourceRequest) GetResourcePermalink() string {
	if x != nil {
		return x.ResourcePermalink
	}
	return ""
}

// DeleteResourceResponse represents an empty response
type DeleteResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResourceResponse) Reset() {
	*x = DeleteResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_controller_v1alpha_controller_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceResponse) ProtoMessage() {}

func (x *DeleteResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_controller_v1alpha_controller_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteResourceResponse) Descriptor() ([]byte, []int) {
	return file_model_controller_v1alpha_controller_proto_rawDescGZIP(), []int{10}
}

var File_model_controller_v1alpha_controller_proto protoreflect.FileDescriptor

var file_model_controller_v1alpha_controller_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x95, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x14, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x12, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x17,
	0x0a, 0x15, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x10, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x13, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x96, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x14, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x12, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x17, 0x0a, 0x15, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x15, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x13, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa7, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0x43, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x63, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x0c,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x48, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x02, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x46, 0xea,
	0x41, 0x43, 0x0a, 0x19, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b,
	0x74, 0x79, 0x70, 0x65, 0x7d, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x74, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45,
	0x92, 0x41, 0x21, 0xca, 0x3e, 0x1e, 0xfa, 0x02, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x6e, 0x6b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x55, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x5c, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x58, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x74, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45, 0x92,
	0x41, 0x21, 0xca, 0x3e, 0x1e, 0xfa, 0x02, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x6e, 0x6b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xff, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x4d, 0x43, 0x58, 0xaa, 0x02, 0x18, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x18, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2,
	0x02, 0x24, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_controller_v1alpha_controller_proto_rawDescOnce sync.Once
	file_model_controller_v1alpha_controller_proto_rawDescData = file_model_controller_v1alpha_controller_proto_rawDesc
)

func file_model_controller_v1alpha_controller_proto_rawDescGZIP() []byte {
	file_model_controller_v1alpha_controller_proto_rawDescOnce.Do(func() {
		file_model_controller_v1alpha_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_controller_v1alpha_controller_proto_rawDescData)
	})
	return file_model_controller_v1alpha_controller_proto_rawDescData
}

var file_model_controller_v1alpha_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_model_controller_v1alpha_controller_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),                       // 0: model.controller.v1alpha.LivenessRequest
	(*LivenessResponse)(nil),                      // 1: model.controller.v1alpha.LivenessResponse
	(*ReadinessRequest)(nil),                      // 2: model.controller.v1alpha.ReadinessRequest
	(*ReadinessResponse)(nil),                     // 3: model.controller.v1alpha.ReadinessResponse
	(*Resource)(nil),                              // 4: model.controller.v1alpha.Resource
	(*GetResourceRequest)(nil),                    // 5: model.controller.v1alpha.GetResourceRequest
	(*GetResourceResponse)(nil),                   // 6: model.controller.v1alpha.GetResourceResponse
	(*UpdateResourceRequest)(nil),                 // 7: model.controller.v1alpha.UpdateResourceRequest
	(*UpdateResourceResponse)(nil),                // 8: model.controller.v1alpha.UpdateResourceResponse
	(*DeleteResourceRequest)(nil),                 // 9: model.controller.v1alpha.DeleteResourceRequest
	(*DeleteResourceResponse)(nil),                // 10: model.controller.v1alpha.DeleteResourceResponse
	(*v1beta.HealthCheckRequest)(nil),             // 11: common.healthcheck.v1beta.HealthCheckRequest
	(*v1beta.HealthCheckResponse)(nil),            // 12: common.healthcheck.v1beta.HealthCheckResponse
	(v1alpha.Model_State)(0),                      // 13: model.model.v1alpha.Model.State
	(v1beta.HealthCheckResponse_ServingStatus)(0), // 14: common.healthcheck.v1beta.HealthCheckResponse.ServingStatus
}
var file_model_controller_v1alpha_controller_proto_depIdxs = []int32{
	11, // 0: model.controller.v1alpha.LivenessRequest.health_check_request:type_name -> common.healthcheck.v1beta.HealthCheckRequest
	12, // 1: model.controller.v1alpha.LivenessResponse.health_check_response:type_name -> common.healthcheck.v1beta.HealthCheckResponse
	11, // 2: model.controller.v1alpha.ReadinessRequest.health_check_request:type_name -> common.healthcheck.v1beta.HealthCheckRequest
	12, // 3: model.controller.v1alpha.ReadinessResponse.health_check_response:type_name -> common.healthcheck.v1beta.HealthCheckResponse
	13, // 4: model.controller.v1alpha.Resource.model_state:type_name -> model.model.v1alpha.Model.State
	14, // 5: model.controller.v1alpha.Resource.backend_state:type_name -> common.healthcheck.v1beta.HealthCheckResponse.ServingStatus
	4,  // 6: model.controller.v1alpha.GetResourceResponse.resource:type_name -> model.controller.v1alpha.Resource
	4,  // 7: model.controller.v1alpha.UpdateResourceRequest.resource:type_name -> model.controller.v1alpha.Resource
	4,  // 8: model.controller.v1alpha.UpdateResourceResponse.resource:type_name -> model.controller.v1alpha.Resource
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_model_controller_v1alpha_controller_proto_init() }
func file_model_controller_v1alpha_controller_proto_init() {
	if File_model_controller_v1alpha_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_controller_v1alpha_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRequest); i {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceResponse); i {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceRequest); i {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceResponse); i {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceRequest); i {
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
		file_model_controller_v1alpha_controller_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceResponse); i {
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
	file_model_controller_v1alpha_controller_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_model_controller_v1alpha_controller_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_model_controller_v1alpha_controller_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Resource_ModelState)(nil),
		(*Resource_BackendState)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_controller_v1alpha_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_controller_v1alpha_controller_proto_goTypes,
		DependencyIndexes: file_model_controller_v1alpha_controller_proto_depIdxs,
		MessageInfos:      file_model_controller_v1alpha_controller_proto_msgTypes,
	}.Build()
	File_model_controller_v1alpha_controller_proto = out.File
	file_model_controller_v1alpha_controller_proto_rawDesc = nil
	file_model_controller_v1alpha_controller_proto_goTypes = nil
	file_model_controller_v1alpha_controller_proto_depIdxs = nil
}
