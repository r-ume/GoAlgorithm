package main

import(
  ct "./const"
  "fmt"
) 

func main(){
  fmt.Println(ct.MAX)             // fooパッケージの定数MAXを参照
  fmt.Println(ct.internal_count)  // コンパイルエラー

  fmt.Println(ct.FooFunc())       // fooパッケージの関数FooFuncを参照
  fmt.Println(ct.internalFunc(5)) // コンパイルエラー
}
