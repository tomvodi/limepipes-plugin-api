// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/tune/tune.proto

package tune

import (
	measure "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
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

type Tune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string             `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type       string             `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Composer   string             `protobuf:"bytes,3,opt,name=composer,proto3" json:"composer,omitempty"`
	Arranger   string             `protobuf:"bytes,4,opt,name=arranger,proto3" json:"arranger,omitempty"`
	Footer     []string           `protobuf:"bytes,5,rep,name=footer,proto3" json:"footer,omitempty"`
	Comments   []string           `protobuf:"bytes,6,rep,name=comments,proto3" json:"comments,omitempty"`
	InLineText []string           `protobuf:"bytes,7,rep,name=in_line_text,json=inLineText,proto3" json:"in_line_text,omitempty"`
	Tempo      uint32             `protobuf:"varint,8,opt,name=tempo,proto3" json:"tempo,omitempty"`
	Measures   []*measure.Measure `protobuf:"bytes,9,rep,name=measures,proto3" json:"measures,omitempty"`
}

func (x *Tune) Reset() {
	*x = Tune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicmodel_v1_tune_tune_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tune) ProtoMessage() {}

func (x *Tune) ProtoReflect() protoreflect.Message {
	mi := &file_musicmodel_v1_tune_tune_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tune.ProtoReflect.Descriptor instead.
func (*Tune) Descriptor() ([]byte, []int) {
	return file_musicmodel_v1_tune_tune_proto_rawDescGZIP(), []int{0}
}

func (x *Tune) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Tune) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Tune) GetComposer() string {
	if x != nil {
		return x.Composer
	}
	return ""
}

func (x *Tune) GetArranger() string {
	if x != nil {
		return x.Arranger
	}
	return ""
}

func (x *Tune) GetFooter() []string {
	if x != nil {
		return x.Footer
	}
	return nil
}

func (x *Tune) GetComments() []string {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Tune) GetInLineText() []string {
	if x != nil {
		return x.InLineText
	}
	return nil
}

func (x *Tune) GetTempo() uint32 {
	if x != nil {
		return x.Tempo
	}
	return 0
}

func (x *Tune) GetMeasures() []*measure.Measure {
	if x != nil {
		return x.Measures
	}
	return nil
}

var File_musicmodel_v1_tune_tune_proto protoreflect.FileDescriptor

var file_musicmodel_v1_tune_tune_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x75, 0x6e, 0x65, 0x2f, 0x74, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x75, 0x6e, 0x65, 0x1a, 0x23, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x04, 0x54, 0x75, 0x6e,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x6e, 0x5f, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x12,
	0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x73, 0x42, 0xca, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x75, 0x6e, 0x65, 0x42, 0x09, 0x54, 0x75, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65,
	0x73, 0x2d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x65,
	0xa2, 0x02, 0x03, 0x4d, 0x56, 0x54, 0xaa, 0x02, 0x12, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x65, 0xca, 0x02, 0x12, 0x4d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x54, 0x75, 0x6e, 0x65,
	0xe2, 0x02, 0x1e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31,
	0x5c, 0x54, 0x75, 0x6e, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x14, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x54, 0x75, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_tune_tune_proto_rawDescOnce sync.Once
	file_musicmodel_v1_tune_tune_proto_rawDescData = file_musicmodel_v1_tune_tune_proto_rawDesc
)

func file_musicmodel_v1_tune_tune_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_tune_tune_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_tune_tune_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_tune_tune_proto_rawDescData)
	})
	return file_musicmodel_v1_tune_tune_proto_rawDescData
}

var file_musicmodel_v1_tune_tune_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_musicmodel_v1_tune_tune_proto_goTypes = []any{
	(*Tune)(nil),            // 0: musicmodel.v1.tune.Tune
	(*measure.Measure)(nil), // 1: musicmodel.v1.measure.Measure
}
var file_musicmodel_v1_tune_tune_proto_depIdxs = []int32{
	1, // 0: musicmodel.v1.tune.Tune.measures:type_name -> musicmodel.v1.measure.Measure
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_tune_tune_proto_init() }
func file_musicmodel_v1_tune_tune_proto_init() {
	if File_musicmodel_v1_tune_tune_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_musicmodel_v1_tune_tune_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Tune); i {
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
			RawDescriptor: file_musicmodel_v1_tune_tune_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_tune_tune_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_tune_tune_proto_depIdxs,
		MessageInfos:      file_musicmodel_v1_tune_tune_proto_msgTypes,
	}.Build()
	File_musicmodel_v1_tune_tune_proto = out.File
	file_musicmodel_v1_tune_tune_proto_rawDesc = nil
	file_musicmodel_v1_tune_tune_proto_goTypes = nil
	file_musicmodel_v1_tune_tune_proto_depIdxs = nil
}
