package main

import (
	"fmt"
	decorator "go_design_patterns/Decorator/pkg"
)

var (
	basePc = decorator.BasePc{}
	homePc = decorator.HomePc{
		Cpu:     4,
		Wrapper: basePc,
	}
	serverPc = decorator.ServerPc{
		Cpu:     16,
		Memory:  256,
		Wrapper: homePc,
	}
)

func main() {
	fmt.Printf("Base [%f]\n", basePc.GetPrice())
	fmt.Printf("Home CPU:[%d] Price:[%f]\n", homePc.Cpu, homePc.GetPrice())
	fmt.Printf("Server CPU:[%d] Memory:[%d] Price:[%f]\n",
		serverPc.Cpu, serverPc.Memory, serverPc.GetPrice())
}
