package main

import (
	"fmt"
	"regexp"
)

const context = `my email is shadowzzj@126.com
email1 is snake@qq.com
email2 is ig@sohu.com.cn
email3 is we@apple.org`

func main() {
	reg := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := reg.FindAllStringSubmatch(context, -1)
	//fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
