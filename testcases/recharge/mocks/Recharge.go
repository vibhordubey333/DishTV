// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	//recharge "DishTV/testcases/recharge"

	"DishTV/recharge"

	mock "github.com/stretchr/testify/mock"
)

// Recharge is an autogenerated mock type for the Recharge type
type Recharge struct {
	mock.Mock
}

// CheckBalance provides a mock function with given fields:
func (_m *Recharge) CheckBalance() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DoRecharge provides a mock function with given fields: _a0
func (_m *Recharge) DoRecharge(_a0 int) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InternalDoRecharge provides a mock function with given fields: _a0
func (_m *Recharge) InternalDoRecharge(_a0 int) {
	_m.Called(_a0)
}

// New provides a mock function with given fields:
func (_m *Recharge) New() *recharge.RechargeTokens {
	ret := _m.Called()

	var r0 *recharge.RechargeTokens
	if rf, ok := ret.Get(0).(func() *recharge.RechargeTokens); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recharge.RechargeTokens)
		}
	}

	return r0
}
