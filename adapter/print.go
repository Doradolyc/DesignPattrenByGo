package adapter

type Print interface {
	PrintWeak()
	PrintString()
}

type PrintBanner struct {
	Banner
}

func (pb *PrintBanner) PrintWeak() {
	pb.ShowWithParen()
}

func (pb *PrintBanner) PrintString() {
	pb.ShowWithAster()
}
