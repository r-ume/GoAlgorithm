package iterator

// Aggregate aggregate interface
// コレクションを表すインターフェース
type (
	Aggregate interface {
		Iterator() Iterator
	}
)
