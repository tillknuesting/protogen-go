// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: instill/connector/v1alpha/spec.proto

package connectorv1alpha

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

// SupportedDestinationSyncMode enumerates destination sync mode
type Spec_SupportedDestinationSyncModes int32

const (
	// SupportedDestinationSyncModes: DESTINATION_SYNC_MODE_UNSPECIFIED
	Spec_SUPPORTED_DESTINATION_SYNC_MODES_UNSPECIFIED Spec_SupportedDestinationSyncModes = 0
	// SupportedDestinationSyncModes: DESTINATION_SYNC_MODE_APPEND
	Spec_SUPPORTED_DESTINATION_SYNC_MODES_APPEND Spec_SupportedDestinationSyncModes = 1
	// SupportedDestinationSyncModes: DESTINATION_SYNC_MODE_OVERWRITE
	Spec_SUPPORTED_DESTINATION_SYNC_MODES_OVERWRITE Spec_SupportedDestinationSyncModes = 2
	// SupportedDestinationSyncModes: DESTINATION_SYNC_MODE_APPEND_DEDUP
	Spec_SUPPORTED_DESTINATION_SYNC_MODES_APPEND_DEDUP Spec_SupportedDestinationSyncModes = 3
)

// Enum value maps for Spec_SupportedDestinationSyncModes.
var (
	Spec_SupportedDestinationSyncModes_name = map[int32]string{
		0: "SUPPORTED_DESTINATION_SYNC_MODES_UNSPECIFIED",
		1: "SUPPORTED_DESTINATION_SYNC_MODES_APPEND",
		2: "SUPPORTED_DESTINATION_SYNC_MODES_OVERWRITE",
		3: "SUPPORTED_DESTINATION_SYNC_MODES_APPEND_DEDUP",
	}
	Spec_SupportedDestinationSyncModes_value = map[string]int32{
		"SUPPORTED_DESTINATION_SYNC_MODES_UNSPECIFIED":  0,
		"SUPPORTED_DESTINATION_SYNC_MODES_APPEND":       1,
		"SUPPORTED_DESTINATION_SYNC_MODES_OVERWRITE":    2,
		"SUPPORTED_DESTINATION_SYNC_MODES_APPEND_DEDUP": 3,
	}
)

func (x Spec_SupportedDestinationSyncModes) Enum() *Spec_SupportedDestinationSyncModes {
	p := new(Spec_SupportedDestinationSyncModes)
	*p = x
	return p
}

func (x Spec_SupportedDestinationSyncModes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Spec_SupportedDestinationSyncModes) Descriptor() protoreflect.EnumDescriptor {
	return file_instill_connector_v1alpha_spec_proto_enumTypes[0].Descriptor()
}

func (Spec_SupportedDestinationSyncModes) Type() protoreflect.EnumType {
	return &file_instill_connector_v1alpha_spec_proto_enumTypes[0]
}

func (x Spec_SupportedDestinationSyncModes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Spec_SupportedDestinationSyncModes.Descriptor instead.
func (Spec_SupportedDestinationSyncModes) EnumDescriptor() ([]byte, []int) {
	return file_instill_connector_v1alpha_spec_proto_rawDescGZIP(), []int{0, 0}
}

////////////////////////////////////
// Spec represents a spec data model
type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Spec documentation URL
	DocumentationUrl string `protobuf:"bytes,4,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	// Spec connection specification
	ConnectionSpecification *structpb.Struct `protobuf:"bytes,5,opt,name=connection_specification,json=connectionSpecification,proto3" json:"connection_specification,omitempty"`
	// Spec supports incremental flag, i.e., if the connector supports incremental
	// mode or not
	SupportsIncremental bool `protobuf:"varint,6,opt,name=supports_incremental,json=supportsIncremental,proto3" json:"supports_incremental,omitempty"`
	// Spec supports normalization flag, i.e., if the connector supports
	// normalization or not
	SupportsNormalization bool `protobuf:"varint,7,opt,name=supports_normalization,json=supportsNormalization,proto3" json:"supports_normalization,omitempty"`
	// Spec supports dbt flag, i.e., if the connector supports DBT or not
	SupportsDbt bool `protobuf:"varint,8,opt,name=supports_dbt,json=supportsDbt,proto3" json:"supports_dbt,omitempty"`
	// Spec destination sync mode, i.e., a list of destination sync modes
	// supported by the connector
	SupportedDestinationSyncModes []Spec_SupportedDestinationSyncModes `protobuf:"varint,9,rep,packed,name=supported_destination_sync_modes,json=supportedDestinationSyncModes,proto3,enum=instill.connector.v1alpha.Spec_SupportedDestinationSyncModes" json:"supported_destination_sync_modes,omitempty"`
	// Spec advanced auth, i.e., additional and optional specification object to
	// describe what an 'advanced' Auth flow would need to function
	AdvancedAuth *AdvancedAuth `protobuf:"bytes,10,opt,name=advanced_auth,json=advancedAuth,proto3" json:"advanced_auth,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instill_connector_v1alpha_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_instill_connector_v1alpha_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_instill_connector_v1alpha_spec_proto_rawDescGZIP(), []int{0}
}

