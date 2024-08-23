package file_type

import (
	"database/sql/driver"
	"fmt"
)

func (x Type) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

func (x *Type) UnmarshalYAML(unmarshal func(interface{}) error) error {
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

func (x Type) Value() (driver.Value, error) {
	return x.String(), nil
}

func (x *Type) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of Type: %[1]T(%[1]v)", value)
	}

	val, ok := Type_value[str]
	if !ok {
		return fmt.Errorf("type value %s is not in value map", str)
	}

	*x = Type(val)
	return nil
}
