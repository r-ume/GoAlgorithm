package main

import (
	"fmt"

	"github.com/romukey/GolangPlayground/design_pattern/adapter"
	"github.com/romukey/GolangPlayground/design_pattern/iterator"
	"github.com/romukey/GolangPlayground/design_pattern/strategy"
	"github.com/romukey/GolangPlayground/design_pattern/templatemethod"
)

func main() {
	// Adapter
	fmt.Println("----Adapter-----")
	var printBanner adapter.Print = adapter.NewPrintBanner("Hello") // 型はadapter.Print
	printBanner.PrintWeak()
	printBanner.PrintStrong()

	// Iterator
	fmt.Println("----Iterator-----")
	// Aggregateのインタフェースで定義されているメソッドが実装してある。
	bookShelf := iterator.NewBookShelf(4)
	bookShelf.AppendBook(iterator.NewBook("Around the world in 80 days"))
	bookShelf.AppendBook(iterator.NewBook("Cinderella"))

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.(*iterator.BookImpl).GetName())
	}

	// Strategy
	fmt.Println("----Strategy-----")
	player1 := strategy.NewPlayer("taro", strategy.NewWinningStrategy(314))
	player2 := strategy.NewPlayer("Hana", strategy.NewProbStrategy(15))

	for i := 0; i < 10; i++ {
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

	// TemplateMethod
	fmt.Println("----Template Method-----")
	charDisplay := templatemethod.NewCharDisplay("s")
	fmt.Println(charDisplay.Display(charDisplay))
	stringDisplay := templatemethod.NewStringDisplay("romukey")
	fmt.Println(stringDisplay.Display(stringDisplay))
}
