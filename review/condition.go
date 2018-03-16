package main

import "fmt"

func main() {
	// if文の条件に使えるのは、bool型か比較演算
	hoge := true

	if hoge {
		fmt.Println("hoge")
	}

	piyo := 10
	if piyo > 5 {
		fmt.Println("piyo is larger than 5.")
	}

	// fuga := "fuga"
	// if fuga {
	// 	fmt.Println("string data is true as well.")
	// }
	// => non-bool fuga (type string) used as if condition
}
