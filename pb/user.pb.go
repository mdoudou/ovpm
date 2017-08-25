// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	user.proto
	vpn.proto
	network.proto

It has these top-level messages:
	UserListRequest
	UserCreateRequest
	UserUpdateRequest
	UserDeleteRequest
	UserRenewRequest
	UserGenConfigRequest
	UserResponse
	UserGenConfigResponse
	VPNStatusRequest
	VPNInitRequest
	VPNStatusResponse
	VPNInitResponse
	NetworkCreateRequest
	NetworkListRequest
	NetworkDeleteRequest
	NetworkGetAllTypesRequest
	Network
	NetworkCreateResponse
	NetworkListResponse
	NetworkDeleteResponse
	NetworkGetAllTypesResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type UserUpdateRequest_GWPref int32

const (
	UserUpdateRequest_NOPREF UserUpdateRequest_GWPref = 0
	UserUpdateRequest_NOGW   UserUpdateRequest_GWPref = 1
	UserUpdateRequest_GW     UserUpdateRequest_GWPref = 2
)

var UserUpdateRequest_GWPref_name = map[int32]string{
	0: "NOPREF",
	1: "NOGW",
	2: "GW",
}
var UserUpdateRequest_GWPref_value = map[string]int32{
	"NOPREF": 0,
	"NOGW":   1,
	"GW":     2,
}

func (x UserUpdateRequest_GWPref) String() string {
	return proto.EnumName(UserUpdateRequest_GWPref_name, int32(x))
}
func (UserUpdateRequest_GWPref) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type UserListRequest struct {
}

