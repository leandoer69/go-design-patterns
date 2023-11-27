package fmethod

import "fmt"

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	LaptopType           = "laptop"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case LaptopType:
		return NewLaptop()
	default:
		fmt.Printf("Несуществующий тип объекта: %s\n", typeName)
		return nil
	}
}
