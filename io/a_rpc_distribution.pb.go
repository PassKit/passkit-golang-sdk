// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: io/core/a_rpc_distribution.proto

package io

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_io_core_a_rpc_distribution_proto protoreflect.FileDescriptor

var file_io_core_a_rpc_distribution_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x5f,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xa2, 0x0d, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0xaf, 0x01, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x65, 0x92, 0x41, 0x44,
	0x12, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x20, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x61, 0x20, 0x77, 0x65,
	0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x61,
	0x6e, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x20, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0xdd, 0x01, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x50, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x22, 0xa5, 0x01, 0x92, 0x41, 0x7c,
	0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50, 0x61, 0x73, 0x73, 0x20,
	0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x66, 0x47, 0x65, 0x74, 0x20, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x20, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50, 0x61, 0x73, 0x73, 0x20,
	0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x61, 0x20,
	0x70, 0x61, 0x73, 0x73, 0x20, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x61, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2c, 0x20, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x20, 0x6f, 0x72, 0x20, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x61, 0x73, 0x73, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0xc5, 0x02, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0xf5, 0x01, 0x92, 0x41, 0xd5, 0x01, 0x12, 0x1f, 0x47, 0x65, 0x74, 0x20, 0x44,
	0x61, 0x74, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x50,
	0x61, 0x67, 0x65, 0x20, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x6e, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x20, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64,
	0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x6c,
	0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4a, 0x22, 0x0a, 0x03, 0x34, 0x30,
	0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e,
	0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20,
	0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x98, 0x02, 0x0a, 0x12, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50, 0x61, 0x73, 0x73, 0x43, 0x73,
	0x76, 0x12, 0x1d, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x43, 0x73, 0x76, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xca, 0x01, 0x92, 0x41, 0xa4, 0x01, 0x12,
	0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x20, 0x43, 0x73, 0x76, 0x1a, 0x8b, 0x01, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x50, 0x61, 0x73, 0x73, 0x20, 0x43, 0x73, 0x76, 0x20, 0x73, 0x65, 0x6e,
	0x64, 0x73, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x20, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x66, 0x69, 0x6c,
	0x65, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x2e, 0x20, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20,
	0x73, 0x65, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x70, 0x61, 0x73, 0x73, 0x12, 0xdf, 0x01, 0x0a, 0x11, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x73, 0x76, 0x12, 0x19, 0x2e, 0x69, 0x6f,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x96,
	0x01, 0x92, 0x41, 0x70, 0x12, 0x13, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x20, 0x43, 0x73, 0x76, 0x1a, 0x59, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x20, 0x63, 0x73, 0x76, 0x20, 0x74,
	0x61, 0x6b, 0x65, 0x73, 0x20, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64,
	0x20, 0x43, 0x53, 0x56, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73,
	0x20, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x73, 0x76,
	0x2d, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0xba, 0x03, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x2e, 0x69, 0x6f,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0b, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x8c, 0x03, 0x92, 0x41, 0xe0, 0x02, 0x12, 0x10, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0xb8,
	0x01, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x54, 0x4f, 0x54, 0x50,
	0x20, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x20, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x66, 0x75, 0x6c, 0x6c, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x20,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x7b, 0x7b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x7d, 0x7d, 0x7c, 0x7c, 0x7b, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x7c, 0x7c, 0x7b, 0x7b, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x7d, 0x7d, 0x7c, 0x7c, 0x7b, 0x7b, 0x74, 0x6f, 0x74,
	0x70, 0x7d, 0x7d, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x54, 0x4f, 0x54, 0x50, 0x20, 0x69, 0x73, 0x20,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x2c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x20, 0x54, 0x4f,
	0x54, 0x50, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x2e, 0x4a, 0x4c, 0x0a, 0x03, 0x34, 0x30, 0x30,
	0x12, 0x45, 0x0a, 0x43, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x77, 0x61, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x20, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x50,
	0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x4a, 0x43, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3c,
	0x0a, 0x3a, 0x54, 0x68, 0x65, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x54, 0x4f,
	0x54, 0x50, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x81, 0x07, 0x92, 0x41, 0xb6, 0x06, 0x12, 0xf1, 0x01, 0x0a, 0x18,
	0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x20, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x12, 0x5a, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x50, 0x61, 0x73, 0x73, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66,
	0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x70, 0x61, 0x73, 0x73, 0x65, 0x73, 0x20, 0x76, 0x69, 0x61,
	0x20, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x2e, 0x1a, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70, 0x61,
	0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x72, 0x6d, 0x73, 0x2d, 0x6f, 0x66, 0x2d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x22, 0x3f,
	0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x20, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x17, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x69, 0x6f, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x40, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2a,
	0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x39, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x32, 0x0a,
	0x28, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x73, 0x20, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01,
	0x07, 0x52, 0x34, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x69, 0x73, 0x20, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x2e, 0x52, 0x30, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x29,
	0x0a, 0x27, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33,
	0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34,
	0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e,
	0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x3c, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12,
	0x35, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x72, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x6e, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x12, 0x06,
	0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x57, 0x0a, 0x03, 0x35, 0x30, 0x33, 0x12, 0x50, 0x0a,
	0x4e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x75, 0x6e, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x20, 0x42, 0x61, 0x63, 0x6b, 0x20, 0x6f, 0x66, 0x66,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x32, 0x35, 0x30, 0x6d, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x2e, 0x5a,
	0x3e, 0x0a, 0x3c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x2e, 0x08, 0x02, 0x12, 0x19, 0x4a, 0x57, 0x54, 0x20, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x1a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62,
	0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x00, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b,
	0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_io_core_a_rpc_distribution_proto_goTypes = []any{
	(*EmailDistributionRequest)(nil),  // 0: io.EmailDistributionRequest
	(*SmartPassLinkRequest)(nil),      // 1: io.SmartPassLinkRequest
	(*ClassObjectInput)(nil),          // 2: io.ClassObjectInput
	(*SmartPassCsvUploadRequest)(nil), // 3: io.SmartPassCsvUploadRequest
	(*ImportProtocolRequest)(nil),     // 4: io.ImportProtocolRequest
	(*Payload)(nil),                   // 5: io.Payload
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
	(*Url)(nil),                       // 7: io.Url
	(*DataCollectionFields)(nil),      // 8: io.DataCollectionFields
}
var file_io_core_a_rpc_distribution_proto_depIdxs = []int32{
	0, // 0: io.Distribution.sendWelcomeEmail:input_type -> io.EmailDistributionRequest
	1, // 1: io.Distribution.getSmartPassLink:input_type -> io.SmartPassLinkRequest
	2, // 2: io.Distribution.getDataCollectionPageFields:input_type -> io.ClassObjectInput
	3, // 3: io.Distribution.uploadSmartPassCsv:input_type -> io.SmartPassCsvUploadRequest
	4, // 4: io.Distribution.importProtocolCsv:input_type -> io.ImportProtocolRequest
	5, // 5: io.Distribution.validateBarcode:input_type -> io.Payload
	6, // 6: io.Distribution.sendWelcomeEmail:output_type -> google.protobuf.Empty
	7, // 7: io.Distribution.getSmartPassLink:output_type -> io.Url
	8, // 8: io.Distribution.getDataCollectionPageFields:output_type -> io.DataCollectionFields
	6, // 9: io.Distribution.uploadSmartPassCsv:output_type -> google.protobuf.Empty
	6, // 10: io.Distribution.importProtocolCsv:output_type -> google.protobuf.Empty
	5, // 11: io.Distribution.validateBarcode:output_type -> io.Payload
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_io_core_a_rpc_distribution_proto_init() }
func file_io_core_a_rpc_distribution_proto_init() {
	if File_io_core_a_rpc_distribution_proto != nil {
		return
	}
	file_io_common_common_objects_proto_init()
	file_io_common_distribution_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_io_core_a_rpc_distribution_proto_rawDesc), len(file_io_core_a_rpc_distribution_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_io_core_a_rpc_distribution_proto_goTypes,
		DependencyIndexes: file_io_core_a_rpc_distribution_proto_depIdxs,
	}.Build()
	File_io_core_a_rpc_distribution_proto = out.File
	file_io_core_a_rpc_distribution_proto_goTypes = nil
	file_io_core_a_rpc_distribution_proto_depIdxs = nil
}
