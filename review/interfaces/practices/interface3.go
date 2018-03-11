package practices

import "fmt"

// Gender gender type
type Gender int

const (
	// Female 1
	Female = iota
	// Male 2
	Male
)

// Person person interface
type Person interface {
	Name() string
	Title() string
}

// Student student struct
type Student struct {
	firstName string
	lastName  string
}

// Lady lady struct
type Lady struct {
	*Student
}

// Title returns "Ms."
func (l *Lady) Title() string {
	return "Ms."
}

// Gentleman gentleman struct
type Gentleman struct {
	*Student
}

// Title returns "Mr."
func (g *Gentleman) Title() string {
	return "Mr."
}

func printFullName(p Person) {
	fmt.Println(p.Title())
}

// Name name of the student
func (s *Student) Name() string {
	return s.firstName + " " + s.lastName
}

// NewPerson new person
func NewPerson(gender Gender, firstName, lastName string) Person {
	s := &Student{firstName, lastName}

	if gender == Male {
		return &Gentleman{s}
	} else {
		return &Lady{s}
	}
}

// Third third interface practice
func Third() {
	taro := NewPerson(Male, "Taro", "Yamada")
	printFullName(taro)

	hanako := NewPerson(Female, "Hanako", "Yamada")
	printFullName(hanako)
}
