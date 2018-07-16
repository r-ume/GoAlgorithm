package iterator

// Iterator iterator interface
// 数え上げ、スキャンを行うインターフェース
type (
	Iterator interface {
		HasNext() bool
		Next() interface{}
	}
)
