package errors

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Error struct {
	Id     string `json:"id"`     // 错误标识
	Code   int    `json:"code"`   // 错误标识(数字型)
	Detail string `json:"detail"` // 错误描述
}

// 返回错误描述
func (e *Error) Error() string {
	txt, _ := json.Marshal(e)
	return string(txt)
}

// 返回json格式字符串
func Errorf(code int, format string, args ...interface{}) string {
	err := &Error{
		Id:     strconv.Itoa(code),
		Code:   code,
		Detail: fmt.Sprintf(format, args),
	}
	return err.Error()
}

func New(id string, code int, format string, args ...interface{}) error {
	return &Error{
		Id:     id,
		Code:   code,
		Detail: fmt.Sprintf(format, args),
	}
}

func Wrap(id string, code int, err error) error {
	if ra, ok := err.(*Error); ok {
		ra.Id = id
		return err
	}
	return &Error{
		Id:     id,
		Code:   code,
		Detail: err.Error(),
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
