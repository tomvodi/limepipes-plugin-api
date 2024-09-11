package embellishment

import (
	"fmt"
)

func (x Type) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Type) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Type_value[s]
	if !ok {
		return fmt.Errorf("type value %s is not valid", s)
	}
	*x = Type(up)
	return err
}

func (x Variant) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Variant) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Variant_value[s]
	if !ok {
		return fmt.Errorf("variant value %s is not valid", s)
	}
	*x = Variant(up)
	return err
}
func (x Weight) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Weight) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Weight_value[s]
	if !ok {
		return fmt.Errorf("weight value %s is not valid", s)
	}
	*x = Weight(up)
	return err
}
