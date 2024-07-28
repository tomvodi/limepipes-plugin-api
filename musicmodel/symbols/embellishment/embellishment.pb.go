// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: symbols/embellishment/embellishment.proto

package embellishment

import (
	pitch "github.com/tomvodi/limepipes-music-model/musicmodel/pitch"
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

type EmbellishmentType int32

const (
	EmbellishmentType_NoEmbellishment   EmbellishmentType = 0
	EmbellishmentType_SingleGrace       EmbellishmentType = 1
	EmbellishmentType_Doubling          EmbellishmentType = 2
	EmbellishmentType_Strike            EmbellishmentType = 3
	EmbellishmentType_Grip              EmbellishmentType = 4
	EmbellishmentType_Taorluath         EmbellishmentType = 5
	EmbellishmentType_Bubbly            EmbellishmentType = 6
	EmbellishmentType_GraceBirl         EmbellishmentType = 7
	EmbellishmentType_ABirl             EmbellishmentType = 8
	EmbellishmentType_Birl              EmbellishmentType = 9
	EmbellishmentType_ThrowD            EmbellishmentType = 10
	EmbellishmentType_Pele              EmbellishmentType = 11
	EmbellishmentType_DoubleStrike      EmbellishmentType = 12
	EmbellishmentType_TripleStrike      EmbellishmentType = 13
	EmbellishmentType_GTripleStrike     EmbellishmentType = 14
	EmbellishmentType_ThumbTripleStrike EmbellishmentType = 15
	EmbellishmentType_HalfTripleStrike  EmbellishmentType = 16
	EmbellishmentType_DoubleGrace       EmbellishmentType = 17
)

// Enum value maps for EmbellishmentType.
var (
	EmbellishmentType_name = map[int32]string{
		0:  "NoEmbellishment",
		1:  "SingleGrace",
		2:  "Doubling",
		3:  "Strike",
		4:  "Grip",
		5:  "Taorluath",
		6:  "Bubbly",
		7:  "GraceBirl",
		8:  "ABirl",
		9:  "Birl",
		10: "ThrowD",
		11: "Pele",
		12: "DoubleStrike",
		13: "TripleStrike",
		14: "GTripleStrike",
		15: "ThumbTripleStrike",
		16: "HalfTripleStrike",
		17: "DoubleGrace",
	}
	EmbellishmentType_value = map[string]int32{
		"NoEmbellishment":   0,
		"SingleGrace":       1,
		"Doubling":          2,
		"Strike":            3,
		"Grip":              4,
		"Taorluath":         5,
		"Bubbly":            6,
		"GraceBirl":         7,
		"ABirl":             8,
		"Birl":              9,
		"ThrowD":            10,
		"Pele":              11,
		"DoubleStrike":      12,
		"TripleStrike":      13,
		"GTripleStrike":     14,
		"ThumbTripleStrike": 15,
		"HalfTripleStrike":  16,
		"DoubleGrace":       17,
	}
)

func (x EmbellishmentType) Enum() *EmbellishmentType {
	p := new(EmbellishmentType)
	*p = x
	return p
}

func (x EmbellishmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmbellishmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_symbols_embellishment_embellishment_proto_enumTypes[0].Descriptor()
}

func (EmbellishmentType) Type() protoreflect.EnumType {
	return &file_symbols_embellishment_embellishment_proto_enumTypes[0]
}

func (x EmbellishmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmbellishmentType.Descriptor instead.
func (EmbellishmentType) EnumDescriptor() ([]byte, []int) {
	return file_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{0}
}

type EmbellishmentVariant int32

const (
	EmbellishmentVariant_NoVariant EmbellishmentVariant = 0
	EmbellishmentVariant_G         EmbellishmentVariant = 1
	EmbellishmentVariant_Half      EmbellishmentVariant = 2
	EmbellishmentVariant_Thumb     EmbellishmentVariant = 3
)

// Enum value maps for EmbellishmentVariant.
var (
	EmbellishmentVariant_name = map[int32]string{
		0: "NoVariant",
		1: "G",
		2: "Half",
		3: "Thumb",
	}
	EmbellishmentVariant_value = map[string]int32{
		"NoVariant": 0,
		"G":         1,
		"Half":      2,
		"Thumb":     3,
	}
)

func (x EmbellishmentVariant) Enum() *EmbellishmentVariant {
	p := new(EmbellishmentVariant)
	*p = x
	return p
}

func (x EmbellishmentVariant) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmbellishmentVariant) Descriptor() protoreflect.EnumDescriptor {
	return file_symbols_embellishment_embellishment_proto_enumTypes[1].Descriptor()
}

func (EmbellishmentVariant) Type() protoreflect.EnumType {
	return &file_symbols_embellishment_embellishment_proto_enumTypes[1]
}

func (x EmbellishmentVariant) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmbellishmentVariant.Descriptor instead.
func (EmbellishmentVariant) EnumDescriptor() ([]byte, []int) {
	return file_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{1}
}

type EmbellishmentWeight int32

const (
	EmbellishmentWeight_NoWeight EmbellishmentWeight = 0
	EmbellishmentWeight_Light    EmbellishmentWeight = 1
	EmbellishmentWeight_Heavy    EmbellishmentWeight = 2
)

// Enum value maps for EmbellishmentWeight.
var (
	EmbellishmentWeight_name = map[int32]string{
		0: "NoWeight",
		1: "Light",
		2: "Heavy",
	}
	EmbellishmentWeight_value = map[string]int32{
		"NoWeight": 0,
		"Light":    1,
		"Heavy":    2,
	}
)

func (x EmbellishmentWeight) Enum() *EmbellishmentWeight {
	p := new(EmbellishmentWeight)
	*p = x
	return p
}

func (x EmbellishmentWeight) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmbellishmentWeight) Descriptor() protoreflect.EnumDescriptor {
	return file_symbols_embellishment_embellishment_proto_enumTypes[2].Descriptor()
}

