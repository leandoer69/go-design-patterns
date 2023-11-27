package main

import (
	"go_design_patterns/Adapter/02/pkg/Phones"
	"go_design_patterns/Adapter/02/pkg/User"
)

func main() {
	user := User.User{}
	android := Phones.AndroidSmartphone{}

	// doesn't work!
	// user.ChargeSmartphone(android)

	adapter := Phones.UsbToLightningAdapter{Android: &android}
	user.ChargeSmartphone(&adapter)
}
