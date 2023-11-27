package absFact

import "fmt"

type AsusComputer struct {
	Memory int
	Cpu    int
}

func (c *AsusComputer) PrintInformation() {
	fmt.Printf("[Asus] Pc Cpu:[%d] Memory[%d]\n", c.Cpu, c.Memory)
}
