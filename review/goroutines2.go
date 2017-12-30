package main

import(
  "fmt"
  "log"
  "net/http"
  "sync"
)

func main(){

  wait := new(sync.WaitGroup)
  urls := []string{
    "http://example.com",
    "http://example.net",
    "http://example.org",
  }

  Without GoRoutine
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
    // waitGroupに追加
    wait.Add(1)
    
    // 取得処理をゴルーチンで実行する
    go func(url string) {
      res, err := http.Get(url)
        if err != nil {
          log.Fatal(err)
        }
      defer res.Body.Close()
      fmt.Println(url, res.Status)
      wait.Done()
    }(url)
  }

  wait.Wait()
}

