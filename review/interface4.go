package main

import "fmt"

// interfaces
type Cooker interface {
	Cook() string
}

type Painter interface {
	Paint() string
}

type Worker interface {
	Work() string
}

// struct
type HomeCooker struct{}

func (c *HomeCooker) Cook() string {
	return "簡単な料理を作ります。"
}

type SundayPainter struct{}

func (p *SundayPainter) Paint() string {
	return "日曜大工レベルでペンキを塗ります。"
}

type OfficeWorker struct{}

func (w *OfficeWorker) Work() string {
	return "会社で仕事します。"
}

type Father struct {
	Name    string
	Cooker  // interface
	Painter // interface
	Worker  // interface
}

func NewGenericFather() *Father {
	return &Father{
		"普通のお父さん",
		&HomeCooker{},    // HomeCookerの構造体には、Cook関数が入っている。
		&SundayPainter{}, // SundayPainterの構造体には、Paint関数が入っている。
		&OfficeWorker{},  // OfficeWorkerの構造体には、Work関数が入っている。
	}
}

type CookingPapa struct {
	Name string
	Painter
	Worker
}

func NewCookingPapa() *CookingPapa {
	return &CookingPapa{
		"クッキングパパ",
		&SundayPainter{},
		&OfficeWorker{},
	}
}

func (cp *CookingPapa) Cook() string {
	return "なんか工夫したりしてすごい料理を作ります。"
}

// interfaceを引数に取る
func MakeDinner(c Cooker) string {
	return c.Cook()
}

func RepairRoof(p Painter) string {
	return p.Paint()
}

func GetMoney(w Worker) string {
	return w.Work()
}

func main() {
	GenericFather := NewGenericFather()
	fmt.Println(GenericFather.Name)
	fmt.Println("- ", MakeDinner(GenericFather))
	fmt.Println("- ", RepairRoof(GenericFather))
	fmt.Println("- ", GetMoney(GenericFather))

	CookingPapa := NewCookingPapa()
	fmt.Println(CookingPapa.Name)
	fmt.Println("- ", MakeDinner(CookingPapa))
	fmt.Println("- ", RepairRoof(CookingPapa))
	fmt.Println("- ", GetMoney(CookingPapa))
}
