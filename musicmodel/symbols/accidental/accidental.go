package accidental

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Accidental) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Accidental) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Accidental_value[value.Value]
	if !ok {
		return fmt.Errorf("accidental %s is not a valid pitch", value.Value)
	}

	*i = Accidental(val)
	return nil
}
