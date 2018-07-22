package facade

// facadeパターン
// 大きなプログラムを使って処理を行うためには、関係し合っているたくさんのクラスを制御しないといけない。
// その処理を行うための「窓口」を用意しておくのがいい。
// たくさんのクラスの個別を制御しなくても、その「窓口」に対して、要求を出すだけで処理がすすむ

type (
	// CarImpl car
	CarImpl struct {
		Speed    int64
		Distance int64
	}
)

// NewCar returns new car implemetantion
func NewCar() *CarImpl {
	return &CarImpl{int64(0), int64(0)}
}

// SetSpeed set speed
func (c *CarImpl) SetSpeed(speed int64) {
	c.Speed = speed
}

// Run run
func (c *CarImpl) Run(distance int64) {
	c.Distance = distance
}

// GetDistance get distance
func (c *CarImpl) GetDistance() int64 {
	return c.Distance
}
