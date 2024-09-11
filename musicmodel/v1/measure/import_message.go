package measure

import (
	"fmt"
)

func (x Severity) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Severity) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Severity_value[s]
	if !ok {
		return fmt.Errorf("severity value %s is not valid", s)
	}
	*x = Severity(up)
	return err
}

func (x Fix) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Fix) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Fix_value[s]
	if !ok {
		return fmt.Errorf("fix value %s is not valid", s)
	}
	*x = Fix(up)
	return err
}
