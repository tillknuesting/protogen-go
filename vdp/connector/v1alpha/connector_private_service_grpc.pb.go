// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ConnectorPrivateServiceClient is the client API for ConnectorPrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorPrivateServiceClient interface {
	// ListSourceConnectorAdmin method receives a ListSourceConnectorAdminRequest message
	// and returns a ListSourceConnectorAdminResponse message.
	ListSourceConnectorAdmin(ctx context.Context, in *ListSourceConnectorAdminRequest, opts ...grpc.CallOption) (*ListSourceConnectorAdminResponse, error)
	// GetSourceConnectorAdmin method receives a GetSourceConnectorAdminRequest message and
	// returns a GetSourceConnectorAdminResponse message.
	GetSourceConnectorAdmin(ctx context.Context, in *GetSourceConnectorAdminRequest, opts ...grpc.CallOption) (*GetSourceConnectorAdminResponse, error)
	// LookUpSourceConnectorAdmin method receives a LookUpSourceConnectorAdminRequest
	// message and returns a LookUpSourceConnectorAdminResponse
	LookUpSourceConnectorAdmin(ctx context.Context, in *LookUpSourceConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpSourceConnectorAdminResponse, error)
	// ListDestinationConnectorAdmin method receives a ListDestinationConnectorAdminRequest
	// message and returns a ListDestinationConnectorResponse message.
	ListDestinationConnectorAdmin(ctx context.Context, in *ListDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*ListDestinationConnectorAdminResponse, error)
	// GetDestinationConnectorAdmin method receives a GetDestinationConnectorAdminRequest
	// message and returns a GetDestinationConnectorAdminResponse message.
	GetDestinationConnectorAdmin(ctx context.Context, in *GetDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*GetDestinationConnectorAdminResponse, error)
	// LookUpDestinationConnectorAdmin method receives a
	// LookUpDestinationConnectorAdminRequest message and returns a
	// LookUpDestinationConnectorAdminResponse
	LookUpDestinationConnectorAdmin(ctx context.Context, in *LookUpDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpDestinationConnectorAdminResponse, error)
}

type connectorPrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorPrivateServiceClient(cc grpc.ClientConnInterface) ConnectorPrivateServiceClient {
	return &connectorPrivateServiceClient{cc}
}

