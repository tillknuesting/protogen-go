// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: artifact/artifact/v1alpha/artifact.proto

package artifactv1alpha

import (
	v1beta "github.com/instill-ai/protogen-go/common/healthcheck/v1beta"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessRequest) ProtoMessage() {}

func (x *LivenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[0]
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
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{0}
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
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessResponse) ProtoMessage() {}

func (x *LivenessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[1]
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
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{1}
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
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessRequest) ProtoMessage() {}

func (x *ReadinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[2]
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
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{2}
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
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessResponse) ProtoMessage() {}

func (x *ReadinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[3]
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
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{3}
}

func (x *ReadinessResponse) GetHealthCheckResponse() *v1beta.HealthCheckResponse {
	if x != nil {
		return x.HealthCheckResponse
	}
	return nil
}

// RepositoryTag contains information about the version of some content in a
// repository.
type RepositoryTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the tag, defined by its parent repository and ID.
	// - Format: `repositories/{repository.id}/tags/{tag.id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The tag identifier.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Unique identifier, computed from the manifest the tag refers to.
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Tag update time, i.e. timestamp of the last push.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *RepositoryTag) Reset() {
	*x = RepositoryTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryTag) ProtoMessage() {}

func (x *RepositoryTag) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryTag.ProtoReflect.Descriptor instead.
func (*RepositoryTag) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{4}
}

func (x *RepositoryTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoryTag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RepositoryTag) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *RepositoryTag) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// ListRepositoryTagsRequest represents a request to list the tags of a
// repository.
type ListRepositoryTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of tags to return. The default and cap values are 10
	// and 100, respectively.
	PageSize *int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// Page number.
	Page *int32 `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	// The repository holding the different versions of a given content.
	// - Format: `repositories/{repository.id}`.
	// - Example: `repository/flaming-wombat/llama-2-7b`.
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ListRepositoryTagsRequest) Reset() {
	*x = ListRepositoryTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepositoryTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepositoryTagsRequest) ProtoMessage() {}

func (x *ListRepositoryTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepositoryTagsRequest.ProtoReflect.Descriptor instead.
func (*ListRepositoryTagsRequest) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{5}
}

func (x *ListRepositoryTagsRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListRepositoryTagsRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListRepositoryTagsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// ListRepositoryTagsResponse contains a list of container image tags.
type ListRepositoryTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of repository tags.
	Tags []*RepositoryTag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	// Total number of tags.
	TotalSize int32 `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	// The requested page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The requested page offset.
	Page int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListRepositoryTagsResponse) Reset() {
	*x = ListRepositoryTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepositoryTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepositoryTagsResponse) ProtoMessage() {}

func (x *ListRepositoryTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepositoryTagsResponse.ProtoReflect.Descriptor instead.
func (*ListRepositoryTagsResponse) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{6}
}

func (x *ListRepositoryTagsResponse) GetTags() []*RepositoryTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ListRepositoryTagsResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *ListRepositoryTagsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRepositoryTagsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

// CreateRepositoryTagRequest represents a request to add a tag to a given
// repository.
type CreateRepositoryTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tag information.
	Tag *RepositoryTag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// The repository holding the different versions of a given content.
	// - Format: `repositories/{repository.id}`.
	// - Example: `repository/flaming-wombat/llama-2-7b`.
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *CreateRepositoryTagRequest) Reset() {
	*x = CreateRepositoryTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepositoryTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepositoryTagRequest) ProtoMessage() {}

func (x *CreateRepositoryTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepositoryTagRequest.ProtoReflect.Descriptor instead.
func (*CreateRepositoryTagRequest) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRepositoryTagRequest) GetTag() *RepositoryTag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *CreateRepositoryTagRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// CreateRepositoryTagResponse contains the created tag.
type CreateRepositoryTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The created tag.
	Tag *RepositoryTag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CreateRepositoryTagResponse) Reset() {
	*x = CreateRepositoryTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepositoryTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepositoryTagResponse) ProtoMessage() {}

func (x *CreateRepositoryTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_artifact_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepositoryTagResponse.ProtoReflect.Descriptor instead.
func (*CreateRepositoryTagResponse) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP(), []int{8}
}

func (x *CreateRepositoryTagResponse) GetTag() *RepositoryTag {
	if x != nil {
		return x.Tag
	}
	return nil
}

var File_artifact_artifact_v1alpha_artifact_proto protoreflect.FileDescriptor

var file_artifact_artifact_v1alpha_artifact_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
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
	0x65, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x61, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x23, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x3b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x23, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54,
	0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x42, 0x81, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x19,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x19, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x25, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_artifact_artifact_v1alpha_artifact_proto_rawDescOnce sync.Once
	file_artifact_artifact_v1alpha_artifact_proto_rawDescData = file_artifact_artifact_v1alpha_artifact_proto_rawDesc
)

