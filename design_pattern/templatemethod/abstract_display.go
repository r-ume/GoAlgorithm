package templatemethod

type (
	// AbstractDisplay abstract display
	AbstractDisplay interface {
		Open() string
		Print() string
		Close() string
	}

	// AbstractDisplayImpl abstract display implementation
	AbstractDisplayImpl struct {
	}
)

// NewAbstractDisplay return
func NewAbstractDisplay() *AbstractDisplayImpl {
	return &AbstractDisplayImpl{}
}

// Display abstract display
func (a *AbstractDisplayImpl) Display(abstractDisplay AbstractDisplay) string {
	result := abstractDisplay.Open()
	for i := 0; i < 5; i++ {
		result += abstractDisplay.Print()
	}
	result += abstractDisplay.Close()
	return result
}
