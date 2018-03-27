package main

import (
	"fmt"
	//"math"
	//"math/cmplx"
)

var (
	aa = 8
	ss = "love"
	kk = true
)

func variableZeroValue() {
	a, b, c, s := 2, 3, 4, "abc"

	fmt.Println(a, b, c, s)
}

func main() {
	variableZeroValue()
	fmt.Println(aa, ss, kk)
}
