// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.21.12
// source: proto/process.proto

package __

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

type StartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command       string                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Args          []string               `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	AutoRestart   string                 `protobuf:"bytes,4,opt,name=autoRestart,proto3" json:"autoRestart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	mi := &file_proto_process_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_process_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_proto_process_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StartRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *StartRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *StartRequest) GetAutoRestart() string {
	if x != nil {
		return x.AutoRestart
	}
	return ""
}

type StopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Force         bool                   `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	mi := &file_proto_process_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_process_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_proto_process_proto_rawDescGZIP(), []int{1}
}

func (x *StopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StopRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Verbose       bool                   `protobuf:"varint,1,opt,name=verbose,proto3" json:"verbose,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_proto_process_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_process_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_process_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

type ProcessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessResponse) Reset() {
	*x = ProcessResponse{}
	mi := &file_proto_process_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessResponse) ProtoMessage() {}

func (x *ProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_process_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessResponse.ProtoReflect.Descriptor instead.
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return file_proto_process_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ProcessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ProcessInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pid           int32                  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessInfo) Reset() {
	*x = ProcessInfo{}
	mi := &file_proto_process_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInfo) ProtoMessage() {}

func (x *ProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_process_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInfo.ProtoReflect.Descriptor instead.
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return file_proto_process_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProcessInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Processes     []*ProcessInfo         `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_proto_process_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_process_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_process_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetProcesses() []*ProcessInfo {
	if x != nil {
		return x.Processes
	}
	return nil
}

var File_proto_process_proto protoreflect.FileDescriptor

var file_proto_process_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x37, 0x0a, 0x0b, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4b, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x49, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x32, 0xf6, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4d, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_process_proto_rawDescOnce sync.Once
	file_proto_process_proto_rawDescData = file_proto_process_proto_rawDesc
)

func file_proto_process_proto_rawDescGZIP() []byte {
	file_proto_process_proto_rawDescOnce.Do(func() {
		file_proto_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_process_proto_rawDescData)
	})
	return file_proto_process_proto_rawDescData
}

var file_proto_process_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_process_proto_goTypes = []any{
	(*StartRequest)(nil),    // 0: processmanager.StartRequest
	(*StopRequest)(nil),     // 1: processmanager.StopRequest
	(*ListRequest)(nil),     // 2: processmanager.ListRequest
	(*ProcessResponse)(nil), // 3: processmanager.ProcessResponse
	(*ProcessInfo)(nil),     // 4: processmanager.ProcessInfo
	(*ListResponse)(nil),    // 5: processmanager.ListResponse
}
var file_proto_process_proto_depIdxs = []int32{
	4, // 0: processmanager.ListResponse.processes:type_name -> processmanager.ProcessInfo
	0, // 1: processmanager.ProcessManager.StartProcess:input_type -> processmanager.StartRequest
	1, // 2: processmanager.ProcessManager.StopProcess:input_type -> processmanager.StopRequest
	2, // 3: processmanager.ProcessManager.ListProcess:input_type -> processmanager.ListRequest
	3, // 4: processmanager.ProcessManager.StartProcess:output_type -> processmanager.ProcessResponse
	3, // 5: processmanager.ProcessManager.StopProcess:output_type -> processmanager.ProcessResponse
	5, // 6: processmanager.ProcessManager.ListProcess:output_type -> processmanager.ListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_process_proto_init() }
func file_proto_process_proto_init() {
	if File_proto_process_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_process_proto_goTypes,
		DependencyIndexes: file_proto_process_proto_depIdxs,
		MessageInfos:      file_proto_process_proto_msgTypes,
	}.Build()
	File_proto_process_proto = out.File
	file_proto_process_proto_rawDesc = nil
	file_proto_process_proto_goTypes = nil
	file_proto_process_proto_depIdxs = nil
}
