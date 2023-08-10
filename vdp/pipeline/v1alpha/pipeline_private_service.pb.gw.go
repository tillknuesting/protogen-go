// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: vdp/pipeline/v1alpha/pipeline_private_service.proto

/*
Package pipelinev1alpha is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pipelinev1alpha

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_PipelinePrivateService_ListPipelinesAdmin_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_PipelinePrivateService_ListPipelinesAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListPipelinesAdminRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PipelinePrivateService_ListPipelinesAdmin_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListPipelinesAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PipelinePrivateService_ListPipelinesAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListPipelinesAdminRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PipelinePrivateService_ListPipelinesAdmin_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListPipelinesAdmin(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_PipelinePrivateService_LookUpPipelineAdmin_0 = &utilities.DoubleArray{Encoding: map[string]int{"permalink": 0}, Base: []int{1, 2, 0, 0}, Check: []int{0, 1, 2, 2}}
)

func request_PipelinePrivateService_LookUpPipelineAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LookUpPipelineAdminRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["permalink"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "permalink")
	}

	protoReq.Permalink, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "permalink", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PipelinePrivateService_LookUpPipelineAdmin_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.LookUpPipelineAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PipelinePrivateService_LookUpPipelineAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LookUpPipelineAdminRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["permalink"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "permalink")
	}

	protoReq.Permalink, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "permalink", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PipelinePrivateService_LookUpPipelineAdmin_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.LookUpPipelineAdmin(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0 = &utilities.DoubleArray{Encoding: map[string]int{"permalink": 0}, Base: []int{1, 2, 0, 0}, Check: []int{0, 1, 2, 2}}
)

func request_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LookUpOperatorDefinitionAdminRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["permalink"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "permalink")
	}

	protoReq.Permalink, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "permalink", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.LookUpOperatorDefinitionAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LookUpOperatorDefinitionAdminRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["permalink"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "permalink")
	}

	protoReq.Permalink, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "permalink", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.LookUpOperatorDefinitionAdmin(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterPipelinePrivateServiceHandlerServer registers the http handlers for service PipelinePrivateService to "mux".
// UnaryRPC     :call PipelinePrivateServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterPipelinePrivateServiceHandlerFromEndpoint instead.
func RegisterPipelinePrivateServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PipelinePrivateServiceServer) error {

	mux.Handle("GET", pattern_PipelinePrivateService_ListPipelinesAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/vdp.pipeline.v1alpha.PipelinePrivateService/ListPipelinesAdmin", runtime.WithHTTPPathPattern("/v1alpha/admin/pipelines"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PipelinePrivateService_LookUpPipelineAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/vdp.pipeline.v1alpha.PipelinePrivateService/LookUpPipelineAdmin", runtime.WithHTTPPathPattern("/v1alpha/admin/{permalink=pipelines/*}/lookUp"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/vdp.pipeline.v1alpha.PipelinePrivateService/LookUpOperatorDefinitionAdmin", runtime.WithHTTPPathPattern("/v1alpha/admin/{permalink=operator-definitions/*}/lookUp"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterPipelinePrivateServiceHandlerFromEndpoint is same as RegisterPipelinePrivateServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPipelinePrivateServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPipelinePrivateServiceHandler(ctx, mux, conn)
}

// RegisterPipelinePrivateServiceHandler registers the http handlers for service PipelinePrivateService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPipelinePrivateServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPipelinePrivateServiceHandlerClient(ctx, mux, NewPipelinePrivateServiceClient(conn))
}

// RegisterPipelinePrivateServiceHandlerClient registers the http handlers for service PipelinePrivateService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PipelinePrivateServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PipelinePrivateServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PipelinePrivateServiceClient" to call the correct interceptors.
func RegisterPipelinePrivateServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PipelinePrivateServiceClient) error {

	mux.Handle("GET", pattern_PipelinePrivateService_ListPipelinesAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/vdp.pipeline.v1alpha.PipelinePrivateService/ListPipelinesAdmin", runtime.WithHTTPPathPattern("/v1alpha/admin/pipelines"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PipelinePrivateService_LookUpPipelineAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/vdp.pipeline.v1alpha.PipelinePrivateService/LookUpPipelineAdmin", runtime.WithHTTPPathPattern("/v1alpha/admin/{permalink=pipelines/*}/lookUp"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/vdp.pipeline.v1alpha.PipelinePrivateService/LookUpOperatorDefinitionAdmin", runtime.WithHTTPPathPattern("/v1alpha/admin/{permalink=operator-definitions/*}/lookUp"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PipelinePrivateService_ListPipelinesAdmin_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1alpha", "admin", "pipelines"}, ""))

	pattern_PipelinePrivateService_LookUpPipelineAdmin_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 2, 5, 3, 2, 4}, []string{"v1alpha", "admin", "pipelines", "permalink", "lookUp"}, ""))

	pattern_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 2, 5, 3, 2, 4}, []string{"v1alpha", "admin", "operator-definitions", "permalink", "lookUp"}, ""))
)

var (
	forward_PipelinePrivateService_ListPipelinesAdmin_0 = runtime.ForwardResponseMessage

	forward_PipelinePrivateService_LookUpPipelineAdmin_0 = runtime.ForwardResponseMessage

	forward_PipelinePrivateService_LookUpOperatorDefinitionAdmin_0 = runtime.ForwardResponseMessage
)
