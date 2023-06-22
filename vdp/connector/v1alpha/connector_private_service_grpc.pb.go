// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: vdp/connector/v1alpha/connector_private_service.proto

package connectorv1alpha

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ConnectorPrivateService_ListSourceConnectorsAdmin_FullMethodName       = "/vdp.connector.v1alpha.ConnectorPrivateService/ListSourceConnectorsAdmin"
	ConnectorPrivateService_LookUpSourceConnectorAdmin_FullMethodName      = "/vdp.connector.v1alpha.ConnectorPrivateService/LookUpSourceConnectorAdmin"
	ConnectorPrivateService_CheckSourceConnector_FullMethodName            = "/vdp.connector.v1alpha.ConnectorPrivateService/CheckSourceConnector"
	ConnectorPrivateService_ListDestinationConnectorsAdmin_FullMethodName  = "/vdp.connector.v1alpha.ConnectorPrivateService/ListDestinationConnectorsAdmin"
	ConnectorPrivateService_LookUpDestinationConnectorAdmin_FullMethodName = "/vdp.connector.v1alpha.ConnectorPrivateService/LookUpDestinationConnectorAdmin"
	ConnectorPrivateService_CheckDestinationConnector_FullMethodName       = "/vdp.connector.v1alpha.ConnectorPrivateService/CheckDestinationConnector"
)

// ConnectorPrivateServiceClient is the client API for ConnectorPrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorPrivateServiceClient interface {
	// ListSourceConnectorsAdmin method receives a ListSourceConnectorsAdminRequest message
	// and returns a ListSourceConnectorsAdminResponse message.
	ListSourceConnectorsAdmin(ctx context.Context, in *ListSourceConnectorsAdminRequest, opts ...grpc.CallOption) (*ListSourceConnectorsAdminResponse, error)
	// LookUpSourceConnectorAdmin method receives a
	// LookUpSourceConnectorAdminRequest message and returns a
	// LookUpSourceConnectorAdminResponse
	LookUpSourceConnectorAdmin(ctx context.Context, in *LookUpSourceConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpSourceConnectorAdminResponse, error)
	// CheckSourceConnector method receives a CheckSourceConnectorRequest message and returns a
	// CheckSourceConnectorResponse
	CheckSourceConnector(ctx context.Context, in *CheckSourceConnectorRequest, opts ...grpc.CallOption) (*CheckSourceConnectorResponse, error)
	// ListDestinationConnectorsAdmin method receives a ListDestinationConnectorsAdminRequest
	// message and returns a ListDestinationConnectorsResponse message.
	ListDestinationConnectorsAdmin(ctx context.Context, in *ListDestinationConnectorsAdminRequest, opts ...grpc.CallOption) (*ListDestinationConnectorsAdminResponse, error)
	// LookUpDestinationConnectorAdmin method receives a
	// LookUpDestinationConnectorAdminRequest message and returns a
	// LookUpDestinationConnectorAdminResponse
	LookUpDestinationConnectorAdmin(ctx context.Context, in *LookUpDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpDestinationConnectorAdminResponse, error)
	// CheckDestinationConnector method receives a CheckDestinationConnectorRequest message and returns a
	// CheckDestinationConnectorResponse
	CheckDestinationConnector(ctx context.Context, in *CheckDestinationConnectorRequest, opts ...grpc.CallOption) (*CheckDestinationConnectorResponse, error)
}

type connectorPrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorPrivateServiceClient(cc grpc.ClientConnInterface) ConnectorPrivateServiceClient {
	return &connectorPrivateServiceClient{cc}
}

