// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.0
// source: pkg/monitor/protobuf/monitor.proto

package protobuf

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_monitor_protobuf_monitor_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	Content       string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_pkg_monitor_protobuf_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Telemetry struct {
	Projects                *Entity `protobuf:"bytes,7,opt,name=projects,proto3" json:"projects,omitempty"`
	Metrics                 *Entity `protobuf:"bytes,10,opt,name=metrics,proto3" json:"metrics,omitempty"`
	state                   protoimpl.MessageState
	Envs                    *Entity `protobuf:"bytes,9,opt,name=envs,proto3" json:"envs,omitempty"`
	Refs                    *Entity `protobuf:"bytes,8,opt,name=refs,proto3" json:"refs,omitempty"`
	unknownFields           protoimpl.UnknownFields
	GitlabApiUsage          float64 `protobuf:"fixed64,1,opt,name=gitlab_api_usage,json=gitlabApiUsage,proto3" json:"gitlab_api_usage,omitempty"`
	TasksBufferUsage        float64 `protobuf:"fixed64,5,opt,name=tasks_buffer_usage,json=tasksBufferUsage,proto3" json:"tasks_buffer_usage,omitempty"`
	TasksExecutedCount      uint64  `protobuf:"varint,6,opt,name=tasks_executed_count,json=tasksExecutedCount,proto3" json:"tasks_executed_count,omitempty"`
	GitlabApiLimitRemaining uint64  `protobuf:"varint,4,opt,name=gitlab_api_limit_remaining,json=gitlabApiLimitRemaining,proto3" json:"gitlab_api_limit_remaining,omitempty"`
	GitlabApiRateLimit      float64 `protobuf:"fixed64,3,opt,name=gitlab_api_rate_limit,json=gitlabApiRateLimit,proto3" json:"gitlab_api_rate_limit,omitempty"`
	GitlabApiRequestsCount  uint64  `protobuf:"varint,2,opt,name=gitlab_api_requests_count,json=gitlabApiRequestsCount,proto3" json:"gitlab_api_requests_count,omitempty"`
	sizeCache               protoimpl.SizeCache
}

func (x *Telemetry) Reset() {
	*x = Telemetry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Telemetry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Telemetry) ProtoMessage() {}

func (x *Telemetry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Telemetry.ProtoReflect.Descriptor instead.
func (*Telemetry) Descriptor() ([]byte, []int) {
	return file_pkg_monitor_protobuf_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *Telemetry) GetGitlabApiUsage() float64 {
	if x != nil {
		return x.GitlabApiUsage
	}
	return 0
}

func (x *Telemetry) GetGitlabApiRequestsCount() uint64 {
	if x != nil {
		return x.GitlabApiRequestsCount
	}
	return 0
}

func (x *Telemetry) GetGitlabApiRateLimit() float64 {
	if x != nil {
		return x.GitlabApiRateLimit
	}
	return 0
}

func (x *Telemetry) GetGitlabApiLimitRemaining() uint64 {
	if x != nil {
		return x.GitlabApiLimitRemaining
	}
	return 0
}

func (x *Telemetry) GetTasksBufferUsage() float64 {
	if x != nil {
		return x.TasksBufferUsage
	}
	return 0
}

func (x *Telemetry) GetTasksExecutedCount() uint64 {
	if x != nil {
		return x.TasksExecutedCount
	}
	return 0
}

func (x *Telemetry) GetProjects() *Entity {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *Telemetry) GetRefs() *Entity {
	if x != nil {
		return x.Refs
	}
	return nil
}

func (x *Telemetry) GetEnvs() *Entity {
	if x != nil {
		return x.Envs
	}
	return nil
}

func (x *Telemetry) GetMetrics() *Entity {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState
	LastGc        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_gc,json=lastGc,proto3" json:"last_gc,omitempty"`
	LastPull      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_pull,json=lastPull,proto3" json:"last_pull,omitempty"`
	NextGc        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=next_gc,json=nextGc,proto3" json:"next_gc,omitempty"`
	NextPull      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=next_pull,json=nextPull,proto3" json:"next_pull,omitempty"`
	unknownFields protoimpl.UnknownFields
	Count         int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_monitor_protobuf_monitor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_pkg_monitor_protobuf_monitor_proto_rawDescGZIP(), []int{3}
}

func (x *Entity) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Entity) GetLastGc() *timestamppb.Timestamp {
	if x != nil {
		return x.LastGc
	}
	return nil
}