func (EmbellishmentWeight) Type() protoreflect.EnumType {
	return &file_symbols_embellishment_embellishment_proto_enumTypes[2]
}

func (x EmbellishmentWeight) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmbellishmentWeight.Descriptor instead.
func (EmbellishmentWeight) EnumDescriptor() ([]byte, []int) {
	return file_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{2}
}

type Embellishment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type EmbellishmentType `protobuf:"varint,1,opt,name=type,proto3,enum=EmbellishmentType" json:"type,omitempty"`
	// Pitch is set for embellishments that have a pitch regardless of the melody note
	// following it (e.g. single g-grace)
	// Other embellishments have their pitch defined by the melody note following it
	// (e.g. doubling) because a d-doubling can only precede a d-melody note.
	// In these cases, Pitch is set to NoPitch
	Pitch   pitch.Pitch          `protobuf:"varint,2,opt,name=pitch,proto3,enum=Pitch" json:"pitch,omitempty"`
	Variant EmbellishmentVariant `protobuf:"varint,3,opt,name=variant,proto3,enum=EmbellishmentVariant" json:"variant,omitempty"`
	Weight  EmbellishmentWeight  `protobuf:"varint,4,opt,name=weight,proto3,enum=EmbellishmentWeight" json:"weight,omitempty"`
}

func (x *Embellishment) Reset() {
	*x = Embellishment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_symbols_embellishment_embellishment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embellishment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embellishment) ProtoMessage() {}

