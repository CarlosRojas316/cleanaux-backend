// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// GetSecret provides a mock function with given fields:
func (_m *Cache) GetSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetSecret provides a mock function with given fields: _a0
func (_m *Cache) SetSecret(_a0 string) {
	_m.Called(_a0)
}
