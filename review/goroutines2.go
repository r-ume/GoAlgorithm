package main

import(
  "fmt"
  "log"
  "net/http"
  // "sync"
)

func main(){

  // wait := new(sync.WaitGroup)
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
  
  statusChannel := getStatus(urls)

  // ゴルーチンの中でstatusChanに値が書き込まれるまで，main()の中では値を読み出すことができません。
  // この場合，main()内ではstatusChanの読み出しが3回完了するまで処理がブロックされるため，waitGroupのような待ち合わせ処理は必要がない。
  // これにより，HTTPリクエストを並行して発行し，早く取得されたステータスから順に受け取ることができる。
  for i := 0; i < len(urls); i++ {
    fmt.Println(<-statusChannel)
  }
}

// With GoRoutine
func getStatus(urls []string) <- chan string{

  statusChannel := make(chan string, len(urls))

  for _, url := range urls {
    // waitGroupに追加
    // wait.Add(1)

    // 取得処理をゴルーチンで実行する
    go func(url string) {
      res, err := http.Get(url)
        if err != nil {
          log.Fatal(err)
        }
      defer res.Body.Close()
      statusChannel <- res.Status
      fmt.Println(url, res.Status)
      // wait.Done()
    }(url)
  }

  return statusChannel
}

