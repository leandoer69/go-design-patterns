package Phones

import "fmt"

type IPhoneSmartphone struct {
}

func (ip *IPhoneSmartphone) ChargeWithLightning() {
	fmt.Println("Iphone is charging")
}
