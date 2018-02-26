package handler

import (
  "encoding/json"
  "log"
  "net/http"
)

// http.Handle関数の第2引数は、Handlerというinterfaceを備えた構造体を受け取ります
// MyHandlerはServeHTTPという関数を持っており、Handler interfaceを満たしているため、引数として渡すことができます
type MyHandler struct{}

// MyHandler構造体のポインタにファンクションを定義します
func (t *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // 構造体を作成
  person := Person{
    Name: "p-oka",
    Age:  25,
  }

  response, err := json.Marshal(person)

  // Golangには例外がありません
  // 厳密にはpanicという仕組みがありますが、滅多に使いません
  // その代わり、返値でerrorを返し、その値がnilかどうかで成功/失敗を判断します
  if err != nil {
    log.Fatal("error occurred:", err)
    w.WriteHeader(http.StatusInternalServerError)
    errorResponse := "500 - Something bad happened!"
    w.Write([]byte(errorResponse))
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(response))
}
