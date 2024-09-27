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

func (x *Symbol) CanBeMergedWith(other *Symbol) bool {
	if x.isOnlyEmbellishment() && other.IsValidNote() {
		return true
	}

	if x.isOnlyMovement() && other.IsValidNote() {
		return true
	}

	if x.IsValidNote() && other.IsOnlyDots() {
		return true
	}

	if x.IsValidNote() && other.IsOnlyFermata() {
		return true
	}

	// tie start is merged with the next note
	if x.IsOnlyTieStart() && other.IsValidNote() {
		return true
	}

	// tie end is merged with the previous note
	if x.IsValidNote() && other.IsOnlyTieEnd() {
		return true
	}

	if x.isOnlyAccidental() && other.IsValidNote() {
		return true
	}

	return false
}

func (x *Symbol) MergeWith(other *Symbol) {
	if x.isOnlyEmbellishment() && other.IsValidNote() {
		other.Note.Embellishment = x.Note.Embellishment
		x.Note = other.Note
	}

	if x.isOnlyMovement() && other.IsValidNote() {
		other.Note.Movement = x.Note.Movement
		x.Note = other.Note
	}

	if x.isOnlyAccidental() && other.IsValidNote() {
		other.Note.Accidental = x.Note.Accidental
		x.Note = other.Note
	}

	if x.IsValidNote() && other.IsOnlyDots() {
		x.Note.Dots = other.Note.Dots
	}

	// tie start is merged with the next note
	if x.IsOnlyTieStart() && other.IsValidNote() {
		other.Note.Tie = x.Note.Tie
		x.Note = other.Note
	}

	// tie end is merged with the previous note
	if x.IsValidNote() && other.IsOnlyTieEnd() {
		x.Note.Tie = other.Note.Tie
	}

	if x.IsValidNote() && other.IsOnlyFermata() {
		x.Note.Fermata = other.Note.Fermata
	}
}

func (x *Symbol) isOnlyEmbellishment() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Embellishment != nil
}

func (x *Symbol) isOnlyMovement() bool {
	if x.Note == nil {
		return false
	}

	return !x.Note.IsValid() && x.Note.Movement != nil
}

func (x *Symbol) isOnlyAccidental() bool {
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
