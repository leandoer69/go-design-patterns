package main

import "go_design_patterns/Builder/pkg"

func main() {
	asusCollector := builder.GetCollector(builder.AsusCollectorType)
	hpCollector := builder.GetCollector(builder.HpCollectorType)

	factory := builder.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
