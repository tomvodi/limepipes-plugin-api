package embellishment

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (i EmbellishmentType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *EmbellishmentType) UnmarshalYAML(value *yaml.Node) error {
	val, ok := EmbellishmentType_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = EmbellishmentType(val)
	return nil
}

func (i EmbellishmentVariant) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *EmbellishmentVariant) UnmarshalYAML(value *yaml.Node) error {
	val, ok := EmbellishmentVariant_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = EmbellishmentVariant(val)
	return nil
}

func (i EmbellishmentWeight) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *EmbellishmentWeight) UnmarshalYAML(value *yaml.Node) error {
	val, ok := EmbellishmentWeight_value[value.Value]
	if !ok {
		return fmt.Errorf("%s is not a valid value", value.Value)
	}

	*i = EmbellishmentWeight(val)
	return nil
}
