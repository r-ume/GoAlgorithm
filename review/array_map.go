package main

import "fmt"

func main(){
  people := People{
    {"Nic", "Williams"},
    {"Banjo", "Williams"},
  }

  proc := func(p *Person) string{
    return p.firstName
  }

  fmt.Printf("%q\n", people.Map(proc))
}

type Person struct{
  firstName string
  lastName string
}

type People []Person

func NewPeople() People{
  return People{}
}

func (people People) firstNames() []string {
  firstNames := make([]string, len(people))

  for i, person := range people{
    fmt.Println(person.firstName)
    firstNames[i] = person.firstName
  }

  return firstNames
}

func (people People) Map(f func(*Person) string) []string{
  result := make([]string, len(people))
  for index, person := range people{
    result[index] = f(&person)
  }

  return result
}


