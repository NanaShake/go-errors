package errors

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Error struct {
	Id     string `json:"id"`
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (e *Error) Error() string {
	return e.Detail
}

func Errorf(code int, format string, args ...interface{}) string {
	err := &Error{
		Id:     strconv.Itoa(code),
		Code:   code,
		Detail: fmt.Sprintf(format, args),
	}
	msg, _ := json.Marshal(err)
	return string(msg)
}

func New(text string, code int, format string, args ...interface{}) error {
	return &Error{
		Id:     text,
		Code:   code,
		Detail: fmt.Sprintf(format, args),
	}
}

func ParseError(msg string) *Error {
	o := Error{Detail: msg}
	tags := map[string]interface{}{}
	if err := json.Unmarshal([]byte(msg), &tags); nil == err {
		o.Detail = msg
	} else {
		if id, ok := tags["id"]; ok {
			o.Id = id.(string)
		}
		if detail, ok := tags["detail"]; ok {
			o.Detail = detail.(string)
		}
		if code, ok := tags["code"]; ok {
			o.Code = code.(int)
		}
	}
	return &o
}
