package iterator

// BookShelfIteratorImpl book shelf iterator struct
// 本棚
type (
	BookShelfIteratorImpl struct {
		BookShelf *BookShelfImpl
		Index     int
	}
)

// NewBookShelfIterator returns new bookshelf iterator implemenation
// BookShelf(本棚)クラスにあるデータを使って、全体をスキャンしていく。
// Iterator インターフェースで定義されているメソッドの実装をしている。
func NewBookShelfIterator(bookShelf *BookShelfImpl) Iterator {
	return &BookShelfIteratorImpl{
		BookShelf: bookShelf,
		Index:     0,
	}
}

// HasNext has next
func (b *BookShelfIteratorImpl) HasNext() bool {
	if b.Index < b.BookShelf.GetLength() {
		return true
	}
	return false
}

// Next next
func (b *BookShelfIteratorImpl) Next() interface{} {
	book := b.BookShelf.GetBookAt(b.Index)
	b.Index++
	return book
}
