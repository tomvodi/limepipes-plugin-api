package tie

import (
	"fmt"
)

func (x Tie) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *Tie) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Tie_value[s]
	if !ok {
		return fmt.Errorf("tie value %s is not valid", s)
	}
	*x = Tie(up)
	return err
}
