package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func TryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occured!")
	fmt.Println(4)
}

func Fibonacci() func()int  {
	a,b := 0,1
	return func() int {
		a,b = b, a+b
		return a
	}
}

func WriteFile(filename string)  {
	//file,err := os.Create(filename)
	file,err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//define err yourself
	err = errors.New("This is a custom error!")
	if err != nil{
		//panic(err)
		//fmt.Println("Error:",err.Error()) //打印错误信息
		//fmt.Println("Error:",err)

		if pathError,ok := err.(*os.PathError);!ok{
			panic(err)
		}else{
			fmt.Printf("%s , %s , %s\n",pathError.Op,pathError.Path,pathError.Err)
		}

		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)//先写到内存里
	defer writer.Flush()  //从内存中写入文件

	f := Fibonacci()
	for i := 0; i < 20 ; i++  {
		fmt.Fprintln(writer,f())
	}
}

func main() {
	TryDefer()
	WriteFile("fibonacci.txt")
}
