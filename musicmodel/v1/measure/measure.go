package measure

import (
	"reflect"
)

func (x *Measure) IsNil() bool {
	return reflect.DeepEqual(x, &Measure{})
}

func (x *Measure) AddMessage(msg *ParserMessage) {
	if x.ParserMessages == nil {
		x.ParserMessages = []*ParserMessage{}
	}

	x.ParserMessages = append(x.ParserMessages, msg)
}
