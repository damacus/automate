// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/event_feed/event_feed.proto

package event_feed

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/event_feed/request"
	response "github.com/chef/automate/components/automate-gateway/api/event_feed/response"
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
	proto.RegisterFile("components/automate-gateway/api/event_feed/event_feed.proto", fileDescriptor_fe4f2cb273de4e97)
}

var fileDescriptor_fe4f2cb273de4e97 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0x19, 0xd1, 0xaa, 0x61, 0x5c, 0x35, 0x2a, 0x94, 0xb2, 0xa7, 0xb9, 0x37, 0xc5, 0xd9,
	0x83, 0x30, 0x82, 0xca, 0x2e, 0xb3, 0x9e, 0x45, 0x4f, 0x22, 0xac, 0xd9, 0xcc, 0xdb, 0x6e, 0x68,
	0x9b, 0xc4, 0xe6, 0x55, 0xe9, 0xb5, 0xc7, 0xb9, 0xfa, 0x59, 0xfa, 0x09, 0xfc, 0x08, 0xde, 0xbd,
	0xe8, 0xe7, 0x10, 0x69, 0xd3, 0x3a, 0x53, 0x14, 0xb5, 0xc3, 0x1c, 0xbc, 0x95, 0xf7, 0xfa, 0xff,
	0xe7, 0xff, 0x4b, 0xf2, 0x42, 0x1e, 0x0b, 0x9d, 0x19, 0xad, 0x40, 0xa1, 0x8d, 0x78, 0x81, 0x3a,
	0xe3, 0x08, 0x61, 0xcc, 0x11, 0x3e, 0xf0, 0x32, 0xe2, 0x46, 0x46, 0xf0, 0x1e, 0x14, 0x9e, 0x5d,
	0x00, 0xac, 0xb6, 0x3e, 0x99, 0xc9, 0x35, 0x6a, 0x7a, 0x28, 0x2e, 0xe1, 0x82, 0xf5, 0x32, 0xc6,
	0x8d, 0x64, 0x9b, 0x7f, 0x82, 0x27, 0x23, 0xac, 0x73, 0x78, 0x57, 0x80, 0x45, 0x57, 0x72, 0xee,
	0xc1, 0x72, 0x57, 0xbd, 0xc5, 0x5c, 0xaa, 0xd8, 0x76, 0x36, 0x4f, 0x47, 0xd9, 0x58, 0xa3, 0x95,
	0x85, 0x41, 0x8e, 0xd3, 0x9d, 0x0d, 0x86, 0x41, 0x0e, 0x63, 0xad, 0xe3, 0x14, 0x5a, 0x09, 0x57,
	0x4a, 0x23, 0x47, 0xa9, 0x55, 0xdf, 0x7d, 0xf6, 0xdb, 0x55, 0x72, 0x23, 0xa2, 0xb6, 0x2f, 0xc2,
	0x18, 0x54, 0x68, 0x74, 0x2a, 0x45, 0xb9, 0x07, 0x07, 0xc9, 0xb3, 0x5f, 0x1d, 0xe6, 0xdf, 0x3d,
	0x72, 0x73, 0xd9, 0x04, 0x3f, 0x05, 0x58, 0xd1, 0x4f, 0x13, 0x32, 0x7d, 0x0e, 0xb8, 0x29, 0x3c,
	0x64, 0x7f, 0x3a, 0x6f, 0xd6, 0x9d, 0x01, 0x73, 0x02, 0x99, 0x22, 0xe4, 0x41, 0xf8, 0x37, 0x89,
	0xdb, 0x2e, 0xa7, 0xb1, 0xb3, 0x17, 0x55, 0xed, 0xdf, 0x20, 0x9e, 0xdb, 0xbc, 0xaa, 0xf6, 0x3d,
	0x7a, 0x35, 0x07, 0xbe, 0x5a, 0xd7, 0xfe, 0x01, 0x99, 0xb6, 0xd5, 0x85, 0xeb, 0xad, 0x6b, 0xff,
	0x1e, 0xbd, 0xbb, 0x5d, 0x59, 0xa4, 0xd2, 0x62, 0xf5, 0xf9, 0xdb, 0xc7, 0x2b, 0x53, 0x4a, 0xdc,
	0xf6, 0x37, 0x4b, 0xd0, 0xaf, 0x13, 0x42, 0x7b, 0x8a, 0x57, 0xa5, 0x81, 0x13, 0x5d, 0x28, 0xb4,
	0xf4, 0xd1, 0x08, 0x16, 0x27, 0xe9, 0x88, 0xe6, 0x63, 0x88, 0x9c, 0x72, 0xf6, 0xb6, 0xda, 0x00,
	0xd8, 0x05, 0x96, 0x06, 0x1a, 0xb8, 0xeb, 0xf4, 0x9a, 0x68, 0xfa, 0xe3, 0xe8, 0xee, 0x53, 0xda,
	0x5d, 0xb8, 0xc6, 0xe7, 0x4c, 0x38, 0x9c, 0x01, 0x25, 0xb7, 0xc9, 0x7f, 0x40, 0xc9, 0x6d, 0xb2,
	0x17, 0x4a, 0x6e, 0x93, 0x9e, 0xf2, 0xcb, 0x84, 0x3c, 0xe8, 0x29, 0x5f, 0xb6, 0xb3, 0x75, 0x5c,
	0x88, 0x04, 0xd0, 0xd2, 0xf9, 0x08, 0x50, 0xa7, 0xb4, 0xc1, 0xd1, 0x18, 0xc6, 0x4e, 0x34, 0x7b,
	0x53, 0xd5, 0xfe, 0x1d, 0x72, 0xd0, 0xe5, 0xed, 0x06, 0x7c, 0xd7, 0x9b, 0x7a, 0x9b, 0xde, 0x1a,
	0x3c, 0x14, 0xc7, 0xcb, 0xd7, 0x27, 0xb1, 0xc4, 0xcb, 0xe2, 0x9c, 0x09, 0x9d, 0x45, 0x4d, 0xbc,
	0x9f, 0x93, 0x1c, 0xfd, 0xfb, 0x2b, 0x74, 0xee, 0xb5, 0xe3, 0x7c, 0xf4, 0x23, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x22, 0x51, 0x49, 0xdd, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventFeedClient is the client API for EventFeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventFeedClient interface {
	GetEventFeed(ctx context.Context, in *request.EventFilter, opts ...grpc.CallOption) (*response.Events, error)
	GetEventTypeCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error)
	GetEventTaskCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error)
	GetEventStringBuckets(ctx context.Context, in *request.EventStrings, opts ...grpc.CallOption) (*response.EventStrings, error)
}

