package strategy

const (
	// GUU 0
	GUU = iota
	// CHO 1
	CHO
	// PAA 2
	PAA
)

var (
	// Hands hands
	Hands = []*Hand{
		NewHand(GUU),
		NewHand(CHO),
		NewHand(PAA),
	}

	// HandName hand name
	HandName = []string{"グー", "チョキ", "パー"}
)

type (
	// Hand hand struct
	Hand struct {
		Value int
	}
)

// NewHand new hand
func NewHand(value int) *Hand {
	return &Hand{Value: value}
}

// GetHand get hand (static)
func GetHand(handValue int) *Hand {
	return Hands[handValue]
}

// IsStrongerThan is stronger than another hand
func (h *Hand) IsStrongerThan(anotherHand *Hand) bool {
	return h.Fight(anotherHand) == 1
}

// IsWeakerThan is weaker than another hand
func (h *Hand) IsWeakerThan(anotherHand *Hand) bool {
	return h.Fight(anotherHand) == -1
}

// Fight return the result
func (h *Hand) Fight(anotherHand *Hand) int {
	if h == anotherHand {
		return 0
	} else if (h.Value+1)%3 == h.Value {
		return 1
	}

	return -1
}

// ToString to string
func (h *Hand) ToString() string {
	return HandName[h.Value]
}
