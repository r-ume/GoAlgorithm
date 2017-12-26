package main

import "fmt"

// Go langでは、連想配列のことをmapと呼ぶ
func main() {
    m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
    fmt.Println(m) // => "map[apple:150 banana:300 lemon:300]"
}
