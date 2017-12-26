package main

import "fmt"

func main(){
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  var s []int = primes[1:4]
  fmt.Println(s)

  names := [4]string{"John", "Paul", "George", "Ringo"}
  fmt.Println(names)

  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  w := []struct {
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(w)

  intarr := []int{2, 3, 5, 7, 11, 13}

  s = s[1:4]
  fmt.Println(s)

  s = s[:2]
  fmt.Println(s)

  s = s[1:]
  fmt.Println(s)


}
