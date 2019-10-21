package errors

import (
	"fmt"
	"go-web/utils"
	"net/http"
)

type HttpError struct {
	HttpState int
	Code      int
	ErrMsg    interface{}
	Stack     string
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("errMsg: %s  stack: %s", h.ErrMsg, h.Stack)
}

func ErrPrompt(businessCode int) *HttpError {
	prompt := newHTTPError(http.StatusOK, GetError(businessCode))
	prompt.Code = businessCode
	return prompt
}

//资源不存在
func ErrNotFound(err ...error) *HttpError {
	return newHTTPError(http.StatusNotFound, err)
}

//方法不被允许
func ErrMethodNotAllowed(err ...error) *HttpError {
	return newHTTPError(http.StatusMethodNotAllowed, err)
}

//API无权限
func ErrForbidden(err ...error) *HttpError {
	return newHTTPError(http.StatusForbidden, err)
}

//身份认证错误
func ErrUnauthorized(err ...error) *HttpError {
	return newHTTPError(http.StatusUnauthorized, err)
}

//内部错误
func ErrInternalServerError(err ...error) *HttpError {
	return newHTTPError(http.StatusInternalServerError, err)
}

//数据没有通过validation规则，需要求改投递数据
func ErrValidation(err ...error) *HttpError {
	return newHTTPError(http.StatusUnprocessableEntity, err)
}

func newHTTPError(httpState int, errMsg ...interface{}) *HttpError {
	he := &HttpError{
		HttpState: httpState,
		ErrMsg:    http.StatusText(httpState),
		Stack:     utils.GetStackInfo(),
	}
	if len(errMsg) > 0 {
		he.ErrMsg = errMsg
	}
	return he
}
