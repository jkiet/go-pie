package driver

type IPin interface {
	Output()
	Low()
	High()
}

type FakePin struct{}

func NewPin() IPin {
	return &FakePin{}
}
func (f *FakePin) Output() {}
func (f *FakePin) Low()    {}
func (f *FakePin) High()   {}
