// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: vm.proto

package vmrpc

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

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{0}
}

func (x *GetValueRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetValueResponse) Reset() {
	*x = GetValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueResponse) ProtoMessage() {}

func (x *GetValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueResponse.ProtoReflect.Descriptor instead.
func (*GetValueResponse) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{1}
}

func (x *GetValueResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type VMCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calldata        []string `protobuf:"bytes,1,rep,name=calldata,proto3" json:"calldata,omitempty"`
	CallerAddress   string   `protobuf:"bytes,2,opt,name=caller_address,json=callerAddress,proto3" json:"caller_address,omitempty"`
	ContractAddress string   `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Root            string   `protobuf:"bytes,4,opt,name=root,proto3" json:"root,omitempty"`
	Selector        string   `protobuf:"bytes,5,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *VMCallRequest) Reset() {
	*x = VMCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMCallRequest) ProtoMessage() {}

func (x *VMCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMCallRequest.ProtoReflect.Descriptor instead.
func (*VMCallRequest) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{2}
}

func (x *VMCallRequest) GetCalldata() []string {
	if x != nil {
		return x.Calldata
	}
	return nil
}

func (x *VMCallRequest) GetCallerAddress() string {
	if x != nil {
		return x.CallerAddress
	}
	return ""
}

func (x *VMCallRequest) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *VMCallRequest) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *VMCallRequest) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

type VMCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retdata [][]byte `protobuf:"bytes,1,rep,name=retdata,proto3" json:"retdata,omitempty"`
}

func (x *VMCallResponse) Reset() {
	*x = VMCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMCallResponse) ProtoMessage() {}

func (x *VMCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMCallResponse.ProtoReflect.Descriptor instead.
func (*VMCallResponse) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{3}
}

func (x *VMCallResponse) GetRetdata() [][]byte {
	if x != nil {
		return x.Retdata
	}
	return nil
}

var File_vm_proto protoreflect.FileDescriptor

var file_vm_proto_rawDesc = []byte{
	0x0a, 0x08, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x56, 0x4d,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x6c, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x6c, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x56, 0x4d, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x32, 0x43, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x2f, 0x0a, 0x02, 0x56, 0x4d,
	0x12, 0x29, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x56, 0x4d, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x56, 0x4d, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x64, 0x45, 0x74, 0x68, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x6d, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vm_proto_rawDescOnce sync.Once
	file_vm_proto_rawDescData = file_vm_proto_rawDesc
)

func file_vm_proto_rawDescGZIP() []byte {
	file_vm_proto_rawDescOnce.Do(func() {
		file_vm_proto_rawDescData = protoimpl.X.CompressGZIP(file_vm_proto_rawDescData)
	})
	return file_vm_proto_rawDescData
}

var file_vm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vm_proto_goTypes = []interface{}{
	(*GetValueRequest)(nil),  // 0: GetValueRequest
	(*GetValueResponse)(nil), // 1: GetValueResponse
	(*VMCallRequest)(nil),    // 2: VMCallRequest
	(*VMCallResponse)(nil),   // 3: VMCallResponse
}
var file_vm_proto_depIdxs = []int32{
	0, // 0: StorageAdapter.GetValue:input_type -> GetValueRequest
	2, // 1: VM.Call:input_type -> VMCallRequest
	1, // 2: StorageAdapter.GetValue:output_type -> GetValueResponse
	3, // 3: VM.Call:output_type -> VMCallResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vm_proto_init() }
func file_vm_proto_init() {
	if File_vm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_vm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueResponse); i {
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
		file_vm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMCallRequest); i {
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
		file_vm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMCallResponse); i {
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
			RawDescriptor: file_vm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_vm_proto_goTypes,
		DependencyIndexes: file_vm_proto_depIdxs,
		MessageInfos:      file_vm_proto_msgTypes,
	}.Build()
	File_vm_proto = out.File
	file_vm_proto_rawDesc = nil
	file_vm_proto_goTypes = nil
	file_vm_proto_depIdxs = nil
}
