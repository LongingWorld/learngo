package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix  = "/d/Go/src/GitHub/learngo/"

type userError string  //类型userError实现了userError interface

func (e userError) Error() string  {
	return e.Message()
}

func (e userError) Message() string  {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//比较获取的URL路径和处理的路径是否一致
	if strings.Index(request.URL.Path,prefix) != 0{
		return  userError("Path mt start " + " with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file , err := os.Open(path)
	if err != nil{
		//panic(err)
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all,err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(all)
	return nil
}
