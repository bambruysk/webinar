// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Multier is an autogenerated mock type for the Multier type
type Multier struct {
	mock.Mock
}

// Mul2 provides a mock function with given fields: _a0
func (_m *Multier) Mul2(_a0 int) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewMultier interface {
	mock.TestingT
	Cleanup(func())
}

// NewMultier creates a new instance of Multier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMultier(t mockConstructorTestingTNewMultier) *Multier {
	mock := &Multier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
