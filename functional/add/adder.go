package main

import (
	"fmt"
)

//闭包 Closure
func Add() func(int) int  {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/*  正统函数式编程
~~不可变性：不能有状态，只有常量和函数
~~函数只能有一个参数
*/
type iAdd func(int) (int,iAdd)

func Add2(base int) iAdd  {
	return func(v int) (int, iAdd) {
		return base+v,Add2(base+v)
	}
}


func main() {
	fmt.Println("Golang functional:")
	a := Add()
	//fmt.Printf("func.sum is : %d \n",a(1))
	for  i:= 0;i < 10 ; i++  {
		fmt.Printf("0 + ... + %d =  %d \n",i,a(i))
	}
	fmt.Println()
	fmt.Println("Traditional functional:")
	b := Add2(0)
	for m:= 0 ;m < 10 ; m++  {
		var s int
		s , b= b(m)
		fmt.Printf("0 + ... + %d =  %d \n",m,s)
	}
}
