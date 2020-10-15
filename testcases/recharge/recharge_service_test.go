package recharge

import (
	"github.com/stretchr/testify/mock"
	//	recharge "DishTV/recharge"
	"DishTV/recharge"
	"DishTV/testcases/recharge/mocks"

	//"DishTV/testcases/recharge/mocks"

	"reflect"
	"testing"
)

func TestRechargeTokens_New(t *testing.T) {

	obj1 := new(mocks.Recharge)
	obj1.On("New", mock.AnythingOfType("")).Return(recharge.RechargeTokens{200, 1})
	//rechargeTokObj := new(recharge.RechargeTokens)
	//rechargeObject := rechargeTokObj.New

	//obj1 = obj1.New
	tests := []struct {
		name    string
		rechObj *recharge.RechargeTokens
		want    *recharge.RechargeTokens
	}{
		{
			"Negative testcase",
			&recharge.RechargeTokens{200, -1},
			&recharge.RechargeTokens{100, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//	if got := tt.rechObj.New(); !reflect.DeepEqual(got, tt.want) {
			// if got := obj1.New(); !reflect.DeepEqual(got, tt.want) { Giving error.
			if got := tt.rechObj.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RechargeTokens.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_CheckBalance(t *testing.T) {
	tests := []struct {
		name    string
		rechObj *recharge.RechargeTokens
		want    int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rechObj.CheckBalance(); got != tt.want {
				t.Errorf("RechargeTokens.CheckBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_DoRecharge(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		rechObj *recharge.RechargeTokens
		args    args
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rechObj.DoRecharge(tt.args.amount); got != tt.want {
				t.Errorf("RechargeTokens.DoRecharge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_InternalDoRecharge(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		rechObj *recharge.RechargeTokens
		args    args
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rechObj.InternalDoRecharge(tt.args.amount); got != tt.want {
				t.Errorf("RechargeTokens.InternalDoRecharge() = %v, want %v", got, tt.want)
			}
		})
	}
}
