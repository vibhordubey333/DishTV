package subscriptioncompute

import (
	"errors"
	"fmt"
)

const SILVER int = 50
const GOLD int = 100

type subscription interface {
	ComputeAmount(int) int
	Discount(int, int) int
}

//Accepts pack type and months of subscription. And returns subscription amount , discount , finalprice,error
func ComputeAmount(pack string, months int) (subAmount int, discount int, finalAmount int, err error) {

	//var subAmount int
	//var discount int
	//var finalAmount int

	if months <= 0 {
		return 0, 0, 0, errors.New("Invalid input")
	}
	if pack == "S" {
		subAmount = months * GOLD
		if months >= 3 {
			discount = Discount(months, GOLD)
			finalAmount = subAmount - discount
		}
	} else {
		discount = Discount(months, SILVER)
		finalAmount = subAmount - discount
	}
	fmt.Print(subAmount, discount, finalAmount)
	return subAmount, discount, finalAmount, nil
}
func Discount(months int, pack int) int {
	amount := months * pack
	discount := (amount / 100) * 10
	return discount
}
