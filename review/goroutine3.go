package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- start ---")

	ch := make(chan int)
	go countUpNumbers(ch)
	go catchNumber(ch)

	time.Sleep(100 * time.Millisecond) // 0.1秒遅延
	fmt.Println("--- end ---")

}

func countUpNumbers(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("countup: ", i)
	}
}

func catchNumbers(ch chan int) {
	for i := 0; i < 5; i++ {
		num := <-ch // channelから取り出し
		fmt.Println("catch: ", num)
	}
}
