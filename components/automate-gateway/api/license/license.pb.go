// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/license/license.proto

package license

import (
	context "context"
	fmt "fmt"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ApplyLicenseReq struct {
	License              string   `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyLicenseReq) Reset()         { *m = ApplyLicenseReq{} }
func (m *ApplyLicenseReq) String() string { return proto.CompactTextString(m) }
func (*ApplyLicenseReq) ProtoMessage()    {}
func (*ApplyLicenseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{0}
}

func (m *ApplyLicenseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyLicenseReq.Unmarshal(m, b)
}
func (m *ApplyLicenseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyLicenseReq.Marshal(b, m, deterministic)
}
func (m *ApplyLicenseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyLicenseReq.Merge(m, src)
}
func (m *ApplyLicenseReq) XXX_Size() int {
	return xxx_messageInfo_ApplyLicenseReq.Size(m)
}
func (m *ApplyLicenseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyLicenseReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyLicenseReq proto.InternalMessageInfo

func (m *ApplyLicenseReq) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

type ApplyLicenseResp struct {
	Status               *GetStatusResp `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApplyLicenseResp) Reset()         { *m = ApplyLicenseResp{} }
func (m *ApplyLicenseResp) String() string { return proto.CompactTextString(m) }
func (*ApplyLicenseResp) ProtoMessage()    {}
func (*ApplyLicenseResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{1}
}

func (m *ApplyLicenseResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyLicenseResp.Unmarshal(m, b)
}
func (m *ApplyLicenseResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyLicenseResp.Marshal(b, m, deterministic)
}
func (m *ApplyLicenseResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyLicenseResp.Merge(m, src)
}
func (m *ApplyLicenseResp) XXX_Size() int {
	return xxx_messageInfo_ApplyLicenseResp.Size(m)
}
func (m *ApplyLicenseResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyLicenseResp.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyLicenseResp proto.InternalMessageInfo

func (m *ApplyLicenseResp) GetStatus() *GetStatusResp {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetStatusReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatusReq) Reset()         { *m = GetStatusReq{} }
func (m *GetStatusReq) String() string { return proto.CompactTextString(m) }
func (*GetStatusReq) ProtoMessage()    {}
func (*GetStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{2}
}

func (m *GetStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusReq.Unmarshal(m, b)
}
func (m *GetStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusReq.Marshal(b, m, deterministic)
}
func (m *GetStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusReq.Merge(m, src)
}
func (m *GetStatusReq) XXX_Size() int {
	return xxx_messageInfo_GetStatusReq.Size(m)
}
func (m *GetStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusReq proto.InternalMessageInfo

type GetStatusResp struct {
	LicenseId            string                   `protobuf:"bytes,1,opt,name=license_id,json=licenseId,proto3" json:"license_id,omitempty"`
	ConfiguredAt         *timestamp.Timestamp     `protobuf:"bytes,2,opt,name=configured_at,json=configuredAt,proto3" json:"configured_at,omitempty"`
	LicensedPeriod       *GetStatusResp_DateRange `protobuf:"bytes,3,opt,name=licensed_period,json=licensedPeriod,proto3" json:"licensed_period,omitempty"`
	CustomerName         string                   `protobuf:"bytes,4,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetStatusResp) Reset()         { *m = GetStatusResp{} }
func (m *GetStatusResp) String() string { return proto.CompactTextString(m) }
func (*GetStatusResp) ProtoMessage()    {}
func (*GetStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{3}
}

func (m *GetStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusResp.Unmarshal(m, b)
}
func (m *GetStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusResp.Marshal(b, m, deterministic)
}
func (m *GetStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusResp.Merge(m, src)
}
func (m *GetStatusResp) XXX_Size() int {
	return xxx_messageInfo_GetStatusResp.Size(m)
}
func (m *GetStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusResp proto.InternalMessageInfo

func (m *GetStatusResp) GetLicenseId() string {
	if m != nil {
		return m.LicenseId
	}
	return ""
}

func (m *GetStatusResp) GetConfiguredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ConfiguredAt
	}
	return nil
}

func (m *GetStatusResp) GetLicensedPeriod() *GetStatusResp_DateRange {
	if m != nil {
		return m.LicensedPeriod
	}
	return nil
}

func (m *GetStatusResp) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

type GetStatusResp_DateRange struct {
	Start                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetStatusResp_DateRange) Reset()         { *m = GetStatusResp_DateRange{} }
func (m *GetStatusResp_DateRange) String() string { return proto.CompactTextString(m) }
func (*GetStatusResp_DateRange) ProtoMessage()    {}
func (*GetStatusResp_DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{3, 0}
}

func (m *GetStatusResp_DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusResp_DateRange.Unmarshal(m, b)
}
func (m *GetStatusResp_DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusResp_DateRange.Marshal(b, m, deterministic)
}
func (m *GetStatusResp_DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusResp_DateRange.Merge(m, src)
}
func (m *GetStatusResp_DateRange) XXX_Size() int {
	return xxx_messageInfo_GetStatusResp_DateRange.Size(m)
}
func (m *GetStatusResp_DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusResp_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusResp_DateRange proto.InternalMessageInfo

func (m *GetStatusResp_DateRange) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *GetStatusResp_DateRange) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type RequestLicenseReq struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	GdprAgree            bool     `protobuf:"varint,4,opt,name=gdpr_agree,json=gdprAgree,proto3" json:"gdpr_agree,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestLicenseReq) Reset()         { *m = RequestLicenseReq{} }
func (m *RequestLicenseReq) String() string { return proto.CompactTextString(m) }
func (*RequestLicenseReq) ProtoMessage()    {}
func (*RequestLicenseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{4}
}

func (m *RequestLicenseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestLicenseReq.Unmarshal(m, b)
}
func (m *RequestLicenseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestLicenseReq.Marshal(b, m, deterministic)
}
func (m *RequestLicenseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestLicenseReq.Merge(m, src)
}
func (m *RequestLicenseReq) XXX_Size() int {
	return xxx_messageInfo_RequestLicenseReq.Size(m)
}
func (m *RequestLicenseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestLicenseReq.DiscardUnknown(m)
}

var xxx_messageInfo_RequestLicenseReq proto.InternalMessageInfo

func (m *RequestLicenseReq) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RequestLicenseReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RequestLicenseReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RequestLicenseReq) GetGdprAgree() bool {
	if m != nil {
		return m.GdprAgree
	}
	return false
}

type RequestLicenseResp struct {
	License              string         `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"` // Deprecated: Do not use.
	Status               *GetStatusResp `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RequestLicenseResp) Reset()         { *m = RequestLicenseResp{} }
func (m *RequestLicenseResp) String() string { return proto.CompactTextString(m) }
func (*RequestLicenseResp) ProtoMessage()    {}
func (*RequestLicenseResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ab9eb5fcc9b9807, []int{5}
}

func (m *RequestLicenseResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestLicenseResp.Unmarshal(m, b)
}
func (m *RequestLicenseResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestLicenseResp.Marshal(b, m, deterministic)
}
func (m *RequestLicenseResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestLicenseResp.Merge(m, src)
}
func (m *RequestLicenseResp) XXX_Size() int {
	return xxx_messageInfo_RequestLicenseResp.Size(m)
}
func (m *RequestLicenseResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestLicenseResp.DiscardUnknown(m)
}

var xxx_messageInfo_RequestLicenseResp proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *RequestLicenseResp) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

func (m *RequestLicenseResp) GetStatus() *GetStatusResp {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ApplyLicenseReq)(nil), "chef.automate.api.license.ApplyLicenseReq")
	proto.RegisterType((*ApplyLicenseResp)(nil), "chef.automate.api.license.ApplyLicenseResp")
	proto.RegisterType((*GetStatusReq)(nil), "chef.automate.api.license.GetStatusReq")
	proto.RegisterType((*GetStatusResp)(nil), "chef.automate.api.license.GetStatusResp")
	proto.RegisterType((*GetStatusResp_DateRange)(nil), "chef.automate.api.license.GetStatusResp.DateRange")
	proto.RegisterType((*RequestLicenseReq)(nil), "chef.automate.api.license.RequestLicenseReq")
	proto.RegisterType((*RequestLicenseResp)(nil), "chef.automate.api.license.RequestLicenseResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/license/license.proto", fileDescriptor_1ab9eb5fcc9b9807)
}

var fileDescriptor_1ab9eb5fcc9b9807 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x95, 0xd3, 0x8f, 0xd4, 0xf3, 0xf2, 0xd5, 0x51, 0x55, 0xcd, 0xf3, 0x6b, 0xf5, 0x90, 0x59,
	0x50, 0xb5, 0x8d, 0x83, 0x8a, 0xd8, 0x74, 0x43, 0x5b, 0x21, 0x21, 0x24, 0x84, 0x90, 0xe9, 0x0a,
	0x16, 0x61, 0x6a, 0xdf, 0xb8, 0x56, 0x63, 0xcf, 0xc4, 0x33, 0x16, 0xca, 0x36, 0x62, 0x95, 0x2d,
	0xbf, 0xc5, 0x7b, 0x7e, 0x01, 0x1b, 0x76, 0x2c, 0x58, 0xf1, 0x2b, 0x58, 0x21, 0x8f, 0xc7, 0x6e,
	0x92, 0xd2, 0x2a, 0x65, 0x15, 0xcd, 0xbd, 0xe7, 0x1e, 0x9f, 0x73, 0xef, 0x9d, 0x09, 0x7a, 0xea,
	0xb1, 0x88, 0xb3, 0x18, 0x62, 0x29, 0x7a, 0x34, 0x95, 0x2c, 0xa2, 0x12, 0xba, 0x01, 0x95, 0xf0,
	0x91, 0x8e, 0x7b, 0x94, 0x87, 0xbd, 0x61, 0xe8, 0x41, 0x2c, 0xa0, 0xfc, 0x75, 0x78, 0xc2, 0x24,
	0xc3, 0xff, 0x7a, 0x97, 0x30, 0x70, 0xca, 0x02, 0x87, 0xf2, 0xd0, 0xd1, 0x00, 0x6b, 0x27, 0x60,
	0x2c, 0x18, 0x82, 0x2a, 0xa6, 0x71, 0xcc, 0x24, 0x95, 0x21, 0x8b, 0x45, 0x51, 0x68, 0xfd, 0xaf,
	0xb3, 0xea, 0x74, 0x91, 0x0e, 0x7a, 0x32, 0x8c, 0x40, 0x48, 0x1a, 0x71, 0x0d, 0x38, 0xf9, 0xa3,
	0xa0, 0x84, 0x7b, 0x45, 0x89, 0xd7, 0x0d, 0x20, 0xee, 0x72, 0x36, 0x0c, 0xbd, 0xf1, 0x2d, 0x9f,
	0xb8, 0x0f, 0x43, 0x48, 0xa3, 0x9b, 0x0c, 0xf6, 0x01, 0x6a, 0x9f, 0x72, 0x3e, 0x1c, 0xbf, 0x2a,
	0x2c, 0xb9, 0x30, 0xc2, 0x04, 0xd5, 0xb5, 0x41, 0x62, 0x3c, 0x30, 0xf6, 0x4c, 0xb7, 0x3c, 0xda,
	0xe7, 0xa8, 0x33, 0x0f, 0x16, 0x1c, 0x9f, 0xa0, 0x75, 0x21, 0xa9, 0x4c, 0x85, 0x02, 0xff, 0x73,
	0xb4, 0xe7, 0xdc, 0xda, 0x2f, 0xe7, 0x05, 0xc8, 0xb7, 0x0a, 0x9b, 0x57, 0xba, 0xba, 0xce, 0x6e,
	0xa1, 0xc6, 0x4c, 0x62, 0x64, 0xff, 0xa8, 0xa1, 0xe6, 0x1c, 0x12, 0xef, 0x22, 0xa4, 0x29, 0xfa,
	0xa1, 0xaf, 0x45, 0x99, 0x3a, 0xf2, 0xd2, 0xc7, 0xcf, 0x50, 0xd3, 0x63, 0xf1, 0x20, 0x0c, 0xd2,
	0x04, 0xfc, 0x3e, 0x95, 0xa4, 0xa6, 0x94, 0x58, 0x4e, 0x31, 0x00, 0xa7, 0x1c, 0x80, 0x73, 0x5e,
	0x0e, 0xc0, 0x6d, 0x5c, 0x17, 0x9c, 0x4a, 0xfc, 0x1e, 0xb5, 0x35, 0x9b, 0xdf, 0xe7, 0x90, 0x84,
	0xcc, 0x27, 0x2b, 0x8a, 0xe2, 0x68, 0x59, 0x33, 0xce, 0x73, 0x2a, 0xc1, 0xa5, 0x71, 0x00, 0x6e,
	0xab, 0xa4, 0x7a, 0xa3, 0x98, 0xf0, 0x43, 0xd4, 0xf4, 0x52, 0x21, 0x59, 0x04, 0x49, 0x3f, 0xa6,
	0x11, 0x90, 0x55, 0xa5, 0xbf, 0x51, 0x06, 0x5f, 0xd3, 0x08, 0xac, 0x2b, 0x64, 0x56, 0x0c, 0xf8,
	0x31, 0x5a, 0x13, 0x92, 0x26, 0x52, 0x77, 0xf4, 0x2e, 0x1f, 0x05, 0x10, 0x1f, 0xa2, 0x15, 0x88,
	0xfd, 0x25, 0x7c, 0xe7, 0x30, 0xfb, 0x93, 0x81, 0x36, 0x5d, 0x18, 0xa5, 0x20, 0xe4, 0xcc, 0xd8,
	0x77, 0x11, 0x1a, 0x84, 0x89, 0x90, 0x85, 0x48, 0xdd, 0x64, 0x15, 0xc9, 0x15, 0xe2, 0xff, 0x90,
	0x39, 0xa4, 0x65, 0xb6, 0xa6, 0xb2, 0x1b, 0x79, 0x40, 0x25, 0xb7, 0xd0, 0x1a, 0x44, 0x34, 0x1c,
	0xaa, 0xb6, 0x99, 0x6e, 0x71, 0xc8, 0x19, 0x03, 0x9f, 0x27, 0x7d, 0x1a, 0x24, 0x50, 0xd8, 0xde,
	0x70, 0xcd, 0x3c, 0x72, 0x9a, 0x07, 0x6c, 0x89, 0xf0, 0xa2, 0x0a, 0xc1, 0xf1, 0xce, 0xc2, 0xf6,
	0x9d, 0xd5, 0x88, 0x51, 0x6d, 0xe0, 0xcc, 0xb6, 0xd5, 0xfe, 0x6e, 0xdb, 0x8e, 0x7e, 0xad, 0xa2,
	0xba, 0xfe, 0x1e, 0xfe, 0x6a, 0xa0, 0xc6, 0xec, 0x42, 0xe3, 0xfd, 0x3b, 0xe8, 0x16, 0xae, 0x89,
	0x75, 0xb0, 0x34, 0x56, 0x70, 0x1b, 0x26, 0x19, 0x31, 0x2b, 0x67, 0x93, 0x8c, 0xd4, 0xf1, 0x1a,
	0xcd, 0x41, 0xd3, 0x8c, 0x74, 0x50, 0x4b, 0x8c, 0x85, 0x84, 0xe8, 0x58, 0xa7, 0xa7, 0x19, 0xd9,
	0xc6, 0x5b, 0xf3, 0xb1, 0x63, 0x85, 0x9e, 0x7c, 0xfb, 0xf9, 0xb9, 0xd6, 0xb1, 0x5b, 0xd5, 0x53,
	0x75, 0x1d, 0x5d, 0x39, 0x36, 0xf6, 0xf1, 0x17, 0x03, 0x99, 0x95, 0x6d, 0xfc, 0x68, 0xb9, 0xe6,
	0x8c, 0xac, 0xa5, 0xbb, 0x68, 0x7f, 0x98, 0x28, 0xbd, 0xa5, 0xa8, 0xa2, 0xa7, 0x93, 0x8c, 0xac,
	0xe3, 0xd5, 0x04, 0xa8, 0x3f, 0xcd, 0x48, 0x1b, 0x35, 0xb5, 0xf2, 0x22, 0x39, 0xcd, 0xc8, 0x16,
	0xc6, 0x0b, 0x66, 0x02, 0x90, 0x4a, 0xf4, 0x26, 0x6e, 0x57, 0x56, 0x0a, 0x38, 0xfe, 0x6e, 0xa0,
	0xd6, 0xfc, 0x5a, 0xe0, 0xc3, 0x3b, 0xe4, 0xdd, 0xd8, 0x63, 0xab, 0x7b, 0x0f, 0xb4, 0xe0, 0xf6,
	0xd5, 0xe2, 0x64, 0x4c, 0x5c, 0x4f, 0x0a, 0xd8, 0x2d, 0xb3, 0x21, 0x78, 0x7b, 0xc1, 0x8e, 0xc6,
	0x2b, 0x4b, 0xd8, 0xee, 0x54, 0x96, 0x66, 0xe3, 0xf9, 0x7c, 0xce, 0xce, 0xde, 0x9d, 0x04, 0xa1,
	0xbc, 0x4c, 0x2f, 0x1c, 0x8f, 0x45, 0xbd, 0x5c, 0x67, 0xf5, 0x6c, 0xf7, 0x96, 0xfc, 0x77, 0xba,
	0x58, 0x57, 0xd7, 0xfa, 0xc9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x06, 0x32, 0x74, 0xcf,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LicenseClient is the client API for License service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LicenseClient interface {
	ApplyLicense(ctx context.Context, in *ApplyLicenseReq, opts ...grpc.CallOption) (*ApplyLicenseResp, error)
	GetStatus(ctx context.Context, in *GetStatusReq, opts ...grpc.CallOption) (*GetStatusResp, error)
	RequestLicense(ctx context.Context, in *RequestLicenseReq, opts ...grpc.CallOption) (*RequestLicenseResp, error)
}

type licenseClient struct {
	cc *grpc.ClientConn
}

func NewLicenseClient(cc *grpc.ClientConn) LicenseClient {
	return &licenseClient{cc}
}

func (c *licenseClient) ApplyLicense(ctx context.Context, in *ApplyLicenseReq, opts ...grpc.CallOption) (*ApplyLicenseResp, error) {
	out := new(ApplyLicenseResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.license.License/ApplyLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) GetStatus(ctx context.Context, in *GetStatusReq, opts ...grpc.CallOption) (*GetStatusResp, error) {
	out := new(GetStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.license.License/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) RequestLicense(ctx context.Context, in *RequestLicenseReq, opts ...grpc.CallOption) (*RequestLicenseResp, error) {
	out := new(RequestLicenseResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.license.License/RequestLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicenseServer is the server API for License service.
type LicenseServer interface {
	ApplyLicense(context.Context, *ApplyLicenseReq) (*ApplyLicenseResp, error)
	GetStatus(context.Context, *GetStatusReq) (*GetStatusResp, error)
	RequestLicense(context.Context, *RequestLicenseReq) (*RequestLicenseResp, error)
}

// UnimplementedLicenseServer can be embedded to have forward compatible implementations.
type UnimplementedLicenseServer struct {
}

func (*UnimplementedLicenseServer) ApplyLicense(ctx context.Context, req *ApplyLicenseReq) (*ApplyLicenseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyLicense not implemented")
}
func (*UnimplementedLicenseServer) GetStatus(ctx context.Context, req *GetStatusReq) (*GetStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedLicenseServer) RequestLicense(ctx context.Context, req *RequestLicenseReq) (*RequestLicenseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLicense not implemented")
}

func RegisterLicenseServer(s *grpc.Server, srv LicenseServer) {
	s.RegisterService(&_License_serviceDesc, srv)
}

func _License_ApplyLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).ApplyLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.license.License/ApplyLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).ApplyLicense(ctx, req.(*ApplyLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.license.License/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).GetStatus(ctx, req.(*GetStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_RequestLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).RequestLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.license.License/RequestLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).RequestLicense(ctx, req.(*RequestLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _License_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.license.License",
	HandlerType: (*LicenseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyLicense",
			Handler:    _License_ApplyLicense_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _License_GetStatus_Handler,
		},
		{
			MethodName: "RequestLicense",
			Handler:    _License_RequestLicense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/license/license.proto",
}
