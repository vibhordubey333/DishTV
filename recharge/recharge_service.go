package recharge

import (
	"fmt"
)

var rechargeObject *RechargeTokens

type Recharge interface {
	CheckBalance() int
	DoRecharge(int) string
	InternalDoRecharge(int)
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
	_ = New() // Initializing 100Rs everytime program restarts.
}

func (rechObj *RechargeTokens) CheckBalance() int {
	return rechObj.balance
}

func (rechObj *RechargeTokens) DoRecharge(amount int) string {

	// Doesn't allow user to enter negative amount.
	if amount < 0 {
		return fmt.Sprintf("Recharge failed, enter valid value.")
	}
	rechObj.balance += amount
	return fmt.Sprintf("Recharge completed successfully. Current balance is %d", rechObj.balance)
}

//For internal updation of balance primarily for deducting balance.
func (rechObj *RechargeTokens) InternalDoRecharge(amount int) string {

	rechObj.balance += amount
	return fmt.Sprintf("Recharge completed successfully. Current balance is %d", rechObj.balance)
}
