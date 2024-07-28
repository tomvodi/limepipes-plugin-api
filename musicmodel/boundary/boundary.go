package boundary

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Boundary) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Boundary) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Boundary_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Boundary(val)
	return nil
}
