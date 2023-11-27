package pkg

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     bool
}

func (h *HpCollector) SetCore() {
	h.Core = 4
}

func (h *HpCollector) SetBrand() {
	h.Brand = "HP"
}

func (h *HpCollector) SetMemory() {
	h.Memory = 4
}

func (h *HpCollector) SetMonitor() {
	h.Monitor = false
}

func (h *HpCollector) SetGraphicCard() {
	h.GraphicCard = 1
}

func (h *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        h.Core,
		Brand:       h.Brand,
		Memory:      h.Memory,
		GraphicCard: h.GraphicCard,
		Monitor:     h.Monitor,
	}
}
