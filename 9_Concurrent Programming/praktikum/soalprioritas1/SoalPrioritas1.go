package soalprioritas1

import (
	"fmt"
	"sync"
	"time"
)

func Kelipatan(x int) {
	go func() {
		for i := 1; ; i++ {
			fmt.Println(x * i)
			time.Sleep(3 * time.Second)
		}
	}()
}

func Kelipatan3GoroutineChannel() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			if i%3 == 0 {

				channel <- i
			}
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}
func Kelipatan3GoroutineBuffer() {
	channel := make(chan int, 100)

	go func() {
		for i := 0; i < 100; i++ {
			if i%3 == 0 {

				channel <- i
			}
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func alMutex(n int) {
	x := 1
	var mutex sync.Mutex
	go func() {
		for i := n; i > 0; i-- {
			mutex.Lock()
			x = x * i
			mutex.Unlock()
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println(x)
}
