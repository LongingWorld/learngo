package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf(
			"Unsupported operation: %s", op)
	}
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func applyMethod(op func(int, int) int, a, b int) int { //函数作为参数
	p := reflect.ValueOf(op).Pointer() //反射
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function  %s with args "+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func main() {
	if result, err := eval(3, 4, "*"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("1+2+...+5= ", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("pow(3,4) is: ", applyMethod(
		func(a, b int) int { //匿名函数
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
}
