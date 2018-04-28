package iterator

// Aggrerate aggrerate interface
type Aggrerate interface {
	Iterator() Iterator
}

// Iterator iterator interface
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// BookShelfIterator book shelf iterator struct
type BookShelfIterator struct {
	BookShelf *BookShelf
	last      int
}

// BookShelf book shelf struct
type BookShelf struct {
	Books []*Book
}

// Book book struct
type Book struct {
	Name string
}

// Add add
func (b *BookShelf) Add(book *Book) {
	b.Books = append(b.Books, book)
}

// Iterator returns book shelf iterator
func (b *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: b}
}

// HasNext has next
func (b *BookShelfIterator) HasNext() bool {
	if b.last >= len(b.BookShelf.Books) {
		return false
	}
	return true
}

// Next next
func (b *BookShelfIterator) Next() interface{} {
	book := b.BookShelf.Books[b.last]
	b.last++
	return book
}
