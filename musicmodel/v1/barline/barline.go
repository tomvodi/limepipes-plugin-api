package barline

import (
	"fmt"
)

func (x Type) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *Type) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Type_value[s]
	if !ok {
		return fmt.Errorf("type value %s is not valid", s)
	}
	*x = Type(up)
	return err
}

func (x Time) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *Time) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Time_value[s]
	if !ok {
		return fmt.Errorf("time value %s is not valid", s)
	}
	*x = Time(up)
	return err
}
