// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rules.proto

package api

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Rule_Event int32

const (
	Rule_CCRFailure        Rule_Event = 0
	Rule_CCRSuccess        Rule_Event = 1
	Rule_ComplianceFailure Rule_Event = 2
	Rule_ComplianceSuccess Rule_Event = 3
)

var Rule_Event_name = map[int32]string{
	0: "CCRFailure",
	1: "CCRSuccess",
	2: "ComplianceFailure",
	3: "ComplianceSuccess",
}

var Rule_Event_value = map[string]int32{
	"CCRFailure":        0,
	"CCRSuccess":        1,
	"ComplianceFailure": 2,
	"ComplianceSuccess": 3,
}

func (x Rule_Event) String() string {
	return proto.EnumName(Rule_Event_name, int32(x))
}

func (Rule_Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{4, 0}
}

type URLValidationResponse_Code int32

const (
	URLValidationResponse_OK                        URLValidationResponse_Code = 0
	URLValidationResponse_ERROR                     URLValidationResponse_Code = 1
	URLValidationResponse_INVALID_URL               URLValidationResponse_Code = 4
	URLValidationResponse_NOTIFICATIONS_UNAVAIALBLE URLValidationResponse_Code = 98
	URLValidationResponse_INTERNAL_ERROR            URLValidationResponse_Code = 99
)

var URLValidationResponse_Code_name = map[int32]string{
	0:  "OK",
	1:  "ERROR",
	4:  "INVALID_URL",
	98: "NOTIFICATIONS_UNAVAIALBLE",
	99: "INTERNAL_ERROR",
}

var URLValidationResponse_Code_value = map[string]int32{
	"OK":                        0,
	"ERROR":                     1,
	"INVALID_URL":               4,
	"NOTIFICATIONS_UNAVAIALBLE": 98,
	"INTERNAL_ERROR":            99,
}

func (x URLValidationResponse_Code) String() string {
	return proto.EnumName(URLValidationResponse_Code_name, int32(x))
}

func (URLValidationResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{8, 0}
}

type RuleUpdateResponse_Code int32

const (
	RuleUpdateResponse_OK               RuleUpdateResponse_Code = 0
	RuleUpdateResponse_DUPLICATE_NAME   RuleUpdateResponse_Code = 1
	RuleUpdateResponse_NOT_FOUND        RuleUpdateResponse_Code = 2
	RuleUpdateResponse_VALIDATION_ERROR RuleUpdateResponse_Code = 4
	RuleUpdateResponse_INTERNAL_ERROR   RuleUpdateResponse_Code = 99
)

var RuleUpdateResponse_Code_name = map[int32]string{
	0:  "OK",
	1:  "DUPLICATE_NAME",
	2:  "NOT_FOUND",
	4:  "VALIDATION_ERROR",
	99: "INTERNAL_ERROR",
}

var RuleUpdateResponse_Code_value = map[string]int32{
	"OK":               0,
	"DUPLICATE_NAME":   1,
	"NOT_FOUND":        2,
	"VALIDATION_ERROR": 4,
	"INTERNAL_ERROR":   99,
}

func (x RuleUpdateResponse_Code) String() string {
	return proto.EnumName(RuleUpdateResponse_Code_name, int32(x))
}

func (RuleUpdateResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{10, 0}
}

type RuleDeleteResponse_Code int32

const (
	RuleDeleteResponse_DELETED        RuleDeleteResponse_Code = 0
	RuleDeleteResponse_NOT_FOUND      RuleDeleteResponse_Code = 2
	RuleDeleteResponse_INTERNAL_ERROR RuleDeleteResponse_Code = 99
)

var RuleDeleteResponse_Code_name = map[int32]string{
	0:  "DELETED",
	2:  "NOT_FOUND",
	99: "INTERNAL_ERROR",
}

var RuleDeleteResponse_Code_value = map[string]int32{
	"DELETED":        0,
	"NOT_FOUND":      2,
	"INTERNAL_ERROR": 99,
}

func (x RuleDeleteResponse_Code) String() string {
	return proto.EnumName(RuleDeleteResponse_Code_name, int32(x))
}

func (RuleDeleteResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{11, 0}
}

type RuleAddResponse_Code int32

const (
	RuleAddResponse_ADDED                 RuleAddResponse_Code = 0
	RuleAddResponse_DUPLICATE_NAME        RuleAddResponse_Code = 1
	RuleAddResponse_NOT_FOUND             RuleAddResponse_Code = 2
	RuleAddResponse_INVALID_ACTION_CONFIG RuleAddResponse_Code = 3
	RuleAddResponse_VALIDATION_ERROR      RuleAddResponse_Code = 4
	RuleAddResponse_INTERNAL_ERROR        RuleAddResponse_Code = 99
)

