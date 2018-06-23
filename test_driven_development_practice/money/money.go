package money

type (
	// DollarImpl strcut
	DollarImpl struct {
		Amount int64
	}
)

// NewDollar returns dollar interface
func NewDollar(amount int64) *DollarImpl {
	return &DollarImpl{Amount: amount}
}

// Times returns DollarImpl whose amount is multipler times larger
func (d *DollarImpl) Times(multipler int64) *DollarImpl {
	return &DollarImpl{Amount: d.Amount * multipler}
}

// EqualsTo checkes if given dollar is the same
// Value Object 別名参照をなくす - 5ドルオブジェクトが必要なら、それは永遠にそのままだ。7ドル必要なら、それは作らないといけない。
func (d *DollarImpl) EqualsTo(dollar interface{}) bool {
	givenDollar := dollar.(*DollarImpl)
	return d.Amount == givenDollar.Amount
}
