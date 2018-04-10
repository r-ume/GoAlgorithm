package main

import (
	"reflect"
	"fmt"
)

type (
	Hoge struct {
		Fuga string
	}

	Piyo struct {
		Koko string
	}
)

func main() {
	hoge := Hoge{Fuga: "fuga"}
	piyo := Piyo{Koko: "koko"}

	PrintDataInLog(hoge)
	PrintDataInLog(piyo)
}

// PrintDataInLog モデルの中身を出力させる。
func PrintDataInLog(i interface{}) {
	var output string
	values := reflect.ValueOf(i)
	types := values.Type()

	for index := 0; index < types.NumField(); index++ {
		field := types.Field(index).Name
		output += fmt.Sprintf("%v: %v, ", field, values.FieldByName(field).Interface())
	}

	fmt.Println(output)
}