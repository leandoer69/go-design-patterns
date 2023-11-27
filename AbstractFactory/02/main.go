package main

import (
	"fmt"
	"go_design_patterns/AbstractFactory/02/pkg/brands"
	"go_design_patterns/AbstractFactory/02/pkg/factory"
)

var brnds = []string{brands.Adidas, brands.Nike, brands.Balenciaga, "Puma"}

func main() {
	for _, brand := range brnds {
		factory, err := factory.NewFactory(brand)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		tShirt := factory.GetTShirt()
		tShirt.Present()

		sneakers := factory.GetSneakers()
		sneakers.Present()

	}
}
