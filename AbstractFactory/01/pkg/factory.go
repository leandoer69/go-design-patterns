package absFact

import "errors"

type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func NewFactory(brandType string) (Factory, error) {
	switch brandType {
	case Asus:
		return &AsusFactory{}, nil
	case HP:
		return &HpFactory{}, nil
	default:
		return nil, errors.New("Wrong brand name")
	}
}
