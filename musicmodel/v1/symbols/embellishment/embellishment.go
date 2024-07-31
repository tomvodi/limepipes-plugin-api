package embellishment

import (
	"fmt"
)

func (x EmbellishmentType) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *EmbellishmentType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := EmbellishmentType_value[s]
	if !ok {
		return fmt.Errorf("EmbellishmentType value %s is not valid", s)
	}
	*x = EmbellishmentType(up)
	return err
}
func (x EmbellishmentVariant) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *EmbellishmentVariant) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := EmbellishmentVariant_value[s]
	if !ok {
		return fmt.Errorf("EmbellishmentVariant value %s is not valid", s)
	}
	*x = EmbellishmentVariant(up)
	return err
}

func (x EmbellishmentWeight) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *EmbellishmentWeight) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := EmbellishmentWeight_value[s]
	if !ok {
		return fmt.Errorf("EmbellishmentWeight value %s is not valid", s)
	}
	*x = EmbellishmentWeight(up)
	return err
}
