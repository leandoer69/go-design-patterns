package fmethod

import "fmt"

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  32,
		Monitor: true,
	}
}

func (s PersonalComputer) GetType() string {
	return s.Type
}

func (s PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory:[%d] Monitor:[%v]\n",
		s.Type, s.Core, s.Memory, s.Monitor)
}
