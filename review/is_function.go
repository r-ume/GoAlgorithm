package main

import(
  "fmt"
  "reflect"
)

func DoSomething() {}

type Foo struct {}

func (this *Foo) DoSomething() {}

func IsFunction(value interface{}) bool{
  if reflect.ValueOf(value).Type().Kind() == reflect.Func {
    return true
  }

  return false
}

func main(){
  fmt.Println(IsFunction(true))
  fmt.Println(IsFunction(1))
  fmt.Println(IsFunction("string"))
  fmt.Println(IsFunction(DoSomething))

  foo := &Foo{}
  fmt.Println(IsFunction(foo.DoSomething))  
}
