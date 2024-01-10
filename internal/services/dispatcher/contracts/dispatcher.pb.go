// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: dispatcher.proto

package contracts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionType int32

const (
	ActionType_START_STEP_RUN  ActionType = 0
	ActionType_CANCEL_STEP_RUN ActionType = 1
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "START_STEP_RUN",
		1: "CANCEL_STEP_RUN",
	}
	ActionType_value = map[string]int32{
		"START_STEP_RUN":  0,
		"CANCEL_STEP_RUN": 1,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_dispatcher_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_dispatcher_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{0}
}

type ActionEventType int32

const (
	ActionEventType_STEP_EVENT_TYPE_UNKNOWN   ActionEventType = 0
	ActionEventType_STEP_EVENT_TYPE_STARTED   ActionEventType = 1
	ActionEventType_STEP_EVENT_TYPE_COMPLETED ActionEventType = 2
	ActionEventType_STEP_EVENT_TYPE_FAILED    ActionEventType = 3
)

// Enum value maps for ActionEventType.
var (
	ActionEventType_name = map[int32]string{
		0: "STEP_EVENT_TYPE_UNKNOWN",
		1: "STEP_EVENT_TYPE_STARTED",
		2: "STEP_EVENT_TYPE_COMPLETED",
		3: "STEP_EVENT_TYPE_FAILED",
	}
	ActionEventType_value = map[string]int32{
		"STEP_EVENT_TYPE_UNKNOWN":   0,
		"STEP_EVENT_TYPE_STARTED":   1,
		"STEP_EVENT_TYPE_COMPLETED": 2,
		"STEP_EVENT_TYPE_FAILED":    3,
	}
)

func (x ActionEventType) Enum() *ActionEventType {
	p := new(ActionEventType)
	*p = x
	return p
}

func (x ActionEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_dispatcher_proto_enumTypes[1].Descriptor()
}

func (ActionEventType) Type() protoreflect.EnumType {
	return &file_dispatcher_proto_enumTypes[1]
}

func (x ActionEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionEventType.Descriptor instead.
func (ActionEventType) EnumDescriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{1}
}

type WorkerRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the name of the worker
	WorkerName string `protobuf:"bytes,2,opt,name=workerName,proto3" json:"workerName,omitempty"`
	// a list of actions that this worker can run
	Actions []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	// (optional) the services for this worker
	Services []string `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *WorkerRegisterRequest) Reset() {
	*x = WorkerRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerRegisterRequest) ProtoMessage() {}

func (x *WorkerRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerRegisterRequest.ProtoReflect.Descriptor instead.
func (*WorkerRegisterRequest) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerRegisterRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *WorkerRegisterRequest) GetWorkerName() string {
	if x != nil {
		return x.WorkerName
	}
	return ""
}

func (x *WorkerRegisterRequest) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *WorkerRegisterRequest) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

type WorkerRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the worker
	WorkerId string `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
	// the name of the worker
	WorkerName string `protobuf:"bytes,3,opt,name=workerName,proto3" json:"workerName,omitempty"`
}

func (x *WorkerRegisterResponse) Reset() {
	*x = WorkerRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerRegisterResponse) ProtoMessage() {}

func (x *WorkerRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerRegisterResponse.ProtoReflect.Descriptor instead.
func (*WorkerRegisterResponse) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerRegisterResponse) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *WorkerRegisterResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *WorkerRegisterResponse) GetWorkerName() string {
	if x != nil {
		return x.WorkerName
	}
	return ""
}

type AssignedAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the job id
	JobId string `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId,omitempty"`
	// the job name
	JobName string `protobuf:"bytes,3,opt,name=jobName,proto3" json:"jobName,omitempty"`
	// the job run id
	JobRunId string `protobuf:"bytes,4,opt,name=jobRunId,proto3" json:"jobRunId,omitempty"`
	// the step id
	StepId string `protobuf:"bytes,5,opt,name=stepId,proto3" json:"stepId,omitempty"`
	// the step run id
	StepRunId string `protobuf:"bytes,6,opt,name=stepRunId,proto3" json:"stepRunId,omitempty"`
	// the action id
	ActionId string `protobuf:"bytes,7,opt,name=actionId,proto3" json:"actionId,omitempty"`
	// the action type
	ActionType ActionType `protobuf:"varint,8,opt,name=actionType,proto3,enum=ActionType" json:"actionType,omitempty"`
	// the action payload
	ActionPayload string `protobuf:"bytes,9,opt,name=actionPayload,proto3" json:"actionPayload,omitempty"`
}

