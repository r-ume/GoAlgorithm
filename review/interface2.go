package main

import "fmt"

type Accessor interface{
  GetText()       string
  SetText(string)
}

type Document struct{
  text string
}

type Page struct{
  Document 
  Page     int
}

func (document Document) GetText() string{
  return document.text
}

func (document *Document) SetText(text string){
  document.text = text
}

// Accessor Interfaceを入れている変数なら動く
func SetAndGet(acsr Accessor) {
  acsr.SetText("accessor")
  fmt.Println(acsr.GetText())
}

func main(){
  var doc *Document = &Document{}
  doc.SetText("document")
  fmt.Println(doc.GetText())

  // Accessor Interface を実装しているので
  // Accessor 型に代入可能
  var acsr Accessor = &Document{}
  acsr.SetText("accessor")
  fmt.Println(acsr.GetText())

  // implements Accessor interface
  var acsr Accessor = &Page{}
  acsr.SetText("page")
  fmt.Println(acsr.GetText())

  SetAndGet(&Page{})
  SetAndGet(&Document{})
}
