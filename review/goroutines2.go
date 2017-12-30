package main

import(
  "fmt"
  "log"
  "net/http"
  "time"
)

func main(){

  urls := []string{
    "http://example.com",
    "http://example.net",
    "http://example.org",
  }

  // Without GoRoutine
  for _, url := range urls{
    response, err := http.Get(url)
    if err != nil{
      log.Fatal(err)
    }
    defer response.Body.Close()
    fmt.Println(url, response.Status)
  }

  // With GoRoutine
  for _, url := range urls {
    // 取得処理をゴルーチンで実行する
    go func(url string) {
      res, err := http.Get(url)
        if err != nil {
          log.Fatal(err)
        }
      defer res.Body.Close()
      fmt.Println(url, res.Status)
    }(url)
  }

  time.Sleep(time.Second)
}

