package Phones

import "fmt"

type AndroidSmartphone struct {
}

func (an *AndroidSmartphone) ChargeWithUSB() {
	fmt.Println("AndroidSmartphone is charging")
}

type UsbToLightningAdapter struct {
	Android *AndroidSmartphone
}

func (adapter *UsbToLightningAdapter) ChargeWithLightning() {
	fmt.Println("Adapter converts Lightning signal to USB")
	adapter.Android.ChargeWithUSB()
}
