package done

import (
	"fmt"
	"time"
)

func Worker(id int, c chan int) {
	for {
		if n, ok := <-c; !ok {
			break
		} else {
			fmt.Printf("OneWorker %d received %d\n", id, n)
		}
	}
}

func CreateWorker(id int) chan<- int {
	c := make(chan int)

	go Worker(id, c)
	return c
}

func ChanDemo() {
	//var c chan int  //c == nil
	var channels [10]chan<- int //array channel

	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}

	/*go func() {  //closure
		for {
			n := <-c
			fmt.Println(n)
		}
	}()*/
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Microsecond)

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
