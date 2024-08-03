// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: api/articles/articles.proto

package articles

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Cid      uint32 `protobuf:"varint,2,opt,name=Cid,proto3" json:"Cid,omitempty"`
	Desc     string `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Img      string `protobuf:"bytes,5,opt,name=Img,proto3" json:"Img,omitempty"`
	PageView uint32 `protobuf:"varint,6,opt,name=PageView,proto3" json:"PageView,omitempty"`
	Uid      uint32 `protobuf:"varint,7,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *CreateArticlesRequest) Reset() {
	*x = CreateArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticlesRequest) ProtoMessage() {}

func (x *CreateArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticlesRequest.ProtoReflect.Descriptor instead.
func (*CreateArticlesRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArticlesRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticlesRequest) GetCid() uint32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *CreateArticlesRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateArticlesRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateArticlesRequest) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *CreateArticlesRequest) GetPageView() uint32 {
	if x != nil {
		return x.PageView
	}
	return 0
}

func (x *CreateArticlesRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type CreateArticlesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *CreateArticlesReply) Reset() {
	*x = CreateArticlesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticlesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticlesReply) ProtoMessage() {}

func (x *CreateArticlesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticlesReply.ProtoReflect.Descriptor instead.
func (*CreateArticlesReply) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticlesReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateArticlesReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize uint32 `protobuf:"varint,1,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	PageNum  uint32 `protobuf:"varint,2,opt,name=PageNum,proto3" json:"PageNum,omitempty"`
}

func (x *GetArticlesRequest) Reset() {
	*x = GetArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesRequest) ProtoMessage() {}

func (x *GetArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesRequest.ProtoReflect.Descriptor instead.
func (*GetArticlesRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticlesRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetArticlesRequest) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type GetArticlesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelectedArticles []*GetArticlesReply_ArticleInfo `protobuf:"bytes,1,rep,name=SelectedArticles,proto3" json:"SelectedArticles,omitempty"`
	Code             uint32                          `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetArticlesReply) Reset() {
	*x = GetArticlesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesReply) ProtoMessage() {}

func (x *GetArticlesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesReply.ProtoReflect.Descriptor instead.
func (*GetArticlesReply) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticlesReply) GetSelectedArticles() []*GetArticlesReply_ArticleInfo {
	if x != nil {
		return x.SelectedArticles
	}
	return nil
}

func (x *GetArticlesReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetSingleArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleID uint32 `protobuf:"varint,1,opt,name=ArticleID,proto3" json:"ArticleID,omitempty"`
}

func (x *GetSingleArticleRequest) Reset() {
	*x = GetSingleArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleArticleRequest) ProtoMessage() {}

