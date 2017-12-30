// 指定したLabelに飛ぶ

package main

import "fmt"

func main() {
    s := 0
    i := 1
    LOOP:
        if i > 100 {
            goto BREAK
        }
        s += i
        i += 1
        goto LOOP
    BREAK:
        fmt.Println(s)
}
