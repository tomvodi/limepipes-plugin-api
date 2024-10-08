// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/symbols/movement/movement.proto

package movement

import (
	pitch "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
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
	Type_NoMovement Type = 0
	Type_Cadence    Type = 1
	Type_Embari     Type = 2
	Type_Endari     Type = 3
	Type_Chedari    Type = 4
	Type_Hedari     Type = 5
	Type_Dili       Type = 6
	Type_Tra        Type = 7
	Type_Edre       Type = 8
	Type_Dare       Type = 9
	Type_CheCheRe   Type = 10
	Type_Grip       Type = 11
	Type_Deda       Type = 12
	Type_Enbain     Type = 13
	Type_Otro       Type = 14
	Type_Odro       Type = 15
	Type_Adeda      Type = 16
	Type_EchoBeat   Type = 17
	Type_Darodo     Type = 18
	Type_Hiharin    Type = 19
	Type_Rodin      Type = 20
	Type_Chelalho   Type = 21
	Type_Din        Type = 22
	Type_Lemluath   Type = 23
	Type_Taorluath  Type = 24
	Type_Crunluath  Type = 25
	Type_Tripling   Type = 26
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "NoMovement",
		1:  "Cadence",
		2:  "Embari",
		3:  "Endari",
		4:  "Chedari",
		5:  "Hedari",
		6:  "Dili",
		7:  "Tra",
		8:  "Edre",
		9:  "Dare",
		10: "CheCheRe",
		11: "Grip",
		12: "Deda",
		13: "Enbain",
		14: "Otro",
		15: "Odro",
		16: "Adeda",
		17: "EchoBeat",
		18: "Darodo",
		19: "Hiharin",
		20: "Rodin",
		21: "Chelalho",
		22: "Din",
		23: "Lemluath",
		24: "Taorluath",
		25: "Crunluath",
		26: "Tripling",
	}
	Type_value = map[string]int32{
		"NoMovement": 0,
		"Cadence":    1,
		"Embari":     2,
		"Endari":     3,
		"Chedari":    4,
		"Hedari":     5,
		"Dili":       6,
		"Tra":        7,
		"Edre":       8,
		"Dare":       9,
		"CheCheRe":   10,
		"Grip":       11,
		"Deda":       12,
		"Enbain":     13,
		"Otro":       14,
		"Odro":       15,
		"Adeda":      16,
		"EchoBeat":   17,
		"Darodo":     18,
		"Hiharin":    19,
		"Rodin":      20,
		"Chelalho":   21,
		"Din":        22,
		"Lemluath":   23,
		"Taorluath":  24,
		"Crunluath":  25,
		"Tripling":   26,
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
	return file_musicmodel_v1_symbols_movement_movement_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_symbols_movement_movement_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_movement_movement_proto_rawDescGZIP(), []int{0}
}

type Variant int32

const (
	Variant_NoVariant Variant = 0
	Variant_Half      Variant = 1
	Variant_G         Variant = 2
	Variant_Thumb     Variant = 3
	Variant_LongLowG  Variant = 4
)

// Enum value maps for Variant.
var (
	Variant_name = map[int32]string{
		0: "NoVariant",
		1: "Half",
		2: "G",
		3: "Thumb",
		4: "LongLowG",
	}
	Variant_value = map[string]int32{
		"NoVariant": 0,
		"Half":      1,
		"G":         2,
		"Thumb":     3,
		"LongLowG":  4,
	}
)

func (x Variant) Enum() *Variant {
	p := new(Variant)
	*p = x
	return p
}

func (x Variant) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Variant) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_symbols_movement_movement_proto_enumTypes[1].Descriptor()
}

func (Variant) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_symbols_movement_movement_proto_enumTypes[1]
}

func (x Variant) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Variant.Descriptor instead.
func (Variant) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_movement_movement_proto_rawDescGZIP(), []int{1}
}

