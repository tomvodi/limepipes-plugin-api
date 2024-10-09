package measure

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"slices"
)

func (x *Measure) AddMessage(msg *ParserMessage) {
	if x.ParserMessages == nil {
		x.ParserMessages = []*ParserMessage{}
	}

	x.ParserMessages = append(x.ParserMessages, msg)
}

func (x *Measure) LastSymbol() *symbols.Symbol {
	if x == nil {
		return nil
	}

	for _, s := range slices.Backward(x.Symbols) {
		if s != nil {
			return s
		}
	}

	return nil
}
