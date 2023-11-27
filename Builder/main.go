package main

import "go_design_patterns/Builder/pkg"

func main() {
	asusCollector := pkg.GetCollector(pkg.AsusCollectorType)
	hpCollector := pkg.GetCollector(pkg.HpCollectorType)

	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
