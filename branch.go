package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string  {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score :%s", score))
	case score < 60:
		g = "F"
	case score >= 60 && score < 70:
		g = "D"
	case score >= 70 && score < 80:
		g = "C"
	case score >= 80 && score < 90:
		g = "B"
	case score >=90 && score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	/*contents,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}*/
	if contents,err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("The contents of the file is %s\n",contents)
	}

	/*fmt.Println(grade(59))
	fmt.Println(grade(66))
	fmt.Println(grade(77))
	fmt.Println(grade(88))
	fmt.Println(grade(99))*/
	fmt.Println(
		grade(59),
		grade(66),
		grade(77),
		grade(88),
		grade(99),
	)
	//grade(101)
	//grade(-3)
}
