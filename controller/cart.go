package controller

type Cart struct {
	Base
}

func (m *Cart) Index() {
	m.SetResponse("hello")
}
