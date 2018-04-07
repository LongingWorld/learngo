package main

import (
	"GitHub/learngo/queue"
	"fmt"
)

func main() {
	q :=queue.Queue{1}

	q.Push(2)
	q.Push(5)
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
