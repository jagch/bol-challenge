// Code generated by mockery v2.20.0. DO NOT EDIT.

package usecasemocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CallhistoryUsecase is an autogenerated mock type for the CallhistoryUsecase type
type CallhistoryUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, statusCode, duration, url
func (_m *CallhistoryUsecase) Create(ctx context.Context, statusCode int, duration float64, url string) error {
	ret := _m.Called(ctx, statusCode, duration, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, float64, string) error); ok {
		r0 = rf(ctx, statusCode, duration, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCallhistoryUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewCallhistoryUsecase creates a new instance of CallhistoryUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCallhistoryUsecase(t mockConstructorTestingTNewCallhistoryUsecase) *CallhistoryUsecase {
	mock := &CallhistoryUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
