package factory

import (
	"errors"
	"go_design_patterns/AbstractFactory/02/pkg/brands"
	"go_design_patterns/AbstractFactory/02/pkg/clothes"
)

type Factory interface {
	GetTShirt() clothes.TShirt
	GetSneakers() clothes.Sneakers
}

func NewFactory(brand string) (Factory, error) {
	switch brand {
	case brands.Adidas:
		return &AdidasFactory{}, nil
	case brands.Nike:
		return &NikeFactory{}, nil
	case brands.Balenciaga:
		return &BalenciagaFactory{}, nil
	default:
		return nil, errors.New("Unknown brand")
	}
}
