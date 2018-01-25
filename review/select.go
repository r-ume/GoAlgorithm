package main

import(
  "fmt"
  "strconv"
  "time"
)

func main(){
  fmt.Println("開始")

  ch1 := make(chan int)
  ch2 := make(chan string)
  chend := make(chan struct())

  go func(chint chan<- int, chstr chan<- string, end chan<- struct{}){
    for i := 0; i < 10; i++{
      if i % 2 == 0 {
        fmt.Println("ch1へ送信")
        chint <-1 
      } else {
        fmt.Println("ch2へ送信")
        chstr <- "test" + strconv.Itoa(i)
      }
    }

    time.Sleep(1 * time.Second)
    close(end) // クローズして通知
  }(ch1, ch2, chend)

  for {
    select {
    case val := <-ch1:
      fmt.Println("ch1から受信：", val)
    case str := <-ch2:
      fmt.Println("ch2から受信：", str)
    case <-chend:
      fmt.Println("終了")
      return
    }
  }
}
