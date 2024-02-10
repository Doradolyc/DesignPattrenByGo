package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	}
	return false
}

func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}
