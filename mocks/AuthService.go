// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

type AuthService_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthService) EXPECT() *AuthService_Expecter {
	return &AuthService_Expecter{mock: &_m.Mock}
}

// ComparePassword provides a mock function with given fields: hashed, password
func (_m *AuthService) ComparePassword(hashed []byte, password string) (bool, error) {
	ret := _m.Called(hashed, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte, string) bool); ok {
		r0 = rf(hashed, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(hashed, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_ComparePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ComparePassword'
type AuthService_ComparePassword_Call struct {
	*mock.Call
}

// ComparePassword is a helper method to define mock.On call
//  - hashed []byte
//  - password string
func (_e *AuthService_Expecter) ComparePassword(hashed interface{}, password interface{}) *AuthService_ComparePassword_Call {
	return &AuthService_ComparePassword_Call{Call: _e.mock.On("ComparePassword", hashed, password)}
}

func (_c *AuthService_ComparePassword_Call) Run(run func(hashed []byte, password string)) *AuthService_ComparePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(string))
	})
	return _c
}

func (_c *AuthService_ComparePassword_Call) Return(_a0 bool, _a1 error) *AuthService_ComparePassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByToken provides a mock function with given fields: ctx, token
func (_m *AuthService) GetByToken(ctx context.Context, token string) (domain.Auth, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Auth); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_GetByToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByToken'
type AuthService_GetByToken_Call struct {
	*mock.Call
}

// GetByToken is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
func (_e *AuthService_Expecter) GetByToken(ctx interface{}, token interface{}) *AuthService_GetByToken_Call {
	return &AuthService_GetByToken_Call{Call: _e.mock.On("GetByToken", ctx, token)}
}

func (_c *AuthService_GetByToken_Call) Run(run func(ctx context.Context, token string)) *AuthService_GetByToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthService_GetByToken_Call) Return(_a0 domain.Auth, _a1 error) *AuthService_GetByToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByTokenWithCache provides a mock function with given fields: ctx, token
func (_m *AuthService) GetByTokenWithCache(ctx context.Context, token string) (domain.Auth, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Auth); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_GetByTokenWithCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByTokenWithCache'
type AuthService_GetByTokenWithCache_Call struct {
	*mock.Call
}

// GetByTokenWithCache is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
func (_e *AuthService_Expecter) GetByTokenWithCache(ctx interface{}, token interface{}) *AuthService_GetByTokenWithCache_Call {
	return &AuthService_GetByTokenWithCache_Call{Call: _e.mock.On("GetByTokenWithCache", ctx, token)}
}

func (_c *AuthService_GetByTokenWithCache_Call) Run(run func(ctx context.Context, token string)) *AuthService_GetByTokenWithCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthService_GetByTokenWithCache_Call) Return(_a0 domain.Auth, _a1 error) *AuthService_GetByTokenWithCache_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPermission provides a mock function with given fields: ctx, id
func (_m *AuthService) GetPermission(ctx context.Context, id uint8) (domain.Permission, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Permission
	if rf, ok := ret.Get(0).(func(context.Context, uint8) domain.Permission); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint8) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthService_GetPermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermission'
type AuthService_GetPermission_Call struct {
	*mock.Call
}

// GetPermission is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint8
func (_e *AuthService_Expecter) GetPermission(ctx interface{}, id interface{}) *AuthService_GetPermission_Call {
	return &AuthService_GetPermission_Call{Call: _e.mock.On("GetPermission", ctx, id)}
}

func (_c *AuthService_GetPermission_Call) Run(run func(ctx context.Context, id uint8)) *AuthService_GetPermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint8))
	})
	return _c
}

func (_c *AuthService_GetPermission_Call) Return(_a0 domain.Permission, _a1 error) *AuthService_GetPermission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *AuthService) Login(ctx context.Context, email string, password string) (domain.Auth, bool, error) {
	ret := _m.Called(ctx, email, password)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string, string) domain.Auth); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, string, string) bool); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AuthService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AuthService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//  - ctx context.Context
//  - email string
//  - password string
func (_e *AuthService_Expecter) Login(ctx interface{}, email interface{}, password interface{}) *AuthService_Login_Call {
	return &AuthService_Login_Call{Call: _e.mock.On("Login", ctx, email, password)}
}

func (_c *AuthService_Login_Call) Run(run func(ctx context.Context, email string, password string)) *AuthService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AuthService_Login_Call) Return(_a0 domain.Auth, _a1 bool, _a2 error) *AuthService_Login_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}