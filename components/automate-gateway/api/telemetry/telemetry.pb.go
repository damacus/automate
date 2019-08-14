// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/telemetry/telemetry.proto

package telemetry

import (
	context "context"
	fmt "fmt"
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

type TelemetryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryRequest) Reset()         { *m = TelemetryRequest{} }
func (m *TelemetryRequest) String() string { return proto.CompactTextString(m) }
func (*TelemetryRequest) ProtoMessage()    {}
func (*TelemetryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d8c6a4fd968090, []int{0}
}

func (m *TelemetryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryRequest.Unmarshal(m, b)
}
func (m *TelemetryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryRequest.Marshal(b, m, deterministic)
}
func (m *TelemetryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryRequest.Merge(m, src)
}
func (m *TelemetryRequest) XXX_Size() int {
	return xxx_messageInfo_TelemetryRequest.Size(m)
}
func (m *TelemetryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryRequest proto.InternalMessageInfo

type TelemetryResponse struct {
	LicenseId            string   `protobuf:"bytes,1,opt,name=license_id,json=licenseId,proto3" json:"license_id,omitempty"`
	CustomerName         string   `protobuf:"bytes,4,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	CustomerId           string   `protobuf:"bytes,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	LicenseType          string   `protobuf:"bytes,6,opt,name=license_type,json=licenseType,proto3" json:"license_type,omitempty"`
	TelemetryEnabled     bool     `protobuf:"varint,7,opt,name=telemetry_enabled,json=telemetryEnabled,proto3" json:"telemetry_enabled,omitempty"`
	TelemetryUrl         string   `protobuf:"bytes,8,opt,name=telemetry_url,json=telemetryUrl,proto3" json:"telemetry_url,omitempty"`
	MaxNodes             int64    `protobuf:"varint,9,opt,name=max_nodes,json=maxNodes,proto3" json:"max_nodes,omitempty"`
	DeploymentId         string   `protobuf:"bytes,10,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryResponse) Reset()         { *m = TelemetryResponse{} }
func (m *TelemetryResponse) String() string { return proto.CompactTextString(m) }
func (*TelemetryResponse) ProtoMessage()    {}
func (*TelemetryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d8c6a4fd968090, []int{1}
}

func (m *TelemetryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryResponse.Unmarshal(m, b)
}
func (m *TelemetryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryResponse.Marshal(b, m, deterministic)
}
func (m *TelemetryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryResponse.Merge(m, src)
}
func (m *TelemetryResponse) XXX_Size() int {
	return xxx_messageInfo_TelemetryResponse.Size(m)
}
func (m *TelemetryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryResponse proto.InternalMessageInfo

func (m *TelemetryResponse) GetLicenseId() string {
	if m != nil {
		return m.LicenseId
	}
	return ""
}

func (m *TelemetryResponse) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *TelemetryResponse) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *TelemetryResponse) GetLicenseType() string {
	if m != nil {
		return m.LicenseType
	}
	return ""
}

func (m *TelemetryResponse) GetTelemetryEnabled() bool {
	if m != nil {
		return m.TelemetryEnabled
	}
	return false
}

func (m *TelemetryResponse) GetTelemetryUrl() string {
	if m != nil {
		return m.TelemetryUrl
	}
	return ""
}

func (m *TelemetryResponse) GetMaxNodes() int64 {
	if m != nil {
		return m.MaxNodes
	}
	return 0
}

func (m *TelemetryResponse) GetDeploymentId() string {
	if m != nil {
		return m.DeploymentId
	}
	return ""
}

func init() {
	proto.RegisterType((*TelemetryRequest)(nil), "chef.automate.api.telemetry.TelemetryRequest")
	proto.RegisterType((*TelemetryResponse)(nil), "chef.automate.api.telemetry.TelemetryResponse")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/telemetry/telemetry.proto", fileDescriptor_52d8c6a4fd968090)
}

var fileDescriptor_52d8c6a4fd968090 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbf, 0x6e, 0x13, 0x41,
	0x10, 0xc6, 0x75, 0x21, 0x18, 0xdf, 0x26, 0x11, 0xf1, 0xd2, 0x2c, 0x4e, 0x10, 0xc6, 0x34, 0x96,
	0x90, 0xef, 0x24, 0xa8, 0x48, 0x85, 0x02, 0x08, 0xb9, 0x49, 0x61, 0x85, 0x86, 0xc6, 0x5a, 0xef,
	0x4e, 0x2e, 0x2b, 0xed, 0x3f, 0x6e, 0xe7, 0x44, 0xae, 0x75, 0x99, 0x96, 0x07, 0xa1, 0xf2, 0x93,
	0xf0, 0x0a, 0x14, 0x3c, 0x01, 0x35, 0xba, 0x3b, 0xdf, 0x39, 0xfc, 0x11, 0x4a, 0x41, 0x77, 0xf7,
	0xdb, 0x6f, 0xbe, 0xf9, 0x34, 0x3b, 0x4b, 0x5e, 0x0a, 0x67, 0xbc, 0xb3, 0x60, 0x31, 0xa4, 0xbc,
	0x40, 0x67, 0x38, 0xc2, 0x34, 0xe3, 0x08, 0x9f, 0x78, 0x99, 0x72, 0xaf, 0x52, 0x04, 0x0d, 0x06,
	0x30, 0x2f, 0xb7, 0x5f, 0x89, 0xcf, 0x1d, 0x3a, 0x7a, 0x24, 0x2e, 0xe1, 0x22, 0x69, 0x8b, 0x12,
	0xee, 0x55, 0xd2, 0x49, 0x86, 0xc7, 0x99, 0x73, 0x99, 0x86, 0xda, 0x82, 0x5b, 0xeb, 0x90, 0xa3,
	0x72, 0x36, 0x34, 0xa5, 0xc3, 0x57, 0x7f, 0xed, 0x9a, 0x7b, 0x91, 0xd6, 0xe7, 0x62, 0x9a, 0x81,
	0x9d, 0x7a, 0xa7, 0x95, 0x28, 0xff, 0x83, 0x83, 0xe2, 0xe6, 0x4f, 0x87, 0x31, 0x25, 0x87, 0xe7,
	0x6d, 0xdc, 0x39, 0x7c, 0x2c, 0x20, 0xe0, 0xf8, 0xcb, 0x0e, 0x19, 0xdc, 0x80, 0xc1, 0x3b, 0x1b,
	0x80, 0x3e, 0x22, 0x44, 0x2b, 0x01, 0x36, 0xc0, 0x42, 0x49, 0x16, 0x8d, 0xa2, 0x49, 0x3c, 0x8f,
	0x37, 0x64, 0x26, 0xe9, 0x53, 0x72, 0x20, 0x8a, 0x80, 0xce, 0x40, 0xbe, 0xb0, 0xdc, 0x00, 0xdb,
	0xad, 0x15, 0xfb, 0x2d, 0x3c, 0xe3, 0x06, 0xe8, 0x63, 0xb2, 0xd7, 0x89, 0x94, 0x64, 0x77, 0x6b,
	0x09, 0x69, 0xd1, 0x4c, 0xd2, 0x27, 0x64, 0xbf, 0x6d, 0x82, 0xa5, 0x07, 0xd6, 0xab, 0x15, 0x7b,
	0x1b, 0x76, 0x5e, 0x7a, 0xa0, 0xcf, 0xc8, 0xa0, 0x1b, 0xf0, 0x02, 0x2c, 0x5f, 0x6a, 0x90, 0xec,
	0xde, 0x28, 0x9a, 0xf4, 0xe7, 0x87, 0xdd, 0xc1, 0xdb, 0x86, 0x57, 0xa9, 0xb6, 0xe2, 0x22, 0xd7,
	0xac, 0xdf, 0xa4, 0xea, 0xe0, 0xfb, 0x5c, 0xd3, 0x23, 0x12, 0x1b, 0x7e, 0xb5, 0xb0, 0x4e, 0x42,
	0x60, 0xf1, 0x28, 0x9a, 0xdc, 0x99, 0xf7, 0x0d, 0xbf, 0x3a, 0xab, 0xfe, 0x2b, 0x07, 0x09, 0x5e,
	0xbb, 0xd2, 0x80, 0xc5, 0x2a, 0x34, 0x69, 0x1c, 0xb6, 0x70, 0x26, 0x9f, 0xff, 0x88, 0x48, 0xdc,
	0x4d, 0x8c, 0x7e, 0x8f, 0xc8, 0xc3, 0x77, 0x80, 0x1d, 0x78, 0xed, 0xec, 0x85, 0xca, 0x8a, 0xbc,
	0x1e, 0x3c, 0x9d, 0x26, 0xff, 0xd8, 0x98, 0xe4, 0xf7, 0xcb, 0x18, 0x26, 0xb7, 0x95, 0x37, 0xd7,
	0x34, 0xd6, 0xab, 0x35, 0xa3, 0x64, 0x3b, 0x89, 0x13, 0x51, 0x77, 0x5f, 0xad, 0x59, 0x8f, 0xee,
	0xe6, 0xc0, 0xe5, 0xf5, 0x9a, 0xdd, 0x27, 0x07, 0xa1, 0x0c, 0x08, 0x66, 0x73, 0x78, 0xbd, 0x66,
	0xc7, 0x74, 0xb8, 0x41, 0xf8, 0x6b, 0xec, 0x93, 0x0c, 0x70, 0xf5, 0xf5, 0xdb, 0xe7, 0x9d, 0x07,
	0x74, 0x70, 0xe3, 0x19, 0x34, 0x85, 0xa7, 0x6f, 0x3e, 0x9c, 0x66, 0x0a, 0x2f, 0x8b, 0x65, 0x22,
	0x9c, 0x49, 0xab, 0xa4, 0xdd, 0x1e, 0xa6, 0xb7, 0x7e, 0x53, 0xcb, 0x5e, 0xbd, 0x8b, 0x2f, 0x7e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x99, 0x19, 0xc4, 0x65, 0x87, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TelemetryClient is the client API for Telemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TelemetryClient interface {
	GetTelemetryConfiguration(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
}

type telemetryClient struct {
	cc *grpc.ClientConn
}

func NewTelemetryClient(cc *grpc.ClientConn) TelemetryClient {
	return &telemetryClient{cc}
}

func (c *telemetryClient) GetTelemetryConfiguration(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.telemetry.Telemetry/GetTelemetryConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServer is the server API for Telemetry service.
type TelemetryServer interface {
	GetTelemetryConfiguration(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
}

// UnimplementedTelemetryServer can be embedded to have forward compatible implementations.
type UnimplementedTelemetryServer struct {
}

func (*UnimplementedTelemetryServer) GetTelemetryConfiguration(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTelemetryConfiguration not implemented")
}

func RegisterTelemetryServer(s *grpc.Server, srv TelemetryServer) {
	s.RegisterService(&_Telemetry_serviceDesc, srv)
}

func _Telemetry_GetTelemetryConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetTelemetryConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.telemetry.Telemetry/GetTelemetryConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetTelemetryConfiguration(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Telemetry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.telemetry.Telemetry",
	HandlerType: (*TelemetryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTelemetryConfiguration",
			Handler:    _Telemetry_GetTelemetryConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/telemetry/telemetry.proto",
}
