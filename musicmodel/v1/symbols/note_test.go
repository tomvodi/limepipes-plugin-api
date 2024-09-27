package symbols

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"testing"
)

func TestNote_HasPitchAndLength(t *testing.T) {
	g := NewGomegaWithT(t)
	n := &Note{}
	g.Expect(n.hasPitchAndLength()).To(BeFalse())
	n.Pitch = pitch.Pitch_D
	g.Expect(n.hasPitchAndLength()).To(BeFalse())
	n.Length = length.Length_Sixteenth
	g.Expect(n.hasPitchAndLength()).To(BeTrue())
}

func TestNote_IsOnlyAccidental(t *testing.T) {
	g := NewGomegaWithT(t)
	n := &Note{}
	g.Expect(n.IsOnlyAccidental()).To(BeFalse())

	n.Accidental = accidental.Accidental_Sharp
	g.Expect(n.IsOnlyAccidental()).To(BeTrue())

	n.Pitch = pitch.Pitch_D
	g.Expect(n.IsOnlyAccidental()).To(BeFalse())
}
