package movement

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

func (i Variant) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Variant) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Variant_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Variant(val)
	return nil
}
