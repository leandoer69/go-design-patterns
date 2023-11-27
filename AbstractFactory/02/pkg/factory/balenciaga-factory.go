package factory

import "go_design_patterns/AbstractFactory/02/pkg/clothes"

type BalenciagaFactory struct {
}

func (factory *BalenciagaFactory) GetTShirt() clothes.TShirt {
	return &clothes.BalenciagaTShirt{
		Color:    "White",
		Material: "Polyester",
	}
}

func (factory *BalenciagaFactory) GetSneakers() clothes.Sneakers {
	return &clothes.BalenciagaSneakers{
		Year:  2022,
		Color: "Black",
	}
}
