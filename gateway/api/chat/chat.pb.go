// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: api/chat/chat.proto

package chat

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

type AIChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg         []*AIChatRequest_Message `protobuf:"bytes,1,rep,name=Msg,proto3" json:"Msg,omitempty"`
	Temperature float32                  `protobuf:"fixed32,2,opt,name=Temperature,proto3" json:"Temperature,omitempty"` // Controls the randomness of the output; higher values produce more random results.
	TopK        uint32                   `protobuf:"varint,3,opt,name=TopK,proto3" json:"TopK,omitempty"`                // Limits the AI to choose from the top 'k' most probable words. Lower values make responses more focused; higher values introduce more variety and potential surprises.
	ModelKind   string                   `protobuf:"bytes,4,opt,name=ModelKind,proto3" json:"ModelKind,omitempty"`
}

func (x *AIChatRequest) Reset() {
	*x = AIChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIChatRequest) ProtoMessage() {}

func (x *AIChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIChatRequest.ProtoReflect.Descriptor instead.
func (*AIChatRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *AIChatRequest) GetMsg() []*AIChatRequest_Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *AIChatRequest) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *AIChatRequest) GetTopK() uint32 {
	if x != nil {
		return x.TopK
	}
	return 0
}

func (x *AIChatRequest) GetModelKind() string {
	if x != nil {
		return x.ModelKind
	}
	return ""
}

type AIChatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *AIChatReply) Reset() {
	*x = AIChatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIChatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIChatReply) ProtoMessage() {}

func (x *AIChatReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIChatReply.ProtoReflect.Descriptor instead.
func (*AIChatReply) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *AIChatReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AIPaintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt    string  `protobuf:"bytes,1,opt,name=Prompt,proto3" json:"Prompt,omitempty"`
	Height    uint32  `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Width     uint32  `protobuf:"varint,3,opt,name=Width,proto3" json:"Width,omitempty"`
	Guidance  float32 `protobuf:"fixed32,4,opt,name=Guidance,proto3" json:"Guidance,omitempty"` // Controls how closely the generated image should adhere to the prompt; higher values make the image more aligned with the prompt
	ModelKind string  `protobuf:"bytes,5,opt,name=ModelKind,proto3" json:"ModelKind,omitempty"`
}

func (x *AIPaintRequest) Reset() {
	*x = AIPaintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIPaintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIPaintRequest) ProtoMessage() {}

func (x *AIPaintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIPaintRequest.ProtoReflect.Descriptor instead.
func (*AIPaintRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *AIPaintRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *AIPaintRequest) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *AIPaintRequest) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *AIPaintRequest) GetGuidance() float32 {
	if x != nil {
		return x.Guidance
	}
	return 0
}

func (x *AIPaintRequest) GetModelKind() string {
	if x != nil {
		return x.ModelKind
	}
	return ""
}

type AIPaintReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImgBinary []byte `protobuf:"bytes,1,opt,name=ImgBinary,proto3" json:"ImgBinary,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *AIPaintReply) Reset() {
	*x = AIPaintReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIPaintReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIPaintReply) ProtoMessage() {}

func (x *AIPaintReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIPaintReply.ProtoReflect.Descriptor instead.
func (*AIPaintReply) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *AIPaintReply) GetImgBinary() []byte {
	if x != nil {
		return x.ImgBinary
	}
	return nil
}

func (x *AIPaintReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AISummarizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleText []byte `protobuf:"bytes,1,opt,name=ArticleText,proto3" json:"ArticleText,omitempty"`
}

func (x *AISummarizationRequest) Reset() {
	*x = AISummarizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AISummarizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AISummarizationRequest) ProtoMessage() {}

func (x *AISummarizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AISummarizationRequest.ProtoReflect.Descriptor instead.
func (*AISummarizationRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *AISummarizationRequest) GetArticleText() []byte {
	if x != nil {
		return x.ArticleText
	}
	return nil
}

type AISummarizationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextAbstract string `protobuf:"bytes,1,opt,name=TextAbstract,proto3" json:"TextAbstract,omitempty"`
}

func (x *AISummarizationReply) Reset() {
	*x = AISummarizationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AISummarizationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AISummarizationReply) ProtoMessage() {}

func (x *AISummarizationReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AISummarizationReply.ProtoReflect.Descriptor instead.
func (*AISummarizationReply) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{5}
}

