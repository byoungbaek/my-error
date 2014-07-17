package my_error

import "fmt"

type MyError struct {
	code int32
	s string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.s)
}
