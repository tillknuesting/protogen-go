// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vdp/model/v1alpha/model_service.proto

package modelv1alpha

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

var File_vdp_model_v1alpha_model_service_proto protoreflect.FileDescriptor

var file_vdp_model_v1alpha_model_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xb8, 0x1e, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x89, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x5a, 0x17,
	0x12, 0x15, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x5f, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x74, 0x0a, 0x09,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x23, 0x2e, 0x76, 0x64, 0x70, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x5f, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x64, 0x70, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xa5, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x32, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x12, 0x23, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x6f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x23, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0xda, 0x41, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x0f, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0xae, 0x01,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x35, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0xda, 0x41,
	0x1b, 0x69, 0x64, 0x2c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x12, 0x7c,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x22, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x27, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x9f, 0x01, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x76,
	0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0xda, 0x41, 0x11,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x32, 0x1e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x85,
	0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0xda,
	0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x96, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x6f, 0x6b, 0x55,
	0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55,
	0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x12,
	0x9c, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x25, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e,
	0xda, 0x41, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x93,
	0x01, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x26, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x32, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a,
	0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x12, 0x9b, 0x01, 0x0a, 0x0e, 0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x6e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x75, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0xa5, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x35, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0xa0, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x2a, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x76, 0x64,
	0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xba, 0x01,
	0x0a, 0x13, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x12, 0xb3, 0x01, 0x0a, 0x13, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x2d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30,
	0x3a, 0x01, 0x2a, 0x22, 0x2b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x12, 0xbb, 0x01, 0x0a, 0x15, 0x55, 0x6e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55,
	0x6e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x76, 0x64,
	0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x55, 0x6e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xda,
	0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22,
	0x2d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x75, 0x6e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0xb3,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2e, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x61,
	0x64, 0x6d, 0x65, 0x7d, 0x12, 0xbe, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45,
	0xda, 0x41, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0xb7, 0x01, 0x0a, 0x24, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3e,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0c, 0xda, 0x41, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x66, 0x69, 0x6c, 0x65, 0x28, 0x01, 0x12,
	0xb2, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x42, 0xda, 0x41, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x12, 0xae, 0x01, 0x0a, 0x21, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3b, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c, 0xda, 0x41, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x66,
	0x69, 0x6c, 0x65, 0x28, 0x01, 0x1a, 0x13, 0xca, 0x41, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x42, 0xd2, 0x01, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x42, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x64, 0x70,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x56, 0x4d,
	0x58, 0xaa, 0x02, 0x11, 0x56, 0x64, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x11, 0x56, 0x64, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1d, 0x56, 0x64, 0x70, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x56, 0x64, 0x70, 0x3a,
	0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_vdp_model_v1alpha_model_service_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),                              // 0: vdp.model.v1alpha.LivenessRequest
	(*ReadinessRequest)(nil),                             // 1: vdp.model.v1alpha.ReadinessRequest
	(*ListModelDefinitionRequest)(nil),                   // 2: vdp.model.v1alpha.ListModelDefinitionRequest
	(*GetModelDefinitionRequest)(nil),                    // 3: vdp.model.v1alpha.GetModelDefinitionRequest
	(*ListModelRequest)(nil),                             // 4: vdp.model.v1alpha.ListModelRequest
	(*CreateModelRequest)(nil),                           // 5: vdp.model.v1alpha.CreateModelRequest
	(*CreateModelBinaryFileUploadRequest)(nil),           // 6: vdp.model.v1alpha.CreateModelBinaryFileUploadRequest
	(*GetModelRequest)(nil),                              // 7: vdp.model.v1alpha.GetModelRequest
	(*UpdateModelRequest)(nil),                           // 8: vdp.model.v1alpha.UpdateModelRequest
	(*DeleteModelRequest)(nil),                           // 9: vdp.model.v1alpha.DeleteModelRequest
	(*LookUpModelRequest)(nil),                           // 10: vdp.model.v1alpha.LookUpModelRequest
	(*RenameModelRequest)(nil),                           // 11: vdp.model.v1alpha.RenameModelRequest
	(*PublishModelRequest)(nil),                          // 12: vdp.model.v1alpha.PublishModelRequest
	(*UnpublishModelRequest)(nil),                        // 13: vdp.model.v1alpha.UnpublishModelRequest
	(*ListModelInstanceRequest)(nil),                     // 14: vdp.model.v1alpha.ListModelInstanceRequest
	(*GetModelInstanceRequest)(nil),                      // 15: vdp.model.v1alpha.GetModelInstanceRequest
	(*LookUpModelInstanceRequest)(nil),                   // 16: vdp.model.v1alpha.LookUpModelInstanceRequest
	(*DeployModelInstanceRequest)(nil),                   // 17: vdp.model.v1alpha.DeployModelInstanceRequest
	(*UndeployModelInstanceRequest)(nil),                 // 18: vdp.model.v1alpha.UndeployModelInstanceRequest
	(*GetModelInstanceCardRequest)(nil),                  // 19: vdp.model.v1alpha.GetModelInstanceCardRequest
	(*TriggerModelInstanceRequest)(nil),                  // 20: vdp.model.v1alpha.TriggerModelInstanceRequest
	(*TriggerModelInstanceBinaryFileUploadRequest)(nil),  // 21: vdp.model.v1alpha.TriggerModelInstanceBinaryFileUploadRequest
	(*TestModelInstanceRequest)(nil),                     // 22: vdp.model.v1alpha.TestModelInstanceRequest
	(*TestModelInstanceBinaryFileUploadRequest)(nil),     // 23: vdp.model.v1alpha.TestModelInstanceBinaryFileUploadRequest
	(*LivenessResponse)(nil),                             // 24: vdp.model.v1alpha.LivenessResponse
	(*ReadinessResponse)(nil),                            // 25: vdp.model.v1alpha.ReadinessResponse
	(*ListModelDefinitionResponse)(nil),                  // 26: vdp.model.v1alpha.ListModelDefinitionResponse
	(*GetModelDefinitionResponse)(nil),                   // 27: vdp.model.v1alpha.GetModelDefinitionResponse
	(*ListModelResponse)(nil),                            // 28: vdp.model.v1alpha.ListModelResponse
	(*CreateModelResponse)(nil),                          // 29: vdp.model.v1alpha.CreateModelResponse
	(*CreateModelBinaryFileUploadResponse)(nil),          // 30: vdp.model.v1alpha.CreateModelBinaryFileUploadResponse
	(*GetModelResponse)(nil),                             // 31: vdp.model.v1alpha.GetModelResponse
	(*UpdateModelResponse)(nil),                          // 32: vdp.model.v1alpha.UpdateModelResponse
	(*DeleteModelResponse)(nil),                          // 33: vdp.model.v1alpha.DeleteModelResponse
	(*LookUpModelResponse)(nil),                          // 34: vdp.model.v1alpha.LookUpModelResponse
	(*RenameModelResponse)(nil),                          // 35: vdp.model.v1alpha.RenameModelResponse
	(*PublishModelResponse)(nil),                         // 36: vdp.model.v1alpha.PublishModelResponse
	(*UnpublishModelResponse)(nil),                       // 37: vdp.model.v1alpha.UnpublishModelResponse
	(*ListModelInstanceResponse)(nil),                    // 38: vdp.model.v1alpha.ListModelInstanceResponse
	(*GetModelInstanceResponse)(nil),                     // 39: vdp.model.v1alpha.GetModelInstanceResponse
	(*LookUpModelInstanceResponse)(nil),                  // 40: vdp.model.v1alpha.LookUpModelInstanceResponse
	(*DeployModelInstanceResponse)(nil),                  // 41: vdp.model.v1alpha.DeployModelInstanceResponse
	(*UndeployModelInstanceResponse)(nil),                // 42: vdp.model.v1alpha.UndeployModelInstanceResponse
	(*GetModelInstanceCardResponse)(nil),                 // 43: vdp.model.v1alpha.GetModelInstanceCardResponse
	(*TriggerModelInstanceResponse)(nil),                 // 44: vdp.model.v1alpha.TriggerModelInstanceResponse
	(*TriggerModelInstanceBinaryFileUploadResponse)(nil), // 45: vdp.model.v1alpha.TriggerModelInstanceBinaryFileUploadResponse
	(*TestModelInstanceResponse)(nil),                    // 46: vdp.model.v1alpha.TestModelInstanceResponse
	(*TestModelInstanceBinaryFileUploadResponse)(nil),    // 47: vdp.model.v1alpha.TestModelInstanceBinaryFileUploadResponse
}
var file_vdp_model_v1alpha_model_service_proto_depIdxs = []int32{
	0,  // 0: vdp.model.v1alpha.ModelService.Liveness:input_type -> vdp.model.v1alpha.LivenessRequest
	1,  // 1: vdp.model.v1alpha.ModelService.Readiness:input_type -> vdp.model.v1alpha.ReadinessRequest
	2,  // 2: vdp.model.v1alpha.ModelService.ListModelDefinition:input_type -> vdp.model.v1alpha.ListModelDefinitionRequest
	3,  // 3: vdp.model.v1alpha.ModelService.GetModelDefinition:input_type -> vdp.model.v1alpha.GetModelDefinitionRequest
	4,  // 4: vdp.model.v1alpha.ModelService.ListModel:input_type -> vdp.model.v1alpha.ListModelRequest
	5,  // 5: vdp.model.v1alpha.ModelService.CreateModel:input_type -> vdp.model.v1alpha.CreateModelRequest
	6,  // 6: vdp.model.v1alpha.ModelService.CreateModelBinaryFileUpload:input_type -> vdp.model.v1alpha.CreateModelBinaryFileUploadRequest
	7,  // 7: vdp.model.v1alpha.ModelService.GetModel:input_type -> vdp.model.v1alpha.GetModelRequest
	8,  // 8: vdp.model.v1alpha.ModelService.UpdateModel:input_type -> vdp.model.v1alpha.UpdateModelRequest
	9,  // 9: vdp.model.v1alpha.ModelService.DeleteModel:input_type -> vdp.model.v1alpha.DeleteModelRequest
	10, // 10: vdp.model.v1alpha.ModelService.LookUpModel:input_type -> vdp.model.v1alpha.LookUpModelRequest
	11, // 11: vdp.model.v1alpha.ModelService.RenameModel:input_type -> vdp.model.v1alpha.RenameModelRequest
	12, // 12: vdp.model.v1alpha.ModelService.PublishModel:input_type -> vdp.model.v1alpha.PublishModelRequest
	13, // 13: vdp.model.v1alpha.ModelService.UnpublishModel:input_type -> vdp.model.v1alpha.UnpublishModelRequest
	14, // 14: vdp.model.v1alpha.ModelService.ListModelInstance:input_type -> vdp.model.v1alpha.ListModelInstanceRequest
	15, // 15: vdp.model.v1alpha.ModelService.GetModelInstance:input_type -> vdp.model.v1alpha.GetModelInstanceRequest
	16, // 16: vdp.model.v1alpha.ModelService.LookUpModelInstance:input_type -> vdp.model.v1alpha.LookUpModelInstanceRequest
	17, // 17: vdp.model.v1alpha.ModelService.DeployModelInstance:input_type -> vdp.model.v1alpha.DeployModelInstanceRequest
	18, // 18: vdp.model.v1alpha.ModelService.UndeployModelInstance:input_type -> vdp.model.v1alpha.UndeployModelInstanceRequest
	19, // 19: vdp.model.v1alpha.ModelService.GetModelInstanceCard:input_type -> vdp.model.v1alpha.GetModelInstanceCardRequest
	20, // 20: vdp.model.v1alpha.ModelService.TriggerModelInstance:input_type -> vdp.model.v1alpha.TriggerModelInstanceRequest
	21, // 21: vdp.model.v1alpha.ModelService.TriggerModelInstanceBinaryFileUpload:input_type -> vdp.model.v1alpha.TriggerModelInstanceBinaryFileUploadRequest
	22, // 22: vdp.model.v1alpha.ModelService.TestModelInstance:input_type -> vdp.model.v1alpha.TestModelInstanceRequest
	23, // 23: vdp.model.v1alpha.ModelService.TestModelInstanceBinaryFileUpload:input_type -> vdp.model.v1alpha.TestModelInstanceBinaryFileUploadRequest
	24, // 24: vdp.model.v1alpha.ModelService.Liveness:output_type -> vdp.model.v1alpha.LivenessResponse
	25, // 25: vdp.model.v1alpha.ModelService.Readiness:output_type -> vdp.model.v1alpha.ReadinessResponse
	26, // 26: vdp.model.v1alpha.ModelService.ListModelDefinition:output_type -> vdp.model.v1alpha.ListModelDefinitionResponse
	27, // 27: vdp.model.v1alpha.ModelService.GetModelDefinition:output_type -> vdp.model.v1alpha.GetModelDefinitionResponse
	28, // 28: vdp.model.v1alpha.ModelService.ListModel:output_type -> vdp.model.v1alpha.ListModelResponse
	29, // 29: vdp.model.v1alpha.ModelService.CreateModel:output_type -> vdp.model.v1alpha.CreateModelResponse
	30, // 30: vdp.model.v1alpha.ModelService.CreateModelBinaryFileUpload:output_type -> vdp.model.v1alpha.CreateModelBinaryFileUploadResponse
	31, // 31: vdp.model.v1alpha.ModelService.GetModel:output_type -> vdp.model.v1alpha.GetModelResponse
	32, // 32: vdp.model.v1alpha.ModelService.UpdateModel:output_type -> vdp.model.v1alpha.UpdateModelResponse
	33, // 33: vdp.model.v1alpha.ModelService.DeleteModel:output_type -> vdp.model.v1alpha.DeleteModelResponse
	34, // 34: vdp.model.v1alpha.ModelService.LookUpModel:output_type -> vdp.model.v1alpha.LookUpModelResponse
	35, // 35: vdp.model.v1alpha.ModelService.RenameModel:output_type -> vdp.model.v1alpha.RenameModelResponse
	36, // 36: vdp.model.v1alpha.ModelService.PublishModel:output_type -> vdp.model.v1alpha.PublishModelResponse
	37, // 37: vdp.model.v1alpha.ModelService.UnpublishModel:output_type -> vdp.model.v1alpha.UnpublishModelResponse
	38, // 38: vdp.model.v1alpha.ModelService.ListModelInstance:output_type -> vdp.model.v1alpha.ListModelInstanceResponse
	39, // 39: vdp.model.v1alpha.ModelService.GetModelInstance:output_type -> vdp.model.v1alpha.GetModelInstanceResponse
	40, // 40: vdp.model.v1alpha.ModelService.LookUpModelInstance:output_type -> vdp.model.v1alpha.LookUpModelInstanceResponse
	41, // 41: vdp.model.v1alpha.ModelService.DeployModelInstance:output_type -> vdp.model.v1alpha.DeployModelInstanceResponse
	42, // 42: vdp.model.v1alpha.ModelService.UndeployModelInstance:output_type -> vdp.model.v1alpha.UndeployModelInstanceResponse
	43, // 43: vdp.model.v1alpha.ModelService.GetModelInstanceCard:output_type -> vdp.model.v1alpha.GetModelInstanceCardResponse
	44, // 44: vdp.model.v1alpha.ModelService.TriggerModelInstance:output_type -> vdp.model.v1alpha.TriggerModelInstanceResponse
	45, // 45: vdp.model.v1alpha.ModelService.TriggerModelInstanceBinaryFileUpload:output_type -> vdp.model.v1alpha.TriggerModelInstanceBinaryFileUploadResponse
	46, // 46: vdp.model.v1alpha.ModelService.TestModelInstance:output_type -> vdp.model.v1alpha.TestModelInstanceResponse
	47, // 47: vdp.model.v1alpha.ModelService.TestModelInstanceBinaryFileUpload:output_type -> vdp.model.v1alpha.TestModelInstanceBinaryFileUploadResponse
	24, // [24:48] is the sub-list for method output_type
	0,  // [0:24] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_vdp_model_v1alpha_model_service_proto_init() }
func file_vdp_model_v1alpha_model_service_proto_init() {
	if File_vdp_model_v1alpha_model_service_proto != nil {
		return
	}
	file_vdp_model_v1alpha_healthcheck_proto_init()
	file_vdp_model_v1alpha_model_definition_proto_init()
	file_vdp_model_v1alpha_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_model_v1alpha_model_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vdp_model_v1alpha_model_service_proto_goTypes,
		DependencyIndexes: file_vdp_model_v1alpha_model_service_proto_depIdxs,
	}.Build()
	File_vdp_model_v1alpha_model_service_proto = out.File
	file_vdp_model_v1alpha_model_service_proto_rawDesc = nil
	file_vdp_model_v1alpha_model_service_proto_goTypes = nil
	file_vdp_model_v1alpha_model_service_proto_depIdxs = nil
}
