// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.3
// source: api/articles/articles.proto

package articles

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
	Articles_CreateArticles_FullMethodName            = "/api.articles.Articles/CreateArticles"
	Articles_UpdateArticles_FullMethodName            = "/api.articles.Articles/UpdateArticles"
	Articles_DeleteArticles_FullMethodName            = "/api.articles.Articles/DeleteArticles"
	Articles_GetArticlesInSameCategory_FullMethodName = "/api.articles.Articles/GetArticlesInSameCategory"
	Articles_GetArticlesByCidAndUid_FullMethodName    = "/api.articles.Articles/GetArticlesByCidAndUid"
	Articles_GetSingleArticle_FullMethodName          = "/api.articles.Articles/GetSingleArticle"
)

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesClient interface {
	CreateArticles(ctx context.Context, in *CreateArticlesRequest, opts ...grpc.CallOption) (*CreateArticlesReply, error)
	UpdateArticles(ctx context.Context, in *UpdateArticlesRequest, opts ...grpc.CallOption) (*UpdateArticlesReply, error)
	DeleteArticles(ctx context.Context, in *DeleteArticlesRequest, opts ...grpc.CallOption) (*DeleteArticlesReply, error)
	GetArticlesInSameCategory(ctx context.Context, in *GetArticlesInSameCategoryRequest, opts ...grpc.CallOption) (*GetArticlesInSameCategoryReply, error)
	GetArticlesByCidAndUid(ctx context.Context, in *GetArticlesByCidAndUidRequest, opts ...grpc.CallOption) (*GetArticlesByCidAndUidReply, error)
	GetSingleArticle(ctx context.Context, in *GetSingleArticleRequest, opts ...grpc.CallOption) (*GetSingleArticleReply, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) CreateArticles(ctx context.Context, in *CreateArticlesRequest, opts ...grpc.CallOption) (*CreateArticlesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateArticlesReply)
	err := c.cc.Invoke(ctx, Articles_CreateArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) UpdateArticles(ctx context.Context, in *UpdateArticlesRequest, opts ...grpc.CallOption) (*UpdateArticlesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateArticlesReply)
	err := c.cc.Invoke(ctx, Articles_UpdateArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) DeleteArticles(ctx context.Context, in *DeleteArticlesRequest, opts ...grpc.CallOption) (*DeleteArticlesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteArticlesReply)
	err := c.cc.Invoke(ctx, Articles_DeleteArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetArticlesInSameCategory(ctx context.Context, in *GetArticlesInSameCategoryRequest, opts ...grpc.CallOption) (*GetArticlesInSameCategoryReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesInSameCategoryReply)
	err := c.cc.Invoke(ctx, Articles_GetArticlesInSameCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetArticlesByCidAndUid(ctx context.Context, in *GetArticlesByCidAndUidRequest, opts ...grpc.CallOption) (*GetArticlesByCidAndUidReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesByCidAndUidReply)
	err := c.cc.Invoke(ctx, Articles_GetArticlesByCidAndUid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetSingleArticle(ctx context.Context, in *GetSingleArticleRequest, opts ...grpc.CallOption) (*GetSingleArticleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSingleArticleReply)
	err := c.cc.Invoke(ctx, Articles_GetSingleArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility.
type ArticlesServer interface {
	CreateArticles(context.Context, *CreateArticlesRequest) (*CreateArticlesReply, error)
	UpdateArticles(context.Context, *UpdateArticlesRequest) (*UpdateArticlesReply, error)
	DeleteArticles(context.Context, *DeleteArticlesRequest) (*DeleteArticlesReply, error)
	GetArticlesInSameCategory(context.Context, *GetArticlesInSameCategoryRequest) (*GetArticlesInSameCategoryReply, error)
	GetArticlesByCidAndUid(context.Context, *GetArticlesByCidAndUidRequest) (*GetArticlesByCidAndUidReply, error)
	GetSingleArticle(context.Context, *GetSingleArticleRequest) (*GetSingleArticleReply, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedArticlesServer struct{}

func (UnimplementedArticlesServer) CreateArticles(context.Context, *CreateArticlesRequest) (*CreateArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticles not implemented")
}
func (UnimplementedArticlesServer) UpdateArticles(context.Context, *UpdateArticlesRequest) (*UpdateArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticles not implemented")
}
func (UnimplementedArticlesServer) DeleteArticles(context.Context, *DeleteArticlesRequest) (*DeleteArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticles not implemented")
}
func (UnimplementedArticlesServer) GetArticlesInSameCategory(context.Context, *GetArticlesInSameCategoryRequest) (*GetArticlesInSameCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesInSameCategory not implemented")
}
func (UnimplementedArticlesServer) GetArticlesByCidAndUid(context.Context, *GetArticlesByCidAndUidRequest) (*GetArticlesByCidAndUidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByCidAndUid not implemented")
}
func (UnimplementedArticlesServer) GetSingleArticle(context.Context, *GetSingleArticleRequest) (*GetSingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleArticle not implemented")
}
func (UnimplementedArticlesServer) mustEmbedUnimplementedArticlesServer() {}
func (UnimplementedArticlesServer) testEmbeddedByValue()                  {}

// UnsafeArticlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesServer will
// result in compilation errors.
type UnsafeArticlesServer interface {
	mustEmbedUnimplementedArticlesServer()
}

func RegisterArticlesServer(s grpc.ServiceRegistrar, srv ArticlesServer) {
	// If the following call pancis, it indicates UnimplementedArticlesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Articles_ServiceDesc, srv)
}

func _Articles_CreateArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).CreateArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Articles_CreateArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).CreateArticles(ctx, req.(*CreateArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_UpdateArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).UpdateArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Articles_UpdateArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).UpdateArticles(ctx, req.(*UpdateArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_DeleteArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).DeleteArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Articles_DeleteArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).DeleteArticles(ctx, req.(*DeleteArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_GetArticlesInSameCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesInSameCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetArticlesInSameCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Articles_GetArticlesInSameCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetArticlesInSameCategory(ctx, req.(*GetArticlesInSameCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_GetArticlesByCidAndUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesByCidAndUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetArticlesByCidAndUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Articles_GetArticlesByCidAndUid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetArticlesByCidAndUid(ctx, req.(*GetArticlesByCidAndUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_GetSingleArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetSingleArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Articles_GetSingleArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetSingleArticle(ctx, req.(*GetSingleArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Articles_ServiceDesc is the grpc.ServiceDesc for Articles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Articles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.articles.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticles",
			Handler:    _Articles_CreateArticles_Handler,
		},
		{
			MethodName: "UpdateArticles",
			Handler:    _Articles_UpdateArticles_Handler,
		},
		{
			MethodName: "DeleteArticles",
			Handler:    _Articles_DeleteArticles_Handler,
		},
		{
			MethodName: "GetArticlesInSameCategory",
			Handler:    _Articles_GetArticlesInSameCategory_Handler,
		},
		{
			MethodName: "GetArticlesByCidAndUid",
			Handler:    _Articles_GetArticlesByCidAndUid_Handler,
		},
		{
			MethodName: "GetSingleArticle",
			Handler:    _Articles_GetSingleArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/articles/articles.proto",
}
