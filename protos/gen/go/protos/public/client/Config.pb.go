// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: client/Config.proto

package client

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Client config file
// all fields optional
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	DevKey     string `protobuf:"bytes,2,opt,name=dev_key,json=devKey,proto3" json:"dev_key,omitempty"`
	Host       string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Workspace  string `protobuf:"bytes,4,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Project    string `protobuf:"bytes,5,opt,name=project,proto3" json:"project,omitempty"`
	Experiment string `protobuf:"bytes,6,opt,name=experiment,proto3" json:"experiment,omitempty"`
	Dataset    string `protobuf:"bytes,7,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_Config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_client_Config_proto_msgTypes[0]
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
	return file_client_Config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Config) GetDevKey() string {
	if x != nil {
		return x.DevKey
	}
	return ""
}

func (x *Config) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Config) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *Config) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Config) GetExperiment() string {
	if x != nil {
		return x.Experiment
	}
	return ""
}

func (x *Config) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

var File_client_Config_proto protoreflect.FileDescriptor

var file_client_Config_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x42, 0x41, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x72, 0x74, 0x61, 0x41, 0x49, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_Config_proto_rawDescOnce sync.Once
	file_client_Config_proto_rawDescData = file_client_Config_proto_rawDesc
)

func file_client_Config_proto_rawDescGZIP() []byte {
	file_client_Config_proto_rawDescOnce.Do(func() {
		file_client_Config_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_Config_proto_rawDescData)
	})
	return file_client_Config_proto_rawDescData
}

var file_client_Config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_Config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: ai.verta.client.Config
}
var file_client_Config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_Config_proto_init() }
func file_client_Config_proto_init() {
	if File_client_Config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_Config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_Config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_Config_proto_goTypes,
		DependencyIndexes: file_client_Config_proto_depIdxs,
		MessageInfos:      file_client_Config_proto_msgTypes,
	}.Build()
	File_client_Config_proto = out.File
	file_client_Config_proto_rawDesc = nil
	file_client_Config_proto_goTypes = nil
	file_client_Config_proto_depIdxs = nil
}
