package length

import "fmt"

func (x Length) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Length) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Length_value[s]
	if !ok {
		return fmt.Errorf("length %s is not a valid Length", s)
	}
	*x = Length(up)
	return err
}