var RuleAddResponse_Code_name = map[int32]string{
	0:  "ADDED",
	1:  "DUPLICATE_NAME",
	2:  "NOT_FOUND",
	3:  "INVALID_ACTION_CONFIG",
	4:  "VALIDATION_ERROR",
	99: "INTERNAL_ERROR",
}

var RuleAddResponse_Code_value = map[string]int32{
	"ADDED":                 0,
	"DUPLICATE_NAME":        1,
	"NOT_FOUND":             2,
	"INVALID_ACTION_CONFIG": 3,
	"VALIDATION_ERROR":      4,
	"INTERNAL_ERROR":        99,
}

func (x RuleAddResponse_Code) String() string {
	return proto.EnumName(RuleAddResponse_Code_name, int32(x))
}

func (RuleAddResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{12, 0}
}

type RuleGetResponse_Code int32

const (
	RuleGetResponse_OK             RuleGetResponse_Code = 0
	RuleGetResponse_NOT_FOUND      RuleGetResponse_Code = 2
	RuleGetResponse_INTERNAL_ERROR RuleGetResponse_Code = 99
)

var RuleGetResponse_Code_name = map[int32]string{
	0:  "OK",
	2:  "NOT_FOUND",
	99: "INTERNAL_ERROR",
}

var RuleGetResponse_Code_value = map[string]int32{
	"OK":             0,
	"NOT_FOUND":      2,
	"INTERNAL_ERROR": 99,
}

func (x RuleGetResponse_Code) String() string {
	return proto.EnumName(RuleGetResponse_Code_name, int32(x))
}

func (RuleGetResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{13, 0}
}

type RuleListResponse_Code int32

const (
	RuleListResponse_OK             RuleListResponse_Code = 0
	RuleListResponse_INTERNAL_ERROR RuleListResponse_Code = 99
)

var RuleListResponse_Code_name = map[int32]string{
	0:  "OK",
	99: "INTERNAL_ERROR",
}

var RuleListResponse_Code_value = map[string]int32{
	"OK":             0,
	"INTERNAL_ERROR": 99,
}

func (x RuleListResponse_Code) String() string {
	return proto.EnumName(RuleListResponse_Code_name, int32(x))
}

func (RuleListResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{14, 0}
}

////
// Rules management messages and responses
////
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type SlackAlert struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlackAlert) Reset()         { *m = SlackAlert{} }
func (m *SlackAlert) String() string { return proto.CompactTextString(m) }
func (*SlackAlert) ProtoMessage()    {}
func (*SlackAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{1}
}

func (m *SlackAlert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlackAlert.Unmarshal(m, b)
}
func (m *SlackAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlackAlert.Marshal(b, m, deterministic)
}
func (m *SlackAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlackAlert.Merge(m, src)
}
func (m *SlackAlert) XXX_Size() int {
	return xxx_messageInfo_SlackAlert.Size(m)
}
func (m *SlackAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_SlackAlert.DiscardUnknown(m)
}

var xxx_messageInfo_SlackAlert proto.InternalMessageInfo

func (m *SlackAlert) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type WebhookAlert struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebhookAlert) Reset()         { *m = WebhookAlert{} }
func (m *WebhookAlert) String() string { return proto.CompactTextString(m) }
func (*WebhookAlert) ProtoMessage()    {}
func (*WebhookAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{2}
}