func file_artifact_artifact_v1alpha_artifact_proto_rawDescGZIP() []byte {
	file_artifact_artifact_v1alpha_artifact_proto_rawDescOnce.Do(func() {
		file_artifact_artifact_v1alpha_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_artifact_v1alpha_artifact_proto_rawDescData)
	})
	return file_artifact_artifact_v1alpha_artifact_proto_rawDescData
}

var file_artifact_artifact_v1alpha_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_artifact_artifact_v1alpha_artifact_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),             // 0: artifact.artifact.v1alpha.LivenessRequest
	(*LivenessResponse)(nil),            // 1: artifact.artifact.v1alpha.LivenessResponse
	(*ReadinessRequest)(nil),            // 2: artifact.artifact.v1alpha.ReadinessRequest
	(*ReadinessResponse)(nil),           // 3: artifact.artifact.v1alpha.ReadinessResponse
	(*RepositoryTag)(nil),               // 4: artifact.artifact.v1alpha.RepositoryTag
	(*ListRepositoryTagsRequest)(nil),   // 5: artifact.artifact.v1alpha.ListRepositoryTagsRequest
	(*ListRepositoryTagsResponse)(nil),  // 6: artifact.artifact.v1alpha.ListRepositoryTagsResponse
	(*CreateRepositoryTagRequest)(nil),  // 7: artifact.artifact.v1alpha.CreateRepositoryTagRequest
	(*CreateRepositoryTagResponse)(nil), // 8: artifact.artifact.v1alpha.CreateRepositoryTagResponse
	(*v1beta.HealthCheckRequest)(nil),   // 9: common.healthcheck.v1beta.HealthCheckRequest
	(*v1beta.HealthCheckResponse)(nil),  // 10: common.healthcheck.v1beta.HealthCheckResponse
	(*timestamppb.Timestamp)(nil),       // 11: google.protobuf.Timestamp
}
var file_artifact_artifact_v1alpha_artifact_proto_depIdxs = []int32{
	9,  // 0: artifact.artifact.v1alpha.LivenessRequest.health_check_request:type_name -> common.healthcheck.v1beta.HealthCheckRequest
	10, // 1: artifact.artifact.v1alpha.LivenessResponse.health_check_response:type_name -> common.healthcheck.v1beta.HealthCheckResponse
	9,  // 2: artifact.artifact.v1alpha.ReadinessRequest.health_check_request:type_name -> common.healthcheck.v1beta.HealthCheckRequest
	10, // 3: artifact.artifact.v1alpha.ReadinessResponse.health_check_response:type_name -> common.healthcheck.v1beta.HealthCheckResponse
	11, // 4: artifact.artifact.v1alpha.RepositoryTag.update_time:type_name -> google.protobuf.Timestamp
	4,  // 5: artifact.artifact.v1alpha.ListRepositoryTagsResponse.tags:type_name -> artifact.artifact.v1alpha.RepositoryTag
	4,  // 6: artifact.artifact.v1alpha.CreateRepositoryTagRequest.tag:type_name -> artifact.artifact.v1alpha.RepositoryTag
	4,  // 7: artifact.artifact.v1alpha.CreateRepositoryTagResponse.tag:type_name -> artifact.artifact.v1alpha.RepositoryTag
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_artifact_artifact_v1alpha_artifact_proto_init() }
func file_artifact_artifact_v1alpha_artifact_proto_init() {
	if File_artifact_artifact_v1alpha_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryTag); i {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepositoryTagsRequest); i {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepositoryTagsResponse); i {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepositoryTagRequest); i {
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
		file_artifact_artifact_v1alpha_artifact_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepositoryTagResponse); i {
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
	file_artifact_artifact_v1alpha_artifact_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_artifact_artifact_v1alpha_artifact_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_artifact_artifact_v1alpha_artifact_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_artifact_artifact_v1alpha_artifact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_artifact_v1alpha_artifact_proto_goTypes,
		DependencyIndexes: file_artifact_artifact_v1alpha_artifact_proto_depIdxs,
		MessageInfos:      file_artifact_artifact_v1alpha_artifact_proto_msgTypes,
	}.Build()
	File_artifact_artifact_v1alpha_artifact_proto = out.File
	file_artifact_artifact_v1alpha_artifact_proto_rawDesc = nil
	file_artifact_artifact_v1alpha_artifact_proto_goTypes = nil
	file_artifact_artifact_v1alpha_artifact_proto_depIdxs = nil
}
