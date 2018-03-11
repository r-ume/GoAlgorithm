package practices

import (
	"fmt"
	"strconv"
)

// Car interface
type Car interface {
	Run(speed int) string
	Stop()
}

// MyCar struct
type MyCar struct {
	name  string
	speed int
}

// Run the car
func (c *MyCar) Run(speed int) string {
	c.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

// Stop the car
func (c *MyCar) Stop() {
	fmt.Println("停止します")
	c.speed = 0
}

// First first interface
func First() {
	myCar := &MyCar{name: "まいかー", speed: 10}

	// implements Car interface
	var i Car = myCar
	fmt.Println(i.Run(50))
	i.Stop()
}
