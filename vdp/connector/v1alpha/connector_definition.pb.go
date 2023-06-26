// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: vdp/connector/v1alpha/connector_definition.proto

package connectorv1alpha

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// View enumerates the definition views
type View int32

const (
	// View: UNSPECIFIED
	View_VIEW_UNSPECIFIED View = 0
	// View: BASIC
	View_VIEW_BASIC View = 1
	// View: FULL
	View_VIEW_FULL View = 2
)

// Enum value maps for View.
var (
	View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "VIEW_BASIC",
		2: "VIEW_FULL",
	}
	View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"VIEW_BASIC":       1,
		"VIEW_FULL":        2,
	}
)

func (x View) Enum() *View {
	p := new(View)
	*p = x
	return p
}

func (x View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (View) Descriptor() protoreflect.EnumDescriptor {
	return file_vdp_connector_v1alpha_connector_definition_proto_enumTypes[0].Descriptor()
}

func (View) Type() protoreflect.EnumType {
	return &file_vdp_connector_v1alpha_connector_definition_proto_enumTypes[0]
}

func (x View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use View.Descriptor instead.
func (View) EnumDescriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{0}
}

// ConnectorType enumerates connector types
type ConnectorType int32

const (
	// ConnectorType: UNSPECIFIED
	ConnectorType_CONNECTOR_TYPE_UNSPECIFIED ConnectorType = 0
	// ConnectorType: SOURCE
	ConnectorType_CONNECTOR_TYPE_SOURCE ConnectorType = 1
	// ConnectorType: DESTINATION
	ConnectorType_CONNECTOR_TYPE_DESTINATION ConnectorType = 2
	// ConnectorType: Model
	ConnectorType_CONNECTOR_TYPE_MODEL ConnectorType = 3
	// ConnectorType: Blockchain
	ConnectorType_CONNECTOR_TYPE_BLOCKCHAIN ConnectorType = 4
)

// Enum value maps for ConnectorType.
var (
	ConnectorType_name = map[int32]string{
		0: "CONNECTOR_TYPE_UNSPECIFIED",
		1: "CONNECTOR_TYPE_SOURCE",
		2: "CONNECTOR_TYPE_DESTINATION",
		3: "CONNECTOR_TYPE_MODEL",
		4: "CONNECTOR_TYPE_BLOCKCHAIN",
	}
	ConnectorType_value = map[string]int32{
		"CONNECTOR_TYPE_UNSPECIFIED": 0,
		"CONNECTOR_TYPE_SOURCE":      1,
		"CONNECTOR_TYPE_DESTINATION": 2,
		"CONNECTOR_TYPE_MODEL":       3,
		"CONNECTOR_TYPE_BLOCKCHAIN":  4,
	}
)

func (x ConnectorType) Enum() *ConnectorType {
	p := new(ConnectorType)
	*p = x
	return p
}

func (x ConnectorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectorType) Descriptor() protoreflect.EnumDescriptor {
	return file_vdp_connector_v1alpha_connector_definition_proto_enumTypes[1].Descriptor()
}

func (ConnectorType) Type() protoreflect.EnumType {
	return &file_vdp_connector_v1alpha_connector_definition_proto_enumTypes[1]
}

func (x ConnectorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectorType.Descriptor instead.
func (ConnectorType) EnumDescriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{1}
}

// ConnectorDefinition represents the connector definition data model
type ConnectorDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ConnectorDefinition resource name. It must have the format of
	// "connector-definitions/*"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ConnectorDefinition UUID
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// ConnectorDefinition resource ID (the last segment of the
	// resource name) used to construct the resource name. This conforms to
	// RFC-1034, which restricts to letters, numbers, and hyphen, with the first
	// character a letter, the last a letter or a number, and a 63 character
	// maximum.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// ConnectorDefinition title
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// ConnectorDefinition documentation URL
	DocumentationUrl string `protobuf:"bytes,5,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	// ConnectorDefinition icon
	Icon string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	// ConnectorDefinition spec
	Spec *Spec `protobuf:"bytes,7,opt,name=spec,proto3" json:"spec,omitempty"`
	// Connector Type
	ConnectorType ConnectorType `protobuf:"varint,8,opt,name=connector_type,json=connectorType,proto3,enum=vdp.connector.v1alpha.ConnectorType" json:"connector_type,omitempty"`
	// ConnectorDefinition tombstone, i.e., if not set or false, the
	// configuration is active, or otherwise, if true, this configuration is
	// permanently off
	Tombstone bool `protobuf:"varint,9,opt,name=tombstone,proto3" json:"tombstone,omitempty"`
	// ConnectorDefinition public flag, i.e., true if this connector
	// definition is available to all workspaces
	Public bool `protobuf:"varint,10,opt,name=public,proto3" json:"public,omitempty"`
	// ConnectorDefinition custom flag, i.e., whether this is a custom
	// connector definition
	Custom bool `protobuf:"varint,11,opt,name=custom,proto3" json:"custom,omitempty"`
	// ConnectorDefinition iconUrl
	IconUrl string `protobuf:"bytes,12,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	// ConnectorDefinition vendor name
	Vendor string `protobuf:"bytes,13,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// ConnectorDefinition vendorAttributes, i.e. the vendor-specific attributes
	VendorAttributes *structpb.Struct `protobuf:"bytes,14,opt,name=vendor_attributes,json=vendorAttributes,proto3" json:"vendor_attributes,omitempty"`
}

func (x *ConnectorDefinition) Reset() {
	*x = ConnectorDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectorDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorDefinition) ProtoMessage() {}

func (x *ConnectorDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorDefinition.ProtoReflect.Descriptor instead.
func (*ConnectorDefinition) Descriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectorDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectorDefinition) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ConnectorDefinition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConnectorDefinition) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ConnectorDefinition) GetDocumentationUrl() string {
	if x != nil {
		return x.DocumentationUrl
	}
	return ""
}

func (x *ConnectorDefinition) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ConnectorDefinition) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ConnectorDefinition) GetConnectorType() ConnectorType {
	if x != nil {
		return x.ConnectorType
	}
	return ConnectorType_CONNECTOR_TYPE_UNSPECIFIED
}

func (x *ConnectorDefinition) GetTombstone() bool {
	if x != nil {
		return x.Tombstone
	}
	return false
}

func (x *ConnectorDefinition) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *ConnectorDefinition) GetCustom() bool {
	if x != nil {
		return x.Custom
	}
	return false
}

func (x *ConnectorDefinition) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *ConnectorDefinition) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *ConnectorDefinition) GetVendorAttributes() *structpb.Struct {
	if x != nil {
		return x.VendorAttributes
	}
	return nil
}

// ListConnectorDefinitionsRequest represents a request to list
// ConnectorDefinitions
type ListConnectorDefinitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of ConnectorDefinitions to return. The
	// service may return fewer than this value. If unspecified, at most 10
	// ConnectorDefinitions will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize *int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// Page token
	PageToken *string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3,oneof" json:"page_token,omitempty"`
	// Definition view (default is DEFINITION_VIEW_BASIC)
	View *View `protobuf:"varint,3,opt,name=view,proto3,enum=vdp.connector.v1alpha.View,oneof" json:"view,omitempty"`
	// Filter expression to list connector definitions
	Filter *string `protobuf:"bytes,4,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
}

func (x *ListConnectorDefinitionsRequest) Reset() {
	*x = ListConnectorDefinitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConnectorDefinitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConnectorDefinitionsRequest) ProtoMessage() {}

func (x *ListConnectorDefinitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConnectorDefinitionsRequest.ProtoReflect.Descriptor instead.
func (*ListConnectorDefinitionsRequest) Descriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{1}
}

func (x *ListConnectorDefinitionsRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListConnectorDefinitionsRequest) GetPageToken() string {
	if x != nil && x.PageToken != nil {
		return *x.PageToken
	}
	return ""
}

func (x *ListConnectorDefinitionsRequest) GetView() View {
	if x != nil && x.View != nil {
		return *x.View
	}
	return View_VIEW_UNSPECIFIED
}

func (x *ListConnectorDefinitionsRequest) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

// ListConnectorDefinitionsResponse represents a response for a list
// of ConnectorDefinitions
type ListConnectorDefinitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of ConnectorDefinition resources
	ConnectorDefinitions []*ConnectorDefinition `protobuf:"bytes,1,rep,name=connector_definitions,json=connectorDefinitions,proto3" json:"connector_definitions,omitempty"`
	// Next page token
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total count of ConnectorDefinition resources
	TotalSize int64 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListConnectorDefinitionsResponse) Reset() {
	*x = ListConnectorDefinitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConnectorDefinitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConnectorDefinitionsResponse) ProtoMessage() {}

func (x *ListConnectorDefinitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConnectorDefinitionsResponse.ProtoReflect.Descriptor instead.
func (*ListConnectorDefinitionsResponse) Descriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{2}
}

func (x *ListConnectorDefinitionsResponse) GetConnectorDefinitions() []*ConnectorDefinition {
	if x != nil {
		return x.ConnectorDefinitions
	}
	return nil
}

func (x *ListConnectorDefinitionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListConnectorDefinitionsResponse) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

// GetConnectorDefinitionRequest represents a request to query a
// ConnectorDefinition resource
type GetConnectorDefinitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ConnectorDefinition resource name. It must have the format of
	// "connector-definitions/*"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ConnectorDefinition resource view (default is
	// DEFINITION_VIEW_BASIC)
	View *View `protobuf:"varint,2,opt,name=view,proto3,enum=vdp.connector.v1alpha.View,oneof" json:"view,omitempty"`
}

func (x *GetConnectorDefinitionRequest) Reset() {
	*x = GetConnectorDefinitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectorDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectorDefinitionRequest) ProtoMessage() {}

func (x *GetConnectorDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectorDefinitionRequest.ProtoReflect.Descriptor instead.
func (*GetConnectorDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{3}
}

func (x *GetConnectorDefinitionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetConnectorDefinitionRequest) GetView() View {
	if x != nil && x.View != nil {
		return *x.View
	}
	return View_VIEW_UNSPECIFIED
}

// GetConnectorDefinitionResponse represents a
// ConnectorDefinition response
type GetConnectorDefinitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A ConnectorDefinition resource
	ConnectorDefinition *ConnectorDefinition `protobuf:"bytes,1,opt,name=connector_definition,json=connectorDefinition,proto3" json:"connector_definition,omitempty"`
}

func (x *GetConnectorDefinitionResponse) Reset() {
	*x = GetConnectorDefinitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectorDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectorDefinitionResponse) ProtoMessage() {}

func (x *GetConnectorDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectorDefinitionResponse.ProtoReflect.Descriptor instead.
func (*GetConnectorDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP(), []int{4}
}

func (x *GetConnectorDefinitionResponse) GetConnectorDefinition() *ConnectorDefinition {
	if x != nil {
		return x.ConnectorDefinition
	}
	return nil
}

var File_vdp_connector_v1alpha_connector_definition_proto protoreflect.FileDescriptor

var file_vdp_connector_v1alpha_connector_definition_proto_rawDesc = []byte{
	0x0a, 0x30, 0x76, 0x64, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x76, 0x64, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a,
	0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12,
	0x17, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x70, 0x65, 0x63, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x50,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x09, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x1b, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x1e, 0x0a,
	0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x11, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x10, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x57, 0xea, 0x41, 0x54, 0x0a, 0x24, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x22, 0xff,
	0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x48, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x39, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x48, 0x02, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x48, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x76, 0x69, 0x65, 0x77, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0xca, 0x01, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x14, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xc7, 0x01,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x62, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4e, 0x92,
	0x41, 0x1f, 0xca, 0x3e, 0x1c, 0xfa, 0x02, 0x19, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x22, 0x7f, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x3b, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x42,
	0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x46,
	0x55, 0x4c, 0x4c, 0x10, 0x02, 0x2a, 0xa3, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x10, 0x04, 0x42, 0xf5, 0x01, 0x0a, 0x19,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x18, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x64, 0x70, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2,
	0x02, 0x03, 0x56, 0x43, 0x58, 0xaa, 0x02, 0x15, 0x56, 0x64, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x15,
	0x56, 0x64, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x21, 0x56, 0x64, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x56, 0x64, 0x70, 0x3a,
	0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vdp_connector_v1alpha_connector_definition_proto_rawDescOnce sync.Once
	file_vdp_connector_v1alpha_connector_definition_proto_rawDescData = file_vdp_connector_v1alpha_connector_definition_proto_rawDesc
)

func file_vdp_connector_v1alpha_connector_definition_proto_rawDescGZIP() []byte {
	file_vdp_connector_v1alpha_connector_definition_proto_rawDescOnce.Do(func() {
		file_vdp_connector_v1alpha_connector_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_vdp_connector_v1alpha_connector_definition_proto_rawDescData)
	})
	return file_vdp_connector_v1alpha_connector_definition_proto_rawDescData
}

var file_vdp_connector_v1alpha_connector_definition_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vdp_connector_v1alpha_connector_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vdp_connector_v1alpha_connector_definition_proto_goTypes = []interface{}{
	(View)(0),                                // 0: vdp.connector.v1alpha.View
	(ConnectorType)(0),                       // 1: vdp.connector.v1alpha.ConnectorType
	(*ConnectorDefinition)(nil),              // 2: vdp.connector.v1alpha.ConnectorDefinition
	(*ListConnectorDefinitionsRequest)(nil),  // 3: vdp.connector.v1alpha.ListConnectorDefinitionsRequest
	(*ListConnectorDefinitionsResponse)(nil), // 4: vdp.connector.v1alpha.ListConnectorDefinitionsResponse
	(*GetConnectorDefinitionRequest)(nil),    // 5: vdp.connector.v1alpha.GetConnectorDefinitionRequest
	(*GetConnectorDefinitionResponse)(nil),   // 6: vdp.connector.v1alpha.GetConnectorDefinitionResponse
	(*Spec)(nil),                             // 7: vdp.connector.v1alpha.Spec
	(*structpb.Struct)(nil),                  // 8: google.protobuf.Struct
}
var file_vdp_connector_v1alpha_connector_definition_proto_depIdxs = []int32{
	7, // 0: vdp.connector.v1alpha.ConnectorDefinition.spec:type_name -> vdp.connector.v1alpha.Spec
	1, // 1: vdp.connector.v1alpha.ConnectorDefinition.connector_type:type_name -> vdp.connector.v1alpha.ConnectorType
	8, // 2: vdp.connector.v1alpha.ConnectorDefinition.vendor_attributes:type_name -> google.protobuf.Struct
	0, // 3: vdp.connector.v1alpha.ListConnectorDefinitionsRequest.view:type_name -> vdp.connector.v1alpha.View
	2, // 4: vdp.connector.v1alpha.ListConnectorDefinitionsResponse.connector_definitions:type_name -> vdp.connector.v1alpha.ConnectorDefinition
	0, // 5: vdp.connector.v1alpha.GetConnectorDefinitionRequest.view:type_name -> vdp.connector.v1alpha.View
	2, // 6: vdp.connector.v1alpha.GetConnectorDefinitionResponse.connector_definition:type_name -> vdp.connector.v1alpha.ConnectorDefinition
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_vdp_connector_v1alpha_connector_definition_proto_init() }
func file_vdp_connector_v1alpha_connector_definition_proto_init() {
	if File_vdp_connector_v1alpha_connector_definition_proto != nil {
		return
	}
	file_vdp_connector_v1alpha_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectorDefinition); i {
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
		file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConnectorDefinitionsRequest); i {
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
		file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConnectorDefinitionsResponse); i {
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
		file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectorDefinitionRequest); i {
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
		file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectorDefinitionResponse); i {
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
	file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_vdp_connector_v1alpha_connector_definition_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_connector_v1alpha_connector_definition_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vdp_connector_v1alpha_connector_definition_proto_goTypes,
		DependencyIndexes: file_vdp_connector_v1alpha_connector_definition_proto_depIdxs,
		EnumInfos:         file_vdp_connector_v1alpha_connector_definition_proto_enumTypes,
		MessageInfos:      file_vdp_connector_v1alpha_connector_definition_proto_msgTypes,
	}.Build()
	File_vdp_connector_v1alpha_connector_definition_proto = out.File
	file_vdp_connector_v1alpha_connector_definition_proto_rawDesc = nil
	file_vdp_connector_v1alpha_connector_definition_proto_goTypes = nil
	file_vdp_connector_v1alpha_connector_definition_proto_depIdxs = nil
}
