// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: plugin/v1/fileformat/fileformat.proto

package fileformat

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

// Format of the file.
// This means the file format of the music data. Every file format may be stored
// with different file extensions. E.g. Bagpipe Music Writer Gold files are stored
// with the extension .bww or .bmw.
type Format int32

const (
	Format_Unknown Format = 0
	// The LimePipes music model data format
	Format_MUSIC_MODEL Format = 1
	// Bagpipe Music Writer Gold and Bagpipe Player format
	Format_BWW Format = 2
	// MusicXML format
	Format_MUSIC_XML Format = 3
	// ABC format
	Format_ABC Format = 4
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "Unknown",
		1: "MUSIC_MODEL",
		2: "BWW",
		3: "MUSIC_XML",
		4: "ABC",
	}
	Format_value = map[string]int32{
		"Unknown":     0,
		"MUSIC_MODEL": 1,
		"BWW":         2,
		"MUSIC_XML":   3,
		"ABC":         4,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_v1_fileformat_fileformat_proto_enumTypes[0].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_plugin_v1_fileformat_fileformat_proto_enumTypes[0]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_plugin_v1_fileformat_fileformat_proto_rawDescGZIP(), []int{0}
}

var File_plugin_v1_fileformat_fileformat_proto protoreflect.FileDescriptor

var file_plugin_v1_fileformat_fileformat_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2a, 0x47, 0x0a,
	0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x57, 0x57, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f, 0x58, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x42, 0x43, 0x10, 0x04, 0x42, 0xdb, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x42, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70,
	0x69, 0x70, 0x65, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x50, 0x56, 0x46, 0xaa, 0x02, 0x14, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0xca, 0x02, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x46, 0x69,
	0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0xe2, 0x02, 0x20, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x46, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x46, 0x69, 0x6c, 0x65, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_v1_fileformat_fileformat_proto_rawDescOnce sync.Once
	file_plugin_v1_fileformat_fileformat_proto_rawDescData = file_plugin_v1_fileformat_fileformat_proto_rawDesc
)

func file_plugin_v1_fileformat_fileformat_proto_rawDescGZIP() []byte {
	file_plugin_v1_fileformat_fileformat_proto_rawDescOnce.Do(func() {
		file_plugin_v1_fileformat_fileformat_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_v1_fileformat_fileformat_proto_rawDescData)
	})
	return file_plugin_v1_fileformat_fileformat_proto_rawDescData
}

var file_plugin_v1_fileformat_fileformat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plugin_v1_fileformat_fileformat_proto_goTypes = []any{
	(Format)(0), // 0: plugin.v1.fileformat.Format
}
var file_plugin_v1_fileformat_fileformat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_v1_fileformat_fileformat_proto_init() }
func file_plugin_v1_fileformat_fileformat_proto_init() {
	if File_plugin_v1_fileformat_fileformat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugin_v1_fileformat_fileformat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_v1_fileformat_fileformat_proto_goTypes,
		DependencyIndexes: file_plugin_v1_fileformat_fileformat_proto_depIdxs,
		EnumInfos:         file_plugin_v1_fileformat_fileformat_proto_enumTypes,
	}.Build()
	File_plugin_v1_fileformat_fileformat_proto = out.File
	file_plugin_v1_fileformat_fileformat_proto_rawDesc = nil
	file_plugin_v1_fileformat_fileformat_proto_goTypes = nil
	file_plugin_v1_fileformat_fileformat_proto_depIdxs = nil
}
