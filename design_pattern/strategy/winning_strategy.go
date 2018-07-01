package strategy

import "math/rand"

type (
	// WinningStrategy winning strategy
	WinningStrategy struct {
		Won      bool
		PrevHand *Hand
		Random   *rand.Rand
	}
)

// NewWinningStrategy returns new winning strategy
// 前回の勝負に勝ったのであれば、次も同じ手を出す。
// 前回の勝負に負けたのであれば、次の手は乱数を使って出す
func NewWinningStrategy(seed int) Strategy {
	random := rand.New(rand.NewSource(1))
	return &WinningStrategy{
		Random: random,
	}
}

// NextHand hand
func (w *WinningStrategy) NextHand() *Hand {
	if !w.Won {
		w.PrevHand = GetHand(rand.Intn(3))
	}
	return w.PrevHand
}

// Study study
func (w *WinningStrategy) Study(win bool) {
	w.Won = win
}
