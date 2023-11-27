package types

import "fmt"

type Hp struct {
}

func (s Hp) ScanFile() {
	fmt.Println("HP scan file")
}
