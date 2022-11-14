// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: config/module/featureflag/v1/featureflag.proto

package featureflagv1

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

type Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flags map[string]bool `protobuf:"bytes,1,rep,name=flags,proto3" json:"flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Simple) Reset() {
	*x = Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_module_featureflag_v1_featureflag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple) ProtoMessage() {}

func (x *Simple) ProtoReflect() protoreflect.Message {
	mi := &file_config_module_featureflag_v1_featureflag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple.ProtoReflect.Descriptor instead.
func (*Simple) Descriptor() ([]byte, []int) {
	return file_config_module_featureflag_v1_featureflag_proto_rawDescGZIP(), []int{0}
}

func (x *Simple) GetFlags() map[string]bool {
	if x != nil {
		return x.Flags
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Config_Simple
	Type isConfig_Type `protobuf_oneof:"type"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_module_featureflag_v1_featureflag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_module_featureflag_v1_featureflag_proto_msgTypes[1]
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
	return file_config_module_featureflag_v1_featureflag_proto_rawDescGZIP(), []int{1}
}

func (m *Config) GetType() isConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Config) GetSimple() *Simple {
	if x, ok := x.GetType().(*Config_Simple); ok {
		return x.Simple
	}
	return nil
}

type isConfig_Type interface {
	isConfig_Type()
}

type Config_Simple struct {
	Simple *Simple `protobuf:"bytes,1,opt,name=simple,proto3,oneof"`
}

func (*Config_Simple) isConfig_Type() {}

var File_config_module_featureflag_v1_featureflag_proto protoreflect.FileDescriptor

var file_config_module_featureflag_v1_featureflag_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c,
	0x61, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x4c, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c,
	0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x1a, 0x38,
	0x0a, 0x0a, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c, 0x61,
	0x67, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x67,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_module_featureflag_v1_featureflag_proto_rawDescOnce sync.Once
	file_config_module_featureflag_v1_featureflag_proto_rawDescData = file_config_module_featureflag_v1_featureflag_proto_rawDesc
)

func file_config_module_featureflag_v1_featureflag_proto_rawDescGZIP() []byte {
	file_config_module_featureflag_v1_featureflag_proto_rawDescOnce.Do(func() {
		file_config_module_featureflag_v1_featureflag_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_module_featureflag_v1_featureflag_proto_rawDescData)
	})
	return file_config_module_featureflag_v1_featureflag_proto_rawDescData
}

var file_config_module_featureflag_v1_featureflag_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_module_featureflag_v1_featureflag_proto_goTypes = []interface{}{
	(*Simple)(nil), // 0: clutch.config.module.featureflag.v1.Simple
	(*Config)(nil), // 1: clutch.config.module.featureflag.v1.Config
	nil,            // 2: clutch.config.module.featureflag.v1.Simple.FlagsEntry
}
var file_config_module_featureflag_v1_featureflag_proto_depIdxs = []int32{
	2, // 0: clutch.config.module.featureflag.v1.Simple.flags:type_name -> clutch.config.module.featureflag.v1.Simple.FlagsEntry
	0, // 1: clutch.config.module.featureflag.v1.Config.simple:type_name -> clutch.config.module.featureflag.v1.Simple
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_config_module_featureflag_v1_featureflag_proto_init() }
func file_config_module_featureflag_v1_featureflag_proto_init() {
	if File_config_module_featureflag_v1_featureflag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_module_featureflag_v1_featureflag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simple); i {
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
		file_config_module_featureflag_v1_featureflag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	file_config_module_featureflag_v1_featureflag_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Config_Simple)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_module_featureflag_v1_featureflag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_module_featureflag_v1_featureflag_proto_goTypes,
		DependencyIndexes: file_config_module_featureflag_v1_featureflag_proto_depIdxs,
		MessageInfos:      file_config_module_featureflag_v1_featureflag_proto_msgTypes,
	}.Build()
	File_config_module_featureflag_v1_featureflag_proto = out.File
	file_config_module_featureflag_v1_featureflag_proto_rawDesc = nil
	file_config_module_featureflag_v1_featureflag_proto_goTypes = nil
	file_config_module_featureflag_v1_featureflag_proto_depIdxs = nil
}
