// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/article/article.proto

package article

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateArticleRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId             int64    `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateArticleRequest) Reset()         { *m = CreateArticleRequest{} }
func (m *CreateArticleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArticleRequest) ProtoMessage()    {}
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{0}
}

func (m *CreateArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArticleRequest.Unmarshal(m, b)
}
func (m *CreateArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArticleRequest.Marshal(b, m, deterministic)
}
func (m *CreateArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArticleRequest.Merge(m, src)
}
func (m *CreateArticleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArticleRequest.Size(m)
}
func (m *CreateArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArticleRequest proto.InternalMessageInfo

func (m *CreateArticleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateArticleRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateArticleRequest) GetAuthorId() int64 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type ListArticleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArticleRequest) Reset()         { *m = ListArticleRequest{} }
func (m *ListArticleRequest) String() string { return proto.CompactTextString(m) }
func (*ListArticleRequest) ProtoMessage()    {}
func (*ListArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{1}
}

func (m *ListArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArticleRequest.Unmarshal(m, b)
}
func (m *ListArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArticleRequest.Marshal(b, m, deterministic)
}
func (m *ListArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArticleRequest.Merge(m, src)
}
func (m *ListArticleRequest) XXX_Size() int {
	return xxx_messageInfo_ListArticleRequest.Size(m)
}
func (m *ListArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArticleRequest proto.InternalMessageInfo

type DetailRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{2}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DetailRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type DeleteArticleRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArticleRequest) Reset()         { *m = DeleteArticleRequest{} }
func (m *DeleteArticleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArticleRequest) ProtoMessage()    {}
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{3}
}

func (m *DeleteArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArticleRequest.Unmarshal(m, b)
}
func (m *DeleteArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArticleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArticleRequest.Merge(m, src)
}
func (m *DeleteArticleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteArticleRequest.Size(m)
}
func (m *DeleteArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArticleRequest proto.InternalMessageInfo

func (m *DeleteArticleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Article struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author               *Author  `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{4}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Article) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Article) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Author struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{5}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Result struct {
	StatusCode           int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Article              *Article   `protobuf:"bytes,3,opt,name=article,proto3" json:"article,omitempty"`
	Articles             []*Article `protobuf:"bytes,4,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{6}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Result) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

func (m *Result) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateArticleRequest)(nil), "CreateArticleRequest")
	proto.RegisterType((*ListArticleRequest)(nil), "ListArticleRequest")
	proto.RegisterType((*DetailRequest)(nil), "DetailRequest")
	proto.RegisterType((*DeleteArticleRequest)(nil), "DeleteArticleRequest")
	proto.RegisterType((*Article)(nil), "Article")
	proto.RegisterType((*Author)(nil), "Author")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() {
	proto.RegisterFile("proto/article/article.proto", fileDescriptor_6bfa7f65df4b480f)
}

var fileDescriptor_6bfa7f65df4b480f = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0xaa, 0xda, 0x40,
	0x10, 0xc6, 0x49, 0xa2, 0x89, 0x8e, 0xe8, 0xc5, 0x36, 0xc2, 0x52, 0x29, 0x0d, 0xa1, 0x14, 0x6f,
	0x1a, 0x41, 0x9f, 0x20, 0x28, 0x85, 0x42, 0xaf, 0xd2, 0x07, 0x90, 0x34, 0x3b, 0xe8, 0x42, 0x6a,
	0x6c, 0x76, 0x72, 0xde, 0xe2, 0x3c, 0xc4, 0xb9, 0x3b, 0x8f, 0x79, 0x70, 0xff, 0x08, 0x39, 0xf1,
	0x5c, 0x25, 0xf3, 0x9b, 0x9d, 0x9d, 0xf9, 0xf6, 0x1b, 0x58, 0x5d, 0xdb, 0x86, 0x9a, 0x4d, 0xd9,
	0x92, 0xac, 0x6a, 0x74, 0xdf, 0x4c, 0xd3, 0xb4, 0x82, 0x78, 0xdf, 0x62, 0x49, 0x98, 0x1b, 0x5c,
	0xe0, 0xff, 0x0e, 0x15, 0xb1, 0x18, 0xc6, 0x24, 0xa9, 0x46, 0xee, 0x25, 0xde, 0x7a, 0x5a, 0x98,
	0x80, 0x71, 0x88, 0xaa, 0xe6, 0x42, 0x78, 0x21, 0xee, 0x6b, 0xee, 0x42, 0xb6, 0x82, 0x69, 0xd9,
	0xd1, 0xb9, 0x69, 0x8f, 0x52, 0xf0, 0x20, 0xf1, 0xd6, 0x41, 0x31, 0x31, 0xe0, 0x97, 0x48, 0x63,
	0x60, 0xbf, 0xa5, 0xa2, 0x7e, 0x8b, 0x74, 0x07, 0xf3, 0x03, 0x52, 0x29, 0x6b, 0xd7, 0x73, 0x01,
	0xbe, 0x14, 0xba, 0x61, 0x50, 0xf8, 0x52, 0x30, 0x06, 0x23, 0x55, 0x77, 0x27, 0xdb, 0x4a, 0xff,
	0xa7, 0xdf, 0x21, 0x3e, 0x60, 0x8d, 0x83, 0x79, 0xdf, 0xd5, 0xa6, 0xaf, 0x1e, 0x44, 0xf6, 0xc8,
	0xe0, 0xde, 0xbb, 0x36, 0xff, 0x03, 0x6d, 0x41, 0x5f, 0xdb, 0x57, 0x08, 0x8d, 0x14, 0x3e, 0x4a,
	0xbc, 0xf5, 0x6c, 0x1b, 0x65, 0xb9, 0x0e, 0x0b, 0x8b, 0xd9, 0x17, 0x80, 0x4a, 0x3f, 0xa2, 0x38,
	0x96, 0xc4, 0xc7, 0xba, 0x7a, 0x6a, 0x49, 0x4e, 0xb7, 0x74, 0x77, 0x15, 0x2e, 0x1d, 0x9a, 0xb4,
	0x25, 0x39, 0xa5, 0x1c, 0x42, 0x73, 0xdf, 0x40, 0xc4, 0xb3, 0x07, 0x61, 0x81, 0xaa, 0xab, 0x6f,
	0x33, 0xcc, 0x14, 0x95, 0xd4, 0xa9, 0x63, 0xd5, 0x08, 0xe3, 0xca, 0xb8, 0x00, 0x83, 0xf6, 0x8d,
	0xd0, 0xe3, 0xff, 0x43, 0xa5, 0xca, 0x93, 0x93, 0xe5, 0x42, 0x96, 0x42, 0x64, 0x3d, 0xd7, 0xc2,
	0x66, 0xdb, 0x49, 0xe6, 0x1e, 0xcf, 0x25, 0xd8, 0x37, 0x98, 0xd8, 0x5f, 0xc5, 0x47, 0x49, 0xd0,
	0x3b, 0x74, 0xcf, 0x6c, 0x5f, 0x3c, 0x58, 0x58, 0xfa, 0x07, 0xdb, 0x27, 0x59, 0x21, 0xfb, 0x01,
	0xf3, 0x9f, 0x48, 0xd5, 0xd9, 0x62, 0xc5, 0x3e, 0x65, 0x43, 0xab, 0x3f, 0x47, 0x99, 0x95, 0xb1,
	0x81, 0x79, 0x6f, 0xdd, 0xd8, 0x32, 0x7b, 0xb4, 0x7e, 0xbd, 0x82, 0x9e, 0xdf, 0x6c, 0x99, 0x3d,
	0xf2, 0xff, 0x5e, 0xf0, 0x37, 0xd4, 0x7b, 0xbd, 0x7b, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x0a,
	0xb4, 0x40, 0xf6, 0x02, 0x00, 0x00,
}
