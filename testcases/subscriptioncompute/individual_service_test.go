package subscriptioncompute

import (
	compute "DishTV/subscriptioncompute"

	mocks "DishTV/testcases/subscriptioncompute/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_panicHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compute.PanicHandler()
		})
	}
}

func TestSubscribeChannel(t *testing.T) {

	obj := new(mocks.IndividualService)
	obj.On("SubscribeChannel", mock.AnythingOfType("NatGeo")).Return(true, nil)
	obj.On("SubscribeChannel", mock.AnythingOfType("")).Return(false, errors.New("Invalid Input."))
	type args struct {
		channels string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus bool
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			"Positive",
			args{"NatGeo"},
			true,
			false,
		},
		{
			"Negative",
			args{""},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, err := compute.SubscribeChannel(tt.args.channels)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscribeChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("SubscribeChannel() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestSubscribeService(t *testing.T) {

	obj := new(mocks.IndividualService)

	obj.On("SubscribeService", mock.AnythingOfType("LearnEnglish")).Return(true, nil)
	obj.On("SubscribeService", mock.AnythingOfType("LearnEnglish")).Return(true, errors.New("Invalid Input."))
	type args struct {
		services string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus bool
		wantErr    bool
	}{
		{
			"Positive",
			args{"LearnEnglish"},
			false,
			false,
		},
		{
			"Negative",
			args{""},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, err := compute.SubscribeService(tt.args.services)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscribeService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("SubscribeService() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
