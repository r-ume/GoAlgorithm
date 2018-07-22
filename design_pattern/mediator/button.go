package mediator

type (
	// Button button
	Button struct {
		Enabled  bool
		Mediator Mediator
	}
)

// NewButton returns new button
func NewButton() Colleague {
	return &Button{}
}

// SetMediator set mediator
func (b *Button) SetMediator(mediator Mediator) {
	b.Mediator = mediator
}

// SetColleagueEnabled set colleague enabled
func (b *Button) SetColleagueEnabled(enabled bool) {
	b.Enabled = enabled
}
