package subscriptioncompute

import (
	compute "DishTV/subscriptioncompute"
	"DishTV/testcases/subscriptioncompute/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestComputeAmount(t *testing.T) {

	obj := new(mocks.Subscription)
	obj.On("ComputeAmount", mock.AnythingOfType("G"), mock.AnythingOfType("1")).Return(100, 0, 100, 0, 100, "GOLD", true, nil)
	type args struct {
		pack   string
		months int
	}
	tests := []struct {
		name                 string
		args                 args
		wantSubAmount        int
		wantDiscount         int
		wantFinalAmount      int
		wantRemainingBalance int
		wantMonthlyPrice     int
		wantPackType         string
		wantStatus           bool
		wantErr              bool
	}{
		{
			"TestCase",
			args{"G", 1},
			100,
			0,
			100,
			0,
			100,
			"GOLD",
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSubAmount, gotDiscount, gotFinalAmount, gotRemainingBalance, gotMonthlyPrice, gotPackType, gotStatus, err := compute.ComputeAmount(tt.args.pack, tt.args.months)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSubAmount != tt.wantSubAmount {
				t.Errorf("ComputeAmount() gotSubAmount = %v, want %v", gotSubAmount, tt.wantSubAmount)
			}
			if gotDiscount != tt.wantDiscount {
				t.Errorf("ComputeAmount() gotDiscount = %v, want %v", gotDiscount, tt.wantDiscount)
			}
			if gotFinalAmount != tt.wantFinalAmount {
				t.Errorf("ComputeAmount() gotFinalAmount = %v, want %v", gotFinalAmount, tt.wantFinalAmount)
			}
			if gotRemainingBalance != tt.wantRemainingBalance {
				t.Errorf("ComputeAmount() gotRemainingBalance = %v, want %v", gotRemainingBalance, tt.wantRemainingBalance)
			}
			if gotMonthlyPrice != tt.wantMonthlyPrice {
				t.Errorf("ComputeAmount() gotMonthlyPrice = %v, want %v", gotMonthlyPrice, tt.wantMonthlyPrice)
			}
			if gotPackType != tt.wantPackType {
				t.Errorf("ComputeAmount() gotPackType = %v, want %v", gotPackType, tt.wantPackType)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("ComputeAmount() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestDiscount(t *testing.T) {
	type args struct {
		months int
		pack   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute.Discount(tt.args.months, tt.args.pack); got != tt.want {
				t.Errorf("Discount() = %v, want %v", got, tt.want)
			}
		})
	}
}
