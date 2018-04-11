package main

import (
	"fmt"
	"sync"
)

func DoWorker(id int, c chan int, wg *sync.WaitGroup) {
	for {
		if n, ok := <-c; !ok { //从worker.in goroutine中取消息
			break
		} else {
			fmt.Printf("OneWorker %d received %c\n", id, n)
			wg.Done()
			//done <- true //往worker.done goroutine中发消息
			//go func() { //start a goroutine in DoWoiker goroutine
			//	worker.done <- true
			//}()
		}
	}
}

func CreateWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		wg: wg,
	}

	go DoWorker(id, w.in, wg) //start a goroutine
	return w
}

type Worker struct {
	in chan int
	wg *sync.WaitGroup
}

func ChanDemo() {
	//var c chan int  //c == nil
	var wg sync.WaitGroup
	var workers [10]Worker

	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}

	/*go func() {  //closure
		for {
			n := <-c
			fmt.Println(n)
		}
	}()*/

	//wg.Add(20)

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i //往worker.in goroutine中发消息
		//<-worker.done  get message
	}

	//for _, wk := range workers {
	//	<-wk.done
	//}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
		//<-worker.done get message
	}
	wg.Wait() //wait until all job done
	//for _, wk := range workers {
	//	<-wk.done
	//}
	//for _, wk := range workers {
	//	<-wk.done
	//	<-wk.done
	//}

	//time.Sleep(time.Microsecond)

}

/*func BufferedChannel() {
	c := make(chan int, 3)
	go Worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Microsecond)
}

func ChannelClose() {
	c := make(chan int, 3)
	go Worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Microsecond)
}*/

func main() {
	fmt.Println("Channel as first-class citizen")
	ChanDemo()
	/*fmt.Println("Buffered channel")
	BufferedChannel()
	fmt.Println("Channel close and range")
	ChannelClose()*/
}
