// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "taxi_order_service/models"
)

// iLocationService is an autogenerated mock type for the iLocationService type
type iLocationService struct {
	mock.Mock
}

// StoreLocation provides a mock function with given fields: ctx, point
func (_m *iLocationService) StoreLocation(ctx context.Context, point models.Point) error {
	ret := _m.Called(ctx, point)

	if len(ret) == 0 {
		panic("no return value specified for StoreLocation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Point) error); ok {
		r0 = rf(ctx, point)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newILocationService creates a new instance of iLocationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newILocationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *iLocationService {
	mock := &iLocationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
