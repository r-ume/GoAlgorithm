package strategy

import "math/rand"

const (
	// GUU 0
	GUU = iota
	// CHO 1
	CHO
	// PAA 2
	PAA
)

type (
	// Hand hand struct
	Hand struct {
		HandValue int
	}

	// Strategy interface
	Strategy interface {
		NextHand() *Hand
		Study(win bool)
	}

	// WinningStrategy winning strategy
	WinningStrategy struct {
		Won      bool
		PrevHand *Hand
	}

	// Player struct
	Player struct {
		Name      string
		Strategy  Strategy
		Wincount  int
		Losecount int
		Gamecount int
	}
)

var hands []*Hand

func init() {
	hands = []*Hand{
		&Hand{GUU},
		&Hand{CHO},
		&Hand{PAA},
	}
}

func getHand(handValue int) *Hand {
	return hands[handValue]
}

// IsStrongerThan is stronger than another hand
func (h *Hand) IsStrongerThan(anotherHand *Hand) bool {
	return h.Fight(anotherHand) == 1
}

// IsWeakerThan is weaker than another hand
func (h *Hand) IsWeakerThan(anotherHand *Hand) bool {
	return h.Fight(anotherHand) == -1
}

// Fight return the result
func (h *Hand) Fight(anotherHand *Hand) int {
	if h == anotherHand {
		return 0
	} else if (h.HandValue+1)%3 == h.HandValue {
		return 1
	} else {
		return -1
	}
}

// NextHand hand
func (w *WinningStrategy) NextHand() *Hand {
	if w.Won {
		w.PrevHand = getHand(rand.Intn(3))
	}
	return w.PrevHand
}

// Study study
func (w *WinningStrategy) Study(win bool) {
	w.Won = win
}

// NextHand *hand
func (p *Player) NextHand() *Hand {
	return p.Strategy.NextHand()
}

// Win win result
func (p *Player) Win() {
	p.Wincount++
	p.Gamecount++
}

// Lose result
func (p *Player) Lose() {
	p.Losecount++
	p.Gamecount++
}

// Even result
func (p *Player) Even() {
	p.Gamecount++
}
