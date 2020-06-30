package subscriptioncompute

import (
	compute "DishTV/subscriptioncompute"
	mocks "DishTV/testcases/subscriptioncompute/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetSubscribeList(t *testing.T) {
	listObj := new(mocks.Subscribed)
	objSlice1 := make([]string, 0)
	objSlice2 := make([]string, 0)
	objSlice2 = []string{"NatGeo"}

	listObj.On("GetSubscribeList", mock.AnythingOfType("NatGeo")).Return(objSlice2)
	listObj.On("GetSubscribeList", mock.AnythingOfType("")).Return(objSlice1)

	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
		{
			"Positive",
			objSlice2,
		},
		{
			"Negative",
			objSlice1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute.GetSubscribeList(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscribeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubscribeServiceList(t *testing.T) {

	subObj := new(mocks.Subscribed)
	wantSlice := make([]string, 0)
	want2Slice := make([]string, 0)
	want2Slice = []string{"ABC"}
	subObj.On("SubscribeServiceList", mock.AnythingOfType("LearnEnglish")).Return(wantSlice)
	subObj.On("SubscribeServiceList", mock.AnythingOfType("")).Return(want2Slice)
	tests := []struct {
		name string
		want []string
	}{
		{
			"Negative",
			wantSlice,
		},
		{
			"Positive",
			want2Slice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute.GetSubscribeServiceList(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscribeServiceList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribedChannel(t *testing.T) {
	type args struct {
		str string
	}
	subObj := new(mocks.Subscribed)
	subObj.On("SubscribedChannel", mock.AnythingOfType("NatGeo")).Return(subObj.TestData().Value().IsNil)
	tests := []struct {
		name string
		args args
	}{
		{
			"TestCase",
			args{"NatGeo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compute.SubscribedChannel(tt.args.str)
		})
	}
}

func TestSubscribedServices(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"TestCase",
			args{"LearnEnglish"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compute.SubscribedServices(tt.args.str)
		})
	}
}
