// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import kurima "github.com/nac-project/kurima"
import mock "github.com/stretchr/testify/mock"

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) GetByEmail(ctx context.Context, email string) (kurima.User, error) {
	ret := _m.Called(ctx, email)

	var r0 kurima.User
	if rf, ok := ret.Get(0).(func(context.Context, string) kurima.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(kurima.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *UserRepository) GetByID(ctx context.Context, ID string) (kurima.User, error) {
	ret := _m.Called(ctx, ID)

	var r0 kurima.User
	if rf, ok := ret.Get(0).(func(context.Context, string) kurima.User); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(kurima.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, user
func (_m *UserRepository) Login(ctx context.Context, user kurima.User) (kurima.User, error) {
	ret := _m.Called(ctx, user)

	var r0 kurima.User
	if rf, ok := ret.Get(0).(func(context.Context, kurima.User) kurima.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(kurima.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kurima.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, user
func (_m *UserRepository) Register(ctx context.Context, user kurima.User) (kurima.User, error) {
	ret := _m.Called(ctx, user)

	var r0 kurima.User
	if rf, ok := ret.Get(0).(func(context.Context, kurima.User) kurima.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(kurima.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kurima.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: ctx, ID, user
func (_m *UserRepository) UpdatePassword(ctx context.Context, ID string, user kurima.User) (kurima.User, error) {
	ret := _m.Called(ctx, ID, user)

	var r0 kurima.User
	if rf, ok := ret.Get(0).(func(context.Context, string, kurima.User) kurima.User); ok {
		r0 = rf(ctx, ID, user)
	} else {
		r0 = ret.Get(0).(kurima.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, kurima.User) error); ok {
		r1 = rf(ctx, ID, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
