// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/v1/inference_server_worker.proto

package v1

import (
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

type EngineStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EngineId   string                   `protobuf:"bytes,1,opt,name=engine_id,json=engineId,proto3" json:"engine_id,omitempty"`
	ModelIds   []string                 `protobuf:"bytes,2,rep,name=model_ids,json=modelIds,proto3" json:"model_ids,omitempty"`
	SyncStatus *EngineStatus_SyncStatus `protobuf:"bytes,3,opt,name=sync_status,json=syncStatus,proto3" json:"sync_status,omitempty"`
}

func (x *EngineStatus) Reset() {
	*x = EngineStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngineStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngineStatus) ProtoMessage() {}

func (x *EngineStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngineStatus.ProtoReflect.Descriptor instead.
func (*EngineStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{0}
}

func (x *EngineStatus) GetEngineId() string {
	if x != nil {
		return x.EngineId
	}
	return ""
}

func (x *EngineStatus) GetModelIds() []string {
	if x != nil {
		return x.ModelIds
	}
	return nil
}

func (x *EngineStatus) GetSyncStatus() *EngineStatus_SyncStatus {
	if x != nil {
		return x.SyncStatus
	}
	return nil
}

type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderValue) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type HttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32                   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Status     string                  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Header     map[string]*HeaderValue `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// body is empty for server sent events.
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HttpResponse) Reset() {
	*x = HttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpResponse) ProtoMessage() {}

func (x *HttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpResponse.ProtoReflect.Descriptor instead.
func (*HttpResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{2}
}

func (x *HttpResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *HttpResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HttpResponse) GetHeader() map[string]*HeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HttpResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ServerSentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	IsLastEvent bool   `protobuf:"varint,2,opt,name=is_last_event,json=isLastEvent,proto3" json:"is_last_event,omitempty"`
}

func (x *ServerSentEvent) Reset() {
	*x = ServerSentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerSentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerSentEvent) ProtoMessage() {}

func (x *ServerSentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerSentEvent.ProtoReflect.Descriptor instead.
func (*ServerSentEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{3}
}

func (x *ServerSentEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ServerSentEvent) GetIsLastEvent() bool {
	if x != nil {
		return x.IsLastEvent
	}
	return false
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// Types that are assignable to Message:
	//
	//	*TaskResult_HttpResponse
	//	*TaskResult_ServerSentEvent
	Message isTaskResult_Message `protobuf_oneof:"message"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{4}
}

func (x *TaskResult) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (m *TaskResult) GetMessage() isTaskResult_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *TaskResult) GetHttpResponse() *HttpResponse {
	if x, ok := x.GetMessage().(*TaskResult_HttpResponse); ok {
		return x.HttpResponse
	}
	return nil
}

func (x *TaskResult) GetServerSentEvent() *ServerSentEvent {
	if x, ok := x.GetMessage().(*TaskResult_ServerSentEvent); ok {
		return x.ServerSentEvent
	}
	return nil
}

type isTaskResult_Message interface {
	isTaskResult_Message()
}

type TaskResult_HttpResponse struct {
	HttpResponse *HttpResponse `protobuf:"bytes,2,opt,name=http_response,json=httpResponse,proto3,oneof"`
}

type TaskResult_ServerSentEvent struct {
	ServerSentEvent *ServerSentEvent `protobuf:"bytes,3,opt,name=server_sent_event,json=serverSentEvent,proto3,oneof"`
}

func (*TaskResult_HttpResponse) isTaskResult_Message() {}

func (*TaskResult_ServerSentEvent) isTaskResult_Message() {}

type ProcessTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ProcessTasksRequest_EngineStatus
	//	*ProcessTasksRequest_TaskResult
	Message isProcessTasksRequest_Message `protobuf_oneof:"message"`
}

func (x *ProcessTasksRequest) Reset() {
	*x = ProcessTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessTasksRequest) ProtoMessage() {}

func (x *ProcessTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessTasksRequest.ProtoReflect.Descriptor instead.
func (*ProcessTasksRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{5}
}

func (m *ProcessTasksRequest) GetMessage() isProcessTasksRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ProcessTasksRequest) GetEngineStatus() *EngineStatus {
	if x, ok := x.GetMessage().(*ProcessTasksRequest_EngineStatus); ok {
		return x.EngineStatus
	}
	return nil
}

