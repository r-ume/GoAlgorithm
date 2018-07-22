package mediator

type (
	// LoginForm login form
	LoginForm struct {
		RadioButton RadioButton
		Button      Button
	}
)

// NewLoginForm returns new login form
func NewLoginForm() Mediator {
	return &LoginForm{}
}

// CreateColleagues create colleagues
func (l *LoginForm) CreateColleagues() {
	l.RadioButton = NewRadioButton()
	l.Button = NewButton()
	l.RadioButton.SetMediator(l)
	l.Button.SetMediator(l)
}

// ColleagueChanged colleague changed
func (l *LoginForm) ColleagueChanged() {
	if self.Button.Checked {
		self.Button.SetColleagueEnabled(true)
	} else {
		self.Button.SetColleagueEnabled(false)
	}
}
