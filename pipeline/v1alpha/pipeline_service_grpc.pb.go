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

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineServiceClient interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// CreatePipeline method receives a CreatePipelineRequest message and returns
	// a CreatePipelineResponse message.
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error)
	// ListPipeline method receives a ListPipelineRequest message and returns a
	// ListPipelineResponse message.
	ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*ListPipelineResponse, error)
	// GetPipeline method receives a GetPipelineRequest message and returns a
	// GetPipelineResponse message.
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error)
	// UpdatePipeline method receives a UpdatePipelineRequest message and returns
	// a UpdatePipelineResponse message.
	UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*UpdatePipelineResponse, error)
	// DeletePipeline method receives a DeletePipelineRequest message and returns
	// a DeletePipelineResponse message.
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*DeletePipelineResponse, error)
	// LookUpPipeline method receives a LookUpPipelineRequest message and returns
	// a LookUpPipelineResponse
	LookUpPipeline(ctx context.Context, in *LookUpPipelineRequest, opts ...grpc.CallOption) (*LookUpPipelineResponse, error)
	// Activate a pipeline.
	// The `state` of the pipeline after activating is `ACTIVE`.
	// ActivatePipeline` can be called on Pipelines in the state `INACTIVE`;
	// Pipelines in a different state (including `ACTIVE`) returns an error.
	ActivatePipeline(ctx context.Context, in *ActivatePipelineRequest, opts ...grpc.CallOption) (*ActivatePipelineResponse, error)
	// Deactivate a pipeline.
	// The `state` of the pipeline after inactivating is `INACTIVE`.
	// DeactivatePipeline` can be called on Pipelines in the state `ACTIVE`;
	// Pipelines in a different state (including `INACTIVE`) returns an error.
	DeactivatePipeline(ctx context.Context, in *DeactivatePipelineRequest, opts ...grpc.CallOption) (*DeactivatePipelineResponse, error)
	// RenamePipeline method receives a RenamePipelineRequest message and returns
	// a RenamePipelineResponse message.
	RenamePipeline(ctx context.Context, in *RenamePipelineRequest, opts ...grpc.CallOption) (*RenamePipelineResponse, error)
	// TriggerPipeline method receives a TriggerPipelineRequest message and
	// returns a TriggerPipelineResponse.
	TriggerPipeline(ctx context.Context, in *TriggerPipelineRequest, opts ...grpc.CallOption) (*TriggerPipelineResponse, error)
	// TriggerPipelineBinaryFileUpload method receives a
	// TriggerPipelineBinaryFileUploadRequest message and returns a
	// TriggerPipelineBinaryFileUploadResponse message.
	//
	// Endpoint: "POST /v1alpha/{name=pipelines/*}:trigger-multipart"
	TriggerPipelineBinaryFileUpload(ctx context.Context, opts ...grpc.CallOption) (PipelineService_TriggerPipelineBinaryFileUploadClient, error)
}

type pipelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineServiceClient(cc grpc.ClientConnInterface) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/Readiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error) {
	out := new(CreatePipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/CreatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*ListPipelineResponse, error) {
	out := new(ListPipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/ListPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error) {
	out := new(GetPipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/GetPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*UpdatePipelineResponse, error) {
	out := new(UpdatePipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/UpdatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*DeletePipelineResponse, error) {
	out := new(DeletePipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/DeletePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) LookUpPipeline(ctx context.Context, in *LookUpPipelineRequest, opts ...grpc.CallOption) (*LookUpPipelineResponse, error) {
	out := new(LookUpPipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/LookUpPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) ActivatePipeline(ctx context.Context, in *ActivatePipelineRequest, opts ...grpc.CallOption) (*ActivatePipelineResponse, error) {
	out := new(ActivatePipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/ActivatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) DeactivatePipeline(ctx context.Context, in *DeactivatePipelineRequest, opts ...grpc.CallOption) (*DeactivatePipelineResponse, error) {
	out := new(DeactivatePipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/DeactivatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) RenamePipeline(ctx context.Context, in *RenamePipelineRequest, opts ...grpc.CallOption) (*RenamePipelineResponse, error) {
	out := new(RenamePipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/RenamePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) TriggerPipeline(ctx context.Context, in *TriggerPipelineRequest, opts ...grpc.CallOption) (*TriggerPipelineResponse, error) {
	out := new(TriggerPipelineResponse)
	err := c.cc.Invoke(ctx, "/instill.pipeline.v1alpha.PipelineService/TriggerPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) TriggerPipelineBinaryFileUpload(ctx context.Context, opts ...grpc.CallOption) (PipelineService_TriggerPipelineBinaryFileUploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &PipelineService_ServiceDesc.Streams[0], "/instill.pipeline.v1alpha.PipelineService/TriggerPipelineBinaryFileUpload", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineServiceTriggerPipelineBinaryFileUploadClient{stream}
	return x, nil
}

type PipelineService_TriggerPipelineBinaryFileUploadClient interface {
	Send(*TriggerPipelineBinaryFileUploadRequest) error
	CloseAndRecv() (*TriggerPipelineBinaryFileUploadResponse, error)
	grpc.ClientStream
}

type pipelineServiceTriggerPipelineBinaryFileUploadClient struct {
	grpc.ClientStream
}

func (x *pipelineServiceTriggerPipelineBinaryFileUploadClient) Send(m *TriggerPipelineBinaryFileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipelineServiceTriggerPipelineBinaryFileUploadClient) CloseAndRecv() (*TriggerPipelineBinaryFileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TriggerPipelineBinaryFileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipelineServiceServer is the server API for PipelineService service.
// All implementations should embed UnimplementedPipelineServiceServer
// for forward compatibility
type PipelineServiceServer interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// CreatePipeline method receives a CreatePipelineRequest message and returns
	// a CreatePipelineResponse message.
	CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error)
	// ListPipeline method receives a ListPipelineRequest message and returns a
	// ListPipelineResponse message.
	ListPipeline(context.Context, *ListPipelineRequest) (*ListPipelineResponse, error)
	// GetPipeline method receives a GetPipelineRequest message and returns a
	// GetPipelineResponse message.
	GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error)
	// UpdatePipeline method receives a UpdatePipelineRequest message and returns
	// a UpdatePipelineResponse message.
	UpdatePipeline(context.Context, *UpdatePipelineRequest) (*UpdatePipelineResponse, error)
	// DeletePipeline method receives a DeletePipelineRequest message and returns
	// a DeletePipelineResponse message.
	DeletePipeline(context.Context, *DeletePipelineRequest) (*DeletePipelineResponse, error)
	// LookUpPipeline method receives a LookUpPipelineRequest message and returns
	// a LookUpPipelineResponse
	LookUpPipeline(context.Context, *LookUpPipelineRequest) (*LookUpPipelineResponse, error)
	// Activate a pipeline.
	// The `state` of the pipeline after activating is `ACTIVE`.
	// ActivatePipeline` can be called on Pipelines in the state `INACTIVE`;
	// Pipelines in a different state (including `ACTIVE`) returns an error.
	ActivatePipeline(context.Context, *ActivatePipelineRequest) (*ActivatePipelineResponse, error)
	// Deactivate a pipeline.
	// The `state` of the pipeline after inactivating is `INACTIVE`.
	// DeactivatePipeline` can be called on Pipelines in the state `ACTIVE`;
	// Pipelines in a different state (including `INACTIVE`) returns an error.
	DeactivatePipeline(context.Context, *DeactivatePipelineRequest) (*DeactivatePipelineResponse, error)
	// RenamePipeline method receives a RenamePipelineRequest message and returns
	// a RenamePipelineResponse message.
	RenamePipeline(context.Context, *RenamePipelineRequest) (*RenamePipelineResponse, error)
	// TriggerPipeline method receives a TriggerPipelineRequest message and
	// returns a TriggerPipelineResponse.
	TriggerPipeline(context.Context, *TriggerPipelineRequest) (*TriggerPipelineResponse, error)
	// TriggerPipelineBinaryFileUpload method receives a
	// TriggerPipelineBinaryFileUploadRequest message and returns a
	// TriggerPipelineBinaryFileUploadResponse message.
	//
	// Endpoint: "POST /v1alpha/{name=pipelines/*}:trigger-multipart"
	TriggerPipelineBinaryFileUpload(PipelineService_TriggerPipelineBinaryFileUploadServer) error
}

// UnimplementedPipelineServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPipelineServiceServer struct {
}

func (UnimplementedPipelineServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedPipelineServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (UnimplementedPipelineServiceServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) ListPipeline(context.Context, *ListPipelineRequest) (*ListPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) UpdatePipeline(context.Context, *UpdatePipelineRequest) (*UpdatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) DeletePipeline(context.Context, *DeletePipelineRequest) (*DeletePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) LookUpPipeline(context.Context, *LookUpPipelineRequest) (*LookUpPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) ActivatePipeline(context.Context, *ActivatePipelineRequest) (*ActivatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) DeactivatePipeline(context.Context, *DeactivatePipelineRequest) (*DeactivatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) RenamePipeline(context.Context, *RenamePipelineRequest) (*RenamePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenamePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) TriggerPipeline(context.Context, *TriggerPipelineRequest) (*TriggerPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) TriggerPipelineBinaryFileUpload(PipelineService_TriggerPipelineBinaryFileUploadServer) error {
	return status.Errorf(codes.Unimplemented, "method TriggerPipelineBinaryFileUpload not implemented")
}

// UnsafePipelineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineServiceServer will
// result in compilation errors.
type UnsafePipelineServiceServer interface {
	mustEmbedUnimplementedPipelineServiceServer()
}

func RegisterPipelineServiceServer(s grpc.ServiceRegistrar, srv PipelineServiceServer) {
	s.RegisterService(&PipelineService_ServiceDesc, srv)
}

func _PipelineService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/Readiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_ListPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).ListPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/ListPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).ListPipeline(ctx, req.(*ListPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/GetPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipeline(ctx, req.(*GetPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_UpdatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).UpdatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/UpdatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).UpdatePipeline(ctx, req.(*UpdatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/DeletePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_LookUpPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).LookUpPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/LookUpPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).LookUpPipeline(ctx, req.(*LookUpPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_ActivatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).ActivatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/ActivatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).ActivatePipeline(ctx, req.(*ActivatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_DeactivatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).DeactivatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/DeactivatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).DeactivatePipeline(ctx, req.(*DeactivatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_RenamePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenamePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).RenamePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/RenamePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).RenamePipeline(ctx, req.(*RenamePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_TriggerPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).TriggerPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.pipeline.v1alpha.PipelineService/TriggerPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).TriggerPipeline(ctx, req.(*TriggerPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_TriggerPipelineBinaryFileUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipelineServiceServer).TriggerPipelineBinaryFileUpload(&pipelineServiceTriggerPipelineBinaryFileUploadServer{stream})
}

type PipelineService_TriggerPipelineBinaryFileUploadServer interface {
	SendAndClose(*TriggerPipelineBinaryFileUploadResponse) error
	Recv() (*TriggerPipelineBinaryFileUploadRequest, error)
	grpc.ServerStream
}

type pipelineServiceTriggerPipelineBinaryFileUploadServer struct {
	grpc.ServerStream
}

func (x *pipelineServiceTriggerPipelineBinaryFileUploadServer) SendAndClose(m *TriggerPipelineBinaryFileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipelineServiceTriggerPipelineBinaryFileUploadServer) Recv() (*TriggerPipelineBinaryFileUploadRequest, error) {
	m := new(TriggerPipelineBinaryFileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipelineService_ServiceDesc is the grpc.ServiceDesc for PipelineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "instill.pipeline.v1alpha.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _PipelineService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _PipelineService_Readiness_Handler,
		},
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineService_CreatePipeline_Handler,
		},
		{
			MethodName: "ListPipeline",
			Handler:    _PipelineService_ListPipeline_Handler,
		},
		{
			MethodName: "GetPipeline",
			Handler:    _PipelineService_GetPipeline_Handler,
		},
		{
			MethodName: "UpdatePipeline",
			Handler:    _PipelineService_UpdatePipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineService_DeletePipeline_Handler,
		},
		{
			MethodName: "LookUpPipeline",
			Handler:    _PipelineService_LookUpPipeline_Handler,
		},
		{
			MethodName: "ActivatePipeline",
			Handler:    _PipelineService_ActivatePipeline_Handler,
		},
		{
			MethodName: "DeactivatePipeline",
			Handler:    _PipelineService_DeactivatePipeline_Handler,
		},
		{
			MethodName: "RenamePipeline",
			Handler:    _PipelineService_RenamePipeline_Handler,
		},
		{
			MethodName: "TriggerPipeline",
			Handler:    _PipelineService_TriggerPipeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TriggerPipelineBinaryFileUpload",
			Handler:       _PipelineService_TriggerPipelineBinaryFileUpload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "instill/pipeline/v1alpha/pipeline_service.proto",
}
