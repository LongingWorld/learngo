package main

import (
	"fmt"
	"runtime"
	"time"
)

//协程 Coroutine
// 轻量级“线程”，非抢占式多任务处理，由协程主动交出控制权，编译器/解释器/虚拟机层面的多任务，多个协程可以在一个或多个线程上运行

func main() {
	/*for i := 0; i < 1000; i++ {
		go func() {//闭包
			for {
				fmt.Printf("Hello goroutine %d\n", i)  //I/O操作存在协程切换
			}
		}()
	}*/
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() //收到交出控制权
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a) //go run -race   race condition
}
