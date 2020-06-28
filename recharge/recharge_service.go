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
		fmt.Println("Flag value zero")
		rechargeObject.flag = 1
		fmt.Println("Recharge Token called.")
		//rechObj := new(RechargeTokens)
		rechargeObject.balance = 100 // Initializing 100 Rs everytime program starts.
		return rechargeObject
	} else {
		fmt.Println("Recharge token else part")
		return rechargeObject
	}

}

func init() {
	fmt.Println("Init Recharge ")
	rechargeObject = new(RechargeTokens) // Creating the object for first time so we don't get nil pointer
	// execption when we're accessing flag variable in New() function.
	_ = New()
}

func (rechObj *RechargeTokens) CheckBalance() int {
	return rechObj.balance
}
func (rechObj *RechargeTokens) DoRecharge(amount int) string {

	if amount <= 0 {
		return fmt.Sprintf("Enter amount bigger than zero.Try again.")
	} else {
		fmt.Println("Balance received is : ", amount)
		rechObj.balance += amount
	}
	return fmt.Sprintf("Recharge completed successfully. Current balance is %d", rechObj.balance)
}
