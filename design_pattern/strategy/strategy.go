package strategy

type (
	// Strategy interface
	Strategy interface {
		NextHand() *Hand
		Study(win bool)
	}
)
