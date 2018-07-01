package templatemethod

type (
	// StringDisplayImpl string display implementation
	StringDisplayImpl struct {
		Str   string
		Width int
		*AbstractDisplayImpl
	}
)

// NewStringDisplay new string display
func NewStringDisplay(str string) *StringDisplayImpl {
	return &StringDisplayImpl{
		Str:   str,
		Width: len(str),
	}
}

// Open open
func (s *StringDisplayImpl) Open() string {
	return s.printLine()
}

// Print print
func (s *StringDisplayImpl) Print() string {
	return "| " + s.Str + " |\n"
}

// Close Close
func (s *StringDisplayImpl) Close() string {
	return s.printLine()
}

func (s *StringDisplayImpl) printLine() string {
	line := "+-"
	for i := 0; i < s.Width; i++ {
		line += "-"
	}
	line += "-+\n"
	return line
}
