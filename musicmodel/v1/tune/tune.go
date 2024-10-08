package tune

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
)

func (x *Tune) ImportMessages() []*measure.ParserMessage {
	var messages []*measure.ParserMessage

	for _, m := range x.Measures {
		messages = append(messages, m.ParserMessages...)
	}

	return messages
}

func (x *Tune) FirstTimeSignature() *measure.TimeSignature {
	for _, m := range x.Measures {
		if m.Time != nil {
			return m.Time
		}
	}

	return nil
}
