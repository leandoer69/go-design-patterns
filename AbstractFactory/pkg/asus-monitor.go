package absFact

import "fmt"

type AsusMonitor struct {
	Size float64
}

func (m *AsusMonitor) PrintInformation() {
	fmt.Printf("[Asus] Monitor Size:[%f]\n", m.Size)
}
