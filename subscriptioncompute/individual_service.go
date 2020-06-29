package subscriptioncompute

import (
	balanceinfo "DishTV/recharge"
	"errors"
	"fmt"
	"strings"
)

type IndividualService interface {
	SubscribeChannel()
}

func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}
}

func SubscribeChannel(channels string) (status bool, err error) {
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
			SubscribedChannel(v)                                               // Saving all the channels which user subscribed to display them later on .
		}
	}
	if userBalance < finalAmount {
		return false, nil // Caller which fetch and let user know funds are insufficient.
	}
	finalAmount = finalAmount * (-1)
	rechargeObject.DoRecharge(finalAmount)

	fmt.Println("Indi Service", indiService, "\nChan", arrChannels)
	return true, nil
}

func SubscribeService(services string) (status bool, err error) {
	//Handling , expecting user to subscribe for multiple services.
	//To save program from crashing.
	defer panicHandler()

	rechargeObject = balanceinfo.New()
	userBalance := rechargeObject.CheckBalance() // Save existing balance and compare it with final billing price .
	subscribedServices := make([]string, 0)      //Save all the valid service name user is subscribing . To display them later on.
	var finalAmount int                          // Save total amount .
	if services == "" {
		return false, errors.New("Invalid Input.")
	}

	arrServices := strings.Split(services, ",")

	indiService := map[string]int{
		"LearnEnglish": 200,
		"LearnCooking": 100,
	}
	for _, v := range arrServices {
		if val, ok := indiService[v]; ok {
			finalAmount += val
			subscribedServices = append(subscribedServices, v) // Appending channel name
			SubscribedServices(v)                              // Saving all the services which user subscribed to display them later on .
		}
	}
	if userBalance < finalAmount {
		fmt.Println("Insufficient bal")
		return false, nil // Caller which fetch and let user know funds are insufficient.
	}
	finalAmount = finalAmount * (-1)
	//balanceinfo.RechargeTokens.DoRecharge(finalAmount)
	rechargeObject.DoRecharge(finalAmount)

	//fmt.Println("\nIndi Service", indiService, "\nChan", arrServices)

	return true, nil
}
