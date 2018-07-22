package mediator

type (
	// RadioButton radio button
	RadioButton struct {
		Enabled  bool
		Checked  bool
		Mediator Mediator
	}
)

// NewRadioButton returns new radio button
func NewRadioButton() Colleague {
	return &RadioButton{}
}

// SetMediator set mediator
func (r *RadioButton) SetMediator(mediator Mediator) {
	r.Mediator = mediator
}

// SetColleagueEnabled set colleage enabled
func (r *RadioButton) SetColleagueEnabled(enabled bool) {
	r.Enabled = enabled
}

// Check check
func (r *RadioButton) Check(checked bool) {
	r.Checked = checked
	// これで仲介者に通知を言う
	r.Mediator.ColleagueChanged()
}
