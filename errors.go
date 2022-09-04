package i18naddress

import (
	"encoding/json"
	"fmt"
)

// InvalidCodeErr indicate given country code is invalid
type InvalidCodeErr struct {
	value interface{}
	msg   string
}

func (i *InvalidCodeErr) Error() string {
	return fmt.Sprintf(i.msg, i.value)
}

func newInvalidCodeErr(value interface{}) *InvalidCodeErr {
	return &InvalidCodeErr{
		msg:   "%s is not a valid code",
		value: value,
	}
}

// ErrorMap is map
type ErrorMap map[string]string

func (e ErrorMap) Error() string {
	if len(e) == 0 {
		return "{}"
	}
	b, _ := json.Marshal(e)
	return string(b)
}
