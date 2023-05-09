// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pipelinev1alpha

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

// PipelinePrivateServiceClient is the client API for PipelinePrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelinePrivateServiceClient interface {
	// ListPipelinesAdmin method receives a ListPipelinesAdminRequest message and
	// returns a ListPipelinesAdminResponse message.
	ListPipelinesAdmin(ctx context.Context, in *ListPipelinesAdminRequest, opts ...grpc.CallOption) (*ListPipelinesAdminResponse, error)
	// LookUpPipelineAdmin method receives a LookUpPipelineAdminRequest message
	// and returns a LookUpPipelineAdminResponse
	LookUpPipelineAdmin(ctx context.Context, in *LookUpPipelineAdminRequest, opts ...grpc.CallOption) (*LookUpPipelineAdminResponse, error)
}

type pipelinePrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelinePrivateServiceClient(cc grpc.ClientConnInterface) PipelinePrivateServiceClient {
	return &pipelinePrivateServiceClient{cc}
}

func (c *pipelinePrivateServiceClient) ListPipelinesAdmin(ctx context.Context, in *ListPipelinesAdminRequest, opts ...grpc.CallOption) (*ListPipelinesAdminResponse, error) {
	out := new(ListPipelinesAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.pipeline.v1alpha.PipelinePrivateService/ListPipelinesAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinePrivateServiceClient) LookUpPipelineAdmin(ctx context.Context, in *LookUpPipelineAdminRequest, opts ...grpc.CallOption) (*LookUpPipelineAdminResponse, error) {
	out := new(LookUpPipelineAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.pipeline.v1alpha.PipelinePrivateService/LookUpPipelineAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelinePrivateServiceServer is the server API for PipelinePrivateService service.
// All implementations should embed UnimplementedPipelinePrivateServiceServer
// for forward compatibility
type PipelinePrivateServiceServer interface {
	// ListPipelinesAdmin method receives a ListPipelinesAdminRequest message and
	// returns a ListPipelinesAdminResponse message.
	ListPipelinesAdmin(context.Context, *ListPipelinesAdminRequest) (*ListPipelinesAdminResponse, error)
	// LookUpPipelineAdmin method receives a LookUpPipelineAdminRequest message
	// and returns a LookUpPipelineAdminResponse
	LookUpPipelineAdmin(context.Context, *LookUpPipelineAdminRequest) (*LookUpPipelineAdminResponse, error)
}

// UnimplementedPipelinePrivateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPipelinePrivateServiceServer struct {
}

func (UnimplementedPipelinePrivateServiceServer) ListPipelinesAdmin(context.Context, *ListPipelinesAdminRequest) (*ListPipelinesAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelinesAdmin not implemented")
}
func (UnimplementedPipelinePrivateServiceServer) LookUpPipelineAdmin(context.Context, *LookUpPipelineAdminRequest) (*LookUpPipelineAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpPipelineAdmin not implemented")
}

// UnsafePipelinePrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelinePrivateServiceServer will
// result in compilation errors.
type UnsafePipelinePrivateServiceServer interface {
	mustEmbedUnimplementedPipelinePrivateServiceServer()
}

func RegisterPipelinePrivateServiceServer(s grpc.ServiceRegistrar, srv PipelinePrivateServiceServer) {
	s.RegisterService(&PipelinePrivateService_ServiceDesc, srv)
}

func _PipelinePrivateService_ListPipelinesAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelinesAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinePrivateServiceServer).ListPipelinesAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.pipeline.v1alpha.PipelinePrivateService/ListPipelinesAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinePrivateServiceServer).ListPipelinesAdmin(ctx, req.(*ListPipelinesAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelinePrivateService_LookUpPipelineAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpPipelineAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinePrivateServiceServer).LookUpPipelineAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.pipeline.v1alpha.PipelinePrivateService/LookUpPipelineAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinePrivateServiceServer).LookUpPipelineAdmin(ctx, req.(*LookUpPipelineAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelinePrivateService_ServiceDesc is the grpc.ServiceDesc for PipelinePrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelinePrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vdp.pipeline.v1alpha.PipelinePrivateService",
	HandlerType: (*PipelinePrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPipelinesAdmin",
			Handler:    _PipelinePrivateService_ListPipelinesAdmin_Handler,
		},
		{
			MethodName: "LookUpPipelineAdmin",
			Handler:    _PipelinePrivateService_LookUpPipelineAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vdp/pipeline/v1alpha/pipeline_private_service.proto",
}
