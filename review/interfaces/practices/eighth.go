package practices

import "fmt"

// Eighth eighth interface practice
func Eighth() {
	anyExec(12)
	anyExec("hello")
	anyExec([]string{"cat", "dog"})
	anyExec([2]string{"hello", "world"})
}

func anyExec(any interface{}) {
	fmt.Println(any)
}

// 12
// hello
// [cat dog]
// [hello world]
