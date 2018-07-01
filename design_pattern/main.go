package main

import "github.com/romukey/GolangPlayground/design_pattern/adapter"

func main() {
	printBanner := adapter.NewPrintBanner("Hello")
	printBanner.PrintWeak()
	printBanner.PrintStrong()
}
