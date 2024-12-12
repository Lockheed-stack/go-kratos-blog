// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.3
// source: api/AI_Cloudflare/AI_Cloudflare.proto

package AI_Cloudflare

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AICloudflare_StreamAISummarization_FullMethodName = "/api.AI_Cloudflare.AICloudflare/StreamAISummarization"
	AICloudflare_StreamAIChat_FullMethodName          = "/api.AI_Cloudflare.AICloudflare/StreamAIChat"
	AICloudflare_AIPaint_FullMethodName               = "/api.AI_Cloudflare.AICloudflare/AIPaint"
)

// AICloudflareClient is the client API for AICloudflare service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AICloudflareClient interface {
	StreamAISummarization(ctx context.Context, in *AISummarizationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AISummarizationReply], error)
	StreamAIChat(ctx context.Context, in *AIChatRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AIChatReply], error)
	AIPaint(ctx context.Context, in *AIPaintRequest, opts ...grpc.CallOption) (*AIPaintReply, error)
}

type aICloudflareClient struct {
	cc grpc.ClientConnInterface
}

func NewAICloudflareClient(cc grpc.ClientConnInterface) AICloudflareClient {
	return &aICloudflareClient{cc}
}

func (c *aICloudflareClient) StreamAISummarization(ctx context.Context, in *AISummarizationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AISummarizationReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AICloudflare_ServiceDesc.Streams[0], AICloudflare_StreamAISummarization_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AISummarizationRequest, AISummarizationReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AICloudflare_StreamAISummarizationClient = grpc.ServerStreamingClient[AISummarizationReply]

func (c *aICloudflareClient) StreamAIChat(ctx context.Context, in *AIChatRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AIChatReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AICloudflare_ServiceDesc.Streams[1], AICloudflare_StreamAIChat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AIChatRequest, AIChatReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AICloudflare_StreamAIChatClient = grpc.ServerStreamingClient[AIChatReply]

func (c *aICloudflareClient) AIPaint(ctx context.Context, in *AIPaintRequest, opts ...grpc.CallOption) (*AIPaintReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AIPaintReply)
	err := c.cc.Invoke(ctx, AICloudflare_AIPaint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AICloudflareServer is the server API for AICloudflare service.
// All implementations must embed UnimplementedAICloudflareServer
// for forward compatibility.
type AICloudflareServer interface {
	StreamAISummarization(*AISummarizationRequest, grpc.ServerStreamingServer[AISummarizationReply]) error
	StreamAIChat(*AIChatRequest, grpc.ServerStreamingServer[AIChatReply]) error
	AIPaint(context.Context, *AIPaintRequest) (*AIPaintReply, error)
	mustEmbedUnimplementedAICloudflareServer()
}

// UnimplementedAICloudflareServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAICloudflareServer struct{}

func (UnimplementedAICloudflareServer) StreamAISummarization(*AISummarizationRequest, grpc.ServerStreamingServer[AISummarizationReply]) error {
	return status.Errorf(codes.Unimplemented, "method StreamAISummarization not implemented")
}
func (UnimplementedAICloudflareServer) StreamAIChat(*AIChatRequest, grpc.ServerStreamingServer[AIChatReply]) error {
	return status.Errorf(codes.Unimplemented, "method StreamAIChat not implemented")
}
func (UnimplementedAICloudflareServer) AIPaint(context.Context, *AIPaintRequest) (*AIPaintReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AIPaint not implemented")
}
func (UnimplementedAICloudflareServer) mustEmbedUnimplementedAICloudflareServer() {}
func (UnimplementedAICloudflareServer) testEmbeddedByValue()                      {}

// UnsafeAICloudflareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AICloudflareServer will
// result in compilation errors.
type UnsafeAICloudflareServer interface {
	mustEmbedUnimplementedAICloudflareServer()
}

func RegisterAICloudflareServer(s grpc.ServiceRegistrar, srv AICloudflareServer) {
	// If the following call pancis, it indicates UnimplementedAICloudflareServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AICloudflare_ServiceDesc, srv)
}

func _AICloudflare_StreamAISummarization_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AISummarizationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AICloudflareServer).StreamAISummarization(m, &grpc.GenericServerStream[AISummarizationRequest, AISummarizationReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AICloudflare_StreamAISummarizationServer = grpc.ServerStreamingServer[AISummarizationReply]

func _AICloudflare_StreamAIChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AIChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AICloudflareServer).StreamAIChat(m, &grpc.GenericServerStream[AIChatRequest, AIChatReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AICloudflare_StreamAIChatServer = grpc.ServerStreamingServer[AIChatReply]

func _AICloudflare_AIPaint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AIPaintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AICloudflareServer).AIPaint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AICloudflare_AIPaint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AICloudflareServer).AIPaint(ctx, req.(*AIPaintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AICloudflare_ServiceDesc is the grpc.ServiceDesc for AICloudflare service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AICloudflare_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AI_Cloudflare.AICloudflare",
	HandlerType: (*AICloudflareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AIPaint",
			Handler:    _AICloudflare_AIPaint_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAISummarization",
			Handler:       _AICloudflare_StreamAISummarization_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamAIChat",
			Handler:       _AICloudflare_StreamAIChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/AI_Cloudflare/AI_Cloudflare.proto",
}