func (c *connectorPrivateServiceClient) ListSourceConnectorsAdmin(ctx context.Context, in *ListSourceConnectorsAdminRequest, opts ...grpc.CallOption) (*ListSourceConnectorsAdminResponse, error) {
	out := new(ListSourceConnectorsAdminResponse)
	err := c.cc.Invoke(ctx, ConnectorPrivateService_ListSourceConnectorsAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) LookUpSourceConnectorAdmin(ctx context.Context, in *LookUpSourceConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpSourceConnectorAdminResponse, error) {
	out := new(LookUpSourceConnectorAdminResponse)
	err := c.cc.Invoke(ctx, ConnectorPrivateService_LookUpSourceConnectorAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) CheckSourceConnector(ctx context.Context, in *CheckSourceConnectorRequest, opts ...grpc.CallOption) (*CheckSourceConnectorResponse, error) {
	out := new(CheckSourceConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorPrivateService_CheckSourceConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) ListDestinationConnectorsAdmin(ctx context.Context, in *ListDestinationConnectorsAdminRequest, opts ...grpc.CallOption) (*ListDestinationConnectorsAdminResponse, error) {
	out := new(ListDestinationConnectorsAdminResponse)
	err := c.cc.Invoke(ctx, ConnectorPrivateService_ListDestinationConnectorsAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) LookUpDestinationConnectorAdmin(ctx context.Context, in *LookUpDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpDestinationConnectorAdminResponse, error) {
	out := new(LookUpDestinationConnectorAdminResponse)
	err := c.cc.Invoke(ctx, ConnectorPrivateService_LookUpDestinationConnectorAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) CheckDestinationConnector(ctx context.Context, in *CheckDestinationConnectorRequest, opts ...grpc.CallOption) (*CheckDestinationConnectorResponse, error) {
	out := new(CheckDestinationConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorPrivateService_CheckDestinationConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorPrivateServiceServer is the server API for ConnectorPrivateService service.
// All implementations should embed UnimplementedConnectorPrivateServiceServer
// for forward compatibility
type ConnectorPrivateServiceServer interface {
	// ListSourceConnectorsAdmin method receives a ListSourceConnectorsAdminRequest message
	// and returns a ListSourceConnectorsAdminResponse message.
	ListSourceConnectorsAdmin(context.Context, *ListSourceConnectorsAdminRequest) (*ListSourceConnectorsAdminResponse, error)
	// LookUpSourceConnectorAdmin method receives a
	// LookUpSourceConnectorAdminRequest message and returns a
	// LookUpSourceConnectorAdminResponse
	LookUpSourceConnectorAdmin(context.Context, *LookUpSourceConnectorAdminRequest) (*LookUpSourceConnectorAdminResponse, error)
	// CheckSourceConnector method receives a CheckSourceConnectorRequest message and returns a
	// CheckSourceConnectorResponse
	CheckSourceConnector(context.Context, *CheckSourceConnectorRequest) (*CheckSourceConnectorResponse, error)
	// ListDestinationConnectorsAdmin method receives a ListDestinationConnectorsAdminRequest
	// message and returns a ListDestinationConnectorsResponse message.
	ListDestinationConnectorsAdmin(context.Context, *ListDestinationConnectorsAdminRequest) (*ListDestinationConnectorsAdminResponse, error)
	// LookUpDestinationConnectorAdmin method receives a
	// LookUpDestinationConnectorAdminRequest message and returns a
	// LookUpDestinationConnectorAdminResponse
	LookUpDestinationConnectorAdmin(context.Context, *LookUpDestinationConnectorAdminRequest) (*LookUpDestinationConnectorAdminResponse, error)
	// CheckDestinationConnector method receives a CheckDestinationConnectorRequest message and returns a
	// CheckDestinationConnectorResponse
	CheckDestinationConnector(context.Context, *CheckDestinationConnectorRequest) (*CheckDestinationConnectorResponse, error)
}

// UnimplementedConnectorPrivateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConnectorPrivateServiceServer struct {
}

func (UnimplementedConnectorPrivateServiceServer) ListSourceConnectorsAdmin(context.Context, *ListSourceConnectorsAdminRequest) (*ListSourceConnectorsAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSourceConnectorsAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) LookUpSourceConnectorAdmin(context.Context, *LookUpSourceConnectorAdminRequest) (*LookUpSourceConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpSourceConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) CheckSourceConnector(context.Context, *CheckSourceConnectorRequest) (*CheckSourceConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSourceConnector not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) ListDestinationConnectorsAdmin(context.Context, *ListDestinationConnectorsAdminRequest) (*ListDestinationConnectorsAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDestinationConnectorsAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) LookUpDestinationConnectorAdmin(context.Context, *LookUpDestinationConnectorAdminRequest) (*LookUpDestinationConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpDestinationConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) CheckDestinationConnector(context.Context, *CheckDestinationConnectorRequest) (*CheckDestinationConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDestinationConnector not implemented")
}

// UnsafeConnectorPrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorPrivateServiceServer will
// result in compilation errors.
type UnsafeConnectorPrivateServiceServer interface {
	mustEmbedUnimplementedConnectorPrivateServiceServer()
}

func RegisterConnectorPrivateServiceServer(s grpc.ServiceRegistrar, srv ConnectorPrivateServiceServer) {
	s.RegisterService(&ConnectorPrivateService_ServiceDesc, srv)
}

func _ConnectorPrivateService_ListSourceConnectorsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSourceConnectorsAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).ListSourceConnectorsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorPrivateService_ListSourceConnectorsAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).ListSourceConnectorsAdmin(ctx, req.(*ListSourceConnectorsAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_LookUpSourceConnectorAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpSourceConnectorAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).LookUpSourceConnectorAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorPrivateService_LookUpSourceConnectorAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).LookUpSourceConnectorAdmin(ctx, req.(*LookUpSourceConnectorAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_CheckSourceConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSourceConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).CheckSourceConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorPrivateService_CheckSourceConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).CheckSourceConnector(ctx, req.(*CheckSourceConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_ListDestinationConnectorsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDestinationConnectorsAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).ListDestinationConnectorsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorPrivateService_ListDestinationConnectorsAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).ListDestinationConnectorsAdmin(ctx, req.(*ListDestinationConnectorsAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_LookUpDestinationConnectorAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpDestinationConnectorAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).LookUpDestinationConnectorAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorPrivateService_LookUpDestinationConnectorAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).LookUpDestinationConnectorAdmin(ctx, req.(*LookUpDestinationConnectorAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_CheckDestinationConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDestinationConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).CheckDestinationConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorPrivateService_CheckDestinationConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).CheckDestinationConnector(ctx, req.(*CheckDestinationConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectorPrivateService_ServiceDesc is the grpc.ServiceDesc for ConnectorPrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectorPrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vdp.connector.v1alpha.ConnectorPrivateService",
	HandlerType: (*ConnectorPrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSourceConnectorsAdmin",
			Handler:    _ConnectorPrivateService_ListSourceConnectorsAdmin_Handler,
		},
		{
			MethodName: "LookUpSourceConnectorAdmin",
			Handler:    _ConnectorPrivateService_LookUpSourceConnectorAdmin_Handler,
		},
		{
			MethodName: "CheckSourceConnector",
			Handler:    _ConnectorPrivateService_CheckSourceConnector_Handler,
		},
		{
			MethodName: "ListDestinationConnectorsAdmin",
			Handler:    _ConnectorPrivateService_ListDestinationConnectorsAdmin_Handler,
		},
		{
			MethodName: "LookUpDestinationConnectorAdmin",
			Handler:    _ConnectorPrivateService_LookUpDestinationConnectorAdmin_Handler,
		},
		{
			MethodName: "CheckDestinationConnector",
			Handler:    _ConnectorPrivateService_CheckDestinationConnector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vdp/connector/v1alpha/connector_private_service.proto",
}
