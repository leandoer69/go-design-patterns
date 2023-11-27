package clothes

import "fmt"

type BalenciagaTShirt struct {
	Color    string
	Material string
}

func (ts *BalenciagaTShirt) Present() {
	fmt.Printf("Balenciaga T-Shirt Color:[%s] Material:[%s]\n", ts.Color, ts.Material)
}