func (x *GetSingleArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleArticleRequest.ProtoReflect.Descriptor instead.
func (*GetSingleArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{4}
}

func (x *GetSingleArticleRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

type GetSingleArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *GetSingleArticleReply_RespondMsg `protobuf:"bytes,1,opt,name=Article,proto3" json:"Article,omitempty"`
	Code    uint32                            `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string                            `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetSingleArticleReply) Reset() {
	*x = GetSingleArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleArticleReply) ProtoMessage() {}

func (x *GetSingleArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleArticleReply.ProtoReflect.Descriptor instead.
func (*GetSingleArticleReply) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{5}
}

func (x *GetSingleArticleReply) GetArticle() *GetSingleArticleReply_RespondMsg {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *GetSingleArticleReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetSingleArticleReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleID uint32 `protobuf:"varint,1,opt,name=ArticleID,proto3" json:"ArticleID,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Cid       uint32 `protobuf:"varint,3,opt,name=Cid,proto3" json:"Cid,omitempty"` // category id
	Desc      string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Content   string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Thumbnail string `protobuf:"bytes,6,opt,name=Thumbnail,proto3" json:"Thumbnail,omitempty"`
}

func (x *UpdateArticlesRequest) Reset() {
	*x = UpdateArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticlesRequest) ProtoMessage() {}

func (x *UpdateArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticlesRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticlesRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateArticlesRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

func (x *UpdateArticlesRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateArticlesRequest) GetCid() uint32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *UpdateArticlesRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *UpdateArticlesRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateArticlesRequest) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

type UpdateArticlesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *UpdateArticlesReply) Reset() {
	*x = UpdateArticlesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticlesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticlesReply) ProtoMessage() {}

func (x *UpdateArticlesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticlesReply.ProtoReflect.Descriptor instead.
func (*UpdateArticlesReply) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateArticlesReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateArticlesReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleID uint32 `protobuf:"varint,1,opt,name=ArticleID,proto3" json:"ArticleID,omitempty"`
}

func (x *DeleteArticlesRequest) Reset() {
	*x = DeleteArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticlesRequest) ProtoMessage() {}

func (x *DeleteArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticlesRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticlesRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteArticlesRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

type DeleteArticlesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *DeleteArticlesReply) Reset() {
	*x = DeleteArticlesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticlesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticlesReply) ProtoMessage() {}

func (x *DeleteArticlesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticlesReply.ProtoReflect.Descriptor instead.
func (*DeleteArticlesReply) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteArticlesReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteArticlesReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetArticlesReply_ArticleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   uint32 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"` // article title
	Name  string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`   // category name
	Img   string `protobuf:"bytes,4,opt,name=Img,proto3" json:"Img,omitempty"`
	Desc  string `protobuf:"bytes,5,opt,name=Desc,proto3" json:"Desc,omitempty"` // article description
}

func (x *GetArticlesReply_ArticleInfo) Reset() {
	*x = GetArticlesReply_ArticleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesReply_ArticleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesReply_ArticleInfo) ProtoMessage() {}

func (x *GetArticlesReply_ArticleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesReply_ArticleInfo.ProtoReflect.Descriptor instead.
func (*GetArticlesReply_ArticleInfo) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetArticlesReply_ArticleInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetArticlesReply_ArticleInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetArticlesReply_ArticleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetArticlesReply_ArticleInfo) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *GetArticlesReply_ArticleInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type GetSingleArticleReply_RespondMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt string `protobuf:"bytes,1,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,2,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Desc      string `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Content   string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	PageView  uint32 `protobuf:"varint,6,opt,name=PageView,proto3" json:"PageView,omitempty"`
}

func (x *GetSingleArticleReply_RespondMsg) Reset() {
	*x = GetSingleArticleReply_RespondMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_articles_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleArticleReply_RespondMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleArticleReply_RespondMsg) ProtoMessage() {}

func (x *GetSingleArticleReply_RespondMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_articles_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleArticleReply_RespondMsg.ProtoReflect.Descriptor instead.
func (*GetSingleArticleReply_RespondMsg) Descriptor() ([]byte, []int) {
	return file_api_articles_articles_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetSingleArticleReply_RespondMsg) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetSingleArticleReply_RespondMsg) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetSingleArticleReply_RespondMsg) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetSingleArticleReply_RespondMsg) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *GetSingleArticleReply_RespondMsg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetSingleArticleReply_RespondMsg) GetPageView() uint32 {
	if x != nil {
		return x.PageView
	}
	return 0
}

var File_api_articles_articles_proto protoreflect.FileDescriptor

var file_api_articles_articles_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x43, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x49,
	0x6d, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x6d, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x50, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x22, 0xef, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x56, 0x0a, 0x10, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x10, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x6f, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x49, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49,
	0x6d, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x22, 0x37, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x22,
	0xb2, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x07, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x07, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0xa8, 0x01, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x22, 0xa9, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x43, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x22, 0x3b, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x35, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x32, 0xc9, 0x03, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x58,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1c, 0x5a,
	0x1a, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_articles_articles_proto_rawDescOnce sync.Once
	file_api_articles_articles_proto_rawDescData = file_api_articles_articles_proto_rawDesc
)

