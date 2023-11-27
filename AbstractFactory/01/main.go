package main

import (
	"fmt"
	absFact2 "go_design_patterns/AbstractFactory/01/pkg"
)

var brands = []string{absFact2.Asus, absFact2.HP, "dell"}

func main() {
	for _, brand := range brands {
		factory, err := absFact2.NewFactory(brand)
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
