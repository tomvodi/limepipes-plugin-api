package measure

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
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

func TestMeasure_LastSymbol(t *testing.T) {
	g := NewGomegaWithT(t)
	m := &Measure{}
	g.Expect(m.LastSymbol()).To(BeNil())
	s := &symbols.Symbol{}
	m.Symbols = append(m.Symbols, s)
	g.Expect(m.LastSymbol()).To(Equal(s))
	s2 := &symbols.Symbol{}
	m.Symbols = append(m.Symbols, s2)
	g.Expect(m.LastSymbol()).To(Equal(s2))
}
