// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/teams.proto

package v2beta

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
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
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/teams.proto", fileDescriptor_2decbd3669833c4c)
}

var fileDescriptor_2decbd3669833c4c = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xbf, 0x6f, 0xd4, 0x3e,
	0x1c, 0x55, 0xfa, 0xfd, 0x72, 0xb4, 0x06, 0xda, 0x9e, 0xfb, 0x83, 0xe8, 0x54, 0x21, 0x64, 0x84,
	0x04, 0x85, 0x26, 0xe8, 0x2a, 0x15, 0xe9, 0x06, 0xa0, 0xb4, 0xc0, 0x42, 0x97, 0xaa, 0xed, 0xc0,
	0xe6, 0x4b, 0x4c, 0xce, 0xd2, 0x25, 0x76, 0x63, 0x5f, 0x51, 0xa9, 0x90, 0xd0, 0xb1, 0xdd, 0xca,
	0xdf, 0x92, 0x8d, 0x91, 0x0d, 0xa6, 0x32, 0xb2, 0xf2, 0x37, 0x20, 0xd8, 0x50, 0xec, 0xeb, 0x25,
	0x69, 0x72, 0x69, 0xae, 0x63, 0xf2, 0x79, 0xcf, 0xf9, 0xbc, 0xcf, 0x7b, 0xb1, 0x0d, 0x36, 0x1c,
	0xe6, 0x73, 0x16, 0x90, 0x40, 0x0a, 0x1b, 0xf7, 0x24, 0xf3, 0xb1, 0x24, 0x6b, 0x1e, 0x96, 0xe4,
	0x1d, 0x3e, 0xb6, 0x31, 0xa7, 0x36, 0xc5, 0xbe, 0x7d, 0xd4, 0x6c, 0x13, 0x89, 0x6d, 0x49, 0xb0,
	0x2f, 0x2c, 0x1e, 0x32, 0xc9, 0xe0, 0x8a, 0xd3, 0x21, 0x6f, 0xad, 0x33, 0x86, 0x85, 0x39, 0xb5,
	0x28, 0xf6, 0x2d, 0x8d, 0x6c, 0xac, 0x78, 0x8c, 0x79, 0x5d, 0xa2, 0x16, 0xc0, 0x41, 0xc0, 0x24,
	0x96, 0x94, 0x05, 0x43, 0x6e, 0xe3, 0xc9, 0x04, 0xdf, 0x0c, 0xc9, 0x61, 0x8f, 0x08, 0x99, 0xfe,
	0x76, 0xe3, 0xe9, 0x44, 0x7c, 0xc1, 0x59, 0x20, 0x48, 0x66, 0x81, 0x67, 0x85, 0x0b, 0x84, 0xdc,
	0xb1, 0x55, 0xdd, 0x59, 0xf3, 0x48, 0xb0, 0xc6, 0x59, 0x97, 0x3a, 0xc7, 0x63, 0x24, 0x4c, 0xb2,
	0x42, 0xdc, 0x4d, 0x6e, 0x85, 0xe6, 0x69, 0x1d, 0x5c, 0xd9, 0x8b, 0x7b, 0x82, 0x5f, 0x0c, 0x30,
	0xf3, 0x9a, 0x0a, 0xa9, 0x9f, 0x56, 0xad, 0xb2, 0xc9, 0x5a, 0x23, 0xe0, 0x2e, 0x39, 0x6c, 0x3c,
	0xa8, 0x8c, 0x15, 0x1c, 0x1d, 0xf4, 0x23, 0xf3, 0x3a, 0x00, 0xb8, 0x27, 0x3b, 0x2d, 0x35, 0x8c,
	0x7e, 0x64, 0xd6, 0xe0, 0xff, 0x21, 0xc1, 0xee, 0x20, 0x32, 0xaf, 0x81, 0x19, 0x8a, 0x7d, 0x5d,
	0x18, 0x44, 0xe6, 0x3c, 0x9c, 0x1d, 0x3d, 0xb6, 0xba, 0x54, 0xc8, 0xfe, 0x8f, 0x5f, 0x9f, 0xa7,
	0x16, 0x60, 0x3d, 0x97, 0x07, 0xf8, 0xd5, 0x00, 0x57, 0x5f, 0x11, 0xf5, 0x21, 0x78, 0xaf, 0xbc,
	0xa1, 0x21, 0x2c, 0x6e, 0xfd, 0x7e, 0x45, 0xa4, 0xe0, 0xc8, 0xed, 0x47, 0x66, 0x1d, 0xcc, 0x25,
	0x8d, 0xb7, 0x4e, 0xa8, 0xfb, 0x21, 0xd3, 0xfd, 0x3c, 0x48, 0xb5, 0x1b, 0x57, 0x07, 0x91, 0x39,
	0x07, 0x6f, 0x24, 0xef, 0x3c, 0xa2, 0x15, 0x98, 0x70, 0x39, 0xa7, 0xc0, 0x8e, 0x09, 0xf0, 0xbb,
	0x01, 0xc0, 0x56, 0x48, 0xb0, 0x24, 0x4a, 0xc9, 0x05, 0xa3, 0x4d, 0x90, 0xb1, 0x98, 0x87, 0xd5,
	0xc1, 0x82, 0x23, 0xa7, 0xc0, 0x88, 0x69, 0x58, 0x73, 0x14, 0x2a, 0x6f, 0x05, 0x84, 0xf3, 0x89,
	0x0e, 0x0d, 0xd2, 0x66, 0xa0, 0xbc, 0x19, 0xaa, 0xf0, 0x5f, 0xcb, 0x58, 0x85, 0x3f, 0x0d, 0x00,
	0xf6, 0xb9, 0x5b, 0x51, 0x4e, 0x82, 0xac, 0x20, 0x27, 0x0d, 0x16, 0x1c, 0x85, 0xe3, 0xec, 0x99,
	0x86, 0xb5, 0x9e, 0x82, 0x8e, 0x31, 0x28, 0x23, 0x4c, 0x23, 0xb5, 0x47, 0x8d, 0x31, 0x1e, 0x25,
	0xea, 0x4e, 0x0d, 0x00, 0xb6, 0x49, 0x97, 0x54, 0x53, 0x97, 0x20, 0x2b, 0xa8, 0x4b, 0x83, 0x05,
	0x47, 0xb4, 0x44, 0x9d, 0xab, 0xa0, 0x55, 0xd4, 0x69, 0xa4, 0x56, 0xb7, 0x3a, 0x2e, 0x81, 0x7f,
	0x0d, 0x50, 0x1f, 0xe6, 0x7e, 0x87, 0xf8, 0x6d, 0x12, 0x8a, 0x0e, 0xe5, 0xb0, 0x59, 0xe9, 0x47,
	0x49, 0x08, 0xb1, 0xc4, 0xf5, 0x89, 0x39, 0x82, 0xa3, 0x93, 0x7e, 0x64, 0xde, 0x04, 0x4b, 0xe7,
	0x94, 0xb6, 0x7a, 0x82, 0x84, 0xd9, 0xad, 0x62, 0x19, 0x2c, 0x66, 0xd5, 0x6a, 0xcc, 0x20, 0x32,
	0x17, 0x21, 0x3c, 0xab, 0xec, 0xc7, 0xaf, 0x92, 0x9d, 0xe3, 0x16, 0x5c, 0x29, 0x56, 0x6d, 0x2b,
	0x2e, 0xfc, 0x6d, 0x80, 0xd9, 0x4d, 0xd7, 0x4d, 0xb5, 0x05, 0xed, 0x72, 0x11, 0x59, 0x74, 0xac,
	0xfa, 0xd1, 0x64, 0x04, 0xc1, 0xd1, 0x47, 0xa3, 0xc4, 0xdd, 0xd1, 0xff, 0x58, 0xe4, 0xee, 0x32,
	0x5c, 0xcc, 0x2a, 0x4d, 0xfd, 0x98, 0x08, 0xdd, 0x2e, 0xd3, 0xda, 0xc2, 0xae, 0x9b, 0x24, 0xb9,
	0x3f, 0x05, 0xea, 0xbb, 0xc4, 0x67, 0x47, 0x24, 0xad, 0xfd, 0x02, 0xd3, 0x73, 0x84, 0x0a, 0xa6,
	0x17, 0x70, 0x04, 0x47, 0x9f, 0x8c, 0xcb, 0xe6, 0x3b, 0x37, 0x81, 0x54, 0xc6, 0xef, 0xa2, 0x3b,
	0xa5, 0x13, 0x08, 0x55, 0x2f, 0xc9, 0x10, 0xfe, 0x24, 0xc9, 0x17, 0x2f, 0x59, 0xa8, 0x1b, 0xac,
	0x98, 0xfc, 0x84, 0x50, 0x3d, 0xf9, 0x69, 0x8e, 0xe0, 0xe8, 0x7d, 0x2a, 0xf9, 0xba, 0x41, 0x95,
	0xea, 0xfc, 0x21, 0x39, 0x4c, 0xfe, 0x79, 0xcc, 0x20, 0x32, 0x17, 0x60, 0xfd, 0xac, 0xb2, 0x97,
	0x39, 0x70, 0xce, 0x05, 0x5f, 0x51, 0xf5, 0x28, 0xf4, 0xe9, 0xf9, 0xcd, 0x00, 0x4b, 0x9b, 0x9c,
	0x77, 0x8f, 0x0f, 0x9a, 0xdb, 0x58, 0xe2, 0x1d, 0xea, 0x85, 0xfa, 0x9a, 0x00, 0x37, 0x2e, 0x88,
	0x73, 0x11, 0x29, 0x1e, 0xc1, 0xe3, 0x4b, 0xf1, 0x04, 0x47, 0x9b, 0xc5, 0xe7, 0xd2, 0x68, 0x0f,
	0x2f, 0x3b, 0x97, 0x34, 0xe8, 0xf9, 0x8b, 0x37, 0x5b, 0x1e, 0x95, 0x9d, 0x5e, 0xdb, 0x72, 0x98,
	0x6f, 0xc7, 0x7d, 0x8c, 0x6e, 0x47, 0x76, 0xf5, 0x4b, 0x5b, 0xbb, 0xa6, 0xae, 0x48, 0xeb, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x13, 0xee, 0x22, 0x9d, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeamsClient is the client API for Teams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamsClient interface {
	ListTeams(ctx context.Context, in *request.ListTeamsReq, opts ...grpc.CallOption) (*response.ListTeamsResp, error)
	GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error)
	CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error)
	UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error)
	DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error)
	GetTeamMembership(ctx context.Context, in *request.GetTeamMembershipReq, opts ...grpc.CallOption) (*response.GetTeamMembershipResp, error)
	AddTeamMembers(ctx context.Context, in *request.AddTeamMembersReq, opts ...grpc.CallOption) (*response.AddTeamMembersResp, error)
	RemoveTeamMembers(ctx context.Context, in *request.RemoveTeamMembersReq, opts ...grpc.CallOption) (*response.RemoveTeamMembersResp, error)
	GetTeamsForMember(ctx context.Context, in *request.GetTeamsForMemberReq, opts ...grpc.CallOption) (*response.GetTeamsForMemberResp, error)
	// Expose on GRPC API only so we don't expose this to the enduser.
	// Just want to be able to trigger this via automate-cli.
	ApplyV2DataMigrations(ctx context.Context, in *request.ApplyV2DataMigrationsReq, opts ...grpc.CallOption) (*response.ApplyV2DataMigrationsResp, error)
}

