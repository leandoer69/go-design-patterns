package main

import "go_design_patterns/FactoryMethod/pkg"

var types = []string{fmethod.PersonalComputerType, fmethod.LaptopType, fmethod.ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := fmethod.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
