package adapter

import "fmt"

type (
	// BannerImpl banner
	BannerImpl struct {
		Str string
	}
)

// NewBanner returns new banner implementation
func NewBanner(str string) *BannerImpl {
	return &BannerImpl{
		Str: str,
	}
}

// ShowWithParen show with parenthesis
func (b *BannerImpl) ShowWithParen() {
	fmt.Println("[" + b.Str + "]")
}

// ShowWithAster show with aster
func (b *BannerImpl) ShowWithAster() {
	fmt.Println("*" + b.Str + "*")
}
