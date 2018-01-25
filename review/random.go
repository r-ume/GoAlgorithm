package main

import(
  "fmt"
  "math/rand"
  "time"
)

func main(){
  m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

  rand.Seed(time.Now().UnixNano())

  fmt.Println(choice(m))
  fmt.Println(choice(m))
  fmt.Println(choice(m))
}

func choice(m map[string]int) string{
  l := len(m)
  i := 0

  index := rand.Intn(l)
  fmt.Println("index", index)

  answer := ""

  for key, _ := range m{
    if index == i {
      answer = key
      break
    } else{
      i++
    }
  }

  return answer
}
