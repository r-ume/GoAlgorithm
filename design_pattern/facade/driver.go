package facade

type (
	// DriverImpl driver
	DriverImpl struct {
		Car *CarImpl
	}
)

// NewDriver returns new driver implementation
func NewDriver(car *CarImpl) *DriverImpl {
	return &DriverImpl{Car: car}
}

// PushPedal pushes pedal
func (d *DriverImpl) PushPedal(speed int64) {
	d.Car.SetSpeed(speed)
}

// Drive drive
func (d *DriverImpl) Drive(distance int64) {
	d.Car.Run(distance)
}
