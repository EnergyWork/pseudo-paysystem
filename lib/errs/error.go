package errs

import "fmt"

// Error representation of error struct
//
//swagger:model Error
type Error struct {
	Code string `json:"Code"`
	Msg  string `json:"Message"`
}

func New() *Error {
	return &Error{}
}

func (e *Error) SetCode(code string) *Error {
	e.Code = code
	return e
}

func (e *Error) SetMsg(msg string, args ...interface{}) *Error {
	e.Msg = fmt.Sprintf(msg, args...)
	return e
}

func (e *Error) String() string {
	if e == nil {
		return ""
	}
	str := fmt.Sprintf("error: %s", e.Code)
	if e.Msg != "" {
		str += fmt.Sprintf("- %s", e.Msg)
	}
	return str
}

func (e *Error) Error() error {
	return fmt.Errorf(e.String())
}
