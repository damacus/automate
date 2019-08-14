// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/data_lifecycle/manageable.proto

package data_lifecycle

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PurgeStatus_StatusCode int32

const (
	PurgeStatus_FAILED  PurgeStatus_StatusCode = 0
	PurgeStatus_SUCCESS PurgeStatus_StatusCode = 1
)

var PurgeStatus_StatusCode_name = map[int32]string{
	0: "FAILED",
	1: "SUCCESS",
}

var PurgeStatus_StatusCode_value = map[string]int32{
	"FAILED":  0,
	"SUCCESS": 1,
}

func (x PurgeStatus_StatusCode) String() string {
	return proto.EnumName(PurgeStatus_StatusCode_name, int32(x))
}

func (PurgeStatus_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39f9ec77ebc52179, []int{2, 0}
}

type PurgeRequest struct {
	// This id must be managed. If a request with a duplicate id comes in,
	// the down stream service should prefer waiting for that id to complete
	// rather than redoing the work.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PurgeRequest) Reset()         { *m = PurgeRequest{} }
func (m *PurgeRequest) String() string { return proto.CompactTextString(m) }
func (*PurgeRequest) ProtoMessage()    {}
func (*PurgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f9ec77ebc52179, []int{0}
}

func (m *PurgeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeRequest.Unmarshal(m, b)
}
func (m *PurgeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeRequest.Marshal(b, m, deterministic)
}
func (m *PurgeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeRequest.Merge(m, src)
}
func (m *PurgeRequest) XXX_Size() int {
	return xxx_messageInfo_PurgeRequest.Size(m)
}
func (m *PurgeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeRequest proto.InternalMessageInfo

func (m *PurgeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PurgeResponse struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	ComponentStatus      map[string]*PurgeStatus `protobuf:"bytes,2,rep,name=component_status,json=componentStatus,proto3" json:"component_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" toml:"component_status,omitempty" mapstructure:"component_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                   `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PurgeResponse) Reset()         { *m = PurgeResponse{} }
func (m *PurgeResponse) String() string { return proto.CompactTextString(m) }
func (*PurgeResponse) ProtoMessage()    {}
func (*PurgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f9ec77ebc52179, []int{1}
}

func (m *PurgeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeResponse.Unmarshal(m, b)
}
func (m *PurgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeResponse.Marshal(b, m, deterministic)
}
func (m *PurgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeResponse.Merge(m, src)
}
func (m *PurgeResponse) XXX_Size() int {
	return xxx_messageInfo_PurgeResponse.Size(m)
}
func (m *PurgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeResponse proto.InternalMessageInfo

func (m *PurgeResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PurgeResponse) GetComponentStatus() map[string]*PurgeStatus {
	if m != nil {
		return m.ComponentStatus
	}
	return nil
}

type PurgeStatus struct {
	Status               PurgeStatus_StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=chef.automate.domain.data_lifecycle.PurgeStatus_StatusCode" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
	Msg                  string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" toml:"msg,omitempty" mapstructure:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PurgeStatus) Reset()         { *m = PurgeStatus{} }
func (m *PurgeStatus) String() string { return proto.CompactTextString(m) }
func (*PurgeStatus) ProtoMessage()    {}
func (*PurgeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f9ec77ebc52179, []int{2}
}

func (m *PurgeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeStatus.Unmarshal(m, b)
}
func (m *PurgeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeStatus.Marshal(b, m, deterministic)
}
func (m *PurgeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeStatus.Merge(m, src)
}
func (m *PurgeStatus) XXX_Size() int {
	return xxx_messageInfo_PurgeStatus.Size(m)
}
func (m *PurgeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeStatus proto.InternalMessageInfo

func (m *PurgeStatus) GetStatus() PurgeStatus_StatusCode {
	if m != nil {
		return m.Status
	}
	return PurgeStatus_FAILED
}

func (m *PurgeStatus) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("chef.automate.domain.data_lifecycle.PurgeStatus_StatusCode", PurgeStatus_StatusCode_name, PurgeStatus_StatusCode_value)
	proto.RegisterType((*PurgeRequest)(nil), "chef.automate.domain.data_lifecycle.PurgeRequest")
	proto.RegisterType((*PurgeResponse)(nil), "chef.automate.domain.data_lifecycle.PurgeResponse")
	proto.RegisterMapType((map[string]*PurgeStatus)(nil), "chef.automate.domain.data_lifecycle.PurgeResponse.ComponentStatusEntry")
	proto.RegisterType((*PurgeStatus)(nil), "chef.automate.domain.data_lifecycle.PurgeStatus")
}

func init() {
	proto.RegisterFile("api/interservice/data_lifecycle/manageable.proto", fileDescriptor_39f9ec77ebc52179)
}

var fileDescriptor_39f9ec77ebc52179 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6b, 0xe2, 0x40,
	0x14, 0xdf, 0x44, 0x74, 0xf1, 0x65, 0xd7, 0x0d, 0xc3, 0xc2, 0x8a, 0x87, 0x45, 0xb2, 0x2c, 0x78,
	0x9a, 0xb8, 0xd9, 0x8b, 0xb8, 0xa7, 0x6d, 0xd4, 0x52, 0xb0, 0x50, 0x12, 0x7a, 0xe9, 0x45, 0xc6,
	0xe4, 0x19, 0x87, 0x26, 0x33, 0x69, 0x32, 0x11, 0xfc, 0x08, 0x7e, 0x86, 0x7e, 0xd9, 0x92, 0x18,
	0xdb, 0x5a, 0x0a, 0xd5, 0x53, 0x1e, 0x99, 0xdf, 0xbf, 0xc7, 0xfb, 0xc1, 0x90, 0xa5, 0xdc, 0xe6,
	0x42, 0x61, 0x96, 0x63, 0xb6, 0xe1, 0x01, 0xda, 0x21, 0x53, 0x6c, 0x11, 0xf3, 0x15, 0x06, 0xdb,
	0x20, 0x46, 0x3b, 0x61, 0x82, 0x45, 0xc8, 0x96, 0x31, 0xd2, 0x34, 0x93, 0x4a, 0x92, 0x5f, 0xc1,
	0x1a, 0x57, 0x94, 0x15, 0x4a, 0x26, 0x4c, 0x21, 0x0d, 0x65, 0xc2, 0xb8, 0xa0, 0xc7, 0x2c, 0xeb,
	0x27, 0x7c, 0xb9, 0x29, 0xb2, 0x08, 0x3d, 0x7c, 0x28, 0x30, 0x57, 0xa4, 0x03, 0x3a, 0x0f, 0xbb,
	0x5a, 0x5f, 0x1b, 0xb4, 0x3d, 0x9d, 0x87, 0xd6, 0x4e, 0x87, 0xaf, 0x35, 0x20, 0x4f, 0xa5, 0xc8,
	0xf1, 0x2d, 0x82, 0x64, 0x60, 0x06, 0x32, 0x49, 0xa5, 0x40, 0xa1, 0x16, 0xb9, 0x62, 0xaa, 0xc8,
	0xbb, 0x7a, 0xbf, 0x31, 0x30, 0x9c, 0x4b, 0x7a, 0x42, 0x02, 0x7a, 0xa4, 0x4e, 0xdd, 0x83, 0x94,
	0x5f, 0x29, 0x4d, 0x85, 0xca, 0xb6, 0xde, 0xb7, 0xe0, 0xf8, 0x6f, 0x4f, 0xc1, 0xf7, 0xf7, 0x80,
	0xc4, 0x84, 0xc6, 0x3d, 0x6e, 0xeb, 0x70, 0xe5, 0x48, 0x66, 0xd0, 0xdc, 0xb0, 0xb8, 0xc0, 0xae,
	0xde, 0xd7, 0x06, 0x86, 0x33, 0x3c, 0x3d, 0xd2, 0x5e, 0xd7, 0xdb, 0xd3, 0xc7, 0xfa, 0x48, 0xb3,
	0x1e, 0x35, 0x30, 0x5e, 0x3d, 0x11, 0x1f, 0x5a, 0xf5, 0xbe, 0xa5, 0x61, 0xc7, 0xf9, 0x77, 0xae,
	0x38, 0xdd, 0x7f, 0x5c, 0x19, 0xa2, 0x57, 0x4b, 0x95, 0x2b, 0x24, 0x79, 0x54, 0xc5, 0x6d, 0x7b,
	0xe5, 0x68, 0xfd, 0x06, 0x78, 0xc1, 0x11, 0x80, 0xd6, 0xec, 0xff, 0xd5, 0x7c, 0x3a, 0x31, 0x3f,
	0x11, 0x03, 0x3e, 0xfb, 0xb7, 0xae, 0x3b, 0xf5, 0x7d, 0x53, 0x73, 0x76, 0x1a, 0xfc, 0x98, 0x30,
	0xc5, 0xe6, 0x07, 0xa7, 0xeb, 0xe7, 0x42, 0x10, 0x01, 0xcd, 0xca, 0x96, 0xfc, 0x39, 0xe7, 0x24,
	0x55, 0x23, 0x7a, 0xce, 0xf9, 0x57, 0xbc, 0x18, 0xdf, 0x8d, 0x22, 0xae, 0xd6, 0xc5, 0x92, 0x06,
	0x32, 0xb1, 0x4b, 0xbe, 0x7d, 0xe0, 0xdb, 0x1f, 0xf4, 0x78, 0xd9, 0xaa, 0xda, 0xfb, 0xf7, 0x29,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0x23, 0x24, 0x91, 0xf1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataLifecycleManageableClient is the client API for DataLifecycleManageable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataLifecycleManageableClient interface {
	// We could make this streaming and return how much is
	// finished
	Purge(ctx context.Context, in *PurgeRequest, opts ...grpc.CallOption) (*PurgeResponse, error)
}

type dataLifecycleManageableClient struct {
	cc *grpc.ClientConn
}

func NewDataLifecycleManageableClient(cc *grpc.ClientConn) DataLifecycleManageableClient {
	return &dataLifecycleManageableClient{cc}
}

func (c *dataLifecycleManageableClient) Purge(ctx context.Context, in *PurgeRequest, opts ...grpc.CallOption) (*PurgeResponse, error) {
	out := new(PurgeResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.data_lifecycle.DataLifecycleManageable/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataLifecycleManageableServer is the server API for DataLifecycleManageable service.
type DataLifecycleManageableServer interface {
	// We could make this streaming and return how much is
	// finished
	Purge(context.Context, *PurgeRequest) (*PurgeResponse, error)
}

// UnimplementedDataLifecycleManageableServer can be embedded to have forward compatible implementations.
type UnimplementedDataLifecycleManageableServer struct {
}

func (*UnimplementedDataLifecycleManageableServer) Purge(ctx context.Context, req *PurgeRequest) (*PurgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}

func RegisterDataLifecycleManageableServer(s *grpc.Server, srv DataLifecycleManageableServer) {
	s.RegisterService(&_DataLifecycleManageable_serviceDesc, srv)
}

func _DataLifecycleManageable_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLifecycleManageableServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.data_lifecycle.DataLifecycleManageable/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLifecycleManageableServer).Purge(ctx, req.(*PurgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataLifecycleManageable_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.data_lifecycle.DataLifecycleManageable",
	HandlerType: (*DataLifecycleManageableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Purge",
			Handler:    _DataLifecycleManageable_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/data_lifecycle/manageable.proto",
}
