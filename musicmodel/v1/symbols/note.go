package symbols

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
)

func (x *Note) IsValid() bool {
	if x.Pitch == pitch.Pitch_NoPitch && x.Length == length.Length_NoLength {
		return false
	}
	return true
}

// IsIncomplete returns true, if the note has no pitch and length but already other
// properties like an embellishment or an accidental
// this is the case, when the bww symbols which modify the note are preceding the note in
// bww code
func (x *Note) IsIncomplete() bool {
	return !x.HasPitchOrLength()
}

func (x *Note) HasPitchOrLength() bool {
	return x.Pitch != pitch.Pitch_NoPitch || x.Length != length.Length_NoLength
}

func (x *Note) HasPitchAndLength() bool {
	return x.Pitch != pitch.Pitch_NoPitch && x.Length != length.Length_NoLength
}

func (x *Note) IsOnlyAccidental() bool {
	if x.Pitch == pitch.Pitch_NoPitch && x.Length == length.Length_NoLength {
		if x.Accidental != accidental.Accidental_NoAccidental {
			return true
		}
	}
	return false
}
