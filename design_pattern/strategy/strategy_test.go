package strategy

import (
	"testing"
)

// TestStrategy strategy pattern test
func TestStrategy(t *testing.T) {
	player1 := Player{Name: "A", Strategy: &WinningStrategy{}}
	player2 := Player{Name: "B", Strategy: &WinningStrategy{}}

	hand1 := player1.NextHand()
	hand2 := player2.NextHand()

	if hand1.IsStrongerThan(hand2) {
		player1.Win()
		player2.Lose()
	} else if hand1.IsWeakerThan(hand2) {
		player2.Lose()
		player1.Win()
	} else {
		player1.Even()
		player2.Even()
	}
}
