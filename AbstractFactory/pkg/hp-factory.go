package absFact

type HpFactory struct {
}

func (factory *HpFactory) GetComputer() Computer {
	return &HpComputer{
		Memory: 8,
		Cpu:    4,
	}
}

func (factory *HpFactory) GetMonitor() Monitor {
	return &AsusMonitor{Size: 13.1}
}