func (m *UserListRequest) Reset()                    { *m = UserListRequest{} }
func (m *UserListRequest) String() string            { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()               {}
func (*UserListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UserCreateRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
	NoGW     bool   `protobuf:"varint,3,opt,name=NoGW" json:"NoGW,omitempty"`
	HostID   uint32 `protobuf:"varint,4,opt,name=HostID" json:"HostID,omitempty"`
}

func (m *UserCreateRequest) Reset()                    { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()               {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserCreateRequest) GetNoGW() bool {
	if m != nil {
		return m.NoGW
	}
	return false
}

func (m *UserCreateRequest) GetHostID() uint32 {
	if m != nil {
		return m.HostID
	}
	return 0
}

type UserUpdateRequest struct {
	Username string                   `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string                   `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
	Gwpref   UserUpdateRequest_GWPref `protobuf:"varint,3,opt,name=gwpref,enum=pb.UserUpdateRequest_GWPref" json:"gwpref,omitempty"`
	HostID   uint32                   `protobuf:"varint,4,opt,name=HostID" json:"HostID,omitempty"`
}

func (m *UserUpdateRequest) Reset()                    { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()               {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserUpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserUpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserUpdateRequest) GetGwpref() UserUpdateRequest_GWPref {
	if m != nil {
		return m.Gwpref
	}
	return UserUpdateRequest_NOPREF
}

func (m *UserUpdateRequest) GetHostID() uint32 {
	if m != nil {
		return m.HostID
	}
	return 0
}

type UserDeleteRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
}

func (m *UserDeleteRequest) Reset()                    { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()               {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserDeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserRenewRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
}

func (m *UserRenewRequest) Reset()                    { *m = UserRenewRequest{} }
func (m *UserRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRenewRequest) ProtoMessage()               {}
func (*UserRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserRenewRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserGenConfigRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
}

func (m *UserGenConfigRequest) Reset()                    { *m = UserGenConfigRequest{} }
func (m *UserGenConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*UserGenConfigRequest) ProtoMessage()               {}
func (*UserGenConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserGenConfigRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserResponse struct {
	Users []*UserResponse_User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserResponse) GetUsers() []*UserResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserResponse_User struct {
	Username           string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	ServerSerialNumber string `protobuf:"bytes,2,opt,name=ServerSerialNumber" json:"ServerSerialNumber,omitempty"`
	Cert               string `protobuf:"bytes,3,opt,name=Cert" json:"Cert,omitempty"`
	CreatedAt          string `protobuf:"bytes,4,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	IPNet              string `protobuf:"bytes,5,opt,name=IPNet" json:"IPNet,omitempty"`
	NoGW               bool   `protobuf:"varint,6,opt,name=NoGW" json:"NoGW,omitempty"`
	HostID             uint32 `protobuf:"varint,7,opt,name=HostID" json:"HostID,omitempty"`
}

func (m *UserResponse_User) Reset()                    { *m = UserResponse_User{} }
func (m *UserResponse_User) String() string            { return proto.CompactTextString(m) }
func (*UserResponse_User) ProtoMessage()               {}
func (*UserResponse_User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *UserResponse_User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse_User) GetServerSerialNumber() string {
	if m != nil {
		return m.ServerSerialNumber
	}
	return ""
}

func (m *UserResponse_User) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *UserResponse_User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserResponse_User) GetIPNet() string {
	if m != nil {
		return m.IPNet
	}
	return ""
}

func (m *UserResponse_User) GetNoGW() bool {
	if m != nil {
		return m.NoGW
	}
	return false
}

func (m *UserResponse_User) GetHostID() uint32 {
	if m != nil {
		return m.HostID
	}
	return 0
}

type UserGenConfigResponse struct {
	ClientConfig string `protobuf:"bytes,1,opt,name=ClientConfig" json:"ClientConfig,omitempty"`
}

func (m *UserGenConfigResponse) Reset()                    { *m = UserGenConfigResponse{} }
func (m *UserGenConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*UserGenConfigResponse) ProtoMessage()               {}
func (*UserGenConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserGenConfigResponse) GetClientConfig() string {
	if m != nil {
		return m.ClientConfig
	}
	return ""
}

func init() {
	proto.RegisterType((*UserListRequest)(nil), "pb.UserListRequest")
	proto.RegisterType((*UserCreateRequest)(nil), "pb.UserCreateRequest")
	proto.RegisterType((*UserUpdateRequest)(nil), "pb.UserUpdateRequest")
	proto.RegisterType((*UserDeleteRequest)(nil), "pb.UserDeleteRequest")
	proto.RegisterType((*UserRenewRequest)(nil), "pb.UserRenewRequest")
	proto.RegisterType((*UserGenConfigRequest)(nil), "pb.UserGenConfigRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
	proto.RegisterType((*UserResponse_User)(nil), "pb.UserResponse.User")
	proto.RegisterType((*UserGenConfigResponse)(nil), "pb.UserGenConfigResponse")
	proto.RegisterEnum("pb.UserUpdateRequest_GWPref", UserUpdateRequest_GWPref_name, UserUpdateRequest_GWPref_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error) {
	out := new(UserGenConfigResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/GenConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	List(context.Context, *UserListRequest) (*UserResponse, error)
	Create(context.Context, *UserCreateRequest) (*UserResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserResponse, error)
	Delete(context.Context, *UserDeleteRequest) (*UserResponse, error)
	Renew(context.Context, *UserRenewRequest) (*UserResponse, error)
	GenConfig(context.Context, *UserGenConfigRequest) (*UserGenConfigResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Renew(ctx, req.(*UserRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGenConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GenConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenConfig(ctx, req.(*UserGenConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _UserService_Renew_Handler,
		},
		{
			MethodName: "GenConfig",
			Handler:    _UserService_GenConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x5d, 0xc7, 0xc4, 0xd3, 0x02, 0xee, 0x90, 0x22, 0x13, 0xf5, 0x10, 0xf9, 0x80, 0x22,
	0x21, 0xb9, 0x22, 0xe5, 0xc6, 0x09, 0x52, 0x30, 0x95, 0x90, 0x1b, 0x6d, 0x55, 0xe5, 0x9c, 0x90,
	0x49, 0x65, 0x29, 0xb1, 0xcd, 0xee, 0x86, 0xbc, 0x00, 0x2f, 0xc6, 0x85, 0x27, 0xe0, 0x81, 0xd0,
	0x7a, 0xfd, 0x83, 0x83, 0x83, 0x72, 0xe0, 0x36, 0x7f, 0xdf, 0x8e, 0xe7, 0x9b, 0x6f, 0x0c, 0xb0,
	0x11, 0xc4, 0x83, 0x8c, 0xa7, 0x32, 0x45, 0x33, 0x9b, 0xfb, 0x67, 0xf0, 0xf4, 0x5e, 0x10, 0xff,
	0x1c, 0x0b, 0xc9, 0xe8, 0xeb, 0x86, 0x84, 0xf4, 0xb7, 0x70, 0xa6, 0x42, 0x63, 0x4e, 0x33, 0x49,
	0x45, 0x10, 0xfb, 0xd0, 0x55, 0xc1, 0x64, 0xb6, 0x26, 0xcf, 0x18, 0x18, 0x43, 0x87, 0x55, 0xbe,
	0xca, 0x4d, 0x66, 0x42, 0x6c, 0x53, 0xbe, 0xf0, 0x4c, 0x9d, 0x2b, 0x7d, 0x44, 0xb0, 0xa2, 0x34,
	0x9c, 0x7a, 0xc7, 0x03, 0x63, 0xd8, 0x65, 0xb9, 0x8d, 0xcf, 0xc1, 0xfe, 0x94, 0x0a, 0x79, 0x73,
	0xed, 0x59, 0x03, 0x63, 0xf8, 0x98, 0x15, 0x9e, 0xff, 0xc3, 0xd0, 0x9d, 0xef, 0xb3, 0xc5, 0x7f,
	0xe8, 0xfc, 0x06, 0xec, 0x87, 0x6d, 0xc6, 0x69, 0x99, 0xf7, 0x7e, 0x32, 0xba, 0x08, 0xb2, 0x79,
	0xf0, 0xd7, 0xf3, 0x41, 0x38, 0x9d, 0x70, 0x5a, 0xb2, 0xa2, 0x76, 0xef, 0xb7, 0xbd, 0x04, 0x5b,
	0x57, 0x22, 0x80, 0x1d, 0xdd, 0x4e, 0xd8, 0x87, 0x8f, 0xee, 0x11, 0x76, 0xc1, 0x8a, 0x6e, 0xc3,
	0xa9, 0x6b, 0xa0, 0x0d, 0x66, 0x38, 0x75, 0x4d, 0xff, 0x52, 0x8f, 0x70, 0x4d, 0x2b, 0x3a, 0x68,
	0x04, 0x3f, 0x00, 0x57, 0xd9, 0x8c, 0x12, 0xda, 0x1e, 0x52, 0x3f, 0x82, 0x9e, 0xb2, 0x43, 0x4a,
	0xc6, 0x69, 0xb2, 0x8c, 0x1f, 0x0e, 0xc1, 0x7c, 0x37, 0xe1, 0x54, 0x37, 0x11, 0x59, 0x9a, 0x08,
	0xc2, 0x57, 0xd0, 0x51, 0x3a, 0x10, 0x9e, 0x31, 0x38, 0x1e, 0x9e, 0x8c, 0xce, 0x4b, 0x6a, 0xca,
	0x02, 0xed, 0xe8, 0x9a, 0xfe, 0x4f, 0x03, 0x2c, 0xe5, 0xff, 0x73, 0x13, 0x01, 0xe0, 0x1d, 0xf1,
	0x6f, 0xc4, 0xef, 0x88, 0xc7, 0xb3, 0x55, 0xb4, 0x59, 0xcf, 0x89, 0x17, 0x3b, 0x69, 0xc9, 0x28,
	0x5d, 0x8c, 0x89, 0xcb, 0x7c, 0x37, 0x0e, 0xcb, 0x6d, 0xbc, 0x00, 0x47, 0x8b, 0x6e, 0xf1, 0x4e,
	0xe6, 0xf4, 0x3b, 0xac, 0x0e, 0x60, 0x0f, 0x3a, 0x37, 0x93, 0x88, 0xa4, 0xd7, 0xc9, 0x33, 0xda,
	0xa9, 0xf4, 0x65, 0xb7, 0xea, 0xeb, 0x51, 0x63, 0x87, 0x6f, 0xe1, 0x7c, 0x87, 0xba, 0x82, 0x0e,
	0x1f, 0x4e, 0xc7, 0xab, 0x98, 0x12, 0xa9, 0xe3, 0xc5, 0x70, 0x8d, 0xd8, 0xe8, 0x97, 0x09, 0x27,
	0x0a, 0xad, 0x66, 0x89, 0xbf, 0x10, 0x5e, 0x82, 0xa5, 0x8e, 0x06, 0x9f, 0x95, 0xdc, 0xfd, 0x71,
	0x42, 0x7d, 0x77, 0x97, 0x50, 0xff, 0x08, 0xaf, 0xc0, 0xd6, 0xc3, 0x60, 0x45, 0x77, 0xe3, 0xc4,
	0xf6, 0x81, 0xb4, 0x5c, 0x6b, 0x50, 0x43, 0xbe, 0xfb, 0x40, 0x5a, 0x7f, 0x35, 0xa8, 0xa1, 0xc7,
	0x56, 0xd0, 0x6b, 0xe8, 0xe4, 0x1a, 0xc4, 0x5e, 0x9d, 0xac, 0x25, 0xd9, 0x0a, 0x79, 0x0f, 0x4e,
	0xc5, 0x25, 0x7a, 0x65, 0xc1, 0xae, 0x32, 0xfb, 0x2f, 0x5a, 0x32, 0xe5, 0x1b, 0x73, 0x3b, 0xff,
	0x15, 0x5d, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x90, 0xef, 0x0f, 0xc4, 0x98, 0x04, 0x00, 0x00,
}
