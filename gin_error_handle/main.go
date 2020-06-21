package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIException : APIException
type APIException struct {
	Code    int    `json:"-"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Request string `json:"request"`
}

const (
	SERVER_ERROR    = 1000 // 系统错误
	NOT_FOUND       = 1001 // 401错误
	UNKNOWN_ERROR   = 1002 // 未知错误
	PARAMETER_ERROR = 1003 // 参数错误
	AUTH_ERROR      = 1004 // 错误
)

// 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

// 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, UNKNOWN_ERROR, message)
}

// 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, PARAMETER_ERROR, message)
}

func (e *APIException) Error() string {
	return e.ErrMsg
}

func newAPIException(code int, errcode int, errmsg string) *APIException {
	return &APIException{
		Code:    code,
		ErrCode: errcode,
		ErrMsg:  errmsg,
	}
}

// HandlerFunc : HandlerFunc
type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			var apiException *APIException
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError()
			}
			apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}

func user(c *gin.Context) error {
	// TODO 逻辑判断
	return ParameterError("userId传参有误")
}

func main() {
	r := gin.Default()

	r.GET("/ping", wrapper(user))
	r.Run(":18080")
}
