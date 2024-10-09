package symbols

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/movement"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tie"
	"testing"
)

func TestSymbol_IsNote(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsNote()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsNote()).To(BeTrue())
}

func TestSymbol_IsValidNote(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsValidNote()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsValidNote()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsValidNote()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsValidNote()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsValidNote()).To(BeTrue())
}

func TestSymbol_IsOnlyEmbellishment(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyEmbellishment()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyEmbellishment()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyEmbellishment()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Embellishment = &embellishment.Embellishment{}
	g.Expect(s.IsOnlyEmbellishment()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyEmbellishment()).To(BeFalse())
}

func TestSymbol_IsOnlyMovement(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyMovement()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyMovement()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyMovement()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Movement = &movement.Movement{}
	g.Expect(s.IsOnlyMovement()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyMovement()).To(BeFalse())
}

func TestSymbol_IsOnlyAccidental(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyAccidental()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyAccidental()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyAccidental()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Accidental = accidental.Accidental_Sharp
	g.Expect(s.IsOnlyAccidental()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyAccidental()).To(BeFalse())
}

func TestSymbol_IsOnlyDots(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyDots()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyDots()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyDots()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Dots = 1
	g.Expect(s.IsOnlyDots()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyDots()).To(BeFalse())
}

func TestSymbol_IsOnlyFermata(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyFermata()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyFermata()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyFermata()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Fermata = true
	g.Expect(s.IsOnlyFermata()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyFermata()).To(BeFalse())
}

func TestSymbol_IsOnlyTieStart(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyTieStart()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyTieStart()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyTieStart()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Tie = tie.Tie_Start
	g.Expect(s.IsOnlyTieStart()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyTieStart()).To(BeFalse())
}

func TestSymbol_IsOnlyTieEnd(t *testing.T) {
	g := NewGomegaWithT(t)
	s := &Symbol{}
	g.Expect(s.IsOnlyTieEnd()).To(BeFalse())
	s.Note = &Note{}
	g.Expect(s.IsOnlyTieEnd()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_D
	g.Expect(s.IsOnlyTieEnd()).To(BeFalse())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Tie = tie.Tie_End
	g.Expect(s.IsOnlyTieEnd()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_D
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsOnlyTieEnd()).To(BeFalse())
}
