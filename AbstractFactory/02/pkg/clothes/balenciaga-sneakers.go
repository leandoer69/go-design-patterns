package clothes

import "fmt"

type BalenciagaSneakers struct {
	Year  uint
	Color string
}

func (sn *BalenciagaSneakers) Present() {
	fmt.Printf("Balenciaga Sneakers Color:[%s] Year:[%d]\n", sn.Color, sn.Year)
}
