package main

import "fmt"

func main(){
  m1 := map[string]string{"key1": "val1", "key2": "val2"}
  m2 := map[string]string{"key3": "val3"}

  fmt.Println(SliceMerge(m1, m2))
}

func SliceMerge(m1, m2 map[string]string) map[string]string {
  result := map[string]string{}

  for key1, value1 := range m1{
    result[key1] = value1
  }

  for key2, value2 := range m2{
    result[key2] = value2
  }

  return result
}
