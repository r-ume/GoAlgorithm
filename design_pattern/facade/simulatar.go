package facade

import "fmt"

type (
	// SimulatorImpl simulator
	SimulatorImpl struct {
	}
)

// NewSimulator returns new simulator
func NewSimulator() *SimulatorImpl {
	return &SimulatorImpl{}
}

// Simulate simulate
func (s *SimulatorImpl) Simulate() {
	car := NewCar()
	driver := NewDriver(car)

	driver.PushPedal(int64(700))
	driver.Drive(int64(30))
	driver.PushPedal(int64(70))
	driver.Drive(int64(20))

	fmt.Println("The travel distance is ", car.GetDistance(), " m.")
}
