package accidental

import (
	"fmt"
)

func (x Accidental) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Accidental) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Accidental_value[s]
	if !ok {
		return fmt.Errorf("accidental value %s is not valid", s)
	}
	*x = Accidental(up)
	return err
}
