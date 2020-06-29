package subscriptioncompute

import (
	balanceinfo "DishTV/recharge"
	//compute "DishTV/subscriptioncompute"
	"errors"
	"fmt"
	"strings"
)

type IndividualService interface {
	SubscribeIndividualService()
}

func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}
}

func SubscribeIndividualService(channels string) (status bool, err error) {
	//To save program from crashing.
	defer panicHandler()

	rechargeObject = balanceinfo.New()
	userBalance := rechargeObject.CheckBalance()    // Save existing balance and compare it with final billing price .
	subscribedChannlelAndPacks := make([]string, 0) //Save all the valid channel name user is subscribing . To display them later on.
	var finalAmount int                             // Save total amount .
	if channels == "" {
		return false, errors.New("Invalid Input.")
	}

	arrChannels := strings.Split(channels, ",")

	indiService := map[string]int{
		"Zee":         10,
		"Sony":        15,
		"Star Plus: ": 20,
		"Discovery":   10,
		"NatGeo":      20,
	}

	for _, v := range arrChannels {
		if val, ok := indiService[v]; ok {
			finalAmount += val
			subscribedChannlelAndPacks = append(subscribedChannlelAndPacks, v) // Appending channel name
			AlreadySubscribedChannel(v)
		}
	}
	if userBalance < finalAmount {
		return false, nil // Caller which fetch and let user know funds are insufficient.
	}
	finalAmount = finalAmount * (-1)
	//balanceinfo.RechargeTokens.DoRecharge(finalAmount)
	rechargeObject.DoRecharge(finalAmount)

	fmt.Println("Indi Service", indiService, "\nChan", arrChannels)

	return true, nil
}
