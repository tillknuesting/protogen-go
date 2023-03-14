// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package modelv1alpha

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

// ModelPrivateServiceClient is the client API for ModelPrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelPrivateServiceClient interface {
	// ListModelAdmin method receives a ListModelAdminRequest message and returns
	// a ListModelAdminResponse
	ListModelAdmin(ctx context.Context, in *ListModelAdminRequest, opts ...grpc.CallOption) (*ListModelAdminResponse, error)
	// GetModelAdmin method receives a GetModelAdminRequest message and returns a
	// GetModelAdminResponse
	GetModelAdmin(ctx context.Context, in *GetModelAdminRequest, opts ...grpc.CallOption) (*GetModelAdminResponse, error)
	// LookUpModelAdmin method receives a LookUpModelAdminRequest message and
	// returns a LookUpModelAdminResponse
	LookUpModelAdmin(ctx context.Context, in *LookUpModelAdminRequest, opts ...grpc.CallOption) (*LookUpModelAdminResponse, error)
}

type modelPrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelPrivateServiceClient(cc grpc.ClientConnInterface) ModelPrivateServiceClient {
	return &modelPrivateServiceClient{cc}
}

func (c *modelPrivateServiceClient) ListModelAdmin(ctx context.Context, in *ListModelAdminRequest, opts ...grpc.CallOption) (*ListModelAdminResponse, error) {
	out := new(ListModelAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.model.v1alpha.ModelPrivateService/ListModelAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelPrivateServiceClient) GetModelAdmin(ctx context.Context, in *GetModelAdminRequest, opts ...grpc.CallOption) (*GetModelAdminResponse, error) {
	out := new(GetModelAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.model.v1alpha.ModelPrivateService/GetModelAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelPrivateServiceClient) LookUpModelAdmin(ctx context.Context, in *LookUpModelAdminRequest, opts ...grpc.CallOption) (*LookUpModelAdminResponse, error) {
	out := new(LookUpModelAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.model.v1alpha.ModelPrivateService/LookUpModelAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelPrivateServiceServer is the server API for ModelPrivateService service.
// All implementations should embed UnimplementedModelPrivateServiceServer
// for forward compatibility
type ModelPrivateServiceServer interface {
	// ListModelAdmin method receives a ListModelAdminRequest message and returns
	// a ListModelAdminResponse
	ListModelAdmin(context.Context, *ListModelAdminRequest) (*ListModelAdminResponse, error)
	// GetModelAdmin method receives a GetModelAdminRequest message and returns a
	// GetModelAdminResponse
	GetModelAdmin(context.Context, *GetModelAdminRequest) (*GetModelAdminResponse, error)
	// LookUpModelAdmin method receives a LookUpModelAdminRequest message and
	// returns a LookUpModelAdminResponse
	LookUpModelAdmin(context.Context, *LookUpModelAdminRequest) (*LookUpModelAdminResponse, error)
}

// UnimplementedModelPrivateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedModelPrivateServiceServer struct {
}

func (UnimplementedModelPrivateServiceServer) ListModelAdmin(context.Context, *ListModelAdminRequest) (*ListModelAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModelAdmin not implemented")
}
func (UnimplementedModelPrivateServiceServer) GetModelAdmin(context.Context, *GetModelAdminRequest) (*GetModelAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelAdmin not implemented")
}
func (UnimplementedModelPrivateServiceServer) LookUpModelAdmin(context.Context, *LookUpModelAdminRequest) (*LookUpModelAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpModelAdmin not implemented")
}

// UnsafeModelPrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelPrivateServiceServer will
// result in compilation errors.
type UnsafeModelPrivateServiceServer interface {
	mustEmbedUnimplementedModelPrivateServiceServer()
}

func RegisterModelPrivateServiceServer(s grpc.ServiceRegistrar, srv ModelPrivateServiceServer) {
	s.RegisterService(&ModelPrivateService_ServiceDesc, srv)
}

func _ModelPrivateService_ListModelAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelPrivateServiceServer).ListModelAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.model.v1alpha.ModelPrivateService/ListModelAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelPrivateServiceServer).ListModelAdmin(ctx, req.(*ListModelAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelPrivateService_GetModelAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelPrivateServiceServer).GetModelAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.model.v1alpha.ModelPrivateService/GetModelAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelPrivateServiceServer).GetModelAdmin(ctx, req.(*GetModelAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelPrivateService_LookUpModelAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpModelAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelPrivateServiceServer).LookUpModelAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.model.v1alpha.ModelPrivateService/LookUpModelAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelPrivateServiceServer).LookUpModelAdmin(ctx, req.(*LookUpModelAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelPrivateService_ServiceDesc is the grpc.ServiceDesc for ModelPrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelPrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vdp.model.v1alpha.ModelPrivateService",
	HandlerType: (*ModelPrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListModelAdmin",
			Handler:    _ModelPrivateService_ListModelAdmin_Handler,
		},
		{
			MethodName: "GetModelAdmin",
			Handler:    _ModelPrivateService_GetModelAdmin_Handler,
		},
		{
			MethodName: "LookUpModelAdmin",
			Handler:    _ModelPrivateService_LookUpModelAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vdp/model/v1alpha/model_private_service.proto",
}
