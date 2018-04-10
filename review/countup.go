package main

import (
	"fmt"
	"time"
)

func main() {
	nextInt := intSeq(3)

	records := []string{"hoge", "piyo", "fuga"}

	for _, record := range records {
		fmt.Println(record)
		nextInt()
	}
}

func intSeq(kazu int) func() {
	i := 0
	return func() {
		i++
		if i%kazu == 0 {
			i = 0
			time.Sleep(2 * time.Second)
		}
	}
}