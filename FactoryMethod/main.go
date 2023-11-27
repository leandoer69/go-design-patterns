package main

import "go_design_patterns/FactoryMethod/pkg"

var types = []string{pkg.PersonalComputerType, pkg.LaptopType, pkg.ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
