// Code generated by mockery v1.1.0. DO NOT EDIT.

package mocks

import (
	service "github.com/ayufan/golang-kardianos-service"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Start provides a mock function with given fields: s
func (_m *Interface) Start(s service.Service) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(service.Service) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: s
func (_m *Interface) Stop(s service.Service) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(service.Service) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
