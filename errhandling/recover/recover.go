package main

import (
	"fmt"
)

func TryRecover()  {
	defer func() {
		r := recover()
		if err , ok := r.(error); ok {
			fmt.Println("Error occured:",err)
		}else {
			panic(fmt.Sprintf("I don't know what to do: %v",r))
		}
	}()   //recover 只能在defer中调用，调用匿名函数func(){}()
	//panic(errors.New("This is an error!"))
	//panic(123)
	b := 0
	a := 5/b
	fmt.Println(a)
	//panic(123)
}

func main() {
	TryRecover()
}
