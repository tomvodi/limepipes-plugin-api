package measure

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestTimeSignature_DisplayString(t *testing.T) {
	g := NewGomegaWithT(t)
	ts := &TimeSignature{
		Beats:    4,
		BeatType: 8,
	}
	g.Expect(ts.DisplayString()).To(Equal("4/8"))
}
