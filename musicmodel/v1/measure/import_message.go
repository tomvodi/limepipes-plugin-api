package measure

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i Severity) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Severity) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Severity_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Severity(val)
	return nil
}

func (i Fix) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Fix) UnmarshalYAML(value *yaml.Node) error {
	val, ok := Fix_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = Fix(val)
	return nil
}
