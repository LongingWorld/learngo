package main

import (
	"fmt"
	"strings" //StringOperation Package
	"unicode/utf8"
)

func main() {
	str := "Yeah我爱HJY!"
	fmt.Println(len(str)) //get byte numbers
	fmt.Println(str)

	for i, ch := range []byte(str) {
		fmt.Printf("(%d,%X)", i, ch)
	}
	fmt.Println()

	for i, ch := range str { //ch is rune
		fmt.Printf("(%d,%X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(str)) //get character numbers

	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(str) {
		fmt.Printf("(%d,%c)", i, ch)
	}
	fmt.Println()

	if i := strings.Compare("abc", "def"); i >= 0 {
		fmt.Println("before is big than after!")
	} else {
		fmt.Println("after is big than before!")
	}

}
