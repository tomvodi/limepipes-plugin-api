package measure

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestMeasure_AddMessage(t *testing.T) {
	g := NewGomegaWithT(t)
	m := &Measure{}
	g.Expect(m.ParserMessages).To(BeNil())
	im := &ParserMessage{
		Symbol:   "sym",
		Severity: 0,
		Text:     "",
		Fix:      0,
	}
	m.AddMessage(im)
	g.Expect(m.ParserMessages).To(HaveLen(1))
	g.Expect(m.ParserMessages[0]).To(Equal(im))
}