func (x *ProcessTasksRequest) GetTaskResult() *TaskResult {
	if x, ok := x.GetMessage().(*ProcessTasksRequest_TaskResult); ok {
		return x.TaskResult
	}
	return nil
}

type isProcessTasksRequest_Message interface {
	isProcessTasksRequest_Message()
}

type ProcessTasksRequest_EngineStatus struct {
	EngineStatus *EngineStatus `protobuf:"bytes,1,opt,name=engine_status,json=engineStatus,proto3,oneof"`
}

type ProcessTasksRequest_TaskResult struct {
	TaskResult *TaskResult `protobuf:"bytes,2,opt,name=task_result,json=taskResult,proto3,oneof"`
}

func (*ProcessTasksRequest_EngineStatus) isProcessTasksRequest_Message() {}

func (*ProcessTasksRequest_TaskResult) isProcessTasksRequest_Message() {}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Request *CreateChatCompletionRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Header  map[string]*HeaderValue      `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{6}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetRequest() *CreateChatCompletionRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Task) GetHeader() map[string]*HeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

type ProcessTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewTask *Task `protobuf:"bytes,1,opt,name=new_task,json=newTask,proto3" json:"new_task,omitempty"`
}

func (x *ProcessTasksResponse) Reset() {
	*x = ProcessTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessTasksResponse) ProtoMessage() {}

func (x *ProcessTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessTasksResponse.ProtoReflect.Descriptor instead.
func (*ProcessTasksResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{7}
}

func (x *ProcessTasksResponse) GetNewTask() *Task {
	if x != nil {
		return x.NewTask
	}
	return nil
}

type EngineStatus_SyncStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in_progress_model_ids is a list of model ids that are currently being synced.
	InProgressModelIds []string `protobuf:"bytes,1,rep,name=in_progress_model_ids,json=inProgressModelIds,proto3" json:"in_progress_model_ids,omitempty"`
}

func (x *EngineStatus_SyncStatus) Reset() {
	*x = EngineStatus_SyncStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_inference_server_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngineStatus_SyncStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngineStatus_SyncStatus) ProtoMessage() {}

func (x *EngineStatus_SyncStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_inference_server_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngineStatus_SyncStatus.ProtoReflect.Descriptor instead.
func (*EngineStatus_SyncStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_inference_server_worker_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EngineStatus_SyncStatus) GetInProgressModelIds() []string {
	if x != nil {
		return x.InProgressModelIds
	}
	return nil
}

var File_api_v1_inference_server_worker_proto protoreflect.FileDescriptor

var file_api_v1_inference_server_worker_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x59, 0x0a, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x3f, 0x0a, 0x0a,
	0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x22, 0x25, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x97, 0x02, 0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x51,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x67, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x49,
	0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73,
	0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x0a, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x54, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0d, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x0c, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x4e, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x04,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x51, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x1a, 0x67, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x58, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x54, 0x61, 0x73, 0x6b, 0x32, 0x9c, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x81, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x34, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x6c, 0x6d, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_inference_server_worker_proto_rawDescOnce sync.Once
	file_api_v1_inference_server_worker_proto_rawDescData = file_api_v1_inference_server_worker_proto_rawDesc
)

func file_api_v1_inference_server_worker_proto_rawDescGZIP() []byte {
	file_api_v1_inference_server_worker_proto_rawDescOnce.Do(func() {
		file_api_v1_inference_server_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_inference_server_worker_proto_rawDescData)
	})
	return file_api_v1_inference_server_worker_proto_rawDescData
}

var file_api_v1_inference_server_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_inference_server_worker_proto_goTypes = []interface{}{
	(*EngineStatus)(nil),                // 0: llmoperator.inference.server.v1.EngineStatus
	(*HeaderValue)(nil),                 // 1: llmoperator.inference.server.v1.HeaderValue
	(*HttpResponse)(nil),                // 2: llmoperator.inference.server.v1.HttpResponse
	(*ServerSentEvent)(nil),             // 3: llmoperator.inference.server.v1.ServerSentEvent
	(*TaskResult)(nil),                  // 4: llmoperator.inference.server.v1.TaskResult
	(*ProcessTasksRequest)(nil),         // 5: llmoperator.inference.server.v1.ProcessTasksRequest
	(*Task)(nil),                        // 6: llmoperator.inference.server.v1.Task
	(*ProcessTasksResponse)(nil),        // 7: llmoperator.inference.server.v1.ProcessTasksResponse
	(*EngineStatus_SyncStatus)(nil),     // 8: llmoperator.inference.server.v1.EngineStatus.SyncStatus
	nil,                                 // 9: llmoperator.inference.server.v1.HttpResponse.HeaderEntry
	nil,                                 // 10: llmoperator.inference.server.v1.Task.HeaderEntry
	(*CreateChatCompletionRequest)(nil), // 11: llmoperator.chat.server.v1.CreateChatCompletionRequest
}
var file_api_v1_inference_server_worker_proto_depIdxs = []int32{
	8,  // 0: llmoperator.inference.server.v1.EngineStatus.sync_status:type_name -> llmoperator.inference.server.v1.EngineStatus.SyncStatus
	9,  // 1: llmoperator.inference.server.v1.HttpResponse.header:type_name -> llmoperator.inference.server.v1.HttpResponse.HeaderEntry
	2,  // 2: llmoperator.inference.server.v1.TaskResult.http_response:type_name -> llmoperator.inference.server.v1.HttpResponse
	3,  // 3: llmoperator.inference.server.v1.TaskResult.server_sent_event:type_name -> llmoperator.inference.server.v1.ServerSentEvent
	0,  // 4: llmoperator.inference.server.v1.ProcessTasksRequest.engine_status:type_name -> llmoperator.inference.server.v1.EngineStatus
	4,  // 5: llmoperator.inference.server.v1.ProcessTasksRequest.task_result:type_name -> llmoperator.inference.server.v1.TaskResult
	11, // 6: llmoperator.inference.server.v1.Task.request:type_name -> llmoperator.chat.server.v1.CreateChatCompletionRequest
	10, // 7: llmoperator.inference.server.v1.Task.header:type_name -> llmoperator.inference.server.v1.Task.HeaderEntry
	6,  // 8: llmoperator.inference.server.v1.ProcessTasksResponse.new_task:type_name -> llmoperator.inference.server.v1.Task
	1,  // 9: llmoperator.inference.server.v1.HttpResponse.HeaderEntry.value:type_name -> llmoperator.inference.server.v1.HeaderValue
	1,  // 10: llmoperator.inference.server.v1.Task.HeaderEntry.value:type_name -> llmoperator.inference.server.v1.HeaderValue
	5,  // 11: llmoperator.inference.server.v1.InferenceWorkerService.ProcessTasks:input_type -> llmoperator.inference.server.v1.ProcessTasksRequest
	7,  // 12: llmoperator.inference.server.v1.InferenceWorkerService.ProcessTasks:output_type -> llmoperator.inference.server.v1.ProcessTasksResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_v1_inference_server_worker_proto_init() }
func file_api_v1_inference_server_worker_proto_init() {
	if File_api_v1_inference_server_worker_proto != nil {
		return
	}
	file_api_v1_inference_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_inference_server_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngineStatus); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpResponse); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerSentEvent); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResult); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessTasksRequest); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessTasksResponse); i {
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
		file_api_v1_inference_server_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngineStatus_SyncStatus); i {
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
	file_api_v1_inference_server_worker_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*TaskResult_HttpResponse)(nil),
		(*TaskResult_ServerSentEvent)(nil),
	}
	file_api_v1_inference_server_worker_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ProcessTasksRequest_EngineStatus)(nil),
		(*ProcessTasksRequest_TaskResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_inference_server_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_inference_server_worker_proto_goTypes,
		DependencyIndexes: file_api_v1_inference_server_worker_proto_depIdxs,
		MessageInfos:      file_api_v1_inference_server_worker_proto_msgTypes,
	}.Build()
	File_api_v1_inference_server_worker_proto = out.File
	file_api_v1_inference_server_worker_proto_rawDesc = nil
	file_api_v1_inference_server_worker_proto_goTypes = nil
	file_api_v1_inference_server_worker_proto_depIdxs = nil
}
