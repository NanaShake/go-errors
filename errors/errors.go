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
	return e.Detail
}

// 返回json格式字符串
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
