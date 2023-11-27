package builder

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     bool
}

func (c *Computer) Print() {
	fmt.Printf("%s Core:[%d] Memory:[%d] Graphic Card:[%d] Has Monitor:[%v]\n",
		c.Brand, c.Core, c.Memory, c.GraphicCard, c.Monitor)
}
