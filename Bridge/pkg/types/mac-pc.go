package types

import (
	"fmt"
	"go_design_patterns/Bridge/pkg/base"
)

type MacPC struct {
	scanner base.Scanner
}

func (pc *MacPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *MacPC) Scan() {
	fmt.Println("Scan pc mac system")
	pc.scanner.ScanFile()
}
