package factory

import "go_design_patterns/AbstractFactory/02/pkg/clothes"

type AdidasFactory struct {
}

func (factory *AdidasFactory) GetTShirt() clothes.TShirt {
	return &clothes.AdidasTShirt{
		Color:    "Black",
		Material: "Cotton",
	}
}

func (factory *AdidasFactory) GetSneakers() clothes.Sneakers {
	return &clothes.AdidasSneakers{
		Year:  2009,
		Color: "Orange",
	}
}
