package iterator

type (
	// BookShelfImpl book shelf implementation
	BookShelfImpl struct {
		Books []*BookImpl
		Last  int
	}
)

// NewBookShelf new book shelf
// オブジェクトのコレクションになるため、Aggregateインターフェースに定義されたメソッド(iterator)を取り組む。
func NewBookShelf(maxsize int64) *BookShelfImpl {
	return &BookShelfImpl{
		Books: make([]*BookImpl, maxsize),
		Last:  0,
	}
}

// GetBookAt get a book at the given index
func (b *BookShelfImpl) GetBookAt(index int) *BookImpl {
	return b.Books[index]
}

// AppendBook append book to book shelf
func (b *BookShelfImpl) AppendBook(book *BookImpl) {
	b.Books[b.Last] = book
	b.Last++
}

// GetLength get the number of books in shelf
func (b *BookShelfImpl) GetLength() int {
	return b.Last
}

// Iterator returns iterator interface
// 本が多く入っている本棚をIteratorクラスに渡す。
func (b *BookShelfImpl) Iterator() Iterator {
	return NewBookShelfIterator(b)
}
