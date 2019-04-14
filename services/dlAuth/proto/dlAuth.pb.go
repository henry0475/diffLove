// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dlAuth.proto

package dlAuth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserLoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975394f7a24c74d, []int{0}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserLoginResponse struct {
	Status               uint32       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Time                 int64        `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	UserInfo             *UserInfoMsg `protobuf:"bytes,4,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975394f7a24c74d, []int{1}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserLoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserLoginResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *UserLoginResponse) GetUserInfo() *UserInfoMsg {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type UserInfoMsg struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender               uint32   `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoMsg) Reset()         { *m = UserInfoMsg{} }
func (m *UserInfoMsg) String() string { return proto.CompactTextString(m) }
func (*UserInfoMsg) ProtoMessage()    {}
func (*UserInfoMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975394f7a24c74d, []int{2}
}

func (m *UserInfoMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoMsg.Unmarshal(m, b)
}
func (m *UserInfoMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoMsg.Marshal(b, m, deterministic)
}
func (m *UserInfoMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoMsg.Merge(m, src)
}
func (m *UserInfoMsg) XXX_Size() int {
	return xxx_messageInfo_UserInfoMsg.Size(m)
}
func (m *UserInfoMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoMsg proto.InternalMessageInfo

func (m *UserInfoMsg) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfoMsg) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfoMsg) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfoMsg) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserRegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Gender               uint32   `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterRequest) Reset()         { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()    {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975394f7a24c74d, []int{3}
}

func (m *UserRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterRequest.Unmarshal(m, b)
}
func (m *UserRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterRequest.Marshal(b, m, deterministic)
}
func (m *UserRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterRequest.Merge(m, src)
}
func (m *UserRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegisterRequest.Size(m)
}
func (m *UserRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterRequest proto.InternalMessageInfo

func (m *UserRegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRegisterRequest) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

type UserRegisterResponse struct {
	Status               uint32       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Time                 int64        `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	UserInfo             *UserInfoMsg `protobuf:"bytes,4,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserRegisterResponse) Reset()         { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()    {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975394f7a24c74d, []int{4}
}

func (m *UserRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResponse.Unmarshal(m, b)
}
func (m *UserRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResponse.Marshal(b, m, deterministic)
}
func (m *UserRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResponse.Merge(m, src)
}
func (m *UserRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResponse.Size(m)
}
func (m *UserRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResponse proto.InternalMessageInfo

func (m *UserRegisterResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserRegisterResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserRegisterResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *UserRegisterResponse) GetUserInfo() *UserInfoMsg {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*UserLoginRequest)(nil), "UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "UserLoginResponse")
	proto.RegisterType((*UserInfoMsg)(nil), "UserInfoMsg")
	proto.RegisterType((*UserRegisterRequest)(nil), "UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "UserRegisterResponse")
}

func init() { proto.RegisterFile("dlAuth.proto", fileDescriptor_e975394f7a24c74d) }

var fileDescriptor_e975394f7a24c74d = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xa5, 0x2d, 0x10, 0x19, 0xc0, 0xc0, 0x88, 0xa6, 0xe1, 0x44, 0xf6, 0xd4, 0xd3, 0x1e, 0xd0,
	0xab, 0x07, 0x8f, 0x1a, 0xbd, 0x6c, 0xc2, 0x07, 0xa0, 0x1d, 0xeb, 0x06, 0xd9, 0xc5, 0xce, 0x36,
	0x1e, 0x8c, 0x5e, 0xfc, 0x71, 0xd3, 0x6d, 0x21, 0x45, 0xb8, 0x99, 0x78, 0xda, 0x79, 0xfb, 0xb2,
	0xf3, 0xde, 0xcc, 0x3e, 0x18, 0xa4, 0xaf, 0x37, 0x85, 0x7b, 0x91, 0x9b, 0xdc, 0x3a, 0x2b, 0xee,
	0x60, 0xb4, 0x60, 0xca, 0xef, 0x6d, 0xa6, 0x8d, 0xa2, 0xb7, 0x82, 0xd8, 0xe1, 0x14, 0x4e, 0x0a,
	0xa6, 0xdc, 0x2c, 0xd7, 0x14, 0x07, 0xb3, 0x20, 0xe9, 0xa9, 0x1d, 0x2e, 0xb9, 0xcd, 0x92, 0xf9,
	0xdd, 0xe6, 0x69, 0x1c, 0x56, 0xdc, 0x16, 0x8b, 0x0f, 0x18, 0x37, 0x7a, 0xf1, 0xc6, 0x1a, 0x26,
	0xbc, 0x80, 0x2e, 0xbb, 0xa5, 0x2b, 0xd8, 0xb7, 0x1a, 0xaa, 0x1a, 0xe1, 0x08, 0xa2, 0x35, 0x67,
	0x75, 0x8f, 0xb2, 0x44, 0x84, 0xb6, 0xd3, 0x6b, 0x8a, 0xa3, 0x59, 0x90, 0x44, 0xca, 0xd7, 0x98,
	0x54, 0x56, 0x6e, 0xcd, 0xb3, 0x8d, 0xdb, 0xb3, 0x20, 0xe9, 0xcf, 0x07, 0x72, 0x51, 0x5f, 0x3c,
	0x70, 0xa6, 0x76, 0xac, 0xf8, 0x0e, 0xa0, 0xdf, 0x60, 0xf0, 0x14, 0x42, 0x9d, 0x7a, 0xcd, 0xb6,
	0x0a, 0x75, 0xba, 0x37, 0x54, 0x78, 0x38, 0x94, 0xd1, 0x4f, 0x2b, 0xcf, 0x45, 0x15, 0xb7, 0xc5,
	0xa5, 0xff, 0x8c, 0x4c, 0x4a, 0xb9, 0xd7, 0x1f, 0xaa, 0x1a, 0xe1, 0x04, 0x3a, 0xce, 0xae, 0xc8,
	0xc4, 0x1d, 0xff, 0xa0, 0x02, 0x82, 0xe0, 0xac, 0x34, 0xa1, 0x28, 0xd3, 0xec, 0xca, 0xf3, 0x4f,
	0x1b, 0x6d, 0x88, 0x47, 0x4d, 0x71, 0xf1, 0x05, 0x93, 0x7d, 0x99, 0xff, 0x5d, 0xf6, 0xfc, 0x13,
	0xba, 0x55, 0x8a, 0xf0, 0x0a, 0x7a, 0xbb, 0x3f, 0xc7, 0xb1, 0xfc, 0x9d, 0xa5, 0x29, 0xca, 0x83,
	0x48, 0x88, 0x16, 0x5e, 0xc3, 0xa0, 0xe9, 0x1f, 0x27, 0xf2, 0xc8, 0xd6, 0xa6, 0xe7, 0xf2, 0xd8,
	0x90, 0xa2, 0xf5, 0xd8, 0xf5, 0xd9, 0xbd, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xcb, 0xb2,
	0x4b, 0xcb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DlAuth service

type DlAuthClient interface {
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error)
}

type dlAuthClient struct {
	c           client.Client
	serviceName string
}

func NewDlAuthClient(serviceName string, c client.Client) DlAuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "dlauth"
	}
	return &dlAuthClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *dlAuthClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DlAuth.UserLogin", in)
	out := new(UserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dlAuthClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DlAuth.UserRegister", in)
	out := new(UserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DlAuth service

type DlAuthHandler interface {
	UserLogin(context.Context, *UserLoginRequest, *UserLoginResponse) error
	UserRegister(context.Context, *UserRegisterRequest, *UserRegisterResponse) error
}

func RegisterDlAuthHandler(s server.Server, hdlr DlAuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DlAuth{hdlr}, opts...))
}

type DlAuth struct {
	DlAuthHandler
}

func (h *DlAuth) UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.DlAuthHandler.UserLogin(ctx, in, out)
}

func (h *DlAuth) UserRegister(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error {
	return h.DlAuthHandler.UserRegister(ctx, in, out)
}
