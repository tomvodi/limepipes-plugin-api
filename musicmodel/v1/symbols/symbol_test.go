package symbols

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
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
	g.Expect(s.IsValidNote()).To(BeTrue())
	s.Note.Pitch = pitch.Pitch_NoPitch
	s.Note.Length = length.Length_Quarter
	g.Expect(s.IsValidNote()).To(BeTrue())
}
