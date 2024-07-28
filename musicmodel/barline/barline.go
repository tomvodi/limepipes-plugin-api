package barline

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Type) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Type) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Type_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Type(val)
	return nil
}

func (i Time) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Time) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Time_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Time(val)
	return nil
}