func (x *Entity) GetLastPull() *timestamppb.Timestamp {
	if x != nil {
		return x.LastPull
	}
	return nil
}

func (x *Entity) GetNextGc() *timestamppb.Timestamp {
	if x != nil {
		return x.NextGc
	}
	return nil
}

func (x *Entity) GetNextPull() *timestamppb.Timestamp {
	if x != nil {
		return x.NextPull
	}
	return nil
}

var File_pkg_monitor_protobuf_monitor_proto protoreflect.FileDescriptor

var file_pkg_monitor_protobuf_monitor_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xe2, 0x03, 0x0a, 0x09,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x41, 0x70, 0x69, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31,
	0x0a, 0x15, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x41, 0x70, 0x69, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x41, 0x70, 0x69,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2c,
	0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x72,
	0x65, 0x66, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73,
	0x12, 0x23, 0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x04, 0x65, 0x6e, 0x76, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x22, 0xfa, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x33, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x6c, 0x61, 0x73, 0x74, 0x47, 0x63, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70,
	0x75, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x12,
	0x33, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x6e, 0x65,
	0x78, 0x74, 0x47, 0x63, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x75, 0x6c,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x32, 0x71, 0x0a,
	0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x76, 0x69, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x61, 0x75, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2d, 0x63, 0x69, 0x2d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2d, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_monitor_protobuf_monitor_proto_rawDescOnce sync.Once
	file_pkg_monitor_protobuf_monitor_proto_rawDescData = file_pkg_monitor_protobuf_monitor_proto_rawDesc
)

func file_pkg_monitor_protobuf_monitor_proto_rawDescGZIP() []byte {
	file_pkg_monitor_protobuf_monitor_proto_rawDescOnce.Do(func() {
		file_pkg_monitor_protobuf_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_monitor_protobuf_monitor_proto_rawDescData)
	})
	return file_pkg_monitor_protobuf_monitor_proto_rawDescData
}

var (
	file_pkg_monitor_protobuf_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_pkg_monitor_protobuf_monitor_proto_goTypes  = []interface{}{
		(*Empty)(nil),                 // 0: monitor.Empty
		(*Config)(nil),                // 1: monitor.Config
		(*Telemetry)(nil),             // 2: monitor.Telemetry
		(*Entity)(nil),                // 3: monitor.Entity
		(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	}
)

var file_pkg_monitor_protobuf_monitor_proto_depIdxs = []int32{
	3,  // 0: monitor.Telemetry.projects:type_name -> monitor.Entity
	3,  // 1: monitor.Telemetry.refs:type_name -> monitor.Entity
	3,  // 2: monitor.Telemetry.envs:type_name -> monitor.Entity
	3,  // 3: monitor.Telemetry.metrics:type_name -> monitor.Entity
	4,  // 4: monitor.Entity.last_gc:type_name -> google.protobuf.Timestamp
	4,  // 5: monitor.Entity.last_pull:type_name -> google.protobuf.Timestamp
	4,  // 6: monitor.Entity.next_gc:type_name -> google.protobuf.Timestamp
	4,  // 7: monitor.Entity.next_pull:type_name -> google.protobuf.Timestamp
	0,  // 8: monitor.Monitor.GetConfig:input_type -> monitor.Empty
	0,  // 9: monitor.Monitor.GetTelemetry:input_type -> monitor.Empty
	1,  // 10: monitor.Monitor.GetConfig:output_type -> monitor.Config
	2,  // 11: monitor.Monitor.GetTelemetry:output_type -> monitor.Telemetry
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_monitor_protobuf_monitor_proto_init() }
func file_pkg_monitor_protobuf_monitor_proto_init() {
	if File_pkg_monitor_protobuf_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_monitor_protobuf_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pkg_monitor_protobuf_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_pkg_monitor_protobuf_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Telemetry); i {
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
		file_pkg_monitor_protobuf_monitor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
			RawDescriptor: file_pkg_monitor_protobuf_monitor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_monitor_protobuf_monitor_proto_goTypes,
		DependencyIndexes: file_pkg_monitor_protobuf_monitor_proto_depIdxs,
		MessageInfos:      file_pkg_monitor_protobuf_monitor_proto_msgTypes,
	}.Build()
	File_pkg_monitor_protobuf_monitor_proto = out.File
	file_pkg_monitor_protobuf_monitor_proto_rawDesc = nil
	file_pkg_monitor_protobuf_monitor_proto_goTypes = nil
	file_pkg_monitor_protobuf_monitor_proto_depIdxs = nil
}
