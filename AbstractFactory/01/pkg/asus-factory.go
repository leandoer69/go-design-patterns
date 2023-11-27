package absFact

type AsusFactory struct {
}

func (factory *AsusFactory) GetComputer() Computer {
	return &AsusComputer{
		Memory: 16,
		Cpu:    8,
	}
}

func (factory *AsusFactory) GetMonitor() Monitor {
	return &AsusMonitor{Size: 16.2}
}
