// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/symbols/note.proto

package symbols

import (
	length "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	pitch "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	accidental "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	embellishment "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	movement "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/movement"
	tie "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tie"
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

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pitch         pitch.Pitch                  `protobuf:"varint,1,opt,name=pitch,proto3,enum=musicmodel.v1.pitch.Pitch" json:"pitch,omitempty"`
	Length        length.Length                `protobuf:"varint,2,opt,name=length,proto3,enum=musicmodel.v1.length.Length" json:"length,omitempty"`
	Dots          uint32                       `protobuf:"varint,3,opt,name=dots,proto3" json:"dots,omitempty"`
	Accidental    accidental.Accidental        `protobuf:"varint,4,opt,name=accidental,proto3,enum=musicmodel.v1.symbols.accidental.Accidental" json:"accidental,omitempty"`
	Fermata       bool                         `protobuf:"varint,5,opt,name=fermata,proto3" json:"fermata,omitempty"`
	Tie           tie.Tie                      `protobuf:"varint,6,opt,name=tie,proto3,enum=musicmodel.v1.symbols.tie.Tie" json:"tie,omitempty"`
	Embellishment *embellishment.Embellishment `protobuf:"bytes,7,opt,name=embellishment,proto3" json:"embellishment,omitempty"`
	Movement      *movement.Movement           `protobuf:"bytes,8,opt,name=movement,proto3" json:"movement,omitempty"`
	Comment       string                       `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicmodel_v1_symbols_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_musicmodel_v1_symbols_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_note_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetPitch() pitch.Pitch {
	if x != nil {
		return x.Pitch
	}
	return pitch.Pitch(0)
}

func (x *Note) GetLength() length.Length {
	if x != nil {
		return x.Length
	}
	return length.Length(0)
}

func (x *Note) GetDots() uint32 {
	if x != nil {
		return x.Dots
	}
	return 0
}

func (x *Note) GetAccidental() accidental.Accidental {
	if x != nil {
		return x.Accidental
	}
	return accidental.Accidental(0)
}

func (x *Note) GetFermata() bool {
	if x != nil {
		return x.Fermata
	}
	return false
}

func (x *Note) GetTie() tie.Tie {
	if x != nil {
		return x.Tie
	}
	return tie.Tie(0)
}

func (x *Note) GetEmbellishment() *embellishment.Embellishment {
	if x != nil {
		return x.Embellishment
	}
	return nil
}

func (x *Note) GetMovement() *movement.Movement {
	if x != nil {
		return x.Movement
	}
	return nil
}

func (x *Note) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_musicmodel_v1_symbols_note_proto protoreflect.FileDescriptor

var file_musicmodel_v1_symbols_note_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x1a, 0x1f, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x70,
	0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x2f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f,
	0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x74, 0x69, 0x65, 0x2f, 0x74, 0x69, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x6d, 0x62,
	0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c,
	0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d,
	0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x03,
	0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x69, 0x74, 0x63,
	0x68, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2e,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x64, 0x6f,
	0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x72, 0x6d, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x66, 0x65, 0x72, 0x6d, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x03, 0x74, 0x69,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e,
	0x74, 0x69, 0x65, 0x2e, 0x54, 0x69, 0x65, 0x52, 0x03, 0x74, 0x69, 0x65, 0x12, 0x58, 0x0a, 0x0d,
	0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x6d, 0x62, 0x65,
	0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c,
	0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69,
	0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xdc, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x73, 0x42, 0x09, 0x4e, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f,
	0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2d,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x73, 0xa2, 0x02, 0x03, 0x4d, 0x56, 0x53, 0xaa, 0x02, 0x15, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0xca,
	0x02, 0x15, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0xe2, 0x02, 0x21, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x4d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_symbols_note_proto_rawDescOnce sync.Once
	file_musicmodel_v1_symbols_note_proto_rawDescData = file_musicmodel_v1_symbols_note_proto_rawDesc
)

func file_musicmodel_v1_symbols_note_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_symbols_note_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_symbols_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_symbols_note_proto_rawDescData)
	})
	return file_musicmodel_v1_symbols_note_proto_rawDescData
}

var file_musicmodel_v1_symbols_note_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_musicmodel_v1_symbols_note_proto_goTypes = []any{
	(*Note)(nil),                        // 0: musicmodel.v1.symbols.Note
	(pitch.Pitch)(0),                    // 1: musicmodel.v1.pitch.Pitch
	(length.Length)(0),                  // 2: musicmodel.v1.length.Length
	(accidental.Accidental)(0),          // 3: musicmodel.v1.symbols.accidental.Accidental
	(tie.Tie)(0),                        // 4: musicmodel.v1.symbols.tie.Tie
	(*embellishment.Embellishment)(nil), // 5: musicmodel.v1.symbols.embellishment.Embellishment
	(*movement.Movement)(nil),           // 6: musicmodel.v1.symbols.movement.Movement
}
var file_musicmodel_v1_symbols_note_proto_depIdxs = []int32{
	1, // 0: musicmodel.v1.symbols.Note.pitch:type_name -> musicmodel.v1.pitch.Pitch
	2, // 1: musicmodel.v1.symbols.Note.length:type_name -> musicmodel.v1.length.Length
	3, // 2: musicmodel.v1.symbols.Note.accidental:type_name -> musicmodel.v1.symbols.accidental.Accidental
	4, // 3: musicmodel.v1.symbols.Note.tie:type_name -> musicmodel.v1.symbols.tie.Tie
	5, // 4: musicmodel.v1.symbols.Note.embellishment:type_name -> musicmodel.v1.symbols.embellishment.Embellishment
	6, // 5: musicmodel.v1.symbols.Note.movement:type_name -> musicmodel.v1.symbols.movement.Movement
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_symbols_note_proto_init() }
func file_musicmodel_v1_symbols_note_proto_init() {
	if File_musicmodel_v1_symbols_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_musicmodel_v1_symbols_note_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Note); i {
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
			RawDescriptor: file_musicmodel_v1_symbols_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_symbols_note_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_symbols_note_proto_depIdxs,
		MessageInfos:      file_musicmodel_v1_symbols_note_proto_msgTypes,
	}.Build()
	File_musicmodel_v1_symbols_note_proto = out.File
	file_musicmodel_v1_symbols_note_proto_rawDesc = nil
	file_musicmodel_v1_symbols_note_proto_goTypes = nil
	file_musicmodel_v1_symbols_note_proto_depIdxs = nil
}
