package main

import(
  "log"
  "os/exec"
)

func main(){
  log.Print("これはコマンドを実行するためのプログラムです。")
  out, error := exec.Command("date").Output()

  if error != nil{
    log.Printf("エラー: %s", error)
    return
  }

  log.Printf("現在時刻: %s", out)

  out_two, error_two := exec.Command("false").Output()

  if error_two != nil{
    log.Printf("エラー: %s", error_two)
    return
  }

  log.Printf("%s", out_two)
}
