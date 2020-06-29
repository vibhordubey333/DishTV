package subscriptioncompute

import (
	_ "fmt"
)

/*
var subscribedObj *ExistingMembers

type ExistingMembers struct {
	subscribed []string
	flag       int
}

func New() []string {
	subSlice := make([]string, 0)
	return subSlice
}
func init() {
	subscribedObj := New()
}
*/
type AlreadySubscribed interface {
	AlreadySubscribedChannel() string
}

var existing []string

func AlreadySubscribedChannel(str string) {
	existing = append(existing, str)
}
func GetSubscribeList() []string {
	return existing
}
