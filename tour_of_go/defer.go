// defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。

// defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません。

package main

import(
  "fmt"
  "os"
)

func main(){
  defer fmt.Println("world")

  fmt.Println("hello")

  file, err := os.Open("hoge.txt")
  if err != nil {
      fmt.Println("File open error: ", err)
      return
  }

  // ここでCloseを遅延実行する
  // deferが行われた以降の行のどこでreturnしても、必ずfile.Closeが実行されます
  defer file.Close()

  buf := make([]byte, 1024)
  for {
      // 省略
  }

}

