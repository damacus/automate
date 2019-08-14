// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/auth/users/users.proto

package users

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/auth/users/request"
	response "github.com/chef/automate/components/automate-gateway/api/auth/users/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("components/automate-gateway/api/auth/users/users.proto", fileDescriptor_a3ccb1f707d1e4d8)
}

var fileDescriptor_a3ccb1f707d1e4d8 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe5, 0x82, 0x42, 0x7a, 0x6d, 0x21, 0x79, 0x1b, 0x52, 0xcb, 0x62, 0xca, 0x50, 0x44,
	0x21, 0x36, 0x02, 0x89, 0x21, 0x03, 0xa0, 0x96, 0x8a, 0x09, 0x81, 0x0a, 0x5d, 0x90, 0x18, 0xae,
	0xce, 0x5b, 0xc7, 0x92, 0xed, 0x73, 0xed, 0xb3, 0x50, 0xc4, 0x9f, 0xc1, 0x53, 0x95, 0x95, 0x6f,
	0xc1, 0x8a, 0xbc, 0x31, 0xf0, 0x1d, 0x58, 0xd8, 0x10, 0x12, 0x13, 0x1f, 0x02, 0x21, 0xfb, 0x92,
	0x9c, 0xad, 0x26, 0xb6, 0xc3, 0x12, 0x29, 0x77, 0xcf, 0xfb, 0xde, 0xf3, 0xbb, 0xc7, 0x77, 0x47,
	0x1e, 0x98, 0xcc, 0xf5, 0x99, 0x87, 0x1e, 0x0f, 0x0d, 0x1a, 0x71, 0xe6, 0x52, 0x8e, 0x7d, 0x8b,
	0x72, 0x7c, 0x4b, 0xc7, 0x06, 0xf5, 0xed, 0x74, 0x70, 0x64, 0x44, 0x21, 0x06, 0xa1, 0xf8, 0xd5,
	0xfd, 0x80, 0x71, 0x06, 0x3b, 0xe6, 0x08, 0x4f, 0xf5, 0x59, 0x85, 0x4e, 0x7d, 0x5b, 0xcf, 0xa6,
	0xb5, 0x1b, 0x16, 0x63, 0x96, 0x83, 0xa2, 0xd6, 0xf3, 0x18, 0xa7, 0xdc, 0x66, 0xde, 0xb4, 0x4c,
	0x7b, 0xb8, 0xc2, 0x72, 0x01, 0x9e, 0x45, 0x18, 0xf2, 0xfc, 0xb2, 0xda, 0xa3, 0x95, 0xea, 0x43,
	0x9f, 0x79, 0x21, 0x16, 0x1a, 0x3c, 0x5e, 0xd8, 0x20, 0xf0, 0x4d, 0x23, 0x9b, 0x37, 0xfb, 0x16,
	0x7a, 0x7d, 0x9f, 0x39, 0xb6, 0x39, 0x5e, 0x82, 0xb0, 0x4a, 0x07, 0x9b, 0xba, 0x17, 0x3b, 0xdc,
	0xfb, 0xbb, 0x49, 0xd6, 0x8f, 0x53, 0x4f, 0xcf, 0x2c, 0x97, 0xc3, 0x17, 0x85, 0x34, 0x9f, 0x22,
	0xcf, 0x06, 0xe0, 0x8e, 0xbe, 0x64, 0x5f, 0xf5, 0xe9, 0x6e, 0xe8, 0x33, 0xe9, 0x11, 0x9e, 0x69,
	0x37, 0x4b, 0xd4, 0x82, 0x5d, 0xcf, 0xb4, 0xbd, 0x17, 0x71, 0xa2, 0x6e, 0x12, 0x92, 0x6e, 0xce,
	0x20, 0x13, 0xc4, 0x89, 0xda, 0x80, 0xcb, 0x01, 0xd2, 0xe1, 0x24, 0x51, 0x37, 0xc8, 0xba, 0x4d,
	0x5d, 0x31, 0x31, 0x49, 0xd4, 0x16, 0x5c, 0x9d, 0xff, 0x1d, 0x38, 0x76, 0xc8, 0xe3, 0xef, 0xbf,
	0x3f, 0xad, 0x6d, 0xc1, 0x46, 0x6e, 0x73, 0xe1, 0x87, 0x42, 0xda, 0x53, 0x2b, 0xfb, 0xe3, 0xf4,
	0xd7, 0xa3, 0x2e, 0xc2, 0xad, 0x4a, 0xfb, 0x33, 0xa9, 0xb6, 0x5b, 0xcf, 0x7b, 0xcf, 0x8f, 0x13,
	0x75, 0x87, 0x5c, 0x97, 0xd6, 0x07, 0xef, 0xa2, 0x69, 0x8f, 0x0f, 0x05, 0x8a, 0x2e, 0xe9, 0x48,
	0xdb, 0x52, 0x33, 0x49, 0xd4, 0x6b, 0xb0, 0x25, 0x67, 0x2c, 0x14, 0x3c, 0x2a, 0x74, 0xf3, 0x1f,
	0x8b, 0x2c, 0x80, 0x6f, 0x0a, 0x21, 0x07, 0x01, 0x52, 0x8e, 0xa9, 0x01, 0xb8, 0x5d, 0xc9, 0x24,
	0xc5, 0xb5, 0xa9, 0xde, 0x2c, 0x08, 0xa4, 0x09, 0x0d, 0x33, 0xeb, 0x74, 0x31, 0x12, 0x80, 0x96,
	0x24, 0x10, 0x22, 0x11, 0x4a, 0x2f, 0x1f, 0x4a, 0x36, 0x74, 0x69, 0xa0, 0xec, 0xc1, 0x1f, 0x85,
	0x74, 0x9e, 0xa0, 0x83, 0xc2, 0xd5, 0xff, 0x05, 0x74, 0xb7, 0x1a, 0x45, 0x2e, 0x71, 0x84, 0xa1,
	0xdf, 0x8b, 0xca, 0xa3, 0x6a, 0x42, 0x63, 0x98, 0x15, 0x94, 0x86, 0x55, 0x40, 0x15, 0x7a, 0x91,
	0xd7, 0xde, 0xb2, 0xbc, 0x7e, 0x29, 0x84, 0x1c, 0xfb, 0xc3, 0xfa, 0x79, 0x49, 0x71, 0xed, 0xbc,
	0x3e, 0x56, 0xa2, 0x45, 0x59, 0xd3, 0xfa, 0x68, 0x42, 0x2f, 0xd0, 0xb4, 0x25, 0x68, 0x32, 0xd0,
	0x9f, 0x73, 0xc6, 0x97, 0xe8, 0x9c, 0xd6, 0x66, 0x4c, 0xc5, 0xb5, 0x19, 0xdf, 0xc7, 0x89, 0x0a,
	0xa4, 0x55, 0x81, 0xa7, 0x11, 0x75, 0x0e, 0x91, 0xb6, 0x2f, 0x22, 0x76, 0xa1, 0x53, 0x9c, 0xcd,
	0x61, 0x6e, 0x6b, 0xed, 0x12, 0xc2, 0xcf, 0x0a, 0xb9, 0x32, 0xbd, 0x50, 0x60, 0xb7, 0x12, 0xef,
	0xd0, 0xa5, 0xb6, 0x53, 0x9b, 0xec, 0x79, 0x9c, 0xa8, 0x1d, 0x02, 0xf9, 0xf4, 0x30, 0x6d, 0x50,
	0xbc, 0x40, 0xb6, 0x49, 0x3b, 0x17, 0x9c, 0x10, 0x2c, 0xb8, 0x3d, 0xce, 0xd7, 0x14, 0xf8, 0xaa,
	0x10, 0x22, 0x3f, 0xfe, 0xda, 0x7e, 0x57, 0x3f, 0x52, 0xaf, 0x4a, 0x9c, 0xe7, 0xcf, 0xd3, 0x12,
	0xef, 0x0b, 0x0e, 0xd3, 0xf9, 0x9a, 0xb2, 0x7f, 0xf8, 0xfa, 0xc0, 0xb2, 0xf9, 0x28, 0x3a, 0xd1,
	0x4d, 0xe6, 0x1a, 0xa9, 0xa7, 0xf9, 0x4b, 0x66, 0xd4, 0x7f, 0x60, 0x4f, 0x1a, 0xd9, 0x73, 0x76,
	0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x02, 0x98, 0x0a, 0x44, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersMgmtClient is the client API for UsersMgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersMgmtClient interface {
	GetUsers(ctx context.Context, in *request.GetUsersReq, opts ...grpc.CallOption) (*response.Users, error)
	GetUserByUsername(ctx context.Context, in *request.Username, opts ...grpc.CallOption) (*response.User, error)
	CreateUser(ctx context.Context, in *request.CreateUser, opts ...grpc.CallOption) (*response.User, error)
	DeleteUserByUsername(ctx context.Context, in *request.Username, opts ...grpc.CallOption) (*response.DeleteUserResp, error)
	UpdateUser(ctx context.Context, in *request.UpdateUser, opts ...grpc.CallOption) (*response.User, error)
	UpdateSelf(ctx context.Context, in *request.UpdateSelf, opts ...grpc.CallOption) (*response.User, error)
	// deprecated API
	GetUser(ctx context.Context, in *request.Email, opts ...grpc.CallOption) (*response.User, error)
	DeleteUser(ctx context.Context, in *request.Email, opts ...grpc.CallOption) (*response.DeleteUserResp, error)
}

