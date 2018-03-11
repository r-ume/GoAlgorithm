package practices

import "fmt"

// Accessor accssor interface
type Accessor interface {
	GetText() string
	SetText(string)
}

// Document document interface
type Document struct {
	text string
}

// Page struct
type Page struct {
	Document
	Page int
}

// GetText text getter
func (document Document) GetText() string {
	return document.text
}

// SetText text setter
func (document *Document) SetText(text string) {
	document.text = text
}

// SetAndGet Accssor Interfaceを入れている構造体なら動く
func SetAndGet(acsr Accessor) {
	acsr.SetText("accessor")
	fmt.Println(acsr.GetText())
}

// Second main fnc
func Second() {
	var doc = new(Document)
	doc.SetText("document")
	fmt.Println(doc.GetText())

	// Accessor Interface を実装しているので
	// Accessor 型に代入可能
	var acsr Accessor = &Document{}
	acsr.SetText("accessor")
	fmt.Println(acsr.GetText())

	// implements Accessor interface
	var secondAcsr Accessor = &Page{}
	secondAcsr.SetText("page")
	fmt.Println(secondAcsr.GetText())

	SetAndGet(&Page{})
	SetAndGet(&Document{})
}
