// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Environment is an autogenerated mock type for the Environment type
type Environment struct {
	mock.Mock
}

// export provides a mock function with given fields:
func (_m *Environment) export() {
	_m.Called()
}

// load provides a mock function with given fields:
func (_m *Environment) load() {
	_m.Called()
}

// must provides a mock function with given fields:
func (_m *Environment) must() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
