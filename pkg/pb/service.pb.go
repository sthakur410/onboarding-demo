// Code generated by protoc-gen-go. DO NOT EDIT.
// source: onboarding-demo/pkg/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// TODO: Structure your own protobuf messages. Each protocol buffer message is a
// small logical record of information, containing a series of name-value pairs.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4496b8d0625f272, []int{0}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dob                  string   `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4496b8d0625f272, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDob() string {
	if m != nil {
		return m.Dob
	}
	return ""
}

type CreateUserReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4496b8d0625f272, []int{2}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserRes struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRes) Reset()         { *m = CreateUserRes{} }
func (m *CreateUserRes) String() string { return proto.CompactTextString(m) }
func (*CreateUserRes) ProtoMessage()    {}
func (*CreateUserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4496b8d0625f272, []int{3}
}

func (m *CreateUserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRes.Unmarshal(m, b)
}
func (m *CreateUserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRes.Marshal(b, m, deterministic)
}
func (m *CreateUserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRes.Merge(m, src)
}
func (m *CreateUserRes) XXX_Size() int {
	return xxx_messageInfo_CreateUserRes.Size(m)
}
func (m *CreateUserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRes proto.InternalMessageInfo

func (m *CreateUserRes) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ReadUserReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserReq) Reset()         { *m = ReadUserReq{} }
func (m *ReadUserReq) String() string { return proto.CompactTextString(m) }
func (*ReadUserReq) ProtoMessage()    {}
func (*ReadUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4496b8d0625f272, []int{4}
}

func (m *ReadUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserReq.Unmarshal(m, b)
}
func (m *ReadUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserReq.Marshal(b, m, deterministic)
}
func (m *ReadUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserReq.Merge(m, src)
}
func (m *ReadUserReq) XXX_Size() int {
	return xxx_messageInfo_ReadUserReq.Size(m)
}
func (m *ReadUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserReq proto.InternalMessageInfo

func (m *ReadUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadUserRes struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserRes) Reset()         { *m = ReadUserRes{} }
func (m *ReadUserRes) String() string { return proto.CompactTextString(m) }
func (*ReadUserRes) ProtoMessage()    {}
func (*ReadUserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4496b8d0625f272, []int{5}
}

func (m *ReadUserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserRes.Unmarshal(m, b)
}
func (m *ReadUserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserRes.Marshal(b, m, deterministic)
}
func (m *ReadUserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserRes.Merge(m, src)
}
func (m *ReadUserRes) XXX_Size() int {
	return xxx_messageInfo_ReadUserRes.Size(m)
}
func (m *ReadUserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserRes proto.InternalMessageInfo

func (m *ReadUserRes) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
	proto.RegisterType((*User)(nil), "service.User")
	proto.RegisterType((*CreateUserReq)(nil), "service.CreateUserReq")
	proto.RegisterType((*CreateUserRes)(nil), "service.CreateUserRes")
	proto.RegisterType((*ReadUserReq)(nil), "service.ReadUserReq")
	proto.RegisterType((*ReadUserRes)(nil), "service.ReadUserRes")
}

func init() {
	proto.RegisterFile("onboarding-demo/pkg/pb/service.proto", fileDescriptor_a4496b8d0625f272)
}

var fileDescriptor_a4496b8d0625f272 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd5, 0xec, 0xb2, 0x5b, 0x66, 0xe9, 0xb2, 0x32, 0x68, 0x95, 0xcd, 0x82, 0x04, 0x11,
	0x07, 0x24, 0x68, 0x8c, 0xca, 0xad, 0x9c, 0x80, 0x45, 0x48, 0x5c, 0x80, 0x4a, 0x70, 0xe0, 0xe6,
	0x34, 0xb3, 0xc6, 0xa2, 0xf1, 0x18, 0x3b, 0x2d, 0x5b, 0x21, 0x2e, 0xdc, 0x38, 0xf3, 0x34, 0xf0,
	0x1a, 0xbc, 0x02, 0x0f, 0x82, 0xea, 0x38, 0x21, 0x94, 0x1e, 0x7a, 0xb3, 0xff, 0x3f, 0xf3, 0x8d,
	0xfd, 0x7b, 0x02, 0x77, 0x48, 0xe7, 0x24, 0x6c, 0xa1, 0xb4, 0x1c, 0x16, 0x58, 0x12, 0x37, 0x1f,
	0x24, 0x37, 0x39, 0x77, 0x68, 0x17, 0x6a, 0x8a, 0x99, 0xb1, 0x54, 0x11, 0xdb, 0x0f, 0xdb, 0xe4,
	0x54, 0x12, 0xc9, 0x19, 0x72, 0x2f, 0xe7, 0xf3, 0x73, 0x8e, 0xa5, 0xa9, 0x96, 0xf5, 0x57, 0xc9,
	0x8d, 0x60, 0x0a, 0xa3, 0xb8, 0xd0, 0x9a, 0x2a, 0x51, 0x29, 0xd2, 0x2e, 0xb8, 0x8f, 0xa5, 0xaa,
	0xde, 0xcf, 0xf3, 0x6c, 0x4a, 0x25, 0x47, 0xbd, 0xa0, 0xa5, 0xb1, 0x74, 0xb1, 0xac, 0x49, 0xd3,
	0xa1, 0x44, 0x3d, 0x5c, 0x88, 0x99, 0x2a, 0x44, 0x85, 0xfc, 0xbf, 0x45, 0x40, 0xdc, 0xef, 0x7c,
	0xec, 0x3e, 0x09, 0x29, 0xd1, 0x72, 0x32, 0xbe, 0xc9, 0x86, 0x86, 0xe3, 0x4e, 0x43, 0xa5, 0xcf,
	0x29, 0x9f, 0xd1, 0x05, 0x19, 0xd4, 0xdd, 0x96, 0x92, 0x6c, 0xd9, 0x22, 0x56, 0x9b, 0xba, 0x36,
	0xbd, 0x07, 0x57, 0xdf, 0xa2, 0x75, 0x8a, 0xf4, 0x04, 0x9d, 0x21, 0xed, 0x90, 0xc5, 0xb0, 0xbf,
	0xa8, 0xa5, 0xb8, 0x77, 0xab, 0x77, 0xf7, 0xf2, 0xa4, 0xd9, 0xa6, 0x67, 0xb0, 0xfb, 0xc6, 0xa1,
	0x65, 0x87, 0x10, 0xa9, 0x22, 0x98, 0x91, 0x2a, 0x18, 0x83, 0x5d, 0x2d, 0x4a, 0x8c, 0x23, 0xaf,
	0xf8, 0x35, 0x3b, 0x82, 0x9d, 0x82, 0xf2, 0x78, 0xc7, 0x4b, 0xab, 0xe5, 0x78, 0xef, 0xe7, 0x8f,
	0x93, 0xa8, 0xdf, 0x4b, 0x47, 0x30, 0x78, 0x6a, 0x51, 0x54, 0xb8, 0x62, 0x4d, 0xf0, 0x23, 0xbb,
	0x5d, 0x63, 0x3d, 0xf0, 0x60, 0x34, 0xc8, 0x9a, 0x27, 0xf1, 0xbe, 0xb7, 0xd6, 0x6b, 0xdc, 0x36,
	0x35, 0x37, 0xe1, 0x60, 0x82, 0xa2, 0x68, 0xba, 0xac, 0x1d, 0x3a, 0x7d, 0xd0, 0xb5, 0xb7, 0x01,
	0x8e, 0xbe, 0x45, 0x70, 0xf8, 0xb2, 0x9d, 0xa2, 0x33, 0x2c, 0x89, 0xbd, 0x02, 0x78, 0x8e, 0x55,
	0x48, 0x90, 0x1d, 0x67, 0xf5, 0x60, 0x64, 0xcd, 0xd4, 0x64, 0xcf, 0x56, 0x53, 0x93, 0xc4, 0x2d,
	0x6d, 0x2d, 0xeb, 0xf4, 0xe8, 0xeb, 0xaf, 0xdf, 0xdf, 0x23, 0x60, 0x7d, 0x1e, 0x32, 0x66, 0x2f,
	0xa0, 0xdf, 0x1c, 0x8b, 0x5d, 0x6f, 0xeb, 0x3a, 0x17, 0x49, 0x36, 0xa9, 0x2e, 0x65, 0x9e, 0x74,
	0x85, 0x01, 0x9f, 0x3b, 0xb4, 0xfc, 0xb3, 0x2a, 0xbe, 0xb0, 0xd7, 0x00, 0x7f, 0x53, 0x63, 0xc7,
	0x6d, 0xdd, 0x3f, 0xf1, 0x27, 0x9b, 0x75, 0x97, 0x5e, 0xf3, 0xc4, 0x41, 0x7a, 0xc9, 0x13, 0xc7,
	0x3e, 0x83, 0x24, 0x3c, 0xe2, 0x93, 0xd3, 0x77, 0x27, 0x9b, 0x7f, 0xa8, 0x47, 0x26, 0xcf, 0xf7,
	0xfc, 0xfd, 0x1f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x08, 0x3d, 0x15, 0x74, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnboardingDemoClient is the client API for OnboardingDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnboardingDemoClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	ReadUser(ctx context.Context, in *ReadUserReq, opts ...grpc.CallOption) (*ReadUserRes, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error)
}

