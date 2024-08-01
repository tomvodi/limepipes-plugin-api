package boundary

import (
	"fmt"
)

func (x Boundary) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *Boundary) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Boundary_value[s]
	if !ok {
		return fmt.Errorf("boundary value %s is not a valid", s)
	}
	*x = Boundary(up)
	return err
}
