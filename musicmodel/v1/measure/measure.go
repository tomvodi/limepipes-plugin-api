package measure

import (
	"reflect"
)

func (x *Measure) IsNil() bool {
	return reflect.DeepEqual(x, &Measure{})
}

func (x *Measure) AddMessage(msg *ImportMessage) {
	if x.ImportMessages == nil {
		x.ImportMessages = []*ImportMessage{}
	}

	x.ImportMessages = append(x.ImportMessages, msg)
}
