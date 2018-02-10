package practices
  
import "fmt"

type Animal struct{
  Name string
  Age int
}

type Animals []Animal

func Second(){
  fmt.Println("patternA")
  patternA()

  fmt.Println("patternB")
  patternB()

  fmt.Println("patternC")
  patternC()

  fmt.Println("patternD")
  patternD()

  fmt.Println("patternE")
  patternE()
}

func patternA(){
  cat1 := &Animal{}
  fmt.Println(cat1)
  fmt.Println(*cat1) 
}

func patternB(){
  cat1 := &Animal{}

  fmt.Println(cat1)  // &ついてるんで、ポインタ
  fmt.Println(*cat1) // * つけてアクセスしたので値を表す

  cat2 := cat1
  fmt.Println(cat1) // &ついてるんで、ポインタ
  fmt.Println(cat2) // cat1を代入した変数なので、ポインタ

  cat2.Name = "catty" // ポインタを渡した別変数でフィールド書き換え
  fmt.Println(cat1)
  fmt.Println(cat2)
} 

func patternC() {
  cat1 := &Animal{}

  fmt.Println(cat1) // &ついてるんで、ポインタ
  fmt.Println(*cat1) // *つけてアクセスしているので、値を表す

  cat2 := *cat1 // 値をcat2に代入
  fmt.Println(cat1) // ポインタ
  fmt.Println(cat2) // 値

  cat2.Name = "catty" // フィールドを書き換える

  fmt.Println(cat1) // 値を渡した元の変数は変化しない
  fmt.Println(cat2)

  // =>
  // &{ 0}
  // { 0}
  // &{ 0}
  // { 0}
  // &{ 0}
  // {catty 0}
}

func patternD() {
  // アドレス演算子をつけて初期化したパターン(cat1 := &Animal{})と同じ挙動
  cat1 := new(Animal) // ゼロ初期化。

  fmt.Println(cat1) // ＆ついてるんで、ポインタ
  fmt.Println(*cat1) // *つけてアクセスしてるんで、値

  cat2 :=cat1 // ポインタを別変数に代入

  fmt.Println(cat1)
  fmt.Println(cat2)

  cat2.Name = "catty" // ポインタを渡した別変数でフィールド書き換え
  fmt.Println(cat1)
  fmt.Println(cat2)

  // =>
  // &{ 0}
  // { 0}
  // &{ 0}
  // &{ 0}
  // &{catty 0}
  // &{catty 0}
}

func patternE() {
  // アドレス演算子をつけて初期化したパターン(cat1 := &Animal{})と同じ挙動
  cat1 := new(Animal) // ゼロ初期化。

  fmt.Println(cat1) // &ついてるんで、ポインタ
  fmt.Println(*cat1) // *つけてアクセスしてるんで、値

  cat2 := *cat1 // 値を別変数に代入

  fmt.Println(cat1) // ポインタ
  fmt.Println(cat2) // 値

  cat2.Name = "catty" // フィールド書き換え
  fmt.Println(cat1) // ポインタは変わらない
  fmt.Println(cat2) // 値は変わる

  // =>
  // &{ 0}
  // { 0}
  // &{ 0}
  // { 0}
  // &{ 0}
  // {catty 0}
}
