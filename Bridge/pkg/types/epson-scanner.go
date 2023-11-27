package types

import "fmt"

type Epson struct {
}

func (s Epson) ScanFile() {
	fmt.Println("Epson scan file")
}
