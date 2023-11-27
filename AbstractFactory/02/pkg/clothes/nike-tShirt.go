package clothes

import "fmt"

type NikeTShirt struct {
	Color    string
	Material string
}

func (ts *NikeTShirt) Present() {
	fmt.Printf("Nike T-Shirt Color:[%s] Material:[%s]\n", ts.Color, ts.Material)
}
