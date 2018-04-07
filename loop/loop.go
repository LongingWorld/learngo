package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	PrintFileContents(file)
}

func PrintFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

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
	printfile("abc.txt")

	s := `heng
	jin
	yu
	123
	love`
	PrintFileContents(strings.NewReader(s))
	//forever()
}
