package main

import(
  "fmt"
  "strconv"
)

type Car interface{
  run(int) string
  stop()
}

type MyCar struct{
  name  string
  speed int
}

func (u *MyCar) run (speed int) string{
  u.speed = speed
  return strconv.Itoa(speed) + "kmで走ります"
}

func (u *MyCar) stop(){
  fmt.Println("停止します")
  u.speed = 0
}

func main(){
  myCar := &MyCar{name: "まいかー", speed: 10}

  // implements Car interface
  var i Car = myCar
  fmt.Println(i.run(50))
  i.stop()
}

