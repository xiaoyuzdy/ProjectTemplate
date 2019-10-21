package errors

import (
	"errors"
	"strconv"
)

const (
	Code1000 = iota + 1000
	Code1001
)

var mapping = map[int]string{
	Code1000: "账户已存在",
}

func GetError(code int) error {
	if msg, ok := mapping[code]; ok {
		return errors.New(msg)
	}
	return errors.New(strconv.FormatInt(int64(code), 10))
}

func GetMsg(code int) string {
	if msg, ok := mapping[code]; ok {
		return msg
	}
	return "unknown"
}
