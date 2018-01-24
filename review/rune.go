package main

import "fmt"

func main(){
  nihongo := "日本語"
  StringLoop(nihongo)
  RuneLoop(nihongo)
  RuneValue(nihongo)
}

func StringLoop(nihongo string) {
  size := len(nihongo)

  fmt.Printf("nihongo = %d bytes :", size)
  for i := 0; i < size; i++ {
    fmt.Printf(" %x", nihongo[i])
  }
  fmt.Print("\n")
}

func RuneLoop(nihongo string) {
  nihongoRune := []rune(nihongo)
  size := len(nihongoRune)

  fmt.Printf("nihongo = %d characters : ", size)
  for i := 0; i < size; i++ {
    fmt.Printf("%#U ", nihongoRune[i])
  }
  fmt.Printf("\n")
}

func RuneValue(nihongo string){
  for pos, runeValue := range nihongo {
    fmt.Printf("%#U starts at byte position %d\n", runeValue, pos)
  }
}