func (x *AssignedAction) Reset() {
	*x = AssignedAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignedAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignedAction) ProtoMessage() {}

func (x *AssignedAction) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignedAction.ProtoReflect.Descriptor instead.
func (*AssignedAction) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{2}
}

func (x *AssignedAction) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *AssignedAction) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *AssignedAction) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *AssignedAction) GetJobRunId() string {
	if x != nil {
		return x.JobRunId
	}
	return ""
}

func (x *AssignedAction) GetStepId() string {
	if x != nil {
		return x.StepId
	}
	return ""
}

func (x *AssignedAction) GetStepRunId() string {
	if x != nil {
		return x.StepRunId
	}
	return ""
}

func (x *AssignedAction) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *AssignedAction) GetActionType() ActionType {
	if x != nil {
		return x.ActionType
	}
	return ActionType_START_STEP_RUN
}

func (x *AssignedAction) GetActionPayload() string {
	if x != nil {
		return x.ActionPayload
	}
	return ""
}

type WorkerListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the worker
	WorkerId string `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
}

func (x *WorkerListenRequest) Reset() {
	*x = WorkerListenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerListenRequest) ProtoMessage() {}

func (x *WorkerListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerListenRequest.ProtoReflect.Descriptor instead.
func (*WorkerListenRequest) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{3}
}

func (x *WorkerListenRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *WorkerListenRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type WorkerUnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id to unsubscribe from
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the worker
	WorkerId string `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
}

func (x *WorkerUnsubscribeRequest) Reset() {
	*x = WorkerUnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerUnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerUnsubscribeRequest) ProtoMessage() {}

func (x *WorkerUnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerUnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*WorkerUnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{4}
}

func (x *WorkerUnsubscribeRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *WorkerUnsubscribeRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type WorkerUnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id to unsubscribe from
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the worker
	WorkerId string `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
}

func (x *WorkerUnsubscribeResponse) Reset() {
	*x = WorkerUnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerUnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerUnsubscribeResponse) ProtoMessage() {}

func (x *WorkerUnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerUnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*WorkerUnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{5}
}

func (x *WorkerUnsubscribeResponse) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *WorkerUnsubscribeResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type ActionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the worker
	WorkerId string `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
	// the id of the job
	JobId string `protobuf:"bytes,3,opt,name=jobId,proto3" json:"jobId,omitempty"`
	// the job run id
	JobRunId string `protobuf:"bytes,4,opt,name=jobRunId,proto3" json:"jobRunId,omitempty"`
	// the id of the step
	StepId string `protobuf:"bytes,5,opt,name=stepId,proto3" json:"stepId,omitempty"`
	// the step run id
	StepRunId string `protobuf:"bytes,6,opt,name=stepRunId,proto3" json:"stepRunId,omitempty"`
	// the action id
	ActionId       string                 `protobuf:"bytes,7,opt,name=actionId,proto3" json:"actionId,omitempty"`
	EventTimestamp *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	// the step event type
	EventType ActionEventType `protobuf:"varint,9,opt,name=eventType,proto3,enum=ActionEventType" json:"eventType,omitempty"`
	// the event payload
	EventPayload string `protobuf:"bytes,10,opt,name=eventPayload,proto3" json:"eventPayload,omitempty"`
}

func (x *ActionEvent) Reset() {
	*x = ActionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionEvent) ProtoMessage() {}

func (x *ActionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionEvent.ProtoReflect.Descriptor instead.
func (*ActionEvent) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{6}
}

func (x *ActionEvent) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ActionEvent) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *ActionEvent) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *ActionEvent) GetJobRunId() string {
	if x != nil {
		return x.JobRunId
	}
	return ""
}

func (x *ActionEvent) GetStepId() string {
	if x != nil {
		return x.StepId
	}
	return ""
}

func (x *ActionEvent) GetStepRunId() string {
	if x != nil {
		return x.StepRunId
	}
	return ""
}

func (x *ActionEvent) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *ActionEvent) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *ActionEvent) GetEventType() ActionEventType {
	if x != nil {
		return x.EventType
	}
	return ActionEventType_STEP_EVENT_TYPE_UNKNOWN
}

func (x *ActionEvent) GetEventPayload() string {
	if x != nil {
		return x.EventPayload
	}
	return ""
}

type ActionEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the worker
	WorkerId string `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
}

func (x *ActionEventResponse) Reset() {
	*x = ActionEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionEventResponse) ProtoMessage() {}

func (x *ActionEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionEventResponse.ProtoReflect.Descriptor instead.
func (*ActionEventResponse) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{7}
}

func (x *ActionEventResponse) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ActionEventResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

var File_dispatcher_proto protoreflect.FileDescriptor

var file_dispatcher_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22,
	0x70, 0x0a, 0x16, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x9d, 0x02, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x65, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x52, 0x75, 0x6e, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x65, 0x70, 0x52, 0x75, 0x6e,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x4d, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x52, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x19, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0xe1, 0x02, 0x0a, 0x0b, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x65, 0x70, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x65, 0x70, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4d, 0x0a,
	0x13, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x35, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x52, 0x55,
	0x4e, 0x10, 0x01, 0x2a, 0x86, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x45, 0x50, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0x81, 0x02, 0x0a,
	0x0a, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x37, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x14, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dispatcher_proto_rawDescOnce sync.Once
	file_dispatcher_proto_rawDescData = file_dispatcher_proto_rawDesc
)

func file_dispatcher_proto_rawDescGZIP() []byte {
	file_dispatcher_proto_rawDescOnce.Do(func() {
		file_dispatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_dispatcher_proto_rawDescData)
	})
	return file_dispatcher_proto_rawDescData
}

var file_dispatcher_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dispatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dispatcher_proto_goTypes = []interface{}{
	(ActionType)(0),                   // 0: ActionType
	(ActionEventType)(0),              // 1: ActionEventType
	(*WorkerRegisterRequest)(nil),     // 2: WorkerRegisterRequest
	(*WorkerRegisterResponse)(nil),    // 3: WorkerRegisterResponse
	(*AssignedAction)(nil),            // 4: AssignedAction
	(*WorkerListenRequest)(nil),       // 5: WorkerListenRequest
	(*WorkerUnsubscribeRequest)(nil),  // 6: WorkerUnsubscribeRequest
	(*WorkerUnsubscribeResponse)(nil), // 7: WorkerUnsubscribeResponse
	(*ActionEvent)(nil),               // 8: ActionEvent
	(*ActionEventResponse)(nil),       // 9: ActionEventResponse
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
}
var file_dispatcher_proto_depIdxs = []int32{
	0,  // 0: AssignedAction.actionType:type_name -> ActionType
	10, // 1: ActionEvent.eventTimestamp:type_name -> google.protobuf.Timestamp
	1,  // 2: ActionEvent.eventType:type_name -> ActionEventType
	2,  // 3: Dispatcher.Register:input_type -> WorkerRegisterRequest
	5,  // 4: Dispatcher.Listen:input_type -> WorkerListenRequest
	8,  // 5: Dispatcher.SendActionEvent:input_type -> ActionEvent
	6,  // 6: Dispatcher.Unsubscribe:input_type -> WorkerUnsubscribeRequest
	3,  // 7: Dispatcher.Register:output_type -> WorkerRegisterResponse
	4,  // 8: Dispatcher.Listen:output_type -> AssignedAction
	9,  // 9: Dispatcher.SendActionEvent:output_type -> ActionEventResponse
	7,  // 10: Dispatcher.Unsubscribe:output_type -> WorkerUnsubscribeResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_dispatcher_proto_init() }
func file_dispatcher_proto_init() {
	if File_dispatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dispatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerRegisterRequest); i {
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
		file_dispatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerRegisterResponse); i {
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
		file_dispatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignedAction); i {
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
		file_dispatcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerListenRequest); i {
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
		file_dispatcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerUnsubscribeRequest); i {
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
		file_dispatcher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerUnsubscribeResponse); i {
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
		file_dispatcher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionEvent); i {
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
		file_dispatcher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionEventResponse); i {
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
			RawDescriptor: file_dispatcher_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dispatcher_proto_goTypes,
		DependencyIndexes: file_dispatcher_proto_depIdxs,
		EnumInfos:         file_dispatcher_proto_enumTypes,
		MessageInfos:      file_dispatcher_proto_msgTypes,
	}.Build()
	File_dispatcher_proto = out.File
	file_dispatcher_proto_rawDesc = nil
	file_dispatcher_proto_goTypes = nil
	file_dispatcher_proto_depIdxs = nil
}
