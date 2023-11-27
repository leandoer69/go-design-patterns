package absFact

import "fmt"

type HpMonitor struct {
	Size float64
}

func (m *HpMonitor) PrintInformation() {
	fmt.Printf("[HP] Monitor Size:[%f]\n", m.Size)
}
