// generated by jsonenums -type Type; DO NOT EDIT

package tenses

import (
	"encoding/json"
	"fmt"
)

var (
	_TypeNameToValue = map[string]Type{
		"None":    None,
		"Past":    Past,
		"Present": Present,
		"Future":  Future,
	}

	_TypeValueToName = map[Type]string{
		None:    "None",
		Past:    "Past",
		Present: "Present",
		Future:  "Future",
	}
)

func init() {
	var v Type
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TypeNameToValue = map[string]Type{
			interface{}(None).(fmt.Stringer).String():    None,
			interface{}(Past).(fmt.Stringer).String():    Past,
			interface{}(Present).(fmt.Stringer).String(): Present,
			interface{}(Future).(fmt.Stringer).String():  Future,
		}
	}
}

// MarshalJSON is generated so Type satisfies json.Marshaler.
func (r Type) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Type: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Type satisfies json.Unmarshaler.
func (r *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Type should be a string, got %s", data)
	}
	v, ok := _TypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Type %q", s)
	}
	*r = v
	return nil
}
