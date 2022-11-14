// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: config/service/db/postgres/v1/database.proto

package postgresv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Connection_SSLMode int32

const (
	Connection_UNSPECIFIED Connection_SSLMode = 0
	Connection_DISABLE     Connection_SSLMode = 1
	Connection_ALLOW       Connection_SSLMode = 2
	Connection_PREFER      Connection_SSLMode = 3
	Connection_REQUIRE     Connection_SSLMode = 4
	Connection_VERIFY_CA   Connection_SSLMode = 5
	Connection_VERIFY_FULL Connection_SSLMode = 6
)

// Enum value maps for Connection_SSLMode.
var (
	Connection_SSLMode_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "DISABLE",
		2: "ALLOW",
		3: "PREFER",
		4: "REQUIRE",
		5: "VERIFY_CA",
		6: "VERIFY_FULL",
	}
	Connection_SSLMode_value = map[string]int32{
		"UNSPECIFIED": 0,
		"DISABLE":     1,
		"ALLOW":       2,
		"PREFER":      3,
		"REQUIRE":     4,
		"VERIFY_CA":   5,
		"VERIFY_FULL": 6,
	}
)

func (x Connection_SSLMode) Enum() *Connection_SSLMode {
	p := new(Connection_SSLMode)
	*p = x
	return p
}

func (x Connection_SSLMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Connection_SSLMode) Descriptor() protoreflect.EnumDescriptor {
	return file_config_service_db_postgres_v1_database_proto_enumTypes[0].Descriptor()
}

func (Connection_SSLMode) Type() protoreflect.EnumType {
	return &file_config_service_db_postgres_v1_database_proto_enumTypes[0]
}

func (x Connection_SSLMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Connection_SSLMode.Descriptor instead.
func (Connection_SSLMode) EnumDescriptor() ([]byte, []int) {
	return file_config_service_db_postgres_v1_database_proto_rawDescGZIP(), []int{0, 0}
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host    string             `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port    uint32             `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	User    string             `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Dbname  string             `protobuf:"bytes,4,opt,name=dbname,proto3" json:"dbname,omitempty"`
	SslMode Connection_SSLMode `protobuf:"varint,5,opt,name=ssl_mode,json=sslMode,proto3,enum=clutch.config.service.db.postgres.v1.Connection_SSLMode" json:"ssl_mode,omitempty"`
	// TODO: GSSAPI, SSPI, Kerberos
	//
	// Types that are assignable to Authn:
	//
	//	*Connection_Password
	Authn isConnection_Authn `protobuf_oneof:"authn"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_db_postgres_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_db_postgres_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_config_service_db_postgres_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Connection) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Connection) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Connection) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

func (x *Connection) GetSslMode() Connection_SSLMode {
	if x != nil {
		return x.SslMode
	}
	return Connection_UNSPECIFIED
}

func (m *Connection) GetAuthn() isConnection_Authn {
	if m != nil {
		return m.Authn
	}
	return nil
}

func (x *Connection) GetPassword() string {
	if x, ok := x.GetAuthn().(*Connection_Password); ok {
		return x.Password
	}
	return ""
}

type isConnection_Authn interface {
	isConnection_Authn()
}

type Connection_Password struct {
	Password string `protobuf:"bytes,6,opt,name=password,proto3,oneof"`
}

func (*Connection_Password) isConnection_Authn() {}

// TODO: Expose more database/sql tunables.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// The default max idle connections is 2
	// Use -1 to specify zero idle connections
	MaxIdleConnections int32 `protobuf:"varint,2,opt,name=max_idle_connections,json=maxIdleConnections,proto3" json:"max_idle_connections,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_db_postgres_v1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_db_postgres_v1_database_proto_msgTypes[1]
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
	return file_config_service_db_postgres_v1_database_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *Config) GetMaxIdleConnections() int32 {
	if x != nil {
		return x.MaxIdleConnections
	}
	return 0
}

var File_config_service_db_postgres_v1_database_proto protoreflect.FileDescriptor

var file_config_service_db_postgres_v1_database_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x64, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x02,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xa8, 0x01, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18,
	0xff, 0xff, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x53,
	0x0a, 0x08, 0x73, 0x73, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x38, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x53, 0x4c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x73, 0x73, 0x6c, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6b, 0x0a, 0x07, 0x53, 0x53,
	0x4c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x59, 0x5f, 0x43, 0x41, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59,
	0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x06, 0x42, 0x07, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x22, 0x8c, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x14, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78,
	0x49, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79,
	0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_db_postgres_v1_database_proto_rawDescOnce sync.Once
	file_config_service_db_postgres_v1_database_proto_rawDescData = file_config_service_db_postgres_v1_database_proto_rawDesc
)

func file_config_service_db_postgres_v1_database_proto_rawDescGZIP() []byte {
	file_config_service_db_postgres_v1_database_proto_rawDescOnce.Do(func() {
		file_config_service_db_postgres_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_db_postgres_v1_database_proto_rawDescData)
	})
	return file_config_service_db_postgres_v1_database_proto_rawDescData
}

var file_config_service_db_postgres_v1_database_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_service_db_postgres_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_service_db_postgres_v1_database_proto_goTypes = []interface{}{
	(Connection_SSLMode)(0), // 0: clutch.config.service.db.postgres.v1.Connection.SSLMode
	(*Connection)(nil),      // 1: clutch.config.service.db.postgres.v1.Connection
	(*Config)(nil),          // 2: clutch.config.service.db.postgres.v1.Config
}
var file_config_service_db_postgres_v1_database_proto_depIdxs = []int32{
	0, // 0: clutch.config.service.db.postgres.v1.Connection.ssl_mode:type_name -> clutch.config.service.db.postgres.v1.Connection.SSLMode
	1, // 1: clutch.config.service.db.postgres.v1.Config.connection:type_name -> clutch.config.service.db.postgres.v1.Connection
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_config_service_db_postgres_v1_database_proto_init() }
func file_config_service_db_postgres_v1_database_proto_init() {
	if File_config_service_db_postgres_v1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_db_postgres_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
		file_config_service_db_postgres_v1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	file_config_service_db_postgres_v1_database_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Connection_Password)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_service_db_postgres_v1_database_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_db_postgres_v1_database_proto_goTypes,
		DependencyIndexes: file_config_service_db_postgres_v1_database_proto_depIdxs,
		EnumInfos:         file_config_service_db_postgres_v1_database_proto_enumTypes,
		MessageInfos:      file_config_service_db_postgres_v1_database_proto_msgTypes,
	}.Build()
	File_config_service_db_postgres_v1_database_proto = out.File
	file_config_service_db_postgres_v1_database_proto_rawDesc = nil
	file_config_service_db_postgres_v1_database_proto_goTypes = nil
	file_config_service_db_postgres_v1_database_proto_depIdxs = nil
}
