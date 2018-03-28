package main

import (
	"fmt"
	"math/cmplx"
	"math"
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

func EulerTheorem(){
	e := 3 + 4i
	fmt.Println(cmplx.Abs(e))
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1)
	fmt.Printf("%.3f\n",cmplx.Exp(1i*math.Pi)+1)
}

func Triangle(){
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func enums()  {
	const(
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 <<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b,kb,mb,gb,tb,pb)

	fmt.Println(cpp,javascript,python,golang)

}

func main() {
	variableZeroValue()
	fmt.Println(aa, ss, kk)
	EulerTheorem()
	Triangle()
	enums()
}
