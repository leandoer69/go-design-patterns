package User

import "go_design_patterns/Adapter/02/pkg/Phones"

type User struct {
}

func (u *User) ChargeSmartphone(sp Phones.Smartphone) {
	sp.ChargeWithLightning()
}
