package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//Closure闭包
func Fibonacci() IntGen  {
	a , b := 0,1
	return func() int {
		a , b = b , a+b
		return a
	}
}

//函数实现接口
type IntGen func() int  //定义函数类型

func (g IntGen) Read(p []byte) (n int, err error) {  //函数实现Reader接口的Read方法
	//panic("implement me")
	next := g()  //取下一个元素
	if next >10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)  //转换成字符串
	//TODO:incorrect if  p is too small  ???
	return strings.NewReader(s).Read(p)  //读取
}

func PrintFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := Fibonacci()
	PrintFileContents(f)
}
