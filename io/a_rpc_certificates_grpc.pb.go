// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: io/core/a_rpc_certificates.proto

package io

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Certificates_GetAppleCertificateData_FullMethodName          = "/io.Certificates/getAppleCertificateData"
	Certificates_GetCertificateSigningRequest_FullMethodName     = "/io.Certificates/getCertificateSigningRequest"
	Certificates_AddAppleCertificate_FullMethodName              = "/io.Certificates/addAppleCertificate"
	Certificates_UpdateAppleCertificate_FullMethodName           = "/io.Certificates/updateAppleCertificate"
	Certificates_ListAppleCertificatesDeprecated_FullMethodName  = "/io.Certificates/listAppleCertificatesDeprecated"
	Certificates_ListAppleCertificates_FullMethodName            = "/io.Certificates/listAppleCertificates"
	Certificates_CountAppleCertificatesDeprecated_FullMethodName = "/io.Certificates/countAppleCertificatesDeprecated"
	Certificates_CountAppleCertificates_FullMethodName           = "/io.Certificates/countAppleCertificates"
	Certificates_SendNFCSigningCredentials_FullMethodName        = "/io.Certificates/sendNFCSigningCredentials"
)

// CertificatesClient is the client API for Certificates service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificatesClient interface {
	GetAppleCertificateData(ctx context.Context, in *PassTypeIdentifier, opts ...grpc.CallOption) (*CertificateData, error)
	GetCertificateSigningRequest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CertificateSigningRequest, error)
	AddAppleCertificate(ctx context.Context, in *FileBytes, opts ...grpc.CallOption) (*CertificateData, error)
	UpdateAppleCertificate(ctx context.Context, in *FileBytes, opts ...grpc.CallOption) (*CertificateData, error)
	ListAppleCertificatesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CertificateData], error)
	ListAppleCertificates(ctx context.Context, in *Filters, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CertificateData], error)
	CountAppleCertificatesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Count, error)
	CountAppleCertificates(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Count, error)
	SendNFCSigningCredentials(ctx context.Context, in *NFCSigningCredentialsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type certificatesClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificatesClient(cc grpc.ClientConnInterface) CertificatesClient {
	return &certificatesClient{cc}
}

func (c *certificatesClient) GetAppleCertificateData(ctx context.Context, in *PassTypeIdentifier, opts ...grpc.CallOption) (*CertificateData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateData)
	err := c.cc.Invoke(ctx, Certificates_GetAppleCertificateData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) GetCertificateSigningRequest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CertificateSigningRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateSigningRequest)
	err := c.cc.Invoke(ctx, Certificates_GetCertificateSigningRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) AddAppleCertificate(ctx context.Context, in *FileBytes, opts ...grpc.CallOption) (*CertificateData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateData)
	err := c.cc.Invoke(ctx, Certificates_AddAppleCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) UpdateAppleCertificate(ctx context.Context, in *FileBytes, opts ...grpc.CallOption) (*CertificateData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateData)
	err := c.cc.Invoke(ctx, Certificates_UpdateAppleCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) ListAppleCertificatesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CertificateData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Certificates_ServiceDesc.Streams[0], Certificates_ListAppleCertificatesDeprecated_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Pagination, CertificateData]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Certificates_ListAppleCertificatesDeprecatedClient = grpc.ServerStreamingClient[CertificateData]

func (c *certificatesClient) ListAppleCertificates(ctx context.Context, in *Filters, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CertificateData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Certificates_ServiceDesc.Streams[1], Certificates_ListAppleCertificates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Filters, CertificateData]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Certificates_ListAppleCertificatesClient = grpc.ServerStreamingClient[CertificateData]

