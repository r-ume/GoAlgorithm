package prototype

import "fmt"

type (
	// Prototype prototype interface
	Prototype interface {
		S() []int
		PrintValue()
		Copy() Prototype
	}

	// Impl prototype implementation
	Impl struct {
		Values []int
	}
)

// NewPrototype returns new prototype implementation
func NewPrototype(values []int) *Impl {
	return &Impl{Values: values}
}

// S s
func (p *Impl) S() []int {
	return p.Values
}

// Copy instance
func (p *Impl) Copy() Prototype {
	newValues := make([]int, cap(p.Values))
	copy(newValues, p.Values)
	return &Impl{Values: newValues}
}

// PrintValue print value
func (p *Impl) PrintValue() {
	fmt.Println("Value: ", p.Values)
}
