package practices

import "fmt"

// Cooker cooker interface
type Cooker interface {
	Cook() string
}

// Painter painter interface
type Painter interface {
	Paint() string
}

// Worker worker interface
type Worker interface {
	Work() string
}

// HomeCooker home cooker struct
type HomeCooker struct{}

// Cook return "簡単な料理を作ります。"
func (c *HomeCooker) Cook() string {
	return "簡単な料理を作ります。"
}

// SundayPainter sunday painter struct
type SundayPainter struct{}

// Paint return "日曜大工レベルでペンキを塗ります。"
func (p *SundayPainter) Paint() string {
	return "日曜大工レベルでペンキを塗ります。"
}

// OfficeWorker return "会社で仕事します。"
type OfficeWorker struct{}

// Work return "会社で仕事します。"
func (w *OfficeWorker) Work() string {
	return "会社で仕事します。"
}

// Father father struct
type Father struct {
	Name    string
	Cooker  // interface
	Painter // interface
	Worker  // interface
}

// NewGenericFather new father instance
func NewGenericFather() *Father {
	return &Father{
		"普通のお父さん",
		&HomeCooker{},    // HomeCookerの構造体には、Cook関数が実装されている。
		&SundayPainter{}, // SundayPainterの構造体には、Paint関数が実装されている。
		&OfficeWorker{},  // OfficeWorkerの構造体には、Work関数が実装されている。
	}
}

// CookingPapa cooking papa struct
type CookingPapa struct {
	Name string
	Painter
	Worker
}

// NewCookingPapa new cooking papa instance
func NewCookingPapa() *CookingPapa {
	return &CookingPapa{
		"クッキングパパ",
		&SundayPainter{},
		&OfficeWorker{},
	}
}

// Cook returns "なんか工夫したりしてすごい料理を作ります。"
func (cp *CookingPapa) Cook() string {
	return "なんか工夫したりしてすごい料理を作ります。"
}

// MakeDinner uses Cook func in Cooker interface
func MakeDinner(c Cooker) string {
	return c.Cook()
}

// RepairRoof users Paint func in Painter interface
func RepairRoof(p Painter) string {
	return p.Paint()
}

// GetMoney uses Work func in Worker interface
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
