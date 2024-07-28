// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: symbols/timeline/timeline.proto

package timeline

import (
	boundary "github.com/tomvodi/limepipes-music-model/musicmodel/boundary"
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

type Type int32

const (
	Type_NoType        Type = 0
	Type_First         Type = 1
	Type_Singling      Type = 2
	Type_Second        Type = 3
	Type_Doubling      Type = 4
	Type_SecondOf2     Type = 5
	Type_SecondOf3     Type = 6
	Type_SecondOf4     Type = 7
	Type_SecondOf2And4 Type = 8
	Type_SecondOf5     Type = 9
	Type_SecondOf6     Type = 10
	Type_SecondOf7     Type = 11
	Type_SecondOf8     Type = 12
	Type_Bis           Type = 13
	Type_Intro         Type = 14
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "NoType",
		1:  "First",
		2:  "Singling",
		3:  "Second",
		4:  "Doubling",
		5:  "SecondOf2",
		6:  "SecondOf3",
		7:  "SecondOf4",
		8:  "SecondOf2And4",
		9:  "SecondOf5",
		10: "SecondOf6",
		11: "SecondOf7",
		12: "SecondOf8",
		13: "Bis",
		14: "Intro",
	}
	Type_value = map[string]int32{
		"NoType":        0,
		"First":         1,
		"Singling":      2,
		"Second":        3,
		"Doubling":      4,
		"SecondOf2":     5,
		"SecondOf3":     6,
		"SecondOf4":     7,
		"SecondOf2And4": 8,
		"SecondOf5":     9,
		"SecondOf6":     10,
		"SecondOf7":     11,
		"SecondOf8":     12,
		"Bis":           13,
		"Intro":         14,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_symbols_timeline_timeline_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_symbols_timeline_timeline_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_symbols_timeline_timeline_proto_rawDescGZIP(), []int{0}
}

type TimeLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         Type              `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	BoundaryType boundary.Boundary `protobuf:"varint,2,opt,name=boundaryType,proto3,enum=Boundary" json:"boundaryType,omitempty"`
}

func (x *TimeLine) Reset() {
	*x = TimeLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_symbols_timeline_timeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeLine) ProtoMessage() {}

func (x *TimeLine) ProtoReflect() protoreflect.Message {
	mi := &file_symbols_timeline_timeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeLine.ProtoReflect.Descriptor instead.
func (*TimeLine) Descriptor() ([]byte, []int) {
	return file_symbols_timeline_timeline_proto_rawDescGZIP(), []int{0}
}

func (x *TimeLine) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NoType
}

func (x *TimeLine) GetBoundaryType() boundary.Boundary {
	if x != nil {
		return x.BoundaryType
	}
	return boundary.Boundary(0)
}

var File_symbols_timeline_timeline_proto protoreflect.FileDescriptor

var file_symbols_timeline_timeline_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x52, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2a, 0xd5, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x32, 0x10, 0x05, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x33, 0x10, 0x06, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x34, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x32, 0x41, 0x6e, 0x64, 0x34, 0x10, 0x08, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x35, 0x10, 0x09, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x36, 0x10, 0x0a, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x37, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x66, 0x38, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x69, 0x73, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x10, 0x0e, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f,
	0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2d,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_symbols_timeline_timeline_proto_rawDescOnce sync.Once
	file_symbols_timeline_timeline_proto_rawDescData = file_symbols_timeline_timeline_proto_rawDesc
)

func file_symbols_timeline_timeline_proto_rawDescGZIP() []byte {
	file_symbols_timeline_timeline_proto_rawDescOnce.Do(func() {
		file_symbols_timeline_timeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_symbols_timeline_timeline_proto_rawDescData)
	})
	return file_symbols_timeline_timeline_proto_rawDescData
}

var file_symbols_timeline_timeline_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_symbols_timeline_timeline_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_symbols_timeline_timeline_proto_goTypes = []interface{}{
	(Type)(0),              // 0: Type
	(*TimeLine)(nil),       // 1: TimeLine
	(boundary.Boundary)(0), // 2: Boundary
}
var file_symbols_timeline_timeline_proto_depIdxs = []int32{
	0, // 0: TimeLine.type:type_name -> Type
	2, // 1: TimeLine.boundaryType:type_name -> Boundary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_symbols_timeline_timeline_proto_init() }
func file_symbols_timeline_timeline_proto_init() {
	if File_symbols_timeline_timeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_symbols_timeline_timeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeLine); i {
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
			RawDescriptor: file_symbols_timeline_timeline_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_symbols_timeline_timeline_proto_goTypes,
		DependencyIndexes: file_symbols_timeline_timeline_proto_depIdxs,
		EnumInfos:         file_symbols_timeline_timeline_proto_enumTypes,
		MessageInfos:      file_symbols_timeline_timeline_proto_msgTypes,
	}.Build()
	File_symbols_timeline_timeline_proto = out.File
	file_symbols_timeline_timeline_proto_rawDesc = nil
	file_symbols_timeline_timeline_proto_goTypes = nil
	file_symbols_timeline_timeline_proto_depIdxs = nil
}