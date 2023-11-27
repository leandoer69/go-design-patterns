package factory

import "go_design_patterns/AbstractFactory/02/pkg/clothes"

type NikeFactory struct {
}

func (factory *NikeFactory) GetTShirt() clothes.TShirt {
	return &clothes.NikeTShirt{
		Color:    "Green",
		Material: "Linen",
	}
}

func (factory *NikeFactory) GetSneakers() clothes.Sneakers {
	return &clothes.NikeSneakers{
		Year:  2015,
		Color: "Violet",
	}
}
