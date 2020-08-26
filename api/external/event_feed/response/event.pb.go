// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/event_feed/response/event.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetEventFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of events.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// Total count of events.
	TotalEvents int64 `protobuf:"varint,2,opt,name=total_events,json=totalEvents,proto3" json:"total_events,omitempty"`
}

func (x *GetEventFeedResponse) Reset() {
	*x = GetEventFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventFeedResponse) ProtoMessage() {}

func (x *GetEventFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventFeedResponse.ProtoReflect.Descriptor instead.
func (*GetEventFeedResponse) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_event_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventFeedResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GetEventFeedResponse) GetTotalEvents() int64 {
	if x != nil {
		return x.TotalEvents
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of event (cookbook, role, etc).
	EventType string `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	// Type of event task (create, update, delete).
	Task string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	// Event start time.
	StartTime  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EntityName string               `protobuf:"bytes,4,opt,name=entity_name,json=entityName,proto3" json:"entity_name,omitempty"`
	// Event record requestor type.
	RequestorType string `protobuf:"bytes,5,opt,name=requestor_type,json=requestorType,proto3" json:"requestor_type,omitempty"`
	// Event record requestor name.
	RequestorName string `protobuf:"bytes,6,opt,name=requestor_name,json=requestorName,proto3" json:"requestor_name,omitempty"`
	// Hostname from which the record was gathered.
	ServiceHostname string `protobuf:"bytes,7,opt,name=service_hostname,json=serviceHostname,proto3" json:"service_hostname,omitempty"`
	// Used for grouping events together.
	StartId string `protobuf:"bytes,8,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
	// Used for grouping events together.
	EventCount int32 `protobuf:"varint,9,opt,name=event_count,json=eventCount,proto3" json:"event_count,omitempty"`
	// Used for grouping events together.
	ParentName string `protobuf:"bytes,16,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	// Used for grouping events together.
	ParentType string `protobuf:"bytes,17,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty"`
	// Used for grouping events together; equal to start_time if not grouped
	EndTime *timestamp.Timestamp `protobuf:"bytes,18,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Used for grouping events together; equal to start_id if not grouped
	EndId string `protobuf:"bytes,19,opt,name=end_id,json=endId,proto3" json:"end_id,omitempty"`
	// Event's Chef Organization
	ChefOrganization string `protobuf:"bytes,20,opt,name=chef_organization,json=chefOrganization,proto3" json:"chef_organization,omitempty"`
	// Event's Chef Infra Server
	ChefInfraServer string `protobuf:"bytes,21,opt,name=chef_infra_server,json=chefInfraServer,proto3" json:"chef_infra_server,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_event_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Event) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Event) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Event) GetEntityName() string {
	if x != nil {
		return x.EntityName
	}
	return ""
}

func (x *Event) GetRequestorType() string {
	if x != nil {
		return x.RequestorType
	}
	return ""
}

func (x *Event) GetRequestorName() string {
	if x != nil {
		return x.RequestorName
	}
	return ""
}

func (x *Event) GetServiceHostname() string {
	if x != nil {
		return x.ServiceHostname
	}
	return ""
}

func (x *Event) GetStartId() string {
	if x != nil {
		return x.StartId
	}
	return ""
}

func (x *Event) GetEventCount() int32 {
	if x != nil {
		return x.EventCount
	}
	return 0
}

func (x *Event) GetParentName() string {
	if x != nil {
		return x.ParentName
	}
	return ""
}

func (x *Event) GetParentType() string {
	if x != nil {
		return x.ParentType
	}
	return ""
}

func (x *Event) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Event) GetEndId() string {
	if x != nil {
		return x.EndId
	}
	return ""
}

func (x *Event) GetChefOrganization() string {
	if x != nil {
		return x.ChefOrganization
	}
	return ""
}

func (x *Event) GetChefInfraServer() string {
	if x != nil {
		return x.ChefInfraServer
	}
	return ""
}

type GetEventTypeCountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total count of events.
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// Total count of events per type.
	Counts []*EventCount `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts,omitempty"`
}

func (x *GetEventTypeCountsResponse) Reset() {
	*x = GetEventTypeCountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventTypeCountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventTypeCountsResponse) ProtoMessage() {}

func (x *GetEventTypeCountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventTypeCountsResponse.ProtoReflect.Descriptor instead.
func (*GetEventTypeCountsResponse) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_event_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventTypeCountsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetEventTypeCountsResponse) GetCounts() []*EventCount {
	if x != nil {
		return x.Counts
	}
	return nil
}

type GetEventTaskCountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total count of events.
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// Total count of events per type.
	Counts []*EventCount `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts,omitempty"`
}

