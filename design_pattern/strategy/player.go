package strategy

import "strconv"

type (
	// Player player impl
	Player struct {
		Name      string
		Strategy  Strategy
		WinCount  int
		LoseCount int
		GameCount int
	}
)

// NewPlayer returns new player
func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		Name:     name,
		Strategy: strategy,
	}
}

// NextHand *hand
func (p *Player) NextHand() *Hand {
	return p.Strategy.NextHand()
}

// Win win result
func (p *Player) Win() {
	p.Strategy.Study(true)
	p.WinCount++
	p.GameCount++
}

// Lose result
func (p *Player) Lose() {
	p.Strategy.Study(false)
	p.LoseCount++
	p.GameCount++
}

// Even result
func (p *Player) Even() {
	p.GameCount++
}

// ToString to string
func (p *Player) ToString() string {
	return "[" + p.Name + ":" + strconv.Itoa(p.GameCount) + " games" + strconv.Itoa(p.WinCount) + " win " + strconv.Itoa(p.LoseCount) + " lose" + "]"
}
