package symbols

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
)

// IsValid returns true if the note has a pitch and length.
func (x *Note) IsValid() bool {
	if x == nil {
		return false
	}

	if x.hasPitchAndLength() {
		return true
	}

	return false
}

func (x *Note) hasPitchAndLength() bool {
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
