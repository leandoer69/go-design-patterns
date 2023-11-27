package clothes

import "fmt"

type AdidasSneakers struct {
	Year  uint
	Color string
}

func (sn *AdidasSneakers) Present() {
	fmt.Printf("Adidas Sneakers Color:[%s] Year:[%d]\n", sn.Color, sn.Year)
}
