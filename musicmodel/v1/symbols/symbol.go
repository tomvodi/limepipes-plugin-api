package symbols

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tie"
)

func (x *Symbol) IsNote() bool {
	return x.Note != nil
}

func (x *Symbol) IsValidNote() bool {
	return x.Note.IsValid()
}

func (x *Symbol) IsOnlyEmbellishment() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Embellishment != nil
}

func (x *Symbol) IsOnlyMovement() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Movement != nil
}

func (x *Symbol) IsOnlyAccidental() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() &&
		x.Note.Accidental != accidental.Accidental_NoAccidental
}

func (x *Symbol) IsOnlyDots() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Dots != 0
}

func (x *Symbol) IsOnlyFermata() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Fermata
}

func (x *Symbol) IsOnlyTieStart() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Tie == tie.Tie_Start
}

func (x *Symbol) IsOnlyTieEnd() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Tie == tie.Tie_End
}
