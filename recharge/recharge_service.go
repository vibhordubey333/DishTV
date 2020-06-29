package recharge

import (
	"fmt"
)

var rechargeObject *RechargeTokens

type Recharge interface {
	CheckBalance() int
	DoRecharge(int) string
}

type RechargeTokens struct {
	balance int
	flag    int
}

func New() *RechargeTokens {

	if rechargeObject.flag == 0 {
		rechargeObject.flag = 1      // Flag to ensure , remaining balance doesn't get overrided.
		rechargeObject.balance = 100 // Initializing 100 Rs everytime program starts.
		return rechargeObject
	} else {
		return rechargeObject
	}
}

func init() {
	rechargeObject = new(RechargeTokens) // Creating the object for first time so we don't get nil pointer
	// execption when we're accessing flag variable in New() function.
	_ = New()
}

func (rechObj *RechargeTokens) CheckBalance() int {
	return rechObj.balance
}
func (rechObj *RechargeTokens) DoRecharge(amount int) string {
	// Currently not handling -ve amount scenario.

	rechObj.balance += amount
	return fmt.Sprintf("Recharge completed successfully. Current balance is %d", rechObj.balance)
}
