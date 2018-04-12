package main

import (
	"GitHub/learngo/errhandling/filelistingsever/filelist"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

//define 函数类型
type AppHandler func(writer http.ResponseWriter, request *http.Request) error

//函数式编程
func ErrWrapper(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//recover
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occured "+"handling request : %s", err.Error()) //日志打印

			//Type assertion  判断错误类型是否是define userError
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/d/Go/src/GitHub/", ErrWrapper(filelist.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
