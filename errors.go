package i18naddress

import "fmt"

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
