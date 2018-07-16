package adapter

// Role: Adapter 適応する側
// Adapteeで揃っている関数を使う

type (
	// PrintBannerImpl print banner interface
	PrintBannerImpl struct {
		Banner *BannerImpl
	}
)

// NewPrintBanner new print banner
func NewPrintBanner(str string) *PrintBannerImpl {
	return &PrintBannerImpl{
		Banner: NewBanner(str),
	}
}

// PrintWeak print weak
func (p *PrintBannerImpl) PrintWeak() {
	p.Banner.ShowWithParen()
}

// PrintStrong print strong
func (p *PrintBannerImpl) PrintStrong() {
	p.Banner.ShowWithAster()
}
