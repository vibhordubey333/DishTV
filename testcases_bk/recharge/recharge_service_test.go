package recharge

import (
	recharge "DishTV/recharge"
	mocks "DishTV/testcases/recharge/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name string
		want *recharge.RechargeTokens
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recharge.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_CheckBalance(t *testing.T) {

	obj := new(mocks.Recharge)
	obj.On("CheckBalance", mock.AnythingOfType("test")).Return(200)

	obj2 := new(mocks.Recharge)
	obj2.On("CheckBalance", mock.AnythingOfType("Test")).Return(100)

	type fields struct {
		Balance int
		Flag    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Testing",
			fields{200, 1},
			200,
		},
		{
			"Testing",
			fields{100, 1},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rechObj := &recharge.RechargeTokens{
				Balance: tt.fields.Balance,
				Flag:    tt.fields.Flag,
			}
			if got := rechObj.CheckBalance(); got != tt.want {
				t.Errorf("RechargeTokens.CheckBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_DoRecharge(t *testing.T) {

	obj := new(mocks.Recharge)
	successString := "Recharge completed successfully. Current balance is 100"

	obj.On("DoRecharge", mock.AnythingOfType("")).Return(successString)

	type fields struct {
		Balance int
		Flag    int
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"TestCase",
			fields{0, 1},
			args{100},
			successString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rechObj := recharge.RechargeTokens{
				Balance: tt.fields.Balance,
				Flag:    tt.fields.Flag,
			}
			if got := rechObj.DoRecharge(tt.args.amount); got != tt.want {
				t.Errorf("RechargeTokens.DoRecharge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_InternalDoRecharge(t *testing.T) {

	obj := new(mocks.Recharge)
	successString := "Recharge completed successfully. Current balance is 100"
	obj.On("InternalDoRecharge", mock.AnythingOfType("")).Return()

	type fields struct {
		Balance int
		Flag    int
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			"Test",
			fields{0, 1},
			args{100},
			successString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rechObj := recharge.RechargeTokens{
				Balance: tt.fields.Balance,
				Flag:    tt.fields.Flag,
			}
			if got := rechObj.InternalDoRecharge(tt.args.amount); got != tt.want {
				t.Errorf("RechargeTokens.InternalDoRecharge() = %v, want %v", got, tt.want)
			}
		})
	}
}
