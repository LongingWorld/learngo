package main

import (
	"GitHub/learngo/retriever/mooc"
	"GitHub/learngo/retriever/realretriever"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mooc.Retriever{"This is Mooc Retriever!"}
	r = realretriever.Retriever{}
	fmt.Println(download(r))
}
