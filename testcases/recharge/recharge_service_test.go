package recharge

import (
	"reflect"
	"testing"
)

func TestRechargeTokens_New(t *testing.T) {
	tests := []struct {
		name    string
		rechObj *RechargeTokens
		want    *RechargeTokens
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rechObj.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RechargeTokens.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRechargeTokens_CheckBalance(t *testing.T) {
	tests := []struct {
		name    string
		rechObj *RechargeTokens
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
		rechObj *RechargeTokens
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
		rechObj *RechargeTokens
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
