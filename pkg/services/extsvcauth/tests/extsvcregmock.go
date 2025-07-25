// Code generated by mockery v2.53.4. DO NOT EDIT.

package tests

import (
	context "context"

	extsvcauth "github.com/grafana/grafana/pkg/services/extsvcauth"
	mock "github.com/stretchr/testify/mock"
)

// ExternalServiceRegistryMock is an autogenerated mock type for the ExternalServiceRegistry type
type ExternalServiceRegistryMock struct {
	mock.Mock
}

// GetExternalServiceNames provides a mock function with given fields: ctx
func (_m *ExternalServiceRegistryMock) GetExternalServiceNames(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetExternalServiceNames")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasExternalService provides a mock function with given fields: ctx, name
func (_m *ExternalServiceRegistryMock) HasExternalService(ctx context.Context, name string) (bool, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for HasExternalService")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveExternalService provides a mock function with given fields: ctx, name
func (_m *ExternalServiceRegistryMock) RemoveExternalService(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for RemoveExternalService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveExternalService provides a mock function with given fields: ctx, cmd
func (_m *ExternalServiceRegistryMock) SaveExternalService(ctx context.Context, cmd *extsvcauth.ExternalServiceRegistration) (*extsvcauth.ExternalService, error) {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for SaveExternalService")
	}

	var r0 *extsvcauth.ExternalService
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *extsvcauth.ExternalServiceRegistration) (*extsvcauth.ExternalService, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *extsvcauth.ExternalServiceRegistration) *extsvcauth.ExternalService); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*extsvcauth.ExternalService)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *extsvcauth.ExternalServiceRegistration) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExternalServiceRegistryMock creates a new instance of ExternalServiceRegistryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExternalServiceRegistryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExternalServiceRegistryMock {
	mock := &ExternalServiceRegistryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
