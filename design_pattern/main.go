package main

import (
	"fmt"

	"github.com/romukey/GolangPlayground/design_pattern/adapter"
	"github.com/romukey/GolangPlayground/design_pattern/templatemethod"
)

func main() {
	// adapter
	printBanner := adapter.NewPrintBanner("Hello")
	printBanner.PrintWeak()
	printBanner.PrintStrong()

	// templateMethod
	charDisplay := templatemethod.NewCharDisplay("s")
	fmt.Println(charDisplay.Display(charDisplay))
	stringDisplay := templatemethod.NewStringDisplay("romukey")
	fmt.Println(stringDisplay.Display(stringDisplay))
}
