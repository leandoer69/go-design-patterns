package types

import (
	"fmt"
	"go_design_patterns/Bridge/pkg/base"
)

type WinPC struct {
	scanner base.Scanner
}

func (pc *WinPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *WinPC) Scan() {
	fmt.Println("Scan pc windows system")
	pc.scanner.ScanFile()
}
