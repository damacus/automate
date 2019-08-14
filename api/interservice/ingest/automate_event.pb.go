// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/ingest/automate_event.proto

package ingest

import (
	context "context"
	fmt "fmt"
	event "github.com/chef/automate/api/interservice/event"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProjectUpdateStatusReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectUpdateStatusReq) Reset()         { *m = ProjectUpdateStatusReq{} }
func (m *ProjectUpdateStatusReq) String() string { return proto.CompactTextString(m) }
func (*ProjectUpdateStatusReq) ProtoMessage()    {}
func (*ProjectUpdateStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dc888f2a4746a, []int{0}
}

func (m *ProjectUpdateStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectUpdateStatusReq.Unmarshal(m, b)
}
func (m *ProjectUpdateStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectUpdateStatusReq.Marshal(b, m, deterministic)
}
func (m *ProjectUpdateStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectUpdateStatusReq.Merge(m, src)
}
func (m *ProjectUpdateStatusReq) XXX_Size() int {
	return xxx_messageInfo_ProjectUpdateStatusReq.Size(m)
}
func (m *ProjectUpdateStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectUpdateStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectUpdateStatusReq proto.InternalMessageInfo

type ProjectUpdateStatusResp struct {
	State                 string               `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty" toml:"state,omitempty" mapstructure:"state,omitempty"`
	EstimatedTimeComplete *timestamp.Timestamp `protobuf:"bytes,2,opt,name=estimated_time_complete,json=estimatedTimeComplete,proto3" json:"estimated_time_complete,omitempty" toml:"estimated_time_complete,omitempty" mapstructure:"estimated_time_complete,omitempty"`
	PercentageComplete    float32              `protobuf:"fixed32,3,opt,name=percentage_complete,json=percentageComplete,proto3" json:"percentage_complete,omitempty" toml:"percentage_complete,omitempty" mapstructure:"percentage_complete,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized      []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache         int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectUpdateStatusResp) Reset()         { *m = ProjectUpdateStatusResp{} }
func (m *ProjectUpdateStatusResp) String() string { return proto.CompactTextString(m) }
func (*ProjectUpdateStatusResp) ProtoMessage()    {}
func (*ProjectUpdateStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dc888f2a4746a, []int{1}
}

func (m *ProjectUpdateStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectUpdateStatusResp.Unmarshal(m, b)
}
func (m *ProjectUpdateStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectUpdateStatusResp.Marshal(b, m, deterministic)
}
func (m *ProjectUpdateStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectUpdateStatusResp.Merge(m, src)
}
func (m *ProjectUpdateStatusResp) XXX_Size() int {
	return xxx_messageInfo_ProjectUpdateStatusResp.Size(m)
}
func (m *ProjectUpdateStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectUpdateStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectUpdateStatusResp proto.InternalMessageInfo

func (m *ProjectUpdateStatusResp) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ProjectUpdateStatusResp) GetEstimatedTimeComplete() *timestamp.Timestamp {
	if m != nil {
		return m.EstimatedTimeComplete
	}
	return nil
}

func (m *ProjectUpdateStatusResp) GetPercentageComplete() float32 {
	if m != nil {
		return m.PercentageComplete
	}
	return 0
}

func init() {
	proto.RegisterType((*ProjectUpdateStatusReq)(nil), "chef.automate.domain.ingest.ProjectUpdateStatusReq")
	proto.RegisterType((*ProjectUpdateStatusResp)(nil), "chef.automate.domain.ingest.ProjectUpdateStatusResp")
}

func init() {
	proto.RegisterFile("api/interservice/ingest/automate_event.proto", fileDescriptor_b34dc888f2a4746a)
}

var fileDescriptor_b34dc888f2a4746a = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0xe5, 0x7e, 0xfa, 0x90, 0x70, 0x99, 0x5c, 0xa0, 0x55, 0x18, 0xa8, 0x3a, 0x65, 0x80,
	0x33, 0x6a, 0x79, 0x02, 0x10, 0x12, 0x0b, 0x12, 0x0a, 0xb0, 0xb0, 0x54, 0x6e, 0x72, 0x4d, 0x8d,
	0x12, 0xdb, 0xc4, 0x97, 0xce, 0x3c, 0x14, 0x2f, 0xc7, 0x86, 0x92, 0x90, 0x52, 0x89, 0x80, 0x60,
	0x4b, 0xee, 0xee, 0x77, 0xf7, 0xf7, 0xff, 0x8e, 0x9f, 0x28, 0xa7, 0xa5, 0x36, 0x84, 0x85, 0xc7,
	0x62, 0xad, 0x63, 0x94, 0xda, 0xa4, 0xe8, 0x49, 0xaa, 0x92, 0x6c, 0xae, 0x08, 0xe7, 0xb8, 0x46,
	0x43, 0xe0, 0x0a, 0x4b, 0x56, 0x1c, 0xc5, 0x2b, 0x5c, 0x42, 0x9b, 0x82, 0xc4, 0xe6, 0x4a, 0x1b,
	0x68, 0x88, 0xe0, 0x38, 0xb5, 0x36, 0xcd, 0x50, 0xd6, 0xa5, 0x8b, 0x72, 0x29, 0x49, 0xe7, 0xe8,
	0x49, 0xe5, 0xae, 0xa1, 0x83, 0xc9, 0x97, 0x59, 0x75, 0x6f, 0xb9, 0x35, 0x61, 0x32, 0xe2, 0x87,
	0xb7, 0x85, 0x7d, 0xc2, 0x98, 0x1e, 0x5c, 0xa2, 0x08, 0xef, 0x48, 0x51, 0xe9, 0x23, 0x7c, 0x9e,
	0xbc, 0x32, 0x3e, 0xec, 0x4c, 0x79, 0x27, 0xf6, 0xf9, 0x7f, 0x4f, 0x8a, 0x70, 0xc4, 0xc6, 0x2c,
	0xdc, 0x8d, 0x9a, 0x1f, 0x11, 0xf1, 0x21, 0x7a, 0xd2, 0x95, 0xd4, 0x64, 0x5e, 0x89, 0x99, 0xc7,
	0x36, 0x77, 0x19, 0x12, 0x8e, 0x7a, 0x63, 0x16, 0xf6, 0xa7, 0x01, 0x34, 0x92, 0xa1, 0x95, 0x0c,
	0xf7, 0xad, 0xe4, 0xe8, 0x60, 0x83, 0x56, 0xb1, 0xcb, 0x0f, 0x50, 0x48, 0x3e, 0x70, 0x58, 0xc4,
	0x68, 0x48, 0xa5, 0x5b, 0xfd, 0xfe, 0x8d, 0x59, 0xd8, 0x8b, 0xc4, 0x67, 0xaa, 0x05, 0xa6, 0x6f,
	0x8c, 0xef, 0x5d, 0x55, 0x0f, 0xbc, 0x56, 0x26, 0xc9, 0xb0, 0x10, 0x4b, 0xde, 0x6f, 0x3e, 0xeb,
	0xa8, 0x08, 0xa1, 0xd3, 0xd3, 0xc6, 0x13, 0xe5, 0x34, 0xd4, 0x65, 0x37, 0x3e, 0x0d, 0x4e, 0x7f,
	0x55, 0x59, 0xf9, 0x61, 0x8d, 0x47, 0xf1, 0xc2, 0xf8, 0xa0, 0xc3, 0x2f, 0x31, 0x83, 0x1f, 0x96,
	0x08, 0xdd, 0xe6, 0x07, 0xe7, 0x7f, 0x87, 0xbc, 0xbb, 0x98, 0x3e, 0x9e, 0xa5, 0x9a, 0x56, 0xe5,
	0x02, 0x62, 0x9b, 0xcb, 0xaa, 0xc3, 0xe6, 0xac, 0xe4, 0x37, 0x77, 0xb7, 0xd8, 0xa9, 0x77, 0x31,
	0x7b, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x42, 0x88, 0xc8, 0x72, 0x99, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventHandlerClient is the client API for EventHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventHandlerClient interface {
	HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error)
	ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error)
}

type eventHandlerClient struct {
	cc *grpc.ClientConn
}

func NewEventHandlerClient(cc *grpc.ClientConn) EventHandlerClient {
	return &eventHandlerClient{cc}
}

func (c *eventHandlerClient) HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error) {
	out := new(event.EventResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.EventHandler/HandleEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventHandlerClient) ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error) {
	out := new(ProjectUpdateStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.EventHandler/ProjectUpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventHandlerServer is the server API for EventHandler service.
type EventHandlerServer interface {
	HandleEvent(context.Context, *event.EventMsg) (*event.EventResponse, error)
	ProjectUpdateStatus(context.Context, *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error)
}

// UnimplementedEventHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedEventHandlerServer struct {
}

func (*UnimplementedEventHandlerServer) HandleEvent(ctx context.Context, req *event.EventMsg) (*event.EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleEvent not implemented")
}
func (*UnimplementedEventHandlerServer) ProjectUpdateStatus(ctx context.Context, req *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectUpdateStatus not implemented")
}

func RegisterEventHandlerServer(s *grpc.Server, srv EventHandlerServer) {
	s.RegisterService(&_EventHandler_serviceDesc, srv)
}

func _EventHandler_HandleEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(event.EventMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHandlerServer).HandleEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.EventHandler/HandleEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHandlerServer).HandleEvent(ctx, req.(*event.EventMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventHandler_ProjectUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectUpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHandlerServer).ProjectUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.EventHandler/ProjectUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHandlerServer).ProjectUpdateStatus(ctx, req.(*ProjectUpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.ingest.EventHandler",
	HandlerType: (*EventHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleEvent",
			Handler:    _EventHandler_HandleEvent_Handler,
		},
		{
			MethodName: "ProjectUpdateStatus",
			Handler:    _EventHandler_ProjectUpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/ingest/automate_event.proto",
}
