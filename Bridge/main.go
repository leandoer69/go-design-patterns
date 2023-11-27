package main

import (
	"go_design_patterns/Bridge/pkg/types"
)

var (
	/* scanners */
	hpScanner    = types.Hp{}
	epsonScanner = types.Epson{}
	/* computers */
	linuxPc   = types.LinuxPC{}
	windowsPc = types.WinPC{}
	macPc     = types.MacPC{}
)

func main() {
	linuxPc.AddScanner(&hpScanner)
	windowsPc.AddScanner(epsonScanner)
	macPc.AddScanner(&hpScanner)

	linuxPc.Scan()
	windowsPc.Scan()
	macPc.Scan()
}
