package money

type (
	// FrancImpl Franc implementation
	FrancImpl struct {
		Amount int64
	}
)

// NewFranc returns new franc implementation
func NewFranc(amount int64) *FrancImpl {
	return &FrancImpl{
		Amount: amount,
	}
}

// Times times
func (f *FrancImpl) Times(multipler int64) *FrancImpl {
	return &FrancImpl{
		Amount: f.Amount * multipler,
	}
}

// EqualsTo checkes if given dollar is the same
// Value Object 別名参照をなくす - 5ドルオブジェクトが必要なら、それは永遠にそのままだ。7ドル必要なら、それは作らないといけない。
func (f *FrancImpl) EqualsTo(money interface{}) bool {
	givenFranc := money.(*FrancImpl)
	return f.Amount == givenFranc.Amount
}