func (m *WebhookAlert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookAlert.Unmarshal(m, b)
}
func (m *WebhookAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookAlert.Marshal(b, m, deterministic)
}
func (m *WebhookAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookAlert.Merge(m, src)
}
func (m *WebhookAlert) XXX_Size() int {
	return xxx_messageInfo_WebhookAlert.Size(m)
}
func (m *WebhookAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookAlert.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookAlert proto.InternalMessageInfo

func (m *WebhookAlert) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ServiceNowAlert struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	SecretId             string   `protobuf:"bytes,2,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	CriticalControlsOnly bool     `protobuf:"varint,3,opt,name=critical_controls_only,json=criticalControlsOnly,proto3" json:"critical_controls_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceNowAlert) Reset()         { *m = ServiceNowAlert{} }
func (m *ServiceNowAlert) String() string { return proto.CompactTextString(m) }
func (*ServiceNowAlert) ProtoMessage()    {}
func (*ServiceNowAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{3}
}

func (m *ServiceNowAlert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceNowAlert.Unmarshal(m, b)
}
func (m *ServiceNowAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceNowAlert.Marshal(b, m, deterministic)
}
func (m *ServiceNowAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNowAlert.Merge(m, src)
}
func (m *ServiceNowAlert) XXX_Size() int {
	return xxx_messageInfo_ServiceNowAlert.Size(m)
}
func (m *ServiceNowAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNowAlert.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNowAlert proto.InternalMessageInfo

func (m *ServiceNowAlert) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ServiceNowAlert) GetSecretId() string {
	if m != nil {
		return m.SecretId
	}
	return ""
}

func (m *ServiceNowAlert) GetCriticalControlsOnly() bool {
	if m != nil {
		return m.CriticalControlsOnly
	}
	return false
}

type Rule struct {
	Id    string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Event Rule_Event `protobuf:"varint,3,opt,name=event,proto3,enum=notifications.Rule_Event" json:"event,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*Rule_SlackAlert
	//	*Rule_WebhookAlert
	//	*Rule_ServiceNowAlert
	Action               isRule_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{4}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetEvent() Rule_Event {
	if m != nil {
		return m.Event
	}
	return Rule_CCRFailure
}

type isRule_Action interface {
	isRule_Action()
}

type Rule_SlackAlert struct {
	SlackAlert *SlackAlert `protobuf:"bytes,4,opt,name=SlackAlert,proto3,oneof"`
}

type Rule_WebhookAlert struct {
	WebhookAlert *WebhookAlert `protobuf:"bytes,5,opt,name=WebhookAlert,proto3,oneof"`
}

type Rule_ServiceNowAlert struct {
	ServiceNowAlert *ServiceNowAlert `protobuf:"bytes,6,opt,name=ServiceNowAlert,proto3,oneof"`
}

func (*Rule_SlackAlert) isRule_Action() {}

func (*Rule_WebhookAlert) isRule_Action() {}

func (*Rule_ServiceNowAlert) isRule_Action() {}

func (m *Rule) GetAction() isRule_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Rule) GetSlackAlert() *SlackAlert {
	if x, ok := m.GetAction().(*Rule_SlackAlert); ok {
		return x.SlackAlert
	}
	return nil
}

func (m *Rule) GetWebhookAlert() *WebhookAlert {
	if x, ok := m.GetAction().(*Rule_WebhookAlert); ok {
		return x.WebhookAlert
	}
	return nil
}

func (m *Rule) GetServiceNowAlert() *ServiceNowAlert {
	if x, ok := m.GetAction().(*Rule_ServiceNowAlert); ok {
		return x.ServiceNowAlert
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Rule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Rule_SlackAlert)(nil),
		(*Rule_WebhookAlert)(nil),
		(*Rule_ServiceNowAlert)(nil),
	}
}

type UsernamePassword struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsernamePassword) Reset()         { *m = UsernamePassword{} }
func (m *UsernamePassword) String() string { return proto.CompactTextString(m) }
func (*UsernamePassword) ProtoMessage()    {}
func (*UsernamePassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{5}
}

func (m *UsernamePassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsernamePassword.Unmarshal(m, b)
}
func (m *UsernamePassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsernamePassword.Marshal(b, m, deterministic)
}
func (m *UsernamePassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsernamePassword.Merge(m, src)
}
func (m *UsernamePassword) XXX_Size() int {
	return xxx_messageInfo_UsernamePassword.Size(m)
}
func (m *UsernamePassword) XXX_DiscardUnknown() {
	xxx_messageInfo_UsernamePassword.DiscardUnknown(m)
}

var xxx_messageInfo_UsernamePassword proto.InternalMessageInfo

func (m *UsernamePassword) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UsernamePassword) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SecretId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretId) Reset()         { *m = SecretId{} }
func (m *SecretId) String() string { return proto.CompactTextString(m) }
func (*SecretId) ProtoMessage()    {}
func (*SecretId) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{6}
}

func (m *SecretId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretId.Unmarshal(m, b)
}
func (m *SecretId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretId.Marshal(b, m, deterministic)
}
func (m *SecretId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretId.Merge(m, src)
}
func (m *SecretId) XXX_Size() int {
	return xxx_messageInfo_SecretId.Size(m)
}
func (m *SecretId) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretId.DiscardUnknown(m)
}

var xxx_messageInfo_SecretId proto.InternalMessageInfo

func (m *SecretId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type URLValidationRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Types that are valid to be assigned to Credentials:
	//	*URLValidationRequest_UsernamePassword
	//	*URLValidationRequest_SecretId
	//	*URLValidationRequest_None
	Credentials          isURLValidationRequest_Credentials `protobuf_oneof:"credentials"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *URLValidationRequest) Reset()         { *m = URLValidationRequest{} }
func (m *URLValidationRequest) String() string { return proto.CompactTextString(m) }
func (*URLValidationRequest) ProtoMessage()    {}
func (*URLValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{7}
}

func (m *URLValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URLValidationRequest.Unmarshal(m, b)
}
func (m *URLValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URLValidationRequest.Marshal(b, m, deterministic)
}
func (m *URLValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URLValidationRequest.Merge(m, src)
}
func (m *URLValidationRequest) XXX_Size() int {
	return xxx_messageInfo_URLValidationRequest.Size(m)
}
func (m *URLValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_URLValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_URLValidationRequest proto.InternalMessageInfo

func (m *URLValidationRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type isURLValidationRequest_Credentials interface {
	isURLValidationRequest_Credentials()
}

type URLValidationRequest_UsernamePassword struct {
	UsernamePassword *UsernamePassword `protobuf:"bytes,2,opt,name=username_password,json=usernamePassword,proto3,oneof"`
}

type URLValidationRequest_SecretId struct {
	SecretId *SecretId `protobuf:"bytes,3,opt,name=secret_id,json=secretId,proto3,oneof"`
}

type URLValidationRequest_None struct {
	None *Empty `protobuf:"bytes,4,opt,name=none,proto3,oneof"`
}

func (*URLValidationRequest_UsernamePassword) isURLValidationRequest_Credentials() {}

func (*URLValidationRequest_SecretId) isURLValidationRequest_Credentials() {}

func (*URLValidationRequest_None) isURLValidationRequest_Credentials() {}

func (m *URLValidationRequest) GetCredentials() isURLValidationRequest_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *URLValidationRequest) GetUsernamePassword() *UsernamePassword {
	if x, ok := m.GetCredentials().(*URLValidationRequest_UsernamePassword); ok {
		return x.UsernamePassword
	}
	return nil
}

func (m *URLValidationRequest) GetSecretId() *SecretId {
	if x, ok := m.GetCredentials().(*URLValidationRequest_SecretId); ok {
		return x.SecretId
	}
	return nil
}

func (m *URLValidationRequest) GetNone() *Empty {
	if x, ok := m.GetCredentials().(*URLValidationRequest_None); ok {
		return x.None
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*URLValidationRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*URLValidationRequest_UsernamePassword)(nil),
		(*URLValidationRequest_SecretId)(nil),
		(*URLValidationRequest_None)(nil),
	}
}

type URLValidationResponse struct {
	Code                 URLValidationResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=notifications.URLValidationResponse_Code" json:"code,omitempty"`
	Messages             []string                   `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *URLValidationResponse) Reset()         { *m = URLValidationResponse{} }
func (m *URLValidationResponse) String() string { return proto.CompactTextString(m) }
func (*URLValidationResponse) ProtoMessage()    {}
func (*URLValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{8}
}

func (m *URLValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URLValidationResponse.Unmarshal(m, b)
}
func (m *URLValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URLValidationResponse.Marshal(b, m, deterministic)
}
func (m *URLValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URLValidationResponse.Merge(m, src)
}
func (m *URLValidationResponse) XXX_Size() int {
	return xxx_messageInfo_URLValidationResponse.Size(m)
}
func (m *URLValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_URLValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_URLValidationResponse proto.InternalMessageInfo

func (m *URLValidationResponse) GetCode() URLValidationResponse_Code {
	if m != nil {
		return m.Code
	}
	return URLValidationResponse_OK
}

func (m *URLValidationResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type RuleIdentifier struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleIdentifier) Reset()         { *m = RuleIdentifier{} }
func (m *RuleIdentifier) String() string { return proto.CompactTextString(m) }
func (*RuleIdentifier) ProtoMessage()    {}
func (*RuleIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{9}
}

func (m *RuleIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleIdentifier.Unmarshal(m, b)
}
func (m *RuleIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleIdentifier.Marshal(b, m, deterministic)
}
func (m *RuleIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleIdentifier.Merge(m, src)
}
func (m *RuleIdentifier) XXX_Size() int {
	return xxx_messageInfo_RuleIdentifier.Size(m)
}
func (m *RuleIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_RuleIdentifier proto.InternalMessageInfo

func (m *RuleIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RuleUpdateResponse struct {
	Code                 RuleUpdateResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=notifications.RuleUpdateResponse_Code" json:"code,omitempty"`
	Messages             []string                `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RuleUpdateResponse) Reset()         { *m = RuleUpdateResponse{} }
func (m *RuleUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*RuleUpdateResponse) ProtoMessage()    {}
func (*RuleUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{10}
}

func (m *RuleUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleUpdateResponse.Unmarshal(m, b)
}
func (m *RuleUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleUpdateResponse.Marshal(b, m, deterministic)
}
func (m *RuleUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleUpdateResponse.Merge(m, src)
}
func (m *RuleUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_RuleUpdateResponse.Size(m)
}
func (m *RuleUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RuleUpdateResponse proto.InternalMessageInfo

func (m *RuleUpdateResponse) GetCode() RuleUpdateResponse_Code {
	if m != nil {
		return m.Code
	}
	return RuleUpdateResponse_OK
}

func (m *RuleUpdateResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type RuleDeleteResponse struct {
	Code                 RuleDeleteResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=notifications.RuleDeleteResponse_Code" json:"code,omitempty"`
	Messages             []string                `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RuleDeleteResponse) Reset()         { *m = RuleDeleteResponse{} }
func (m *RuleDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*RuleDeleteResponse) ProtoMessage()    {}
func (*RuleDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{11}
}

func (m *RuleDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleDeleteResponse.Unmarshal(m, b)
}
func (m *RuleDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleDeleteResponse.Marshal(b, m, deterministic)
}
func (m *RuleDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleDeleteResponse.Merge(m, src)
}
func (m *RuleDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_RuleDeleteResponse.Size(m)
}
func (m *RuleDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RuleDeleteResponse proto.InternalMessageInfo

func (m *RuleDeleteResponse) GetCode() RuleDeleteResponse_Code {
	if m != nil {
		return m.Code
	}
	return RuleDeleteResponse_DELETED
}

func (m *RuleDeleteResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type RuleAddResponse struct {
	Code                 RuleAddResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=notifications.RuleAddResponse_Code" json:"code,omitempty"`
	Messages             []string             `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Id                   string               `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RuleAddResponse) Reset()         { *m = RuleAddResponse{} }
func (m *RuleAddResponse) String() string { return proto.CompactTextString(m) }
func (*RuleAddResponse) ProtoMessage()    {}
func (*RuleAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{12}
}

func (m *RuleAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleAddResponse.Unmarshal(m, b)
}
func (m *RuleAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleAddResponse.Marshal(b, m, deterministic)
}
func (m *RuleAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleAddResponse.Merge(m, src)
}
func (m *RuleAddResponse) XXX_Size() int {
	return xxx_messageInfo_RuleAddResponse.Size(m)
}
func (m *RuleAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RuleAddResponse proto.InternalMessageInfo

func (m *RuleAddResponse) GetCode() RuleAddResponse_Code {
	if m != nil {
		return m.Code
	}
	return RuleAddResponse_ADDED
}

func (m *RuleAddResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *RuleAddResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RuleGetResponse struct {
	Code                 RuleGetResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=notifications.RuleGetResponse_Code" json:"code,omitempty"`
	Messages             []string             `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Rule                 *Rule                `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RuleGetResponse) Reset()         { *m = RuleGetResponse{} }
func (m *RuleGetResponse) String() string { return proto.CompactTextString(m) }
func (*RuleGetResponse) ProtoMessage()    {}
func (*RuleGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{13}
}

func (m *RuleGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleGetResponse.Unmarshal(m, b)
}
func (m *RuleGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleGetResponse.Marshal(b, m, deterministic)
}
func (m *RuleGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleGetResponse.Merge(m, src)
}
func (m *RuleGetResponse) XXX_Size() int {
	return xxx_messageInfo_RuleGetResponse.Size(m)
}
func (m *RuleGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RuleGetResponse proto.InternalMessageInfo

func (m *RuleGetResponse) GetCode() RuleGetResponse_Code {
	if m != nil {
		return m.Code
	}
	return RuleGetResponse_OK
}

func (m *RuleGetResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *RuleGetResponse) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type RuleListResponse struct {
	Code                 RuleListResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=notifications.RuleListResponse_Code" json:"code,omitempty"`
	Messages             []string              `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Rules                []*Rule               `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RuleListResponse) Reset()         { *m = RuleListResponse{} }
func (m *RuleListResponse) String() string { return proto.CompactTextString(m) }
func (*RuleListResponse) ProtoMessage()    {}
func (*RuleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{14}
}

func (m *RuleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleListResponse.Unmarshal(m, b)
}
func (m *RuleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleListResponse.Marshal(b, m, deterministic)
}
func (m *RuleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleListResponse.Merge(m, src)
}
func (m *RuleListResponse) XXX_Size() int {
	return xxx_messageInfo_RuleListResponse.Size(m)
}
func (m *RuleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RuleListResponse proto.InternalMessageInfo

func (m *RuleListResponse) GetCode() RuleListResponse_Code {
	if m != nil {
		return m.Code
	}
	return RuleListResponse_OK
}

func (m *RuleListResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *RuleListResponse) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterEnum("notifications.Rule_Event", Rule_Event_name, Rule_Event_value)
	proto.RegisterEnum("notifications.URLValidationResponse_Code", URLValidationResponse_Code_name, URLValidationResponse_Code_value)
	proto.RegisterEnum("notifications.RuleUpdateResponse_Code", RuleUpdateResponse_Code_name, RuleUpdateResponse_Code_value)
	proto.RegisterEnum("notifications.RuleDeleteResponse_Code", RuleDeleteResponse_Code_name, RuleDeleteResponse_Code_value)
	proto.RegisterEnum("notifications.RuleAddResponse_Code", RuleAddResponse_Code_name, RuleAddResponse_Code_value)
	proto.RegisterEnum("notifications.RuleGetResponse_Code", RuleGetResponse_Code_name, RuleGetResponse_Code_value)
	proto.RegisterEnum("notifications.RuleListResponse_Code", RuleListResponse_Code_name, RuleListResponse_Code_value)
	proto.RegisterType((*Empty)(nil), "notifications.Empty")
	proto.RegisterType((*SlackAlert)(nil), "notifications.SlackAlert")
	proto.RegisterType((*WebhookAlert)(nil), "notifications.WebhookAlert")
	proto.RegisterType((*ServiceNowAlert)(nil), "notifications.ServiceNowAlert")
	proto.RegisterType((*Rule)(nil), "notifications.Rule")
	proto.RegisterType((*UsernamePassword)(nil), "notifications.UsernamePassword")
	proto.RegisterType((*SecretId)(nil), "notifications.SecretId")
	proto.RegisterType((*URLValidationRequest)(nil), "notifications.URLValidationRequest")
	proto.RegisterType((*URLValidationResponse)(nil), "notifications.URLValidationResponse")
	proto.RegisterType((*RuleIdentifier)(nil), "notifications.RuleIdentifier")
	proto.RegisterType((*RuleUpdateResponse)(nil), "notifications.RuleUpdateResponse")
	proto.RegisterType((*RuleDeleteResponse)(nil), "notifications.RuleDeleteResponse")
	proto.RegisterType((*RuleAddResponse)(nil), "notifications.RuleAddResponse")
	proto.RegisterType((*RuleGetResponse)(nil), "notifications.RuleGetResponse")
	proto.RegisterType((*RuleListResponse)(nil), "notifications.RuleListResponse")
}

func init() { proto.RegisterFile("rules.proto", fileDescriptor_8e722d3e922f0937) }

var fileDescriptor_8e722d3e922f0937 = []byte{
	// 862 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0x86, 0xed, 0xd8, 0xe9, 0x26, 0x27, 0x34, 0xf5, 0x0e, 0x2d, 0xa4, 0xbb, 0x62, 0x89, 0x0c,
	0x82, 0x2e, 0x17, 0x41, 0x14, 0xb4, 0x20, 0x10, 0x17, 0x6e, 0xe2, 0x6e, 0xbc, 0x18, 0x67, 0x35,
	0x49, 0x8a, 0x04, 0x42, 0x96, 0x6b, 0xcf, 0xc2, 0x08, 0xd7, 0x13, 0x3c, 0x76, 0x57, 0x7d, 0x03,
	0x1e, 0x84, 0x77, 0xe0, 0x9e, 0x3b, 0x2e, 0x79, 0x0e, 0xae, 0x78, 0x83, 0x95, 0x27, 0x76, 0x6a,
	0x3b, 0xae, 0x94, 0xde, 0x79, 0xc6, 0xff, 0x39, 0x73, 0xbe, 0xff, 0x9c, 0x49, 0x0c, 0xbd, 0x38,
	0x0d, 0x09, 0x1f, 0xad, 0x62, 0x96, 0x30, 0xb4, 0x1f, 0xb1, 0x84, 0xbe, 0xa2, 0xbe, 0x97, 0x50,
	0x16, 0x71, 0xfd, 0x01, 0xb4, 0xcd, 0xab, 0x55, 0x72, 0xa3, 0x3f, 0x01, 0x98, 0x87, 0x9e, 0xff,
	0x9b, 0x11, 0x92, 0x38, 0x41, 0x1a, 0x28, 0x69, 0x1c, 0x0e, 0xe4, 0xa1, 0x7c, 0xd2, 0xc5, 0xd9,
	0xa3, 0x3e, 0x84, 0xb7, 0x7e, 0x20, 0x97, 0xbf, 0x32, 0x76, 0xa7, 0xe2, 0x1a, 0x0e, 0xe6, 0x24,
	0xbe, 0xa6, 0x3e, 0x71, 0xd8, 0xeb, 0x3b, 0x44, 0xe8, 0x31, 0x74, 0x39, 0xf1, 0x63, 0x92, 0xb8,
	0x34, 0x18, 0xb4, 0xc4, 0x7e, 0x67, 0xbd, 0x61, 0x05, 0xe8, 0x0b, 0x78, 0xc7, 0x8f, 0x69, 0x42,
	0x7d, 0x2f, 0x74, 0x7d, 0x16, 0x25, 0x31, 0x0b, 0xb9, 0xcb, 0xa2, 0xf0, 0x66, 0xa0, 0x0c, 0xe5,
	0x93, 0x0e, 0x3e, 0x2c, 0xde, 0x8e, 0xf3, 0x97, 0xb3, 0x28, 0xbc, 0xd1, 0xff, 0x50, 0x40, 0xc5,
	0x69, 0x48, 0x50, 0x1f, 0x5a, 0x34, 0xc8, 0x0f, 0x6b, 0xd1, 0x00, 0x21, 0x50, 0x23, 0xef, 0x8a,
	0xe4, 0xc7, 0x88, 0x67, 0xf4, 0x29, 0xb4, 0xc9, 0x35, 0x89, 0x12, 0x91, 0xb1, 0x7f, 0x7a, 0x3c,
	0xaa, 0xd8, 0x31, 0xca, 0xf2, 0x8c, 0xcc, 0x4c, 0x80, 0xd7, 0x3a, 0xf4, 0x4d, 0xd9, 0x97, 0x81,
	0x3a, 0x94, 0x4f, 0x7a, 0x5b, 0x51, 0xb7, 0x82, 0xa9, 0x84, 0xcb, 0x36, 0x1a, 0x55, 0xd3, 0x06,
	0x6d, 0x11, 0xfe, 0xb8, 0x16, 0x5e, 0x96, 0x4c, 0x25, 0x5c, 0xf5, 0xf9, 0xc5, 0x96, 0xab, 0x83,
	0x3d, 0x91, 0xe5, 0x49, 0xbd, 0x88, 0xaa, 0x6a, 0x2a, 0xe1, 0x7a, 0xa0, 0xbe, 0x84, 0xb6, 0x60,
	0x43, 0x7d, 0x80, 0xf1, 0x18, 0x9f, 0x7b, 0x34, 0x4c, 0x63, 0xa2, 0x49, 0xf9, 0x7a, 0x9e, 0xfa,
	0x3e, 0xe1, 0x5c, 0x93, 0xd1, 0x11, 0x3c, 0x1c, 0xb3, 0xab, 0x55, 0x48, 0xbd, 0xc8, 0x27, 0x85,
	0xac, 0x55, 0xdd, 0x2e, 0xd4, 0xca, 0x59, 0x07, 0xf6, 0x3c, 0x3f, 0x2b, 0x42, 0x7f, 0x01, 0xda,
	0x92, 0x93, 0x38, 0x73, 0xfa, 0xa5, 0xc7, 0xf9, 0x6b, 0x16, 0x07, 0xe8, 0x11, 0x74, 0xd2, 0x7c,
	0x2f, 0xef, 0xcd, 0x66, 0x9d, 0xbd, 0x5b, 0xe5, 0xba, 0x62, 0x18, 0x8a, 0xb5, 0xfe, 0x08, 0x3a,
	0xf3, 0x62, 0x30, 0x6a, 0x9d, 0xd5, 0xff, 0x93, 0xe1, 0x70, 0x89, 0xed, 0x0b, 0x2f, 0xa4, 0x81,
	0xc0, 0xc7, 0xe4, 0xf7, 0x94, 0xf0, 0xa6, 0x81, 0x73, 0xe0, 0x61, 0x71, 0x9c, 0x5b, 0x39, 0xab,
	0x77, 0xfa, 0x7e, 0xcd, 0xc1, 0x7a, 0xe9, 0x53, 0x09, 0x6b, 0x69, 0x1d, 0xe7, 0x59, 0x79, 0x80,
	0x15, 0x91, 0xe7, 0xdd, 0xad, 0x4e, 0xac, 0xcb, 0x9e, 0x4a, 0xa5, 0xd9, 0xfe, 0x04, 0xd4, 0x88,
	0x45, 0x24, 0x9f, 0xa0, 0xc3, 0x5a, 0x88, 0xb8, 0x83, 0x53, 0x09, 0x0b, 0xcd, 0xd9, 0x3e, 0xf4,
	0xfc, 0x98, 0x04, 0x24, 0x4a, 0xa8, 0x17, 0x72, 0xfd, 0x5f, 0x19, 0x8e, 0x6a, 0xb4, 0x7c, 0xc5,
	0x22, 0x4e, 0xd0, 0xb7, 0xa0, 0xfa, 0x2c, 0x58, 0xfb, 0xda, 0x3f, 0x7d, 0x5a, 0xe7, 0x69, 0x8a,
	0x19, 0x8d, 0x59, 0x40, 0xb0, 0x08, 0xcb, 0xec, 0xbf, 0x22, 0x9c, 0x7b, 0xbf, 0x10, 0x3e, 0x68,
	0x0d, 0x95, 0xcc, 0xfe, 0x62, 0xad, 0xff, 0x0c, 0x6a, 0xa6, 0x44, 0x7b, 0xd0, 0x9a, 0x7d, 0xa7,
	0x49, 0xa8, 0x0b, 0x6d, 0x13, 0xe3, 0x19, 0xd6, 0x64, 0x74, 0x00, 0x3d, 0xcb, 0xb9, 0x30, 0x6c,
	0x6b, 0xe2, 0x2e, 0xb1, 0xad, 0xa9, 0xe8, 0x3d, 0x38, 0x76, 0x66, 0x0b, 0xeb, 0xdc, 0x1a, 0x1b,
	0x0b, 0x6b, 0xe6, 0xcc, 0xdd, 0xa5, 0x63, 0x5c, 0x18, 0x96, 0x61, 0x9f, 0xd9, 0xa6, 0x76, 0x89,
	0x10, 0xf4, 0x2d, 0x67, 0x61, 0x62, 0xc7, 0xb0, 0xdd, 0x75, 0x0e, 0x5f, 0x1f, 0x42, 0x3f, 0xbb,
	0x6b, 0x96, 0x80, 0x7c, 0x45, 0x49, 0xbc, 0xd5, 0xe3, 0x7f, 0x64, 0x40, 0x99, 0x64, 0xb9, 0x0a,
	0xbc, 0x84, 0x6c, 0x90, 0xbf, 0xae, 0x20, 0x7f, 0xd4, 0x70, 0x7f, 0xab, 0x01, 0xbb, 0xf2, 0xfe,
	0x54, 0xe3, 0x45, 0xd0, 0x9f, 0x2c, 0x5f, 0xda, 0x19, 0x92, 0xe9, 0x3a, 0xc6, 0xf7, 0xa6, 0x26,
	0xa3, 0x7d, 0xe8, 0x3a, 0xb3, 0x85, 0x7b, 0x3e, 0x5b, 0x3a, 0x13, 0xad, 0x85, 0x0e, 0x41, 0x13,
	0x2e, 0x08, 0xe8, 0x9c, 0x4c, 0x6d, 0xa4, 0xfd, 0x33, 0x67, 0x99, 0x90, 0x90, 0xdc, 0x8b, 0xa5,
	0x1a, 0xb0, 0x2b, 0xcb, 0xb3, 0x9c, 0xa5, 0x07, 0x0f, 0x26, 0xa6, 0x6d, 0x2e, 0xcc, 0x89, 0x26,
	0xd5, 0x8b, 0x6f, 0x2a, 0xf3, 0x7f, 0x19, 0x0e, 0xb2, 0x53, 0x8d, 0x20, 0xd8, 0xd4, 0xf8, 0x65,
	0xa5, 0xc6, 0x0f, 0x1a, 0x6a, 0x2c, 0xa9, 0x77, 0x2c, 0x30, 0xef, 0xb5, 0xb2, 0xe9, 0xf5, 0x4d,
	0x5e, 0x70, 0x17, 0xda, 0xc6, 0x64, 0x22, 0xca, 0xdd, 0xc1, 0xff, 0x63, 0x38, 0x2a, 0xe6, 0xd0,
	0x18, 0x8b, 0x1e, 0x8c, 0x67, 0xce, 0xb9, 0xf5, 0x5c, 0x53, 0xee, 0xd1, 0x9a, 0xbf, 0x73, 0xe6,
	0xe7, 0x24, 0xb9, 0x07, 0x73, 0x49, 0xbd, 0x2b, 0xf3, 0xc7, 0xa0, 0x66, 0xff, 0xc3, 0xf9, 0x6f,
	0xc6, 0xdb, 0x0d, 0x49, 0xb1, 0x10, 0xe8, 0x9f, 0xd5, 0x26, 0x71, 0x87, 0xc6, 0xfd, 0x25, 0x83,
	0x96, 0x65, 0xb0, 0x29, 0xbf, 0xa5, 0xf8, 0xaa, 0x42, 0xf1, 0x61, 0xc3, 0x81, 0x65, 0xf9, 0xae,
	0x18, 0x4f, 0xa1, 0x2d, 0x3e, 0x27, 0x06, 0xca, 0x50, 0xb9, 0x8b, 0x63, 0xad, 0xd0, 0xf5, 0xed,
	0x2b, 0x55, 0xaf, 0xfc, 0xac, 0xfd, 0xa3, 0xe2, 0xad, 0xe8, 0xe5, 0x9e, 0xf8, 0x38, 0xf9, 0xfc,
	0x4d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x6f, 0x42, 0xe3, 0xab, 0x08, 0x00, 0x00,
}
