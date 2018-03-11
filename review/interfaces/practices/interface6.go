package practices

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct{}

func (d Dog) Cry() {
	fmt.Println("わんわん")
}

type Cat struct{}

func (c Cat) Cry() {
	fmt.Println("ニャーニャー")
}

func MakeSomeoneCry(someone interface{}) {
	fmt.Println("鳴け！")
	a, ok := someone.(Animal)
	if !ok {
		fmt.Println("動物では無いので鳴けません。")
		return
	}
	a.Cry()
}

func Sixth() {
	dog := Dog{}
	cat := Cat{}

	MakeSomeoneCry(dog)
	MakeSomeoneCry(cat)
}
