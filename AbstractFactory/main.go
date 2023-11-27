package main

import (
	"fmt"
	absFact "go_design_patterns/AbstractFactory/pkg"
)

var brands = []string{absFact.Asus, absFact.HP, "dell"}

func main() {
	for _, brand := range brands {
		factory, err := absFact.NewFactory(brand)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		computer := factory.GetComputer()
		computer.PrintInformation()

		monitor := factory.GetMonitor()
		monitor.PrintInformation()
	}
}
