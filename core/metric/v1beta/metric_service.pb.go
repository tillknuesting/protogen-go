//
//This is the POC for metric-backend service that supports
//1. receiving usage reports from pipeline-backend and model-backend clients
//2. uploading usage to InfluxDB
//3. uploading usage to Stripe
//4. responding usage/price enquiries from clients

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: core/metric/v1beta/metric_service.proto

package metricv1beta

import (
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

var File_core_metric_v1beta_metric_service_proto protoreflect.FileDescriptor

var file_core_metric_v1beta_metric_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1,
	0x0c, 0x0a, 0x0f, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x61, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x30,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x31, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x94, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x3e,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0xb2, 0x01, 0x0a, 0x27, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x75, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x42, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x43, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x75, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x9a, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x8e, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xa4, 0x0b, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x12, 0x24, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a,
	0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x73, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x30, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x6c, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9a, 0x01,
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x23, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x43,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x43,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x30, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x88, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x6c, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c,
	0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd9, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x42, 0x12, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0xa2, 0x02, 0x03,
	0x43, 0x4d, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xca, 0x02, 0x12, 0x43, 0x6f, 0x72, 0x65, 0x5c,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xe2, 0x02, 0x1e,
	0x43, 0x6f, 0x72, 0x65, 0x5c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x14, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_core_metric_v1beta_metric_service_proto_goTypes = []interface{}{
	(*GetPipelinesRequest)(nil),                             // 0: core.metric.v1beta.GetPipelinesRequest
	(*ReportPipelineTriggerRequest)(nil),                    // 1: core.metric.v1beta.ReportPipelineTriggerRequest
	(*ReportPipelineTriggersRequest)(nil),                   // 2: core.metric.v1beta.ReportPipelineTriggersRequest
	(*GetPipelineTriggerRecordsRequest)(nil),                // 3: core.metric.v1beta.GetPipelineTriggerRecordsRequest
	(*GetBulkPipelineTriggerRecordsRequest)(nil),            // 4: core.metric.v1beta.GetBulkPipelineTriggerRecordsRequest
	(*GetCumulativePipelineTriggerRecordsRequest)(nil),      // 5: core.metric.v1beta.GetCumulativePipelineTriggerRecordsRequest
	(*GetBulkCumulativePipelineTriggerRecordsRequest)(nil),  // 6: core.metric.v1beta.GetBulkCumulativePipelineTriggerRecordsRequest
	(*GetPipelineTriggerSummaryRequest)(nil),                // 7: core.metric.v1beta.GetPipelineTriggerSummaryRequest
	(*GetBulkPipelineTriggerSummariesRequest)(nil),          // 8: core.metric.v1beta.GetBulkPipelineTriggerSummariesRequest
	(*GetPipelineTriggerPriceRequest)(nil),                  // 9: core.metric.v1beta.GetPipelineTriggerPriceRequest
	(*GetBulkPipelineTriggerPriceRequest)(nil),              // 10: core.metric.v1beta.GetBulkPipelineTriggerPriceRequest
	(*GetModelsRequest)(nil),                                // 11: core.metric.v1beta.GetModelsRequest
	(*ReportModelOnlineRequest)(nil),                        // 12: core.metric.v1beta.ReportModelOnlineRequest
	(*ReportModelOnlinesRequest)(nil),                       // 13: core.metric.v1beta.ReportModelOnlinesRequest
	(*GetModelOnlineRecordsRequest)(nil),                    // 14: core.metric.v1beta.GetModelOnlineRecordsRequest
	(*GetBulkModelOnlineRecordsRequest)(nil),                // 15: core.metric.v1beta.GetBulkModelOnlineRecordsRequest
	(*GetCumulativeModelOnlineRecordsRequest)(nil),          // 16: core.metric.v1beta.GetCumulativeModelOnlineRecordsRequest
	(*GetBulkCumulativeModelOnlineRecordsRequest)(nil),      // 17: core.metric.v1beta.GetBulkCumulativeModelOnlineRecordsRequest
	(*GetModelOnlineSummaryRequest)(nil),                    // 18: core.metric.v1beta.GetModelOnlineSummaryRequest
	(*GetBulkModelOnlineSummaryRequest)(nil),                // 19: core.metric.v1beta.GetBulkModelOnlineSummaryRequest
	(*GetModelOnlinePriceRequest)(nil),                      // 20: core.metric.v1beta.GetModelOnlinePriceRequest
	(*GetBulkModelOnlinePriceRequest)(nil),                  // 21: core.metric.v1beta.GetBulkModelOnlinePriceRequest
	(*GetPipelinesResponse)(nil),                            // 22: core.metric.v1beta.GetPipelinesResponse
	(*ReportPipelineTriggerResponse)(nil),                   // 23: core.metric.v1beta.ReportPipelineTriggerResponse
	(*ReportPipelineTriggersResponse)(nil),                  // 24: core.metric.v1beta.ReportPipelineTriggersResponse
	(*GetPipelineTriggerRecordsResponse)(nil),               // 25: core.metric.v1beta.GetPipelineTriggerRecordsResponse
	(*GetBulkPipelineTriggerRecordsResponse)(nil),           // 26: core.metric.v1beta.GetBulkPipelineTriggerRecordsResponse
	(*GetCumulativePipelineTriggerRecordsResponse)(nil),     // 27: core.metric.v1beta.GetCumulativePipelineTriggerRecordsResponse
	(*GetBulkCumulativePipelineTriggerRecordsResponse)(nil), // 28: core.metric.v1beta.GetBulkCumulativePipelineTriggerRecordsResponse
	(*GetPipelineTriggerSummaryResponse)(nil),               // 29: core.metric.v1beta.GetPipelineTriggerSummaryResponse
	(*GetBulkPipelineTriggerSummariesResponse)(nil),         // 30: core.metric.v1beta.GetBulkPipelineTriggerSummariesResponse
	(*GetPipelineTriggerPriceResponse)(nil),                 // 31: core.metric.v1beta.GetPipelineTriggerPriceResponse
	(*GetBulkPipelineTriggerPriceResponse)(nil),             // 32: core.metric.v1beta.GetBulkPipelineTriggerPriceResponse
	(*GetModelsResponse)(nil),                               // 33: core.metric.v1beta.GetModelsResponse
	(*ReportModelOnlineResponse)(nil),                       // 34: core.metric.v1beta.ReportModelOnlineResponse
	(*ReportModelOnlinesResponse)(nil),                      // 35: core.metric.v1beta.ReportModelOnlinesResponse
	(*GetModelOnlineRecordsResponse)(nil),                   // 36: core.metric.v1beta.GetModelOnlineRecordsResponse
	(*GetBulkModelOnlineRecordsResponse)(nil),               // 37: core.metric.v1beta.GetBulkModelOnlineRecordsResponse
	(*GetCumulativeModelOnlineRecordsResponse)(nil),         // 38: core.metric.v1beta.GetCumulativeModelOnlineRecordsResponse
	(*GetBulkCumulativeModelOnlineRecordsResponse)(nil),     // 39: core.metric.v1beta.GetBulkCumulativeModelOnlineRecordsResponse
	(*GetModelOnlineSummaryResponse)(nil),                   // 40: core.metric.v1beta.GetModelOnlineSummaryResponse
	(*GetBulkModelOnlineSummaryResponse)(nil),               // 41: core.metric.v1beta.GetBulkModelOnlineSummaryResponse
	(*GetModelOnlinePriceResponse)(nil),                     // 42: core.metric.v1beta.GetModelOnlinePriceResponse
	(*GetBulkModelOnlinePriceResponse)(nil),                 // 43: core.metric.v1beta.GetBulkModelOnlinePriceResponse
}
var file_core_metric_v1beta_metric_service_proto_depIdxs = []int32{
	0,  // 0: core.metric.v1beta.PipelineService.GetPipelines:input_type -> core.metric.v1beta.GetPipelinesRequest
	1,  // 1: core.metric.v1beta.PipelineService.ReportPipelineTrigger:input_type -> core.metric.v1beta.ReportPipelineTriggerRequest
	2,  // 2: core.metric.v1beta.PipelineService.ReportPipelineTriggers:input_type -> core.metric.v1beta.ReportPipelineTriggersRequest
	3,  // 3: core.metric.v1beta.PipelineService.GetPipelineTriggerRecords:input_type -> core.metric.v1beta.GetPipelineTriggerRecordsRequest
	4,  // 4: core.metric.v1beta.PipelineService.GetBulkPipelineTriggerRecords:input_type -> core.metric.v1beta.GetBulkPipelineTriggerRecordsRequest
	5,  // 5: core.metric.v1beta.PipelineService.GetCumulativePipelineTriggerRecords:input_type -> core.metric.v1beta.GetCumulativePipelineTriggerRecordsRequest
	6,  // 6: core.metric.v1beta.PipelineService.GetBulkCumulativePipelineTriggerRecords:input_type -> core.metric.v1beta.GetBulkCumulativePipelineTriggerRecordsRequest
	7,  // 7: core.metric.v1beta.PipelineService.GetPipelineTriggerSummary:input_type -> core.metric.v1beta.GetPipelineTriggerSummaryRequest
	8,  // 8: core.metric.v1beta.PipelineService.GetBulkPipelineTriggerSummaries:input_type -> core.metric.v1beta.GetBulkPipelineTriggerSummariesRequest
	9,  // 9: core.metric.v1beta.PipelineService.GetPipelineTriggerPrice:input_type -> core.metric.v1beta.GetPipelineTriggerPriceRequest
	10, // 10: core.metric.v1beta.PipelineService.GetBulkPipelineTriggerPrice:input_type -> core.metric.v1beta.GetBulkPipelineTriggerPriceRequest
	11, // 11: core.metric.v1beta.ModelService.GetModels:input_type -> core.metric.v1beta.GetModelsRequest
	12, // 12: core.metric.v1beta.ModelService.ReportModelOnline:input_type -> core.metric.v1beta.ReportModelOnlineRequest
	13, // 13: core.metric.v1beta.ModelService.ReportModelOnlines:input_type -> core.metric.v1beta.ReportModelOnlinesRequest
	14, // 14: core.metric.v1beta.ModelService.GetModelOnlineRecords:input_type -> core.metric.v1beta.GetModelOnlineRecordsRequest
	15, // 15: core.metric.v1beta.ModelService.GetBulkModelOnlineRecords:input_type -> core.metric.v1beta.GetBulkModelOnlineRecordsRequest
	16, // 16: core.metric.v1beta.ModelService.GetCumulativeModelOnlineRecords:input_type -> core.metric.v1beta.GetCumulativeModelOnlineRecordsRequest
	17, // 17: core.metric.v1beta.ModelService.GetBulkCumulativeModelOnlineRecords:input_type -> core.metric.v1beta.GetBulkCumulativeModelOnlineRecordsRequest
	18, // 18: core.metric.v1beta.ModelService.GetModelOnlineSummary:input_type -> core.metric.v1beta.GetModelOnlineSummaryRequest
	19, // 19: core.metric.v1beta.ModelService.GetBulkModelOnlineSummary:input_type -> core.metric.v1beta.GetBulkModelOnlineSummaryRequest
	20, // 20: core.metric.v1beta.ModelService.GetModelOnlinePrice:input_type -> core.metric.v1beta.GetModelOnlinePriceRequest
	21, // 21: core.metric.v1beta.ModelService.GetBulkModelOnlinePrice:input_type -> core.metric.v1beta.GetBulkModelOnlinePriceRequest
	22, // 22: core.metric.v1beta.PipelineService.GetPipelines:output_type -> core.metric.v1beta.GetPipelinesResponse
	23, // 23: core.metric.v1beta.PipelineService.ReportPipelineTrigger:output_type -> core.metric.v1beta.ReportPipelineTriggerResponse
	24, // 24: core.metric.v1beta.PipelineService.ReportPipelineTriggers:output_type -> core.metric.v1beta.ReportPipelineTriggersResponse
	25, // 25: core.metric.v1beta.PipelineService.GetPipelineTriggerRecords:output_type -> core.metric.v1beta.GetPipelineTriggerRecordsResponse
	26, // 26: core.metric.v1beta.PipelineService.GetBulkPipelineTriggerRecords:output_type -> core.metric.v1beta.GetBulkPipelineTriggerRecordsResponse
	27, // 27: core.metric.v1beta.PipelineService.GetCumulativePipelineTriggerRecords:output_type -> core.metric.v1beta.GetCumulativePipelineTriggerRecordsResponse
	28, // 28: core.metric.v1beta.PipelineService.GetBulkCumulativePipelineTriggerRecords:output_type -> core.metric.v1beta.GetBulkCumulativePipelineTriggerRecordsResponse
	29, // 29: core.metric.v1beta.PipelineService.GetPipelineTriggerSummary:output_type -> core.metric.v1beta.GetPipelineTriggerSummaryResponse
	30, // 30: core.metric.v1beta.PipelineService.GetBulkPipelineTriggerSummaries:output_type -> core.metric.v1beta.GetBulkPipelineTriggerSummariesResponse
	31, // 31: core.metric.v1beta.PipelineService.GetPipelineTriggerPrice:output_type -> core.metric.v1beta.GetPipelineTriggerPriceResponse
	32, // 32: core.metric.v1beta.PipelineService.GetBulkPipelineTriggerPrice:output_type -> core.metric.v1beta.GetBulkPipelineTriggerPriceResponse
	33, // 33: core.metric.v1beta.ModelService.GetModels:output_type -> core.metric.v1beta.GetModelsResponse
	34, // 34: core.metric.v1beta.ModelService.ReportModelOnline:output_type -> core.metric.v1beta.ReportModelOnlineResponse
	35, // 35: core.metric.v1beta.ModelService.ReportModelOnlines:output_type -> core.metric.v1beta.ReportModelOnlinesResponse
	36, // 36: core.metric.v1beta.ModelService.GetModelOnlineRecords:output_type -> core.metric.v1beta.GetModelOnlineRecordsResponse
	37, // 37: core.metric.v1beta.ModelService.GetBulkModelOnlineRecords:output_type -> core.metric.v1beta.GetBulkModelOnlineRecordsResponse
	38, // 38: core.metric.v1beta.ModelService.GetCumulativeModelOnlineRecords:output_type -> core.metric.v1beta.GetCumulativeModelOnlineRecordsResponse
	39, // 39: core.metric.v1beta.ModelService.GetBulkCumulativeModelOnlineRecords:output_type -> core.metric.v1beta.GetBulkCumulativeModelOnlineRecordsResponse
	40, // 40: core.metric.v1beta.ModelService.GetModelOnlineSummary:output_type -> core.metric.v1beta.GetModelOnlineSummaryResponse
	41, // 41: core.metric.v1beta.ModelService.GetBulkModelOnlineSummary:output_type -> core.metric.v1beta.GetBulkModelOnlineSummaryResponse
	42, // 42: core.metric.v1beta.ModelService.GetModelOnlinePrice:output_type -> core.metric.v1beta.GetModelOnlinePriceResponse
	43, // 43: core.metric.v1beta.ModelService.GetBulkModelOnlinePrice:output_type -> core.metric.v1beta.GetBulkModelOnlinePriceResponse
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_core_metric_v1beta_metric_service_proto_init() }
func file_core_metric_v1beta_metric_service_proto_init() {
	if File_core_metric_v1beta_metric_service_proto != nil {
		return
	}
	file_core_metric_v1beta_metric_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_metric_v1beta_metric_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_core_metric_v1beta_metric_service_proto_goTypes,
		DependencyIndexes: file_core_metric_v1beta_metric_service_proto_depIdxs,
	}.Build()
	File_core_metric_v1beta_metric_service_proto = out.File
	file_core_metric_v1beta_metric_service_proto_rawDesc = nil
	file_core_metric_v1beta_metric_service_proto_goTypes = nil
	file_core_metric_v1beta_metric_service_proto_depIdxs = nil
}
