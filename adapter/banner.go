package adapter

import "fmt"

type Banner struct {
	Str string
}

func (b *Banner) ShowWithParen() {
	fmt.Println("(" + b.Str + ")")
}

func (b *Banner) ShowWithAster() {
	fmt.Println("*" + b.Str + "*")
}
