package templatemethod

type (
	// CharDisplayImpl char display struct
	CharDisplayImpl struct {
		Char string
		*AbstractDisplayImpl
	}
)

// NewCharDisplay new char display
func NewCharDisplay(char string) *CharDisplayImpl {
	return &CharDisplayImpl{Char: char}
}

// Open open
func (c *CharDisplayImpl) Open() string {
	return "<<"
}

// Print print
func (c *CharDisplayImpl) Print() string {
	return c.Char
}

// Close close
func (c *CharDisplayImpl) Close() string {
	return ">>"
}
