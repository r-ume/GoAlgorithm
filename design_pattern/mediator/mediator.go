package mediator

// Mediator パターン
// 複雑に絡み合った複数のオブジェクト間の関係を必ず「仲介者」を介して処理を行うようにすることで、単純かつ明快なインタフェースを提供するパターン。
// 管轄下にある複数のオブジェクト各々からの問い合わせを受け, 適宜判断を行い, 管轄下にあるオブジェクト全体, または一部へ指示を出す [仲介人] の役割を果たすクラスを利用するパターン.

// 仲介者
// Colleagueからの相談を受け、それに対して判断をおろし、Colleagueへ指示を出す。
type (
	// Mediator mediator
	Mediator interface {
		CreateColleagues()
		CollagueChanged()
	}
)
