// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.3
// source: api/stat_user/stat_user.proto

package stat_user

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
	StatUser_GetUserSevenDaysStat_FullMethodName = "/api.stat_user.StatUser/GetUserSevenDaysStat"
	StatUser_SetUserStatInfo_FullMethodName      = "/api.stat_user.StatUser/SetUserStatInfo"
)

// StatUserClient is the client API for StatUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatUserClient interface {
	GetUserSevenDaysStat(ctx context.Context, in *GetUserSevenDaysStatRequest, opts ...grpc.CallOption) (*GetUserSevenDaysStatReply, error)
	SetUserStatInfo(ctx context.Context, in *SetUserStatInfoRequest, opts ...grpc.CallOption) (*SetUserStatInfoReply, error)
}

type statUserClient struct {
	cc grpc.ClientConnInterface
}

func NewStatUserClient(cc grpc.ClientConnInterface) StatUserClient {
	return &statUserClient{cc}
}

func (c *statUserClient) GetUserSevenDaysStat(ctx context.Context, in *GetUserSevenDaysStatRequest, opts ...grpc.CallOption) (*GetUserSevenDaysStatReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserSevenDaysStatReply)
	err := c.cc.Invoke(ctx, StatUser_GetUserSevenDaysStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statUserClient) SetUserStatInfo(ctx context.Context, in *SetUserStatInfoRequest, opts ...grpc.CallOption) (*SetUserStatInfoReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetUserStatInfoReply)
	err := c.cc.Invoke(ctx, StatUser_SetUserStatInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatUserServer is the server API for StatUser service.
// All implementations must embed UnimplementedStatUserServer
// for forward compatibility.
type StatUserServer interface {
	GetUserSevenDaysStat(context.Context, *GetUserSevenDaysStatRequest) (*GetUserSevenDaysStatReply, error)
	SetUserStatInfo(context.Context, *SetUserStatInfoRequest) (*SetUserStatInfoReply, error)
	mustEmbedUnimplementedStatUserServer()
}

// UnimplementedStatUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatUserServer struct{}

func (UnimplementedStatUserServer) GetUserSevenDaysStat(context.Context, *GetUserSevenDaysStatRequest) (*GetUserSevenDaysStatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSevenDaysStat not implemented")
}
func (UnimplementedStatUserServer) SetUserStatInfo(context.Context, *SetUserStatInfoRequest) (*SetUserStatInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserStatInfo not implemented")
}
func (UnimplementedStatUserServer) mustEmbedUnimplementedStatUserServer() {}
func (UnimplementedStatUserServer) testEmbeddedByValue()                  {}

// UnsafeStatUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatUserServer will
// result in compilation errors.
type UnsafeStatUserServer interface {
	mustEmbedUnimplementedStatUserServer()
}

func RegisterStatUserServer(s grpc.ServiceRegistrar, srv StatUserServer) {
	// If the following call pancis, it indicates UnimplementedStatUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StatUser_ServiceDesc, srv)
}

func _StatUser_GetUserSevenDaysStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSevenDaysStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatUserServer).GetUserSevenDaysStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatUser_GetUserSevenDaysStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatUserServer).GetUserSevenDaysStat(ctx, req.(*GetUserSevenDaysStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatUser_SetUserStatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserStatInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatUserServer).SetUserStatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatUser_SetUserStatInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatUserServer).SetUserStatInfo(ctx, req.(*SetUserStatInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatUser_ServiceDesc is the grpc.ServiceDesc for StatUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.stat_user.StatUser",
	HandlerType: (*StatUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserSevenDaysStat",
			Handler:    _StatUser_GetUserSevenDaysStat_Handler,
		},
		{
			MethodName: "SetUserStatInfo",
			Handler:    _StatUser_SetUserStatInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/stat_user/stat_user.proto",
}