func (x *GetEventTaskCountsResponse) Reset() {
	*x = GetEventTaskCountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventTaskCountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventTaskCountsResponse) ProtoMessage() {}

func (x *GetEventTaskCountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventTaskCountsResponse.ProtoReflect.Descriptor instead.
func (*GetEventTaskCountsResponse) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventTaskCountsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetEventTaskCountsResponse) GetCounts() []*EventCount {
	if x != nil {
		return x.Counts
	}
	return nil
}

type EventCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Count of events.
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *EventCount) Reset() {
	*x = EventCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCount) ProtoMessage() {}

func (x *EventCount) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCount.ProtoReflect.Descriptor instead.
func (*EventCount) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_event_proto_rawDescGZIP(), []int{4}
}

func (x *EventCount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EventCount) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type EventExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Exported reports in JSON or CSV.
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EventExportResponse) Reset() {
	*x = EventExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventExportResponse) ProtoMessage() {}

func (x *EventExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventExportResponse.ProtoReflect.Descriptor instead.
func (*EventExportResponse) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_event_proto_rawDescGZIP(), []int{5}
}

func (x *EventExportResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_external_event_feed_response_event_proto protoreflect.FileDescriptor

var file_external_event_feed_response_event_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xb4, 0x04, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x68,
	0x65, 0x66, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x66, 0x49,
	0x6e, 0x66, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x49,
	0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x7d, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x49, 0x0a,
	0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x2f, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_event_feed_response_event_proto_rawDescOnce sync.Once
	file_external_event_feed_response_event_proto_rawDescData = file_external_event_feed_response_event_proto_rawDesc
)

func file_external_event_feed_response_event_proto_rawDescGZIP() []byte {
	file_external_event_feed_response_event_proto_rawDescOnce.Do(func() {
		file_external_event_feed_response_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_event_feed_response_event_proto_rawDescData)
	})
	return file_external_event_feed_response_event_proto_rawDescData
}

var file_external_event_feed_response_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_event_feed_response_event_proto_goTypes = []interface{}{
	(*GetEventFeedResponse)(nil),       // 0: chef.automate.api.event_feed.response.GetEventFeedResponse
	(*Event)(nil),                      // 1: chef.automate.api.event_feed.response.Event
	(*GetEventTypeCountsResponse)(nil), // 2: chef.automate.api.event_feed.response.GetEventTypeCountsResponse
	(*GetEventTaskCountsResponse)(nil), // 3: chef.automate.api.event_feed.response.GetEventTaskCountsResponse
	(*EventCount)(nil),                 // 4: chef.automate.api.event_feed.response.EventCount
	(*EventExportResponse)(nil),        // 5: chef.automate.api.event_feed.response.EventExportResponse
	(*timestamp.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_external_event_feed_response_event_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.event_feed.response.GetEventFeedResponse.events:type_name -> chef.automate.api.event_feed.response.Event
	6, // 1: chef.automate.api.event_feed.response.Event.start_time:type_name -> google.protobuf.Timestamp
	6, // 2: chef.automate.api.event_feed.response.Event.end_time:type_name -> google.protobuf.Timestamp
	4, // 3: chef.automate.api.event_feed.response.GetEventTypeCountsResponse.counts:type_name -> chef.automate.api.event_feed.response.EventCount
	4, // 4: chef.automate.api.event_feed.response.GetEventTaskCountsResponse.counts:type_name -> chef.automate.api.event_feed.response.EventCount
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_external_event_feed_response_event_proto_init() }
func file_external_event_feed_response_event_proto_init() {
	if File_external_event_feed_response_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_event_feed_response_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventFeedResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_event_feed_response_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_event_feed_response_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventTypeCountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_event_feed_response_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventTaskCountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_event_feed_response_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_event_feed_response_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventExportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_event_feed_response_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_event_feed_response_event_proto_goTypes,
		DependencyIndexes: file_external_event_feed_response_event_proto_depIdxs,
		MessageInfos:      file_external_event_feed_response_event_proto_msgTypes,
	}.Build()
	File_external_event_feed_response_event_proto = out.File
	file_external_event_feed_response_event_proto_rawDesc = nil
	file_external_event_feed_response_event_proto_goTypes = nil
	file_external_event_feed_response_event_proto_depIdxs = nil
}