func (c *certificatesClient) CountAppleCertificatesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Count, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Count)
	err := c.cc.Invoke(ctx, Certificates_CountAppleCertificatesDeprecated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) CountAppleCertificates(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Count, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Count)
	err := c.cc.Invoke(ctx, Certificates_CountAppleCertificates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) SendNFCSigningCredentials(ctx context.Context, in *NFCSigningCredentialsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Certificates_SendNFCSigningCredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificatesServer is the server API for Certificates service.
// All implementations should embed UnimplementedCertificatesServer
// for forward compatibility.
type CertificatesServer interface {
	GetAppleCertificateData(context.Context, *PassTypeIdentifier) (*CertificateData, error)
	GetCertificateSigningRequest(context.Context, *emptypb.Empty) (*CertificateSigningRequest, error)
	AddAppleCertificate(context.Context, *FileBytes) (*CertificateData, error)
	UpdateAppleCertificate(context.Context, *FileBytes) (*CertificateData, error)
	ListAppleCertificatesDeprecated(*Pagination, grpc.ServerStreamingServer[CertificateData]) error
	ListAppleCertificates(*Filters, grpc.ServerStreamingServer[CertificateData]) error
	CountAppleCertificatesDeprecated(context.Context, *Pagination) (*Count, error)
	CountAppleCertificates(context.Context, *Filters) (*Count, error)
	SendNFCSigningCredentials(context.Context, *NFCSigningCredentialsRequest) (*emptypb.Empty, error)
}

// UnimplementedCertificatesServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCertificatesServer struct{}

func (UnimplementedCertificatesServer) GetAppleCertificateData(context.Context, *PassTypeIdentifier) (*CertificateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppleCertificateData not implemented")
}
func (UnimplementedCertificatesServer) GetCertificateSigningRequest(context.Context, *emptypb.Empty) (*CertificateSigningRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificateSigningRequest not implemented")
}
func (UnimplementedCertificatesServer) AddAppleCertificate(context.Context, *FileBytes) (*CertificateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAppleCertificate not implemented")
}
func (UnimplementedCertificatesServer) UpdateAppleCertificate(context.Context, *FileBytes) (*CertificateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppleCertificate not implemented")
}
func (UnimplementedCertificatesServer) ListAppleCertificatesDeprecated(*Pagination, grpc.ServerStreamingServer[CertificateData]) error {
	return status.Errorf(codes.Unimplemented, "method ListAppleCertificatesDeprecated not implemented")
}
func (UnimplementedCertificatesServer) ListAppleCertificates(*Filters, grpc.ServerStreamingServer[CertificateData]) error {
	return status.Errorf(codes.Unimplemented, "method ListAppleCertificates not implemented")
}
func (UnimplementedCertificatesServer) CountAppleCertificatesDeprecated(context.Context, *Pagination) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAppleCertificatesDeprecated not implemented")
}
func (UnimplementedCertificatesServer) CountAppleCertificates(context.Context, *Filters) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAppleCertificates not implemented")
}
func (UnimplementedCertificatesServer) SendNFCSigningCredentials(context.Context, *NFCSigningCredentialsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNFCSigningCredentials not implemented")
}
func (UnimplementedCertificatesServer) testEmbeddedByValue() {}

// UnsafeCertificatesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificatesServer will
// result in compilation errors.
type UnsafeCertificatesServer interface {
	mustEmbedUnimplementedCertificatesServer()
}

func RegisterCertificatesServer(s grpc.ServiceRegistrar, srv CertificatesServer) {
	// If the following call pancis, it indicates UnimplementedCertificatesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Certificates_ServiceDesc, srv)
}

func _Certificates_GetAppleCertificateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassTypeIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).GetAppleCertificateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_GetAppleCertificateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).GetAppleCertificateData(ctx, req.(*PassTypeIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_GetCertificateSigningRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).GetCertificateSigningRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_GetCertificateSigningRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).GetCertificateSigningRequest(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_AddAppleCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileBytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).AddAppleCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_AddAppleCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).AddAppleCertificate(ctx, req.(*FileBytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_UpdateAppleCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileBytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).UpdateAppleCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_UpdateAppleCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).UpdateAppleCertificate(ctx, req.(*FileBytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_ListAppleCertificatesDeprecated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CertificatesServer).ListAppleCertificatesDeprecated(m, &grpc.GenericServerStream[Pagination, CertificateData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Certificates_ListAppleCertificatesDeprecatedServer = grpc.ServerStreamingServer[CertificateData]

func _Certificates_ListAppleCertificates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filters)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CertificatesServer).ListAppleCertificates(m, &grpc.GenericServerStream[Filters, CertificateData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Certificates_ListAppleCertificatesServer = grpc.ServerStreamingServer[CertificateData]

func _Certificates_CountAppleCertificatesDeprecated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).CountAppleCertificatesDeprecated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_CountAppleCertificatesDeprecated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).CountAppleCertificatesDeprecated(ctx, req.(*Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_CountAppleCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).CountAppleCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_CountAppleCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).CountAppleCertificates(ctx, req.(*Filters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_SendNFCSigningCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NFCSigningCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).SendNFCSigningCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certificates_SendNFCSigningCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).SendNFCSigningCredentials(ctx, req.(*NFCSigningCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Certificates_ServiceDesc is the grpc.ServiceDesc for Certificates service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Certificates_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.Certificates",
	HandlerType: (*CertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAppleCertificateData",
			Handler:    _Certificates_GetAppleCertificateData_Handler,
		},
		{
			MethodName: "getCertificateSigningRequest",
			Handler:    _Certificates_GetCertificateSigningRequest_Handler,
		},
		{
			MethodName: "addAppleCertificate",
			Handler:    _Certificates_AddAppleCertificate_Handler,
		},
		{
			MethodName: "updateAppleCertificate",
			Handler:    _Certificates_UpdateAppleCertificate_Handler,
		},
		{
			MethodName: "countAppleCertificatesDeprecated",
			Handler:    _Certificates_CountAppleCertificatesDeprecated_Handler,
		},
		{
			MethodName: "countAppleCertificates",
			Handler:    _Certificates_CountAppleCertificates_Handler,
		},
		{
			MethodName: "sendNFCSigningCredentials",
			Handler:    _Certificates_SendNFCSigningCredentials_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listAppleCertificatesDeprecated",
			Handler:       _Certificates_ListAppleCertificatesDeprecated_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listAppleCertificates",
			Handler:       _Certificates_ListAppleCertificates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "io/core/a_rpc_certificates.proto",
}
