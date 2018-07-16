package builder

type (
	// DirectorImpl director implementation
	DirectorImpl struct {
		Builder Builder
	}
)

// NewDirector returns new director implementation
func NewDirector(builder Builder) *DirectorImpl {
	return &DirectorImpl{Builder: builder}
}

// Construct construct
func (d *DirectorImpl) Construct() string {
	var result string
	result += d.Builder.MakeTitle("greeting")
	result += d.Builder.MakeString("朝から昼にかけて")
	result += d.Builder.MakeItems([]string{"good morning", "hello"})
	result += d.Builder.Close()
	return result
}
