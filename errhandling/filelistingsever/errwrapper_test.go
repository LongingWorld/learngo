package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func ErrPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type TestuserError string //类型userError实现了userError interface

func (e TestuserError) Error() string {
	return e.Message()
}

func (e TestuserError) Message() string {
	return string(e)
}

func ErrUserError(writer http.ResponseWriter, request *http.Request) error {
	return TestuserError("user error")
}

func ErrUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func ErrNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func ErrNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func NoError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprint(writer, "No Error")
	return nil
}

//构造表格驱动测试
var tests = []struct {
	h       AppHandler
	code    int
	message string
}{
	{ErrPanic, 500, "Internal Server Error"},
	{ErrNoPermission, 403, "Forbidden"},
	{ErrNotFound, 404, "Not Found"},
	{ErrUserError, 400, "user error"},
	{ErrUnknown, 500, "Internal Server Error"},
	{NoError, 200, "No Error"},
}

func VerifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	//body := string(b)

	log.Printf("tt.message : %s \n body : %s \n ", expectedMsg, body)

	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expected (%d , %s);"+"got (%d , %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}

//手动创建HTTPSERVER测试
func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h       AppHandler
		code    int
		message string
	}{
		{ErrPanic, 500, "Internal Server Error"},
		{ErrNoPermission, 403, "Forbidden"},
		{ErrNotFound, 404, "Not Found"},
		{ErrUserError, 400, "user error"},
		{ErrUnknown, 500, "Internal Server Error"},
		{NoError, 200, "No Error"},
	}

	for _, tt := range tests {
		f := ErrWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)
		VerifyResponse(response.Result(), tt.code, tt.message, t)

	}
}

//起HTTP服务测试
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := ErrWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))

		resp, _ := http.Get(server.URL)

		VerifyResponse(resp, tt.code, tt.message, t)

		/*b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b),"\n")
		//body := string(b)

		log.Printf("tt.message : %s \n body : %s \n ",tt.message,body)

		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("expected (%d , %s);" + "got (%d , %s)",tt.code,tt.message,resp.StatusCode,body)
		}*/
	}
}
