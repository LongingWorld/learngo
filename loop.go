package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	a := 2
	s := ""
	var c = 0
	for ; n > 0; n /= 2 {
		c = n % a
		s = strconv.Itoa(c) + s
	}
	return s
}

func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func deadLoop() {
	for {
		fmt.Println("HJY")
	}
}

func main() {
	fmt.Println(convertToBin(13))
	printfile("C:\\Users\\HJY\\Desktop\\abc\\abc121.txt")
	//forever()
}
