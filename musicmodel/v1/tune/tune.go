package tune

import (
	"github.com/tomvodi/limepipes-music-model/musicmodel/v1/measure"
)

func (x *Tune) ImportMessages() []*measure.ImportMessage {
	var messages []*measure.ImportMessage

	for _, measure := range x.Measures {
		messages = append(messages, measure.ImportMessages...)
	}

	return messages
}

func (x *Tune) FirstTimeSignature() *measure.TimeSignature {
	for _, measure := range x.Measures {
		if measure.Time != nil {
			return measure.Time
		}
	}

	return nil
}
