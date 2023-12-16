package decorator

type BasePc struct {
}

func (b BasePc) GetPrice() float64 {
	return 10.0
}
