// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/pitch/pitch.proto

package pitch

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

type Pitch int32

const (
	Pitch_NoPitch Pitch = 0
	Pitch_LowG    Pitch = 1
	Pitch_LowA    Pitch = 2
	Pitch_B       Pitch = 3
	Pitch_C       Pitch = 4
	Pitch_D       Pitch = 5
	Pitch_E       Pitch = 6
	Pitch_F       Pitch = 7
	Pitch_HighG   Pitch = 8
	Pitch_HighA   Pitch = 9
)

// Enum value maps for Pitch.
var (
	Pitch_name = map[int32]string{
		0: "NoPitch",
		1: "LowG",
		2: "LowA",
		3: "B",
		4: "C",
		5: "D",
		6: "E",
		7: "F",
		8: "HighG",
		9: "HighA",
	}
	Pitch_value = map[string]int32{
		"NoPitch": 0,
		"LowG":    1,
		"LowA":    2,
		"B":       3,
		"C":       4,
		"D":       5,
		"E":       6,
		"F":       7,
		"HighG":   8,
		"HighA":   9,
	}
)

func (x Pitch) Enum() *Pitch {
	p := new(Pitch)
	*p = x
	return p
}

func (x Pitch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pitch) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_pitch_pitch_proto_enumTypes[0].Descriptor()
}

func (Pitch) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_pitch_pitch_proto_enumTypes[0]
}

func (x Pitch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pitch.Descriptor instead.
func (Pitch) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_pitch_pitch_proto_rawDescGZIP(), []int{0}
}

var File_musicmodel_v1_pitch_pitch_proto protoreflect.FileDescriptor

var file_musicmodel_v1_pitch_pitch_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2a, 0x61, 0x0a, 0x05, 0x50, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x50, 0x69, 0x74, 0x63, 0x68, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x4c, 0x6f, 0x77, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x6f, 0x77, 0x41, 0x10, 0x02,
	0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x04, 0x12, 0x05,
	0x0a, 0x01, 0x44, 0x10, 0x05, 0x12, 0x05, 0x0a, 0x01, 0x45, 0x10, 0x06, 0x12, 0x05, 0x0a, 0x01,
	0x46, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x69, 0x67, 0x68, 0x47, 0x10, 0x08, 0x12, 0x09,
	0x0a, 0x05, 0x48, 0x69, 0x67, 0x68, 0x41, 0x10, 0x09, 0x42, 0xd0, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x69, 0x74, 0x63, 0x68, 0x42, 0x0a, 0x50, 0x69, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65,
	0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x74, 0x63, 0x68,
	0xa2, 0x02, 0x03, 0x4d, 0x56, 0x50, 0xaa, 0x02, 0x13, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x69, 0x74, 0x63, 0x68, 0xca, 0x02, 0x13, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x69, 0x74,
	0x63, 0x68, 0xe2, 0x02, 0x1f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c,
	0x56, 0x31, 0x5c, 0x50, 0x69, 0x74, 0x63, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x69, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_pitch_pitch_proto_rawDescOnce sync.Once
	file_musicmodel_v1_pitch_pitch_proto_rawDescData = file_musicmodel_v1_pitch_pitch_proto_rawDesc
)

func file_musicmodel_v1_pitch_pitch_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_pitch_pitch_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_pitch_pitch_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_pitch_pitch_proto_rawDescData)
	})
	return file_musicmodel_v1_pitch_pitch_proto_rawDescData
}

var file_musicmodel_v1_pitch_pitch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_musicmodel_v1_pitch_pitch_proto_goTypes = []any{
	(Pitch)(0), // 0: musicmodel.v1.pitch.Pitch
}
var file_musicmodel_v1_pitch_pitch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_pitch_pitch_proto_init() }
func file_musicmodel_v1_pitch_pitch_proto_init() {
	if File_musicmodel_v1_pitch_pitch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_musicmodel_v1_pitch_pitch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_pitch_pitch_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_pitch_pitch_proto_depIdxs,
		EnumInfos:         file_musicmodel_v1_pitch_pitch_proto_enumTypes,
	}.Build()
	File_musicmodel_v1_pitch_pitch_proto = out.File
	file_musicmodel_v1_pitch_pitch_proto_rawDesc = nil
	file_musicmodel_v1_pitch_pitch_proto_goTypes = nil
	file_musicmodel_v1_pitch_pitch_proto_depIdxs = nil
}