func (c *connectorPrivateServiceClient) ListSourceConnectorAdmin(ctx context.Context, in *ListSourceConnectorAdminRequest, opts ...grpc.CallOption) (*ListSourceConnectorAdminResponse, error) {
	out := new(ListSourceConnectorAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.connector.v1alpha.ConnectorPrivateService/ListSourceConnectorAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) GetSourceConnectorAdmin(ctx context.Context, in *GetSourceConnectorAdminRequest, opts ...grpc.CallOption) (*GetSourceConnectorAdminResponse, error) {
	out := new(GetSourceConnectorAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.connector.v1alpha.ConnectorPrivateService/GetSourceConnectorAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) LookUpSourceConnectorAdmin(ctx context.Context, in *LookUpSourceConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpSourceConnectorAdminResponse, error) {
	out := new(LookUpSourceConnectorAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.connector.v1alpha.ConnectorPrivateService/LookUpSourceConnectorAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) ListDestinationConnectorAdmin(ctx context.Context, in *ListDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*ListDestinationConnectorAdminResponse, error) {
	out := new(ListDestinationConnectorAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.connector.v1alpha.ConnectorPrivateService/ListDestinationConnectorAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) GetDestinationConnectorAdmin(ctx context.Context, in *GetDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*GetDestinationConnectorAdminResponse, error) {
	out := new(GetDestinationConnectorAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.connector.v1alpha.ConnectorPrivateService/GetDestinationConnectorAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorPrivateServiceClient) LookUpDestinationConnectorAdmin(ctx context.Context, in *LookUpDestinationConnectorAdminRequest, opts ...grpc.CallOption) (*LookUpDestinationConnectorAdminResponse, error) {
	out := new(LookUpDestinationConnectorAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.connector.v1alpha.ConnectorPrivateService/LookUpDestinationConnectorAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorPrivateServiceServer is the server API for ConnectorPrivateService service.
// All implementations should embed UnimplementedConnectorPrivateServiceServer
// for forward compatibility
type ConnectorPrivateServiceServer interface {
	// ListSourceConnectorAdmin method receives a ListSourceConnectorAdminRequest message
	// and returns a ListSourceConnectorAdminResponse message.
	ListSourceConnectorAdmin(context.Context, *ListSourceConnectorAdminRequest) (*ListSourceConnectorAdminResponse, error)
	// GetSourceConnectorAdmin method receives a GetSourceConnectorAdminRequest message and
	// returns a GetSourceConnectorAdminResponse message.
	GetSourceConnectorAdmin(context.Context, *GetSourceConnectorAdminRequest) (*GetSourceConnectorAdminResponse, error)
	// LookUpSourceConnectorAdmin method receives a LookUpSourceConnectorAdminRequest
	// message and returns a LookUpSourceConnectorAdminResponse
	LookUpSourceConnectorAdmin(context.Context, *LookUpSourceConnectorAdminRequest) (*LookUpSourceConnectorAdminResponse, error)
	// ListDestinationConnectorAdmin method receives a ListDestinationConnectorAdminRequest
	// message and returns a ListDestinationConnectorResponse message.
	ListDestinationConnectorAdmin(context.Context, *ListDestinationConnectorAdminRequest) (*ListDestinationConnectorAdminResponse, error)
	// GetDestinationConnectorAdmin method receives a GetDestinationConnectorAdminRequest
	// message and returns a GetDestinationConnectorAdminResponse message.
	GetDestinationConnectorAdmin(context.Context, *GetDestinationConnectorAdminRequest) (*GetDestinationConnectorAdminResponse, error)
	// LookUpDestinationConnectorAdmin method receives a
	// LookUpDestinationConnectorAdminRequest message and returns a
	// LookUpDestinationConnectorAdminResponse
	LookUpDestinationConnectorAdmin(context.Context, *LookUpDestinationConnectorAdminRequest) (*LookUpDestinationConnectorAdminResponse, error)
}

// UnimplementedConnectorPrivateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConnectorPrivateServiceServer struct {
}

func (UnimplementedConnectorPrivateServiceServer) ListSourceConnectorAdmin(context.Context, *ListSourceConnectorAdminRequest) (*ListSourceConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSourceConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) GetSourceConnectorAdmin(context.Context, *GetSourceConnectorAdminRequest) (*GetSourceConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourceConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) LookUpSourceConnectorAdmin(context.Context, *LookUpSourceConnectorAdminRequest) (*LookUpSourceConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpSourceConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) ListDestinationConnectorAdmin(context.Context, *ListDestinationConnectorAdminRequest) (*ListDestinationConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDestinationConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) GetDestinationConnectorAdmin(context.Context, *GetDestinationConnectorAdminRequest) (*GetDestinationConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationConnectorAdmin not implemented")
}
func (UnimplementedConnectorPrivateServiceServer) LookUpDestinationConnectorAdmin(context.Context, *LookUpDestinationConnectorAdminRequest) (*LookUpDestinationConnectorAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpDestinationConnectorAdmin not implemented")
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

func _ConnectorPrivateService_ListSourceConnectorAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSourceConnectorAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).ListSourceConnectorAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.connector.v1alpha.ConnectorPrivateService/ListSourceConnectorAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).ListSourceConnectorAdmin(ctx, req.(*ListSourceConnectorAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_GetSourceConnectorAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSourceConnectorAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).GetSourceConnectorAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.connector.v1alpha.ConnectorPrivateService/GetSourceConnectorAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).GetSourceConnectorAdmin(ctx, req.(*GetSourceConnectorAdminRequest))
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
		FullMethod: "/vdp.connector.v1alpha.ConnectorPrivateService/LookUpSourceConnectorAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).LookUpSourceConnectorAdmin(ctx, req.(*LookUpSourceConnectorAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_ListDestinationConnectorAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDestinationConnectorAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).ListDestinationConnectorAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.connector.v1alpha.ConnectorPrivateService/ListDestinationConnectorAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).ListDestinationConnectorAdmin(ctx, req.(*ListDestinationConnectorAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorPrivateService_GetDestinationConnectorAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationConnectorAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorPrivateServiceServer).GetDestinationConnectorAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.connector.v1alpha.ConnectorPrivateService/GetDestinationConnectorAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).GetDestinationConnectorAdmin(ctx, req.(*GetDestinationConnectorAdminRequest))
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
		FullMethod: "/vdp.connector.v1alpha.ConnectorPrivateService/LookUpDestinationConnectorAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorPrivateServiceServer).LookUpDestinationConnectorAdmin(ctx, req.(*LookUpDestinationConnectorAdminRequest))
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
			MethodName: "ListSourceConnectorAdmin",
			Handler:    _ConnectorPrivateService_ListSourceConnectorAdmin_Handler,
		},
		{
			MethodName: "GetSourceConnectorAdmin",
			Handler:    _ConnectorPrivateService_GetSourceConnectorAdmin_Handler,
		},
		{
			MethodName: "LookUpSourceConnectorAdmin",
			Handler:    _ConnectorPrivateService_LookUpSourceConnectorAdmin_Handler,
		},
		{
			MethodName: "ListDestinationConnectorAdmin",
			Handler:    _ConnectorPrivateService_ListDestinationConnectorAdmin_Handler,
		},
		{
			MethodName: "GetDestinationConnectorAdmin",
			Handler:    _ConnectorPrivateService_GetDestinationConnectorAdmin_Handler,
		},
		{
			MethodName: "LookUpDestinationConnectorAdmin",
			Handler:    _ConnectorPrivateService_LookUpDestinationConnectorAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vdp/connector/v1alpha/connector_private_service.proto",
}