type Movement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Type          `protobuf:"varint,1,opt,name=type,proto3,enum=musicmodel.v1.symbols.movement.Type" json:"type,omitempty"`
	Pitches []pitch.Pitch `protobuf:"varint,2,rep,packed,name=pitches,proto3,enum=musicmodel.v1.pitch.Pitch" json:"pitches,omitempty"`
	Fermata bool          `protobuf:"varint,3,opt,name=fermata,proto3" json:"fermata,omitempty"`
	Variant Variant       `protobuf:"varint,4,opt,name=variant,proto3,enum=musicmodel.v1.symbols.movement.Variant" json:"variant,omitempty"`
	// Bww symbols for piobaireachd are written with a pitch e.g. edred which
	// indicates that a B note should precede this note.
	// Nevertheless, in some cases, a note with another pitch is preceding the movement
	// so the pitch hint will make sure that no information gets lost in importing bww
	// files
	PitchHint           pitch.Pitch `protobuf:"varint,5,opt,name=pitch_hint,json=pitchHint,proto3,enum=musicmodel.v1.pitch.Pitch" json:"pitch_hint,omitempty"`
	AdditionalPitchHint pitch.Pitch `protobuf:"varint,6,opt,name=additional_pitch_hint,json=additionalPitchHint,proto3,enum=musicmodel.v1.pitch.Pitch" json:"additional_pitch_hint,omitempty"`
	Pitch               pitch.Pitch `protobuf:"varint,7,opt,name=pitch,proto3,enum=musicmodel.v1.pitch.Pitch" json:"pitch,omitempty"` // Pitch, if the movement has a distinct pitch like echo notes
	// Abbreviate is true, when the movement should show as its abbreviation symbol and not
	// as every grace note
	Abbreviate bool `protobuf:"varint,8,opt,name=abbreviate,proto3" json:"abbreviate,omitempty"`
	Breabach   bool `protobuf:"varint,9,opt,name=breabach,proto3" json:"breabach,omitempty"`
	AMach      bool `protobuf:"varint,10,opt,name=a_mach,json=aMach,proto3" json:"a_mach,omitempty"`
}

func (x *Movement) Reset() {
	*x = Movement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicmodel_v1_symbols_movement_movement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movement) ProtoMessage() {}

func (x *Movement) ProtoReflect() protoreflect.Message {
	mi := &file_musicmodel_v1_symbols_movement_movement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movement.ProtoReflect.Descriptor instead.
func (*Movement) Descriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_movement_movement_proto_rawDescGZIP(), []int{0}
}

func (x *Movement) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NoMovement
}

func (x *Movement) GetPitches() []pitch.Pitch {
	if x != nil {
		return x.Pitches
	}
	return nil
}

func (x *Movement) GetFermata() bool {
	if x != nil {
		return x.Fermata
	}
	return false
}

func (x *Movement) GetVariant() Variant {
	if x != nil {
		return x.Variant
	}
	return Variant_NoVariant
}

func (x *Movement) GetPitchHint() pitch.Pitch {
	if x != nil {
		return x.PitchHint
	}
	return pitch.Pitch(0)
}

func (x *Movement) GetAdditionalPitchHint() pitch.Pitch {
	if x != nil {
		return x.AdditionalPitchHint
	}
	return pitch.Pitch(0)
}

func (x *Movement) GetPitch() pitch.Pitch {
	if x != nil {
		return x.Pitch
	}
	return pitch.Pitch(0)
}

func (x *Movement) GetAbbreviate() bool {
	if x != nil {
		return x.Abbreviate
	}
	return false
}

func (x *Movement) GetBreabach() bool {
	if x != nil {
		return x.Breabach
	}
	return false
}

func (x *Movement) GetAMach() bool {
	if x != nil {
		return x.AMach
	}
	return false
}

var File_musicmodel_v1_symbols_movement_movement_proto protoreflect.FileDescriptor

