package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generator() chan int { //MessageProducer
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func Worker(id int, c chan int) { //MessageConsumer
	/*for n := range c {
		fmt.Printf("OneWorker %d received %d\n", id, n)
	}*/
	for {
		if n, ok := <-c; !ok {
			break
		} else {
			time.Sleep(time.Second) //one second get one message
			fmt.Printf("OneWorker %d received %d\n", id, n)
		}
	}
}

func CreateWorkder(id int) chan<- int {
	c := make(chan int)
	/*go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()*/
	go Worker(id, c)
	return c
}

func main() {
	var c1, c2 = Generator(), Generator() //MessageProducer
	var worker = CreateWorkder(0)         //MessageConsumer

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(2 * time.Second) //timer
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker //MessageConsumer
			activeValue = values[0]
		}
		select {
		case n := <-c1: //MessageProducer
			values = append(values, n)
		case n := <-c2: //MessageProducer
			values = append(values, n)
		case activeWorker <- activeValue: //MessageConsumer
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Timeout")
		case <-tick:
			fmt.Printf("values number is %d\n", len(values))
		case <-tm:
			fmt.Println("Byebye select Channel")
			return

			/*default:   //非阻塞的select
			fmt.Println("No data received")*/
		}
	}
}
