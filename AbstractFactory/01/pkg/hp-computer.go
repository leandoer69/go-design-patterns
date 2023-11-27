package absFact

import "fmt"

type HpComputer struct {
	Memory int
	Cpu    int
}

func (c *HpComputer) PrintInformation() {
	fmt.Printf("[HP] Pc Cpu:[%d] Memory[%d]\n", c.Cpu, c.Memory)
}
