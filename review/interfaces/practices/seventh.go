package practices

import "fmt"

type (
	// Human interface
	Human interface {
		Speak() string
		Eat(food string)
	}

	// UniversityStudent struct
	UniversityStudent struct {
		HumanService Human
		Fullname     string
	}

	// HighSchoolStudent struct
	HighSchoolStudent struct {
		Fullname string
	}
)

// Seventh seventh practice
func Seventh() {
	lucky := UniversityStudent{Fullname: "lucky"}
	fmt.Println(lucky.Speak())
}

// Speak dafa
func (u UniversityStudent) Speak() string {
	return u.HumanService.Speak()
}

// Speak hoge
func (h HighSchoolStudent) Speak() string {
	return "hoge"
}
