package tie

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Tie) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Tie) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Tie_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Tie(val)
	return nil
}
