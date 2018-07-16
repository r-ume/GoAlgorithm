package iterator

type (
	// BookImpl book struct
	BookImpl struct {
		Name string
	}
)

// NewBook new book struct
func NewBook(name string) *BookImpl {
	return &BookImpl{Name: name}
}

// GetName get name of the book
func (b *BookImpl) GetName() string {
	return b.Name
}
