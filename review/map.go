package main

import "fmt"

// Go langでは、連想配列のことをmapと呼ぶ
func main() {
  m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
  fmt.Println(m) // => "map[apple:150 banana:300 lemon:300]"

  fmt.Println(keys(m))
  fmt.Println(values(m))
  fmt.Println(to_a(m))

  keys := []string{"apple", "lemon"}
  fmt.Println(indexes(m, keys))
}

func keys(m map[string]int) []string{
  ks := []string{}

  for k, _ := range m {
    ks = append(ks, k)
  }

  return ks
}

func values(m map[string]int) []int {
  vs := []int{}

  for _, v := range m {
    vs = append(vs, v)
  }

  return vs
}

func indexes(m map[string]int, keys []string) []int{
  vs := []int{}
  for _, k = range keys{
    vs = append(vs, m[k])
  }

  return vs
}

