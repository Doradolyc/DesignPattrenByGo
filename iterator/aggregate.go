package iterator

type Aggregate interface {
	GetIterator() Iterator
}

type Book struct {
	Name string
}

func (b *Book) GetName() string {
	return b.Name
}

type BookShelf struct {
	books []*Book
	index int
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book ...*Book) {
	bs.books = append(bs.books, book...)
	bs.index += len(book)
}

func (bs *BookShelf) GetLength() int {
	return bs.index
}

func (bs *BookShelf) GetIterator() Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
		index:     0,
	}
}
