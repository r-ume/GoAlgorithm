package observer

import (
	"math/rand"
)

// Observer interface
type Observer interface {
	Update() int
}

// Number interface
type Number interface {
	GetNumber() int
}

// NumberGenerator number generator
type NumberGenerator struct {
	Observers []Observer
}

// AddObserver add observer struct
func (n *NumberGenerator) AddObserver(observer Observer) {
	n.Observers = append(n.Observers, observer)
}

// NotifyObservers notify observers struct
func (n *NumberGenerator) NotifyObservers() []int {
	var result []int
	for _, observer := range n.Observers {
		result = append(result, observer.Update())
	}

	return result
}

// RandomNumberGenerator random number generator struct
type RandomNumberGenerator struct {
	*NumberGenerator
}

// GetNumber get number
func (r *RandomNumberGenerator) GetNumber() int {
	return rand.Intn(50)
}

// Execute notify observer
func (r *RandomNumberGenerator) Execute() []int {
	return r.NotifyObservers()
}

// DigitObserver digit observer struct
type DigitObserver struct {
	Generator Number
}

// Update update the number by getting a different number
func (d *DigitObserver) Update() int {
	return d.Generator.GetNumber()
}

// NewRandomNumberGenerator return new random number generator instance.
func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{&NumberGenerator{}}
}