func (x *AISummarizationReply) GetTextAbstract() string {
	if x != nil {
		return x.TextAbstract
	}
	return ""
}

type AIChatRequest_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AIChatRequest_Message) Reset() {
	*x = AIChatRequest_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIChatRequest_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIChatRequest_Message) ProtoMessage() {}

func (x *AIChatRequest_Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIChatRequest_Message.ProtoReflect.Descriptor instead.
func (*AIChatRequest_Message) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AIChatRequest_Message) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *AIChatRequest_Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_api_chat_chat_proto protoreflect.FileDescriptor

var file_api_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x22,
	0xcf, 0x01, 0x0a, 0x0d, 0x41, 0x49, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x49, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x70, 0x4b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x54, 0x6f, 0x70, 0x4b, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0x37, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x49, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x41, 0x49, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x47,
	0x75, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x47,
	0x75, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4b, 0x69, 0x6e, 0x64, 0x22, 0x3e, 0x0a, 0x0c, 0x41, 0x49, 0x50, 0x61, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x67, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x49, 0x6d, 0x67, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x3a, 0x0a, 0x16, 0x41, 0x49, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x22, 0x3a, 0x0a, 0x14, 0x41, 0x49, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x65, 0x78,
	0x74, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x54, 0x65, 0x78, 0x74, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x32, 0xe8, 0x01,
	0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x46, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x49, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x49, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x41, 0x49, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x3b,
	0x0a, 0x07, 0x41, 0x49, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x49, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41,
	0x49, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x15, 0x41,
	0x49, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x41, 0x49, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x41, 0x49, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x41, 0x49, 0x43, 0x68,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x3b, 0x63, 0x68, 0x61, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chat_chat_proto_rawDescOnce sync.Once
	file_api_chat_chat_proto_rawDescData = file_api_chat_chat_proto_rawDesc
)

func file_api_chat_chat_proto_rawDescGZIP() []byte {
	file_api_chat_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_chat_proto_rawDescData)
	})
	return file_api_chat_chat_proto_rawDescData
}

var file_api_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_chat_chat_proto_goTypes = []any{
	(*AIChatRequest)(nil),          // 0: api.chat.AIChatRequest
	(*AIChatReply)(nil),            // 1: api.chat.AIChatReply
	(*AIPaintRequest)(nil),         // 2: api.chat.AIPaintRequest
	(*AIPaintReply)(nil),           // 3: api.chat.AIPaintReply
	(*AISummarizationRequest)(nil), // 4: api.chat.AISummarizationRequest
	(*AISummarizationReply)(nil),   // 5: api.chat.AISummarizationReply
	(*AIChatRequest_Message)(nil),  // 6: api.chat.AIChatRequest.Message
}
var file_api_chat_chat_proto_depIdxs = []int32{
	6, // 0: api.chat.AIChatRequest.Msg:type_name -> api.chat.AIChatRequest.Message
	0, // 1: api.chat.Chat.ServerStreamAIChat:input_type -> api.chat.AIChatRequest
	2, // 2: api.chat.Chat.AIPaint:input_type -> api.chat.AIPaintRequest
	4, // 3: api.chat.Chat.AISummarizationStream:input_type -> api.chat.AISummarizationRequest
	1, // 4: api.chat.Chat.ServerStreamAIChat:output_type -> api.chat.AIChatReply
	3, // 5: api.chat.Chat.AIPaint:output_type -> api.chat.AIPaintReply
	5, // 6: api.chat.Chat.AISummarizationStream:output_type -> api.chat.AISummarizationReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_chat_chat_proto_init() }
func file_api_chat_chat_proto_init() {
	if File_api_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chat_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AIChatRequest); i {
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
		file_api_chat_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AIChatReply); i {
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
		file_api_chat_chat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AIPaintRequest); i {
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
		file_api_chat_chat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AIPaintReply); i {
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
		file_api_chat_chat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AISummarizationRequest); i {
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
		file_api_chat_chat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AISummarizationReply); i {
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
		file_api_chat_chat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AIChatRequest_Message); i {
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
			RawDescriptor: file_api_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_chat_proto_depIdxs,
		MessageInfos:      file_api_chat_chat_proto_msgTypes,
	}.Build()
	File_api_chat_chat_proto = out.File
	file_api_chat_chat_proto_rawDesc = nil
	file_api_chat_chat_proto_goTypes = nil
	file_api_chat_chat_proto_depIdxs = nil
}
