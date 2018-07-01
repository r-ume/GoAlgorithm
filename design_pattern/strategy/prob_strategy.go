package strategy

import "math/rand"

var (
	// History history
	History = map[int][]int{
		0: {1, 1, 1},
		1: {1, 1, 1},
		2: {1, 1, 1},
	}

	prevHandValue    int
	currentHandValue int
)

type (
	// ProbStrategy probability strategy
	ProbStrategy struct {
		Random *rand.Rand
	}
)

// NewProbStrategy returns new prob strategy
func NewProbStrategy(seed int) *ProbStrategy {
	random := rand.New(rand.NewSource(1))
	return &ProbStrategy{
		Random: random,
	}
}

// NewHand returns new hand
func (p *ProbStrategy) NextHand() *Hand {
	bet := rand.Intn(p.getNum(currentHandValue))
	handValue := 0

	if bet < History[currentHandValue][0] {
		handValue = 0
	} else if bet < History[currentHandValue][0]+History[currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}

	prevHandValue = currentHandValue
	currentHandValue = handValue
	return GetHand(handValue)
}

// Study study
func (p *ProbStrategy) Study(win bool) {
	if win {
		History[prevHandValue][currentHandValue]++
	}

	History[prevHandValue][(currentHandValue+1)%3]++
	History[prevHandValue][(currentHandValue+2)%3]++
}

func (p *ProbStrategy) getNum(handValue int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += History[handValue][i]
	}
	return sum
}
