package main

import(
  "log"
  "github.com/jmcvetta/napping"
)

func main(){
  urlName := "r-ume" // ここは自分のものに
    password := "BaymaxTakashi" // ここは自分のものに

    // 入力JSONのフォーマットを定義 & 初期化
    input :=  struct {
        UrlName string `json:"url_name"` // JSONのキーと構造体のキーが一致しないものは、ここでマッピングする 
        Password string `json:"password"` // 同上
    } {
        urlName,
        password,
    }

    // 正常時のJSONのフォーマットを定義 & 初期化
    // Qiitaからのレスポンスには url_name も含まれるが、必要ないのでここでは定義しない
    response := struct {
        Token string
    } {}

    // エラー時のJSONのフォーマットを定義 & 初期化
    error := struct {
        Error string
    } {}

    resp, err := napping.Post("https://qiita.com/api/v1/auth", &input, &response, &error) // input, response, errorは参照を渡す

    if err != nil {
        panic(err)
    }

    log.Printf("status: %#v", resp.Status())
    log.Printf("token: %#v", response.Token)
    log.Printf("error message: %#v", error.Error)
}
