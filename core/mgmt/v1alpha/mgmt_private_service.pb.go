// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: core/mgmt/v1alpha/mgmt_private_service.proto

package mgmtv1alpha

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_core_mgmt_v1alpha_mgmt_private_service_proto protoreflect.FileDescriptor

var file_core_mgmt_v1alpha_mgmt_private_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x1c, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf4, 0x07, 0x0a, 0x12, 0x4d, 0x67, 0x6d, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x28, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0xa7, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d,
	0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x12, 0xa3, 0x01,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0xad, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2e, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0xda,
	0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0xc7, 0x01, 0x0a, 0x17, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x31, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x31, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x42, 0xd7, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x17, 0x4d, 0x67, 0x6d, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x6f, 0x72, 0x65,
	0x2e, 0x4d, 0x67, 0x6d, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x11,
	0x43, 0x6f, 0x72, 0x65, 0x5c, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0xe2, 0x02, 0x1d, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_core_mgmt_v1alpha_mgmt_private_service_proto_goTypes = []interface{}{
	(*ListUsersAdminRequest)(nil),           // 0: core.mgmt.v1alpha.ListUsersAdminRequest
	(*GetUserAdminRequest)(nil),             // 1: core.mgmt.v1alpha.GetUserAdminRequest
	(*LookUpUserAdminRequest)(nil),          // 2: core.mgmt.v1alpha.LookUpUserAdminRequest
	(*ListOrganizationsAdminRequest)(nil),   // 3: core.mgmt.v1alpha.ListOrganizationsAdminRequest
	(*GetOrganizationAdminRequest)(nil),     // 4: core.mgmt.v1alpha.GetOrganizationAdminRequest
	(*LookUpOrganizationAdminRequest)(nil),  // 5: core.mgmt.v1alpha.LookUpOrganizationAdminRequest
	(*ListUsersAdminResponse)(nil),          // 6: core.mgmt.v1alpha.ListUsersAdminResponse
	(*GetUserAdminResponse)(nil),            // 7: core.mgmt.v1alpha.GetUserAdminResponse
	(*LookUpUserAdminResponse)(nil),         // 8: core.mgmt.v1alpha.LookUpUserAdminResponse
	(*ListOrganizationsAdminResponse)(nil),  // 9: core.mgmt.v1alpha.ListOrganizationsAdminResponse
	(*GetOrganizationAdminResponse)(nil),    // 10: core.mgmt.v1alpha.GetOrganizationAdminResponse
	(*LookUpOrganizationAdminResponse)(nil), // 11: core.mgmt.v1alpha.LookUpOrganizationAdminResponse
}
var file_core_mgmt_v1alpha_mgmt_private_service_proto_depIdxs = []int32{
	0,  // 0: core.mgmt.v1alpha.MgmtPrivateService.ListUsersAdmin:input_type -> core.mgmt.v1alpha.ListUsersAdminRequest
	1,  // 1: core.mgmt.v1alpha.MgmtPrivateService.GetUserAdmin:input_type -> core.mgmt.v1alpha.GetUserAdminRequest
	2,  // 2: core.mgmt.v1alpha.MgmtPrivateService.LookUpUserAdmin:input_type -> core.mgmt.v1alpha.LookUpUserAdminRequest
	3,  // 3: core.mgmt.v1alpha.MgmtPrivateService.ListOrganizationsAdmin:input_type -> core.mgmt.v1alpha.ListOrganizationsAdminRequest
	4,  // 4: core.mgmt.v1alpha.MgmtPrivateService.GetOrganizationAdmin:input_type -> core.mgmt.v1alpha.GetOrganizationAdminRequest
	5,  // 5: core.mgmt.v1alpha.MgmtPrivateService.LookUpOrganizationAdmin:input_type -> core.mgmt.v1alpha.LookUpOrganizationAdminRequest
	6,  // 6: core.mgmt.v1alpha.MgmtPrivateService.ListUsersAdmin:output_type -> core.mgmt.v1alpha.ListUsersAdminResponse
	7,  // 7: core.mgmt.v1alpha.MgmtPrivateService.GetUserAdmin:output_type -> core.mgmt.v1alpha.GetUserAdminResponse
	8,  // 8: core.mgmt.v1alpha.MgmtPrivateService.LookUpUserAdmin:output_type -> core.mgmt.v1alpha.LookUpUserAdminResponse
	9,  // 9: core.mgmt.v1alpha.MgmtPrivateService.ListOrganizationsAdmin:output_type -> core.mgmt.v1alpha.ListOrganizationsAdminResponse
	10, // 10: core.mgmt.v1alpha.MgmtPrivateService.GetOrganizationAdmin:output_type -> core.mgmt.v1alpha.GetOrganizationAdminResponse
	11, // 11: core.mgmt.v1alpha.MgmtPrivateService.LookUpOrganizationAdmin:output_type -> core.mgmt.v1alpha.LookUpOrganizationAdminResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_core_mgmt_v1alpha_mgmt_private_service_proto_init() }
func file_core_mgmt_v1alpha_mgmt_private_service_proto_init() {
	if File_core_mgmt_v1alpha_mgmt_private_service_proto != nil {
		return
	}
	file_core_mgmt_v1alpha_mgmt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_mgmt_v1alpha_mgmt_private_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_mgmt_v1alpha_mgmt_private_service_proto_goTypes,
		DependencyIndexes: file_core_mgmt_v1alpha_mgmt_private_service_proto_depIdxs,
	}.Build()
	File_core_mgmt_v1alpha_mgmt_private_service_proto = out.File
	file_core_mgmt_v1alpha_mgmt_private_service_proto_rawDesc = nil
	file_core_mgmt_v1alpha_mgmt_private_service_proto_goTypes = nil
	file_core_mgmt_v1alpha_mgmt_private_service_proto_depIdxs = nil
}