type teamsClient struct {
	cc *grpc.ClientConn
}

func NewTeamsClient(cc *grpc.ClientConn) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) ListTeams(ctx context.Context, in *request.ListTeamsReq, opts ...grpc.CallOption) (*response.ListTeamsResp, error) {
	out := new(response.ListTeamsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/ListTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error) {
	out := new(response.GetTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error) {
	out := new(response.CreateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error) {
	out := new(response.UpdateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error) {
	out := new(response.DeleteTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamMembership(ctx context.Context, in *request.GetTeamMembershipReq, opts ...grpc.CallOption) (*response.GetTeamMembershipResp, error) {
	out := new(response.GetTeamMembershipResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/GetTeamMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) AddTeamMembers(ctx context.Context, in *request.AddTeamMembersReq, opts ...grpc.CallOption) (*response.AddTeamMembersResp, error) {
	out := new(response.AddTeamMembersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/AddTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) RemoveTeamMembers(ctx context.Context, in *request.RemoveTeamMembersReq, opts ...grpc.CallOption) (*response.RemoveTeamMembersResp, error) {
	out := new(response.RemoveTeamMembersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/RemoveTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamsForMember(ctx context.Context, in *request.GetTeamsForMemberReq, opts ...grpc.CallOption) (*response.GetTeamsForMemberResp, error) {
	out := new(response.GetTeamsForMemberResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/GetTeamsForMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) ApplyV2DataMigrations(ctx context.Context, in *request.ApplyV2DataMigrationsReq, opts ...grpc.CallOption) (*response.ApplyV2DataMigrationsResp, error) {
	out := new(response.ApplyV2DataMigrationsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Teams/ApplyV2DataMigrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServer is the server API for Teams service.
type TeamsServer interface {
	ListTeams(context.Context, *request.ListTeamsReq) (*response.ListTeamsResp, error)
	GetTeam(context.Context, *request.GetTeamReq) (*response.GetTeamResp, error)
	CreateTeam(context.Context, *request.CreateTeamReq) (*response.CreateTeamResp, error)
	UpdateTeam(context.Context, *request.UpdateTeamReq) (*response.UpdateTeamResp, error)
	DeleteTeam(context.Context, *request.DeleteTeamReq) (*response.DeleteTeamResp, error)
	GetTeamMembership(context.Context, *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error)
	AddTeamMembers(context.Context, *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error)
	RemoveTeamMembers(context.Context, *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error)
	GetTeamsForMember(context.Context, *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error)
	// Expose on GRPC API only so we don't expose this to the enduser.
	// Just want to be able to trigger this via automate-cli.
	ApplyV2DataMigrations(context.Context, *request.ApplyV2DataMigrationsReq) (*response.ApplyV2DataMigrationsResp, error)
}

// UnimplementedTeamsServer can be embedded to have forward compatible implementations.
type UnimplementedTeamsServer struct {
}

func (*UnimplementedTeamsServer) ListTeams(ctx context.Context, req *request.ListTeamsReq) (*response.ListTeamsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeams not implemented")
}
func (*UnimplementedTeamsServer) GetTeam(ctx context.Context, req *request.GetTeamReq) (*response.GetTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (*UnimplementedTeamsServer) CreateTeam(ctx context.Context, req *request.CreateTeamReq) (*response.CreateTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (*UnimplementedTeamsServer) UpdateTeam(ctx context.Context, req *request.UpdateTeamReq) (*response.UpdateTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (*UnimplementedTeamsServer) DeleteTeam(ctx context.Context, req *request.DeleteTeamReq) (*response.DeleteTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (*UnimplementedTeamsServer) GetTeamMembership(ctx context.Context, req *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamMembership not implemented")
}
func (*UnimplementedTeamsServer) AddTeamMembers(ctx context.Context, req *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeamMembers not implemented")
}
func (*UnimplementedTeamsServer) RemoveTeamMembers(ctx context.Context, req *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTeamMembers not implemented")
}
func (*UnimplementedTeamsServer) GetTeamsForMember(ctx context.Context, req *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsForMember not implemented")
}
func (*UnimplementedTeamsServer) ApplyV2DataMigrations(ctx context.Context, req *request.ApplyV2DataMigrationsReq) (*response.ApplyV2DataMigrationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyV2DataMigrations not implemented")
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_ListTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListTeamsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).ListTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/ListTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).ListTeams(ctx, req.(*request.ListTeamsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeam(ctx, req.(*request.GetTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).CreateTeam(ctx, req.(*request.CreateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).UpdateTeam(ctx, req.(*request.UpdateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).DeleteTeam(ctx, req.(*request.DeleteTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamMembershipReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/GetTeamMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamMembership(ctx, req.(*request.GetTeamMembershipReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_AddTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.AddTeamMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).AddTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/AddTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).AddTeamMembers(ctx, req.(*request.AddTeamMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_RemoveTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.RemoveTeamMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).RemoveTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/RemoveTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).RemoveTeamMembers(ctx, req.(*request.RemoveTeamMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamsForMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamsForMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamsForMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/GetTeamsForMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamsForMember(ctx, req.(*request.GetTeamsForMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_ApplyV2DataMigrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyV2DataMigrationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).ApplyV2DataMigrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Teams/ApplyV2DataMigrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).ApplyV2DataMigrations(ctx, req.(*request.ApplyV2DataMigrationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2beta.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTeams",
			Handler:    _Teams_ListTeams_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _Teams_GetTeam_Handler,
		},
		{
			MethodName: "CreateTeam",
			Handler:    _Teams_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Teams_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Teams_DeleteTeam_Handler,
		},
		{
			MethodName: "GetTeamMembership",
			Handler:    _Teams_GetTeamMembership_Handler,
		},
		{
			MethodName: "AddTeamMembers",
			Handler:    _Teams_AddTeamMembers_Handler,
		},
		{
			MethodName: "RemoveTeamMembers",
			Handler:    _Teams_RemoveTeamMembers_Handler,
		},
		{
			MethodName: "GetTeamsForMember",
			Handler:    _Teams_GetTeamsForMember_Handler,
		},
		{
			MethodName: "ApplyV2DataMigrations",
			Handler:    _Teams_ApplyV2DataMigrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2beta/teams.proto",
}
