package subscriptioncompute

import (
	balanceinfo "DishTV/recharge"
	"errors"
	"fmt"
)

const SILVER int = 50
const GOLD int = 100
const DISCOUNT float64 = 0.1

var rechargeObject *balanceinfo.RechargeTokens

type subscription interface {
	ComputeAmount(int) int
	Discount(int, int) int
}

//Accepts pack type and months of subscription. And returns subscription amount , discount , finalprice,error
func ComputeAmount(pack string, months int) (subAmount, discount, finalAmount int, remainingBalance, monthlyPrice int, packType string, status bool, err error) {

	rechargeObject = balanceinfo.New()

	fmt.Println("Recharge from compute : ", rechargeObject.CheckBalance())

	remainingBalance = rechargeObject.CheckBalance()

	if months <= 0 {
		return subAmount, discount, finalAmount, remainingBalance, monthlyPrice, packType, false, errors.New("Invalid input")
	}
	if pack == "G" { // Handling for GOLD pack.

		monthlyPrice = GOLD

		packType = "GOLD"
		subAmount = months * GOLD
		if months >= 3 {
			discount = int(Discount(months, GOLD))
			finalAmount = subAmount - discount
		}
	} else { // Handling for SILVER pack.

		monthlyPrice = SILVER
		packType = "SILVER"
		subAmount = months * SILVER
		if months >= 3 {
			discount = int(Discount(months, SILVER))
			finalAmount = subAmount - discount
		}
	}

	if remainingBalance < subAmount {
		return 0, 0, 0, remainingBalance, monthlyPrice, packType, false, nil
	}
	remainingBalance -= finalAmount
	fmt.Println("Compute", subAmount, discount, finalAmount, "RemainingBalance: ", remainingBalance)
	AlreadySubscribedChannel(packType) // Saving packType
	return subAmount, discount, finalAmount, remainingBalance, monthlyPrice, packType, true, nil
}
func Discount(months int, pack int) float64 {
	amount := months * pack
	discount := float64(amount) * DISCOUNT
	return discount
}
