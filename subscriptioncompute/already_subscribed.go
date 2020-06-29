package subscriptioncompute

import (
	_ "fmt"
)

type Subscribed interface {
	SubscribedChannel(string)
	SubscribedServices(string)
}

var existing []string
var existingService []string

func SubscribedChannel(str string) { // To avoid the redundent entry placing a check.

	for _, v := range existing {
		if v == str {
			return
		}
	}

	existing = append(existing, str)
}
func SubscribedServices(str string) { // To avoid the redundent entry placing a check.

	for _, v := range existingService {
		if v == str {
			return
		}
	}
	existingService = append(existingService, str)
}
func GetSubscribeList() []string {
	return existing
}
func GetSubscribeServiceList() []string {
	return existingService
}
