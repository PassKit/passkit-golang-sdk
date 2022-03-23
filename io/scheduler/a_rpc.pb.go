//*
// Scheduler RPC
//
// The PassKit Scheduler API lets you schedule a job.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: io/scheduler/a_rpc.proto

package scheduler

import (
	reflect "reflect"

	ct "github.com/PassKit/passkit-golang-grpc-sdk/ct"
	io "github.com/PassKit/passkit-golang-grpc-sdk/io"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_io_scheduler_a_rpc_proto protoreflect.FileDescriptor

var file_io_scheduler_a_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x69, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x63, 0x74, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xac, 0x11, 0x0a,
	0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x8f, 0x02, 0x0a, 0x13, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a,
	0x6f, 0x62, 0x12, 0x11, 0x2e, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x1a, 0x19, 0x2e, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xc9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0xab,
	0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f,
	0x62, 0x73, 0x12, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x1a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20,
	0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x30, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x29, 0x0a, 0x27, 0x54,
	0x68, 0x65, 0x72, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4a, 0x34, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x2d, 0x0a,
	0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x12, 0x88, 0x02, 0x0a,
	0x10, 0x67, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f,
	0x62, 0x12, 0x06, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x74, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x22, 0xd8, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0xb8, 0x01, 0x0a,
	0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x73,
	0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x20, 0x4a, 0x6f, 0x62, 0x1a, 0x4d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20,
	0x61, 0x20, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62,
	0x2e, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x6a, 0x6f, 0x62, 0x20, 0x77, 0x69,
	0x6c, 0x6c, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x69, 0x66, 0x20, 0x61,
	0x6e, 0x79, 0x2e, 0x4a, 0x22, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73,
	0x65, 0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17,
	0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0xd5, 0x03, 0x0a, 0x13, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12,
	0x11, 0x2e, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a,
	0x6f, 0x62, 0x1a, 0x19, 0x2e, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x03,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a, 0x0f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0xf1, 0x02, 0x0a, 0x0f,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x73, 0x12,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x1a, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x20,
	0x61, 0x20, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62,
	0x2e, 0x20, 0x46, 0x75, 0x6c, 0x6c, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x4a, 0x6f, 0x62, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x69, 0x73, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x20, 0x41, 0x6e, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x20, 0x6f, 0x72, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x77,
	0x69, 0x6c, 0x6c, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x20, 0x61, 0x6e, 0x79,
	0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x4a, 0x30, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x29, 0x0a, 0x27, 0x54, 0x68, 0x65, 0x72, 0x65,
	0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x20, 0x77, 0x69,
	0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x4a, 0x22, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65,
	0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17, 0x0a,
	0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x4a, 0x57, 0x0a, 0x03, 0x35, 0x30, 0x33, 0x12, 0x50, 0x0a,
	0x4e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x75, 0x6e, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x20, 0x42, 0x61, 0x63, 0x6b, 0x20, 0x6f, 0x66, 0x66,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x32, 0x35, 0x30, 0x6d, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x2e, 0x12,
	0xfa, 0x02, 0x0a, 0x12, 0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x11, 0x2e, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x1a, 0x19, 0x2e, 0x63, 0x74, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb5, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x32, 0x0f, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x97, 0x02, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e,
	0x67, 0x20, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x14, 0x50, 0x61, 0x74, 0x63, 0x68, 0x20, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x1a, 0x1f, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x30, 0x0a,
	0x03, 0x34, 0x30, 0x30, 0x12, 0x29, 0x0a, 0x27, 0x54, 0x68, 0x65, 0x72, 0x65, 0x20, 0x69, 0x73,
	0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4a,
	0x22, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6c,
	0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17, 0x0a, 0x15, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x2e, 0x4a, 0x57, 0x0a, 0x03, 0x35, 0x30, 0x33, 0x12, 0x50, 0x0a, 0x4e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x20, 0x42, 0x61, 0x63, 0x6b, 0x20, 0x6f, 0x66, 0x66, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x32, 0x35, 0x30, 0x6d, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x75, 0x6e, 0x74, 0x69, 0x6c,
	0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x2e, 0x12, 0xaa, 0x02, 0x0a,
	0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e,
	0x67, 0x4a, 0x6f, 0x62, 0x12, 0x06, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0xf2, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x92, 0x41, 0xd2, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x1a,
	0x52, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62, 0x2e, 0x20, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x73, 0x74, 0x69, 0x6c,
	0x6c, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6a,
	0x6f, 0x62, 0x2e, 0x4a, 0x34, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34,
	0x12, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0xfa, 0x01, 0x0a, 0x17, 0x67, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x06, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x63, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xc6, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x92, 0x41, 0xa2, 0x01, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x20, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1a, 0x47, 0x65, 0x74, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x20, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a,
	0x31, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62, 0x20, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x67, 0x20, 0x69,
	0x64, 0x2e, 0x4a, 0x22, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65,
	0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17, 0x0a,
	0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0x82, 0x02, 0x0a, 0x1a, 0x6c, 0x69, 0x73, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x63, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xb9, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x92, 0x01, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x20, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x20, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62, 0x20, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x4a, 0x22, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19,
	0x55, 0x73, 0x65, 0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34,
	0x12, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x30, 0x01, 0x42, 0xdd, 0x02, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5a, 0x2e, 0x73, 0x74, 0x61, 0x73,
	0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0xaa, 0x02, 0x16, 0x50, 0x61, 0x73,
	0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x92, 0x41, 0xf4, 0x01, 0x12, 0xca, 0x01, 0x0a, 0x15, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x69, 0x74, 0x20, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x20, 0x41, 0x50,
	0x49, 0x12, 0x31, 0x54, 0x68, 0x69, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x20, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73,
	0x69, 0x76, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x20, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x1a, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70, 0x61,
	0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x72, 0x6d, 0x73, 0x2d, 0x6f, 0x66, 0x2d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x22, 0x3f,
	0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x20, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x17, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x69, 0x6f, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x40, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_io_scheduler_a_rpc_proto_goTypes = []interface{}{
	(*ct.SchedulingJob)(nil),         // 0: ct.SchedulingJob
	(*io.Id)(nil),                    // 1: io.Id
	(*ListRequest)(nil),              // 2: scheduler.ListRequest
	(*ct.SchedulingJobResponse)(nil), // 3: ct.SchedulingJobResponse
	(*emptypb.Empty)(nil),            // 4: google.protobuf.Empty
	(*ct.JobHistory)(nil),            // 5: ct.JobHistory
}
var file_io_scheduler_a_rpc_proto_depIdxs = []int32{
	0, // 0: scheduler.Scheduler.createSchedulingJob:input_type -> ct.SchedulingJob
	1, // 1: scheduler.Scheduler.getSchedulingJob:input_type -> io.Id
	0, // 2: scheduler.Scheduler.updateSchedulingJob:input_type -> ct.SchedulingJob
	0, // 3: scheduler.Scheduler.patchSchedulingJob:input_type -> ct.SchedulingJob
	1, // 4: scheduler.Scheduler.deleteSchedulingJob:input_type -> io.Id
	1, // 5: scheduler.Scheduler.getSchedulingJobHistory:input_type -> io.Id
	2, // 6: scheduler.Scheduler.listSchedulingJobHistories:input_type -> scheduler.ListRequest
	3, // 7: scheduler.Scheduler.createSchedulingJob:output_type -> ct.SchedulingJobResponse
	0, // 8: scheduler.Scheduler.getSchedulingJob:output_type -> ct.SchedulingJob
	3, // 9: scheduler.Scheduler.updateSchedulingJob:output_type -> ct.SchedulingJobResponse
	3, // 10: scheduler.Scheduler.patchSchedulingJob:output_type -> ct.SchedulingJobResponse
	4, // 11: scheduler.Scheduler.deleteSchedulingJob:output_type -> google.protobuf.Empty
	5, // 12: scheduler.Scheduler.getSchedulingJobHistory:output_type -> ct.JobHistory
	5, // 13: scheduler.Scheduler.listSchedulingJobHistories:output_type -> ct.JobHistory
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_io_scheduler_a_rpc_proto_init() }
func file_io_scheduler_a_rpc_proto_init() {
	if File_io_scheduler_a_rpc_proto != nil {
		return
	}
	file_io_scheduler_scheduler_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_scheduler_a_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_io_scheduler_a_rpc_proto_goTypes,
		DependencyIndexes: file_io_scheduler_a_rpc_proto_depIdxs,
	}.Build()
	File_io_scheduler_a_rpc_proto = out.File
	file_io_scheduler_a_rpc_proto_rawDesc = nil
	file_io_scheduler_a_rpc_proto_goTypes = nil
	file_io_scheduler_a_rpc_proto_depIdxs = nil
}