type onboardingDemoClient struct {
	cc *grpc.ClientConn
}

func NewOnboardingDemoClient(cc *grpc.ClientConn) OnboardingDemoClient {
	return &onboardingDemoClient{cc}
}

func (c *onboardingDemoClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/service.OnboardingDemo/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingDemoClient) ReadUser(ctx context.Context, in *ReadUserReq, opts ...grpc.CallOption) (*ReadUserRes, error) {
	out := new(ReadUserRes)
	err := c.cc.Invoke(ctx, "/service.OnboardingDemo/ReadUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingDemoClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error) {
	out := new(CreateUserRes)
	err := c.cc.Invoke(ctx, "/service.OnboardingDemo/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnboardingDemoServer is the server API for OnboardingDemo service.
type OnboardingDemoServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
	ReadUser(context.Context, *ReadUserReq) (*ReadUserRes, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
}

func RegisterOnboardingDemoServer(s *grpc.Server, srv OnboardingDemoServer) {
	s.RegisterService(&_OnboardingDemo_serviceDesc, srv)
}

func _OnboardingDemo_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingDemoServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.OnboardingDemo/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingDemoServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingDemo_ReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingDemoServer).ReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.OnboardingDemo/ReadUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingDemoServer).ReadUser(ctx, req.(*ReadUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingDemo_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingDemoServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.OnboardingDemo/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingDemoServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnboardingDemo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.OnboardingDemo",
	HandlerType: (*OnboardingDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _OnboardingDemo_GetVersion_Handler,
		},
		{
			MethodName: "ReadUser",
			Handler:    _OnboardingDemo_ReadUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _OnboardingDemo_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onboarding-demo/pkg/pb/service.proto",
}
