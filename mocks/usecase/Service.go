// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	zerolog "github.com/rs/zerolog"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Load provides a mock function with given fields: _a0, _a1
func (_m *Service) Load(_a0 interface{}, _a1 *zerolog.Logger) {
	_m.Called(_a0, _a1)
}