func file_api_articles_articles_proto_rawDescGZIP() []byte {
	file_api_articles_articles_proto_rawDescOnce.Do(func() {
		file_api_articles_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_articles_articles_proto_rawDescData)
	})
	return file_api_articles_articles_proto_rawDescData
}

var file_api_articles_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_articles_articles_proto_goTypes = []any{
	(*CreateArticlesRequest)(nil),            // 0: api.articles.CreateArticlesRequest
	(*CreateArticlesReply)(nil),              // 1: api.articles.CreateArticlesReply
	(*GetArticlesRequest)(nil),               // 2: api.articles.GetArticlesRequest
	(*GetArticlesReply)(nil),                 // 3: api.articles.GetArticlesReply
	(*GetSingleArticleRequest)(nil),          // 4: api.articles.GetSingleArticleRequest
	(*GetSingleArticleReply)(nil),            // 5: api.articles.GetSingleArticleReply
	(*UpdateArticlesRequest)(nil),            // 6: api.articles.UpdateArticlesRequest
	(*UpdateArticlesReply)(nil),              // 7: api.articles.UpdateArticlesReply
	(*DeleteArticlesRequest)(nil),            // 8: api.articles.DeleteArticlesRequest
	(*DeleteArticlesReply)(nil),              // 9: api.articles.DeleteArticlesReply
	(*GetArticlesReply_ArticleInfo)(nil),     // 10: api.articles.GetArticlesReply.ArticleInfo
	(*GetSingleArticleReply_RespondMsg)(nil), // 11: api.articles.GetSingleArticleReply.RespondMsg
}
var file_api_articles_articles_proto_depIdxs = []int32{
	10, // 0: api.articles.GetArticlesReply.SelectedArticles:type_name -> api.articles.GetArticlesReply.ArticleInfo
	11, // 1: api.articles.GetSingleArticleReply.Article:type_name -> api.articles.GetSingleArticleReply.RespondMsg
	0,  // 2: api.articles.Articles.CreateArticles:input_type -> api.articles.CreateArticlesRequest
	6,  // 3: api.articles.Articles.UpdateArticles:input_type -> api.articles.UpdateArticlesRequest
	8,  // 4: api.articles.Articles.DeleteArticles:input_type -> api.articles.DeleteArticlesRequest
	2,  // 5: api.articles.Articles.GetArticles:input_type -> api.articles.GetArticlesRequest
	4,  // 6: api.articles.Articles.GetSingleArticle:input_type -> api.articles.GetSingleArticleRequest
	1,  // 7: api.articles.Articles.CreateArticles:output_type -> api.articles.CreateArticlesReply
	7,  // 8: api.articles.Articles.UpdateArticles:output_type -> api.articles.UpdateArticlesReply
	9,  // 9: api.articles.Articles.DeleteArticles:output_type -> api.articles.DeleteArticlesReply
	3,  // 10: api.articles.Articles.GetArticles:output_type -> api.articles.GetArticlesReply
	5,  // 11: api.articles.Articles.GetSingleArticle:output_type -> api.articles.GetSingleArticleReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_articles_articles_proto_init() }
func file_api_articles_articles_proto_init() {
	if File_api_articles_articles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_articles_articles_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateArticlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateArticlesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetArticlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetArticlesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetSingleArticleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetSingleArticleReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateArticlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateArticlesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteArticlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteArticlesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetArticlesReply_ArticleInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_articles_articles_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetSingleArticleReply_RespondMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_articles_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_articles_articles_proto_goTypes,
		DependencyIndexes: file_api_articles_articles_proto_depIdxs,
		MessageInfos:      file_api_articles_articles_proto_msgTypes,
	}.Build()
	File_api_articles_articles_proto = out.File
	file_api_articles_articles_proto_rawDesc = nil
	file_api_articles_articles_proto_goTypes = nil
	file_api_articles_articles_proto_depIdxs = nil
}
