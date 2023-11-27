package types

import (
	"fmt"
	"go_design_patterns/Bridge/pkg/base"
)

type LinuxPC struct {
	scanner base.Scanner
}

func (pc *LinuxPC) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *LinuxPC) Scan() {
	fmt.Println("Scan pc linux system")
	pc.scanner.ScanFile()
}
