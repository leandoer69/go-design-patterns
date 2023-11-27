package builder

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     bool
}

func (a *AsusCollector) SetCore() {
	a.Core = 4
}

func (a *AsusCollector) SetBrand() {
	a.Brand = "Asus"
}

func (a *AsusCollector) SetMemory() {
	a.Memory = 8
}

func (a *AsusCollector) SetMonitor() {
	a.Monitor = true
}

func (a *AsusCollector) SetGraphicCard() {
	a.GraphicCard = 2
}

func (a *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        a.Core,
		Brand:       a.Brand,
		Memory:      a.Memory,
		GraphicCard: a.GraphicCard,
		Monitor:     a.Monitor,
	}
}
