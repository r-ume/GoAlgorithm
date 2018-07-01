package main

import (
	"fmt"

	"github.com/romukey/GolangPlayground/design_pattern/adapter"
	"github.com/romukey/GolangPlayground/design_pattern/strategy"
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

	// Player
	player1 := strategy.NewPlayer("taro", strategy.NewWinningStrategy(314))
	player2 := strategy.NewPlayer("Hana", strategy.NewProbStrategy(15))

	for i := 0; i < 1000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()

		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Println("Winnger: " + player1.ToString())
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			fmt.Println("Loser: " + player2.ToString())
			player1.Win()
			player2.Lose()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}

		fmt.Println("Total result: ")
		fmt.Println(player1.ToString())
		fmt.Println(player2.ToString())
	}
}
