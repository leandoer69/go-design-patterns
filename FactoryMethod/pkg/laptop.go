package pkg

import "fmt"

type Laptop struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewLaptop() Computer {
	return Laptop{
		Type:    LaptopType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (s Laptop) GetType() string {
	return s.Type
}

func (s Laptop) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory:[%d]\n", s.Type, s.Core, s.Memory)
}