type usersMgmtClient struct {
	cc *grpc.ClientConn
}

func NewUsersMgmtClient(cc *grpc.ClientConn) UsersMgmtClient {
	return &usersMgmtClient{cc}
}

func (c *usersMgmtClient) GetUsers(ctx context.Context, in *request.GetUsersReq, opts ...grpc.CallOption) (*response.Users, error) {
	out := new(response.Users)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersMgmtClient) GetUserByUsername(ctx context.Context, in *request.Username, opts ...grpc.CallOption) (*response.User, error) {
	out := new(response.User)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersMgmtClient) CreateUser(ctx context.Context, in *request.CreateUser, opts ...grpc.CallOption) (*response.User, error) {
	out := new(response.User)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersMgmtClient) DeleteUserByUsername(ctx context.Context, in *request.Username, opts ...grpc.CallOption) (*response.DeleteUserResp, error) {
	out := new(response.DeleteUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/DeleteUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersMgmtClient) UpdateUser(ctx context.Context, in *request.UpdateUser, opts ...grpc.CallOption) (*response.User, error) {
	out := new(response.User)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersMgmtClient) UpdateSelf(ctx context.Context, in *request.UpdateSelf, opts ...grpc.CallOption) (*response.User, error) {
	out := new(response.User)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/UpdateSelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *usersMgmtClient) GetUser(ctx context.Context, in *request.Email, opts ...grpc.CallOption) (*response.User, error) {
	out := new(response.User)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *usersMgmtClient) DeleteUser(ctx context.Context, in *request.Email, opts ...grpc.CallOption) (*response.DeleteUserResp, error) {
	out := new(response.DeleteUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.users.UsersMgmt/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersMgmtServer is the server API for UsersMgmt service.
type UsersMgmtServer interface {
	GetUsers(context.Context, *request.GetUsersReq) (*response.Users, error)
	GetUserByUsername(context.Context, *request.Username) (*response.User, error)
	CreateUser(context.Context, *request.CreateUser) (*response.User, error)
	DeleteUserByUsername(context.Context, *request.Username) (*response.DeleteUserResp, error)
	UpdateUser(context.Context, *request.UpdateUser) (*response.User, error)
	UpdateSelf(context.Context, *request.UpdateSelf) (*response.User, error)
	// deprecated API
	GetUser(context.Context, *request.Email) (*response.User, error)
	DeleteUser(context.Context, *request.Email) (*response.DeleteUserResp, error)
}

// UnimplementedUsersMgmtServer can be embedded to have forward compatible implementations.
type UnimplementedUsersMgmtServer struct {
}

func (*UnimplementedUsersMgmtServer) GetUsers(ctx context.Context, req *request.GetUsersReq) (*response.Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUsersMgmtServer) GetUserByUsername(ctx context.Context, req *request.Username) (*response.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (*UnimplementedUsersMgmtServer) CreateUser(ctx context.Context, req *request.CreateUser) (*response.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUsersMgmtServer) DeleteUserByUsername(ctx context.Context, req *request.Username) (*response.DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByUsername not implemented")
}
func (*UnimplementedUsersMgmtServer) UpdateUser(ctx context.Context, req *request.UpdateUser) (*response.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUsersMgmtServer) UpdateSelf(ctx context.Context, req *request.UpdateSelf) (*response.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSelf not implemented")
}
func (*UnimplementedUsersMgmtServer) GetUser(ctx context.Context, req *request.Email) (*response.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUsersMgmtServer) DeleteUser(ctx context.Context, req *request.Email) (*response.DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUsersMgmtServer(s *grpc.Server, srv UsersMgmtServer) {
	s.RegisterService(&_UsersMgmt_serviceDesc, srv)
}

func _UsersMgmt_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).GetUsers(ctx, req.(*request.GetUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).GetUserByUsername(ctx, req.(*request.Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).CreateUser(ctx, req.(*request.CreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_DeleteUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).DeleteUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/DeleteUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).DeleteUserByUsername(ctx, req.(*request.Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).UpdateUser(ctx, req.(*request.UpdateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_UpdateSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateSelf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).UpdateSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/UpdateSelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).UpdateSelf(ctx, req.(*request.UpdateSelf))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).GetUser(ctx, req.(*request.Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersMgmt_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersMgmtServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.users.UsersMgmt/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersMgmtServer).DeleteUser(ctx, req.(*request.Email))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersMgmt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.users.UsersMgmt",
	HandlerType: (*UsersMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _UsersMgmt_GetUsers_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UsersMgmt_GetUserByUsername_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UsersMgmt_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUserByUsername",
			Handler:    _UsersMgmt_DeleteUserByUsername_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UsersMgmt_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateSelf",
			Handler:    _UsersMgmt_UpdateSelf_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UsersMgmt_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UsersMgmt_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/auth/users/users.proto",
}