func (x *Spec) GetDocumentationUrl() string {
	if x != nil {
		return x.DocumentationUrl
	}
	return ""
}

func (x *Spec) GetConnectionSpecification() *structpb.Struct {
	if x != nil {
		return x.ConnectionSpecification
	}
	return nil
}

func (x *Spec) GetSupportsIncremental() bool {
	if x != nil {
		return x.SupportsIncremental
	}
	return false
}

func (x *Spec) GetSupportsNormalization() bool {
	if x != nil {
		return x.SupportsNormalization
	}
	return false
}

func (x *Spec) GetSupportsDbt() bool {
	if x != nil {
		return x.SupportsDbt
	}
	return false
}

func (x *Spec) GetSupportedDestinationSyncModes() []Spec_SupportedDestinationSyncModes {
	if x != nil {
		return x.SupportedDestinationSyncModes
	}
	return nil
}

func (x *Spec) GetAdvancedAuth() *AdvancedAuth {
	if x != nil {
		return x.AdvancedAuth
	}
	return nil
}

var File_instill_connector_v1alpha_spec_proto protoreflect.FileDescriptor

var file_instill_connector_v1alpha_spec_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x05, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x30, 0x0a, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x10, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x57, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x14, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x13, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x16, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x15, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x62, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x44, 0x62, 0x74, 0x12, 0x8b, 0x01, 0x0a, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x3d, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x73,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x4d,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0d, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x22, 0xe1, 0x01, 0x0a, 0x1d, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x2c, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x53, 0x5f,
	0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x2e, 0x0a, 0x2a, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x31, 0x0a, 0x2d, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x53, 0x5f, 0x41, 0x50, 0x50,
	0x45, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x44, 0x55, 0x50, 0x10, 0x03, 0x42, 0xfe, 0x01, 0x0a, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x09, 0x53,
	0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x49, 0x43, 0x58, 0xaa, 0x02,
	0x19, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x19, 0x49, 0x6e, 0x73,
	0x74, 0x69, 0x6c, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x25, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c,
	0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1b, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instill_connector_v1alpha_spec_proto_rawDescOnce sync.Once
	file_instill_connector_v1alpha_spec_proto_rawDescData = file_instill_connector_v1alpha_spec_proto_rawDesc
)

func file_instill_connector_v1alpha_spec_proto_rawDescGZIP() []byte {
	file_instill_connector_v1alpha_spec_proto_rawDescOnce.Do(func() {
		file_instill_connector_v1alpha_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_instill_connector_v1alpha_spec_proto_rawDescData)
	})
	return file_instill_connector_v1alpha_spec_proto_rawDescData
}

var file_instill_connector_v1alpha_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_instill_connector_v1alpha_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_instill_connector_v1alpha_spec_proto_goTypes = []interface{}{
	(Spec_SupportedDestinationSyncModes)(0), // 0: instill.connector.v1alpha.Spec.SupportedDestinationSyncModes
	(*Spec)(nil),                            // 1: instill.connector.v1alpha.Spec
	(*structpb.Struct)(nil),                 // 2: google.protobuf.Struct
	(*AdvancedAuth)(nil),                    // 3: instill.connector.v1alpha.AdvancedAuth
}
var file_instill_connector_v1alpha_spec_proto_depIdxs = []int32{
	2, // 0: instill.connector.v1alpha.Spec.connection_specification:type_name -> google.protobuf.Struct
	0, // 1: instill.connector.v1alpha.Spec.supported_destination_sync_modes:type_name -> instill.connector.v1alpha.Spec.SupportedDestinationSyncModes
	3, // 2: instill.connector.v1alpha.Spec.advanced_auth:type_name -> instill.connector.v1alpha.AdvancedAuth
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_instill_connector_v1alpha_spec_proto_init() }
func file_instill_connector_v1alpha_spec_proto_init() {
	if File_instill_connector_v1alpha_spec_proto != nil {
		return
	}
	file_instill_connector_v1alpha_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_instill_connector_v1alpha_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_instill_connector_v1alpha_spec_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_instill_connector_v1alpha_spec_proto_goTypes,
		DependencyIndexes: file_instill_connector_v1alpha_spec_proto_depIdxs,
		EnumInfos:         file_instill_connector_v1alpha_spec_proto_enumTypes,
		MessageInfos:      file_instill_connector_v1alpha_spec_proto_msgTypes,
	}.Build()
	File_instill_connector_v1alpha_spec_proto = out.File
	file_instill_connector_v1alpha_spec_proto_rawDesc = nil
	file_instill_connector_v1alpha_spec_proto_goTypes = nil
	file_instill_connector_v1alpha_spec_proto_depIdxs = nil
}
