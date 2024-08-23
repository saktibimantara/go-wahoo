// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	go_wahoo "github.com/saktibimantara/go-wahoo"
	mock "github.com/stretchr/testify/mock"
)

// IWahoo is an autogenerated mock type for the IWahoo type
type IWahoo struct {
	mock.Mock
}

// GetAccessToken provides a mock function with given fields: code, uniqueCode
func (_m *IWahoo) GetAccessToken(code string, uniqueCode string) (*go_wahoo.TokenResponse, *go_wahoo.RequestError) {
	ret := _m.Called(code, uniqueCode)

	if len(ret) == 0 {
		panic("no return value specified for GetAccessToken")
	}

	var r0 *go_wahoo.TokenResponse
	var r1 *go_wahoo.RequestError
	if rf, ok := ret.Get(0).(func(string, string) (*go_wahoo.TokenResponse, *go_wahoo.RequestError)); ok {
		return rf(code, uniqueCode)
	}
	if rf, ok := ret.Get(0).(func(string, string) *go_wahoo.TokenResponse); ok {
		r0 = rf(code, uniqueCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_wahoo.TokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *go_wahoo.RequestError); ok {
		r1 = rf(code, uniqueCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*go_wahoo.RequestError)
		}
	}

	return r0, r1
}

// GetAllWorkout provides a mock function with given fields: token, page, limit
func (_m *IWahoo) GetAllWorkout(token string, page int, limit int) (*go_wahoo.WorkoutsResponse, *go_wahoo.RequestError) {
	ret := _m.Called(token, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetAllWorkout")
	}

	var r0 *go_wahoo.WorkoutsResponse
	var r1 *go_wahoo.RequestError
	if rf, ok := ret.Get(0).(func(string, int, int) (*go_wahoo.WorkoutsResponse, *go_wahoo.RequestError)); ok {
		return rf(token, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *go_wahoo.WorkoutsResponse); ok {
		r0 = rf(token, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_wahoo.WorkoutsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) *go_wahoo.RequestError); ok {
		r1 = rf(token, page, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*go_wahoo.RequestError)
		}
	}

	return r0, r1
}

// GetAuthenticateURL provides a mock function with given fields: uniqueCode
func (_m *IWahoo) GetAuthenticateURL(uniqueCode string) (*string, error) {
	ret := _m.Called(uniqueCode)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthenticateURL")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*string, error)); ok {
		return rf(uniqueCode)
	}
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(uniqueCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uniqueCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshToken provides a mock function with given fields: refreshToken, uniqueCode
func (_m *IWahoo) RefreshToken(refreshToken string, uniqueCode string) (*go_wahoo.TokenResponse, *go_wahoo.RequestError) {
	ret := _m.Called(refreshToken, uniqueCode)

	if len(ret) == 0 {
		panic("no return value specified for RefreshToken")
	}

	var r0 *go_wahoo.TokenResponse
	var r1 *go_wahoo.RequestError
	if rf, ok := ret.Get(0).(func(string, string) (*go_wahoo.TokenResponse, *go_wahoo.RequestError)); ok {
		return rf(refreshToken, uniqueCode)
	}
	if rf, ok := ret.Get(0).(func(string, string) *go_wahoo.TokenResponse); ok {
		r0 = rf(refreshToken, uniqueCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_wahoo.TokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *go_wahoo.RequestError); ok {
		r1 = rf(refreshToken, uniqueCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*go_wahoo.RequestError)
		}
	}

	return r0, r1
}

// NewIWahoo creates a new instance of IWahoo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIWahoo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IWahoo {
	mock := &IWahoo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
