package main

import "fmt"

type Strings []string

func main(){
  fmt.Println(Strings{"A", "B", "C"}.Join(""))
}

func (s Strings) Join(d string) string{
  sum := ""
  for _, v := range s{
    if sum != "" {
      sum += d
    }
    sum += v
  }
  return sum
}


