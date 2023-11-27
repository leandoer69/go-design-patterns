package clothes

import "fmt"

type AdidasTShirt struct {
	Color    string
	Material string
}

func (ts *AdidasTShirt) Present() {
	fmt.Printf("Adidas T-Shirt Color:[%s] Material:[%s]\n", ts.Color, ts.Material)
}
