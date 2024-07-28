package length

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Length) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Length) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Length_value[value.Value]
	if !ok {
		return fmt.Errorf("pitch %s is not a valid pitch", value.Value)
	}

	*i = Length(val)
	return nil
}
