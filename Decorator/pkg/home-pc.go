package decorator

type HomePc struct {
	Cpu     int
	Wrapper Wrapper
}

func (pc HomePc) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.Cpu)
}
