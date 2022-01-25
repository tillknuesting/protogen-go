// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pipeline

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PipelineClient is the client API for Pipeline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineClient interface {
	Liveness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error)
	Readiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error)
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error)
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TriggerPipeline(ctx context.Context, in *TriggerPipelineRequest, opts ...grpc.CallOption) (*structpb.Struct, error)
}

type pipelineClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineClient(cc grpc.ClientConnInterface) PipelineClient {
	return &pipelineClient{cc}
}

func (c *pipelineClient) Liveness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) Readiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/Readiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/CreatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error) {
	out := new(ListPipelinesResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/ListPipelines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/GetPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/UpdatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/DeletePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) TriggerPipeline(ctx context.Context, in *TriggerPipelineRequest, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/instill.pipeline.Pipeline/TriggerPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServer is the server API for Pipeline service.
// All implementations should embed UnimplementedPipelineServer
// for forward compatibility
type PipelineServer interface {
	Liveness(context.Context, *emptypb.Empty) (*structpb.Struct, error)
	Readiness(context.Context, *emptypb.Empty) (*structpb.Struct, error)
	CreatePipeline(context.Context, *CreatePipelineRequest) (*PipelineInfo, error)
	ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error)
	GetPipeline(context.Context, *GetPipelineRequest) (*PipelineInfo, error)
	UpdatePipeline(context.Context, *UpdatePipelineRequest) (*PipelineInfo, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*emptypb.Empty, error)
	TriggerPipeline(context.Context, *TriggerPipelineRequest) (*structpb.Struct, error)
}

// UnimplementedPipelineServer should be embedded to have forward compatible implementations.
type UnimplementedPipelineServer struct {
}

func (UnimplementedPipelineServer) Liveness(context.Context, *emptypb.Empty) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedPipelineServer) Readiness(context.Context, *emptypb.Empty) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (UnimplementedPipelineServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*PipelineInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedPipelineServer) ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelines not implemented")
}
func (UnimplementedPipelineServer) GetPipeline(context.Context, *GetPipelineRequest) (*PipelineInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipeline not implemented")
}
func (UnimplementedPipelineServer) UpdatePipeline(context.Context, *UpdatePipelineRequest) (*PipelineInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePipeline not implemented")
}
func (UnimplementedPipelineServer) DeletePipeline(context.Context, *DeletePipelineRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedPipelineServer) TriggerPipeline(context.Context, *TriggerPipelineRequest) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerPipeline not implemented")
}

// UnsafePipelineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineServer will
// result in compilation errors.
type UnsafePipelineServer interface {
	mustEmbedUnimplementedPipelineServer()
}

func RegisterPipelineServer(s grpc.ServiceRegistrar, srv PipelineServer) {
	s.RegisterService(&Pipeline_ServiceDesc, srv)
}

func _Pipeline_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).Liveness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/Readiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).Readiness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/ListPipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).ListPipelines(ctx, req.(*ListPipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/GetPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).GetPipeline(ctx, req.(*GetPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_UpdatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).UpdatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/UpdatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).UpdatePipeline(ctx, req.(*UpdatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/DeletePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipeline_TriggerPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServer).TriggerPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.Pipeline/TriggerPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServer).TriggerPipeline(ctx, req.(*TriggerPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pipeline_ServiceDesc is the grpc.ServiceDesc for Pipeline service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pipeline_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "instill.pipeline.Pipeline",
	HandlerType: (*PipelineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _Pipeline_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _Pipeline_Readiness_Handler,
		},
		{
			MethodName: "CreatePipeline",
			Handler:    _Pipeline_CreatePipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _Pipeline_ListPipelines_Handler,
		},
		{
			MethodName: "GetPipeline",
			Handler:    _Pipeline_GetPipeline_Handler,
		},
		{
			MethodName: "UpdatePipeline",
			Handler:    _Pipeline_UpdatePipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _Pipeline_DeletePipeline_Handler,
		},
		{
			MethodName: "TriggerPipeline",
			Handler:    _Pipeline_TriggerPipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipeline/pipeline.proto",
}
