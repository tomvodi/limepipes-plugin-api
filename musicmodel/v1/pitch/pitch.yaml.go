package pitch

import "fmt"

func (x Pitch) MarshalYAML() (any, error) {
	return x.String(), nil
}

func (x *Pitch) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	up, ok := Pitch_value[s]
	if !ok {
		return fmt.Errorf("pitch value %s is not valid", s)
	}
	*x = Pitch(up)
	return err
}
