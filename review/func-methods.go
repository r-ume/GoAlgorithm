// 関数
// func 関数名（引数）戻り値の型 {
//     // 処理
//     return
// }

// メソッド
// func (レシーバ　型) 関数名（引数）戻り値の型 {
//     // 処理
//     return
// }

type Calc struct { hoge, fuga int }
type Sums []Calc

func Add(q Calc) int {
  return q.hoge + q.fuga
}

func (p Calc) Add() int {
  return p.hoge + p.fuga
}

func (p Calc) Sub() int {
  return p.hoge - p.fuga
}

func (p Calc) Multi() int {
  return p.hoge * p.fuga
}

func (p Calc) Div() int {
  if p.fuga == 0 {
      return 0
  }
  return p.hoge / p.fuga
}

func (s Sums) Adds() int {
    ans := 0
    for _, s := range s {
        ans += s.Add()
    }
    return ans
}

func main() {
  q := Calc{ 8, 6 }
  fmt.Println(Add(q))

  p := Calc{ 3, 2 }         
  fmt.Println(p.Add()) 

  sums := Sums{
      {1, 3},
      {2, 4},
      {3, 5},
  }

  fmt.Println(sums.Adds())  
}
