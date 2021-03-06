package datatype

import (
	"strings"
)

type FlexibleBool struct {
	Value  bool
	String string
}

func (f *FlexibleBool) UnmarshalJSON(b []byte) error {
	if f.String = strings.Trim(string(b), `"`); f.String == "" {
		f.String = "false"
	}
	f.Value = f.String == "1" || strings.EqualFold(f.String, "true")
	return nil
}
