package pitch

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Pitch) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Pitch) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Pitch_value[value.Value]
	if !ok {
		return fmt.Errorf("pitch %s is not a valid pitch", value.Value)
	}

	*i = Pitch(val)
	return nil
}