type eventFeedClient struct {
	cc *grpc.ClientConn
}

func NewEventFeedClient(cc *grpc.ClientConn) EventFeedClient {
	return &eventFeedClient{cc}
}

func (c *eventFeedClient) GetEventFeed(ctx context.Context, in *request.EventFilter, opts ...grpc.CallOption) (*response.Events, error) {
	out := new(response.Events)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventTypeCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error) {
	out := new(response.EventCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventTypeCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventTaskCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error) {
	out := new(response.EventCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventTaskCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventStringBuckets(ctx context.Context, in *request.EventStrings, opts ...grpc.CallOption) (*response.EventStrings, error) {
	out := new(response.EventStrings)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventStringBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventFeedServer is the server API for EventFeed service.
type EventFeedServer interface {
	GetEventFeed(context.Context, *request.EventFilter) (*response.Events, error)
	GetEventTypeCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	GetEventTaskCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	GetEventStringBuckets(context.Context, *request.EventStrings) (*response.EventStrings, error)
}

// UnimplementedEventFeedServer can be embedded to have forward compatible implementations.
type UnimplementedEventFeedServer struct {
}

func (*UnimplementedEventFeedServer) GetEventFeed(ctx context.Context, req *request.EventFilter) (*response.Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventFeed not implemented")
}
func (*UnimplementedEventFeedServer) GetEventTypeCounts(ctx context.Context, req *request.EventCountsFilter) (*response.EventCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTypeCounts not implemented")
}
func (*UnimplementedEventFeedServer) GetEventTaskCounts(ctx context.Context, req *request.EventCountsFilter) (*response.EventCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTaskCounts not implemented")
}
func (*UnimplementedEventFeedServer) GetEventStringBuckets(ctx context.Context, req *request.EventStrings) (*response.EventStrings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStringBuckets not implemented")
}

func RegisterEventFeedServer(s *grpc.Server, srv EventFeedServer) {
	s.RegisterService(&_EventFeed_serviceDesc, srv)
}

func _EventFeed_GetEventFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventFeed(ctx, req.(*request.EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventTypeCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventCountsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventTypeCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventTypeCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventTypeCounts(ctx, req.(*request.EventCountsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventTaskCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventCountsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventTaskCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventTaskCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventTaskCounts(ctx, req.(*request.EventCountsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventStringBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventStrings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventStringBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventStringBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventStringBuckets(ctx, req.(*request.EventStrings))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventFeed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.event_feed.EventFeed",
	HandlerType: (*EventFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventFeed",
			Handler:    _EventFeed_GetEventFeed_Handler,
		},
		{
			MethodName: "GetEventTypeCounts",
			Handler:    _EventFeed_GetEventTypeCounts_Handler,
		},
		{
			MethodName: "GetEventTaskCounts",
			Handler:    _EventFeed_GetEventTaskCounts_Handler,
		},
		{
			MethodName: "GetEventStringBuckets",
			Handler:    _EventFeed_GetEventStringBuckets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/event_feed/event_feed.proto",
}
