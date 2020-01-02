package micro

import "fmt"

const ERRNO_OK = 200

type Error struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func NewError(errno int, errmsg string) *Error {
	return &Error{Errno: errno, Errmsg: errmsg}
}

func (E *Error) Error() string {
	return fmt.Sprintf("[%d] %s", E.Errno, E.Errmsg)
}