var file_musicmodel_v1_symbols_movement_movement_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x1f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x69, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe7, 0x03, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50,
	0x69, 0x74, 0x63, 0x68, 0x52, 0x07, 0x70, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x65, 0x72, 0x6d, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x66, 0x65, 0x72, 0x6d, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x69,
	0x74, 0x63, 0x68, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x69, 0x74, 0x63, 0x68, 0x52, 0x09, 0x70, 0x69, 0x74, 0x63,
	0x68, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x69, 0x74, 0x63, 0x68,
	0x52, 0x13, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x69, 0x74, 0x63,
	0x68, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x69, 0x74, 0x63, 0x68,
	0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x62, 0x62, 0x72, 0x65,
	0x76, 0x69, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x62, 0x62,
	0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x65, 0x61, 0x62,
	0x61, 0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x72, 0x65, 0x61, 0x62,
	0x61, 0x63, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x4d, 0x61, 0x63, 0x68, 0x2a, 0xcb, 0x02, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x6f, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x6d, 0x62, 0x61, 0x72, 0x69, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x45, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x64,
	0x61, 0x72, 0x69, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x65, 0x64, 0x61, 0x72, 0x69, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x69, 0x6c, 0x69, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x72, 0x61, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x64, 0x72, 0x65, 0x10, 0x08, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x61, 0x72, 0x65, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x43,
	0x68, 0x65, 0x52, 0x65, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x72, 0x69, 0x70, 0x10, 0x0b,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x65, 0x64, 0x61, 0x10, 0x0c, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x6e,
	0x62, 0x61, 0x69, 0x6e, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x74, 0x72, 0x6f, 0x10, 0x0e,
	0x12, 0x08, 0x0a, 0x04, 0x4f, 0x64, 0x72, 0x6f, 0x10, 0x0f, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64,
	0x65, 0x64, 0x61, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x63, 0x68, 0x6f, 0x42, 0x65, 0x61,
	0x74, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x61, 0x72, 0x6f, 0x64, 0x6f, 0x10, 0x12, 0x12,
	0x0b, 0x0a, 0x07, 0x48, 0x69, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x10, 0x13, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x6f, 0x64, 0x69, 0x6e, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x6c, 0x61,
	0x6c, 0x68, 0x6f, 0x10, 0x15, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x69, 0x6e, 0x10, 0x16, 0x12, 0x0c,
	0x0a, 0x08, 0x4c, 0x65, 0x6d, 0x6c, 0x75, 0x61, 0x74, 0x68, 0x10, 0x17, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x61, 0x6f, 0x72, 0x6c, 0x75, 0x61, 0x74, 0x68, 0x10, 0x18, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x72, 0x75, 0x6e, 0x6c, 0x75, 0x61, 0x74, 0x68, 0x10, 0x19, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72,
	0x69, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x1a, 0x2a, 0x42, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x6f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x61, 0x6c, 0x66, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01,
	0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x6f, 0x77, 0x47, 0x10, 0x04, 0x42, 0x97, 0x02, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x0d, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70,
	0x65, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x04, 0x4d,
	0x56, 0x53, 0x4d, 0xaa, 0x02, 0x1e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x1e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5c, 0x4d, 0x6f, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x2a, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5c, 0x4d, 0x6f,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x21, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x3a, 0x3a, 0x4d, 0x6f,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_symbols_movement_movement_proto_rawDescOnce sync.Once
	file_musicmodel_v1_symbols_movement_movement_proto_rawDescData = file_musicmodel_v1_symbols_movement_movement_proto_rawDesc
)

func file_musicmodel_v1_symbols_movement_movement_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_symbols_movement_movement_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_symbols_movement_movement_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_symbols_movement_movement_proto_rawDescData)
	})
	return file_musicmodel_v1_symbols_movement_movement_proto_rawDescData
}

var file_musicmodel_v1_symbols_movement_movement_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_musicmodel_v1_symbols_movement_movement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_musicmodel_v1_symbols_movement_movement_proto_goTypes = []any{
	(Type)(0),        // 0: musicmodel.v1.symbols.movement.Type
	(Variant)(0),     // 1: musicmodel.v1.symbols.movement.Variant
	(*Movement)(nil), // 2: musicmodel.v1.symbols.movement.Movement
	(pitch.Pitch)(0), // 3: musicmodel.v1.pitch.Pitch
}
var file_musicmodel_v1_symbols_movement_movement_proto_depIdxs = []int32{
	0, // 0: musicmodel.v1.symbols.movement.Movement.type:type_name -> musicmodel.v1.symbols.movement.Type
	3, // 1: musicmodel.v1.symbols.movement.Movement.pitches:type_name -> musicmodel.v1.pitch.Pitch
	1, // 2: musicmodel.v1.symbols.movement.Movement.variant:type_name -> musicmodel.v1.symbols.movement.Variant
	3, // 3: musicmodel.v1.symbols.movement.Movement.pitch_hint:type_name -> musicmodel.v1.pitch.Pitch
	3, // 4: musicmodel.v1.symbols.movement.Movement.additional_pitch_hint:type_name -> musicmodel.v1.pitch.Pitch
	3, // 5: musicmodel.v1.symbols.movement.Movement.pitch:type_name -> musicmodel.v1.pitch.Pitch
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_symbols_movement_movement_proto_init() }
func file_musicmodel_v1_symbols_movement_movement_proto_init() {
	if File_musicmodel_v1_symbols_movement_movement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_musicmodel_v1_symbols_movement_movement_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Movement); i {
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
			RawDescriptor: file_musicmodel_v1_symbols_movement_movement_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_symbols_movement_movement_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_symbols_movement_movement_proto_depIdxs,
		EnumInfos:         file_musicmodel_v1_symbols_movement_movement_proto_enumTypes,
		MessageInfos:      file_musicmodel_v1_symbols_movement_movement_proto_msgTypes,
	}.Build()
	File_musicmodel_v1_symbols_movement_movement_proto = out.File
	file_musicmodel_v1_symbols_movement_movement_proto_rawDesc = nil
	file_musicmodel_v1_symbols_movement_movement_proto_goTypes = nil
	file_musicmodel_v1_symbols_movement_movement_proto_depIdxs = nil
}
