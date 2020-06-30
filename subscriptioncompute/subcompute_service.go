package subscriptioncompute

import (
	balanceinfo "DishTV/recharge"
	"errors"
)

const (
	SILVER   int     = 50
	GOLD     int     = 100
	DISCOUNT float64 = 0.1 // 10% is 0.1 .
)

var rechargeObject *balanceinfo.RechargeTokens

type Subscription interface {
	ComputeAmount(int) int
	Discount(int, int) int
}

//Accepts pack type and months of subscription. And returns subscription amount , discount , finalprice,error
func ComputeAmount(pack string, months int) (subAmount, discount, finalAmount int, remainingBalance, monthlyPrice int, packType string, status bool, err error) {

	rechargeObject = balanceinfo.New()
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
		} else {
			finalAmount = subAmount // As no discount is applied it will be similar to pack amount
		}
	} else { // Handling for SILVER pack.

		monthlyPrice = SILVER
		packType = "SILVER"
		subAmount = months * SILVER
		if months >= 3 {
			discount = int(Discount(months, SILVER))
			finalAmount = subAmount - discount
		} else {
			finalAmount = subAmount // As no discount is applied it will be similar to pack amount
		}
	} //End of else brace.

	if remainingBalance < subAmount { // If balance is less than final amount throw error.
		return 0, 0, 0, remainingBalance, monthlyPrice, packType, false, nil
	}

	deductbalance := finalAmount * -1
	rechargeObject.InternalDoRecharge(deductbalance)
	remainingBalance -= finalAmount
	SubscribedChannel(packType) // Saving packType

	return subAmount, discount, finalAmount, remainingBalance, monthlyPrice, packType, true, nil
}

//Compute discount.
func Discount(months int, pack int) float64 {

	amount := months * pack
	discount := float64(amount) * DISCOUNT
	return discount
}
