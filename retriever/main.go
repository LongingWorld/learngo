package main

import (
	"GitHub/learngo/retriever/mooc"
	"GitHub/learngo/retriever/realretriever"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("http://www.baidu.com")
}

func Inspect(r Retriever)  {
	fmt.Println("Inspecting:",r )
	fmt.Printf(" > %T   %v \n",r,r)
	fmt.Print("> Type swith:")
	switch v := r.(type) {
	case *mooc.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *realretriever.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	r = &mooc.Retriever{"This is Fake Mooc Retriever!"}
	//fmt.Printf("%T   %v \n",r,r)
	Inspect(r)
	r = &realretriever.Retriever{
		UserAgent:"Golang/1.10",
		TimeOut:time.Minute,
	}
	Inspect(r)

	//Type assertion
	if moocRetriever, ok := r.(mooc.Retriever);ok{
		fmt.Println(moocRetriever.Contents)
	}else {
		fmt.Println("Not a mooc retriever!")
	}

	//fmt.Printf("%T   %v \n",r,r)
	//fmt.Println(download(r))
}