func (x *Embellishment) ProtoReflect() protoreflect.Message {
	mi := &file_symbols_embellishment_embellishment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embellishment.ProtoReflect.Descriptor instead.
func (*Embellishment) Descriptor() ([]byte, []int) {
	return file_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{0}
}

func (x *Embellishment) GetType() EmbellishmentType {
	if x != nil {
		return x.Type
	}
	return EmbellishmentType_NoEmbellishment
}

func (x *Embellishment) GetPitch() pitch.Pitch {
	if x != nil {
		return x.Pitch
	}
	return pitch.Pitch(0)
}

func (x *Embellishment) GetVariant() EmbellishmentVariant {
	if x != nil {
		return x.Variant
	}
	return EmbellishmentVariant_NoVariant
}

func (x *Embellishment) GetWeight() EmbellishmentWeight {
	if x != nil {
		return x.Weight
	}
	return EmbellishmentWeight_NoWeight
}

var File_symbols_embellishment_embellishment_proto protoreflect.FileDescriptor

var file_symbols_embellishment_embellishment_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c,
	0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x69, 0x74,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x45, 0x6d, 0x62,
	0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x6c,
	0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x06, 0x2e, 0x50, 0x69, 0x74, 0x63, 0x68, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68,
	0x12, 0x2f, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2a,
	0xa7, 0x02, 0x0a, 0x11, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x6f, 0x45, 0x6d, 0x62, 0x65, 0x6c,
	0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x63, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x72,
	0x69, 0x6b, 0x65, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x72, 0x69, 0x70, 0x10, 0x04, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x61, 0x6f, 0x72, 0x6c, 0x75, 0x61, 0x74, 0x68, 0x10, 0x05, 0x12, 0x0a,
	0x0a, 0x06, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x79, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x72,
	0x61, 0x63, 0x65, 0x42, 0x69, 0x72, 0x6c, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x42, 0x69,
	0x72, 0x6c, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x69, 0x72, 0x6c, 0x10, 0x09, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x68, 0x72, 0x6f, 0x77, 0x44, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x65,
	0x6c, 0x65, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6b, 0x65, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x54, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65,
	0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x61, 0x6c, 0x66, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x47, 0x72, 0x61, 0x63, 0x65, 0x10, 0x11, 0x2a, 0x41, 0x0a, 0x14, 0x45, 0x6d, 0x62,
	0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x6f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x10, 0x00,
	0x12, 0x05, 0x0a, 0x01, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x61, 0x6c, 0x66, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x10, 0x03, 0x2a, 0x39, 0x0a, 0x13,
	0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x48, 0x65, 0x61, 0x76, 0x79, 0x10, 0x02, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69,
	0x6d, 0x65, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_symbols_embellishment_embellishment_proto_rawDescOnce sync.Once
	file_symbols_embellishment_embellishment_proto_rawDescData = file_symbols_embellishment_embellishment_proto_rawDesc
)

func file_symbols_embellishment_embellishment_proto_rawDescGZIP() []byte {
	file_symbols_embellishment_embellishment_proto_rawDescOnce.Do(func() {
		file_symbols_embellishment_embellishment_proto_rawDescData = protoimpl.X.CompressGZIP(file_symbols_embellishment_embellishment_proto_rawDescData)
	})
	return file_symbols_embellishment_embellishment_proto_rawDescData
}

var file_symbols_embellishment_embellishment_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_symbols_embellishment_embellishment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_symbols_embellishment_embellishment_proto_goTypes = []interface{}{
	(EmbellishmentType)(0),    // 0: EmbellishmentType
	(EmbellishmentVariant)(0), // 1: EmbellishmentVariant
	(EmbellishmentWeight)(0),  // 2: EmbellishmentWeight
	(*Embellishment)(nil),     // 3: Embellishment
	(pitch.Pitch)(0),          // 4: Pitch
}
var file_symbols_embellishment_embellishment_proto_depIdxs = []int32{
	0, // 0: Embellishment.type:type_name -> EmbellishmentType
	4, // 1: Embellishment.pitch:type_name -> Pitch
	1, // 2: Embellishment.variant:type_name -> EmbellishmentVariant
	2, // 3: Embellishment.weight:type_name -> EmbellishmentWeight
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_symbols_embellishment_embellishment_proto_init() }
func file_symbols_embellishment_embellishment_proto_init() {
	if File_symbols_embellishment_embellishment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_symbols_embellishment_embellishment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embellishment); i {
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
			RawDescriptor: file_symbols_embellishment_embellishment_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_symbols_embellishment_embellishment_proto_goTypes,
		DependencyIndexes: file_symbols_embellishment_embellishment_proto_depIdxs,
		EnumInfos:         file_symbols_embellishment_embellishment_proto_enumTypes,
		MessageInfos:      file_symbols_embellishment_embellishment_proto_msgTypes,
	}.Build()
	File_symbols_embellishment_embellishment_proto = out.File
	file_symbols_embellishment_embellishment_proto_rawDesc = nil
	file_symbols_embellishment_embellishment_proto_goTypes = nil
	file_symbols_embellishment_embellishment_proto_depIdxs = nil
}