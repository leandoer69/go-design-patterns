package clothes

import "fmt"

type NikeSneakers struct {
	Year  uint
	Color string
}

func (sn *NikeSneakers) Present() {
	fmt.Printf("Nike Sneakers Color:[%s] Year:[%d]\n", sn.Color, sn.Year)
}
