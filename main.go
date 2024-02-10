package main

import (
	"dorado.host/DesignPattrenByGo/adapter"
	"dorado.host/DesignPattrenByGo/iterator"
	"fmt"
)

func main() {
	//// Iterator 迭代器模式
	//showIteratorOutPut()

	//// Adapter 适配器模式
	//showAdapterOutPut()
}

func showIteratorOutPut() {
	bookShelf := new(iterator.BookShelf)
	bookShelf.AppendBook(&iterator.Book{Name: "abc"}, &iterator.Book{Name: "bcd"}, &iterator.Book{Name: "cdf"})

	bookShelfIterator := bookShelf.GetIterator()
	for bookShelfIterator.HasNext() {
		book := bookShelfIterator.Next().(*iterator.Book)
		fmt.Println(book.GetName())
	}
}

func showAdapterOutPut() {
	var p adapter.Print
	p = &adapter.PrintBanner{
		Banner: adapter.Banner{
			Str: "Hello World",
		},
	}
	p.PrintWeak()
	p.PrintString()
}
