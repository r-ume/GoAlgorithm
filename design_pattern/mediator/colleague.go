package mediator

// 他の [Colleague] を制御したい場合は, [Mediator] に相談する. [Mediator] からの指示を受ける窓口インタフェースを定義する. (advice) 自身が相談する [Mediator] を格納するためのメソッドを定義する. (setMediator).

type (
	// Colleague collague
	Colleague interface {
		SetMediator(mediator Mediator)
		SetCollagueEnabled(enabled bool)
	}
)
