// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	time "time"
)

// AuthRepo is an autogenerated mock type for the AuthRepo type
type AuthRepo struct {
	mock.Mock
}

type AuthRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthRepo) EXPECT() *AuthRepo_Expecter {
	return &AuthRepo_Expecter{mock: &_m.Mock}
}

// CreateAccessToken provides a mock function with given fields: ctx, userID, name, expiration
func (_m *AuthRepo) CreateAccessToken(ctx context.Context, userID model.UserID, name string, expiration time.Duration) (string, error) {
	ret := _m.Called(ctx, userID, name, expiration)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, string, time.Duration) string); ok {
		r0 = rf(ctx, userID, name, expiration)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, string, time.Duration) error); ok {
		r1 = rf(ctx, userID, name, expiration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_CreateAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccessToken'
type AuthRepo_CreateAccessToken_Call struct {
	*mock.Call
}

// CreateAccessToken is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
//  - name string
//  - expiration time.Duration
func (_e *AuthRepo_Expecter) CreateAccessToken(ctx any, userID any, name any, expiration any) *AuthRepo_CreateAccessToken_Call {
	return &AuthRepo_CreateAccessToken_Call{Call: _e.mock.On("CreateAccessToken", ctx, userID, name, expiration)}
}

func (_c *AuthRepo_CreateAccessToken_Call) Run(run func(ctx context.Context, userID model.UserID, name string, expiration time.Duration)) *AuthRepo_CreateAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(string), args[3].(time.Duration))
	})
	return _c
}

func (_c *AuthRepo_CreateAccessToken_Call) Return(token string, err error) *AuthRepo_CreateAccessToken_Call {
	_c.Call.Return(token, err)
	return _c
}

// DeleteAccessToken provides a mock function with given fields: ctx, tokenID
func (_m *AuthRepo) DeleteAccessToken(ctx context.Context, tokenID uint32) (bool, error) {
	ret := _m.Called(ctx, tokenID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint32) bool); ok {
		r0 = rf(ctx, tokenID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, tokenID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_DeleteAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccessToken'
type AuthRepo_DeleteAccessToken_Call struct {
	*mock.Call
}

// DeleteAccessToken is a helper method to define mock.On call
//  - ctx context.Context
//  - tokenID uint32
func (_e *AuthRepo_Expecter) DeleteAccessToken(ctx any, tokenID any) *AuthRepo_DeleteAccessToken_Call {
	return &AuthRepo_DeleteAccessToken_Call{Call: _e.mock.On("DeleteAccessToken", ctx, tokenID)}
}

func (_c *AuthRepo_DeleteAccessToken_Call) Run(run func(ctx context.Context, tokenID uint32)) *AuthRepo_DeleteAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *AuthRepo_DeleteAccessToken_Call) Return(_a0 bool, _a1 error) *AuthRepo_DeleteAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *AuthRepo) GetByEmail(ctx context.Context, email string) (domain.Auth, []byte, error) {
	ret := _m.Called(ctx, email)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Auth); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AuthRepo_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type AuthRepo_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//  - ctx context.Context
//  - email string
func (_e *AuthRepo_Expecter) GetByEmail(ctx any, email any) *AuthRepo_GetByEmail_Call {
	return &AuthRepo_GetByEmail_Call{Call: _e.mock.On("GetByEmail", ctx, email)}
}

func (_c *AuthRepo_GetByEmail_Call) Run(run func(ctx context.Context, email string)) *AuthRepo_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthRepo_GetByEmail_Call) Return(_a0 domain.Auth, _a1 []byte, _a2 error) *AuthRepo_GetByEmail_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// GetByToken provides a mock function with given fields: ctx, token
func (_m *AuthRepo) GetByToken(ctx context.Context, token string) (domain.Auth, error) {
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

// AuthRepo_GetByToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByToken'
type AuthRepo_GetByToken_Call struct {
	*mock.Call
}

// GetByToken is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
func (_e *AuthRepo_Expecter) GetByToken(ctx any, token any) *AuthRepo_GetByToken_Call {
	return &AuthRepo_GetByToken_Call{Call: _e.mock.On("GetByToken", ctx, token)}
}

func (_c *AuthRepo_GetByToken_Call) Run(run func(ctx context.Context, token string)) *AuthRepo_GetByToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthRepo_GetByToken_Call) Return(_a0 domain.Auth, _a1 error) *AuthRepo_GetByToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPermission provides a mock function with given fields: ctx, groupID
func (_m *AuthRepo) GetPermission(ctx context.Context, groupID uint8) (domain.Permission, error) {
	ret := _m.Called(ctx, groupID)

	var r0 domain.Permission
	if rf, ok := ret.Get(0).(func(context.Context, uint8) domain.Permission); ok {
		r0 = rf(ctx, groupID)
	} else {
		r0 = ret.Get(0).(domain.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint8) error); ok {
		r1 = rf(ctx, groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_GetPermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermission'
type AuthRepo_GetPermission_Call struct {
	*mock.Call
}

// GetPermission is a helper method to define mock.On call
//  - ctx context.Context
//  - groupID uint8
func (_e *AuthRepo_Expecter) GetPermission(ctx any, groupID any) *AuthRepo_GetPermission_Call {
	return &AuthRepo_GetPermission_Call{Call: _e.mock.On("GetPermission", ctx, groupID)}
}

func (_c *AuthRepo_GetPermission_Call) Run(run func(ctx context.Context, groupID uint8)) *AuthRepo_GetPermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint8))
	})
	return _c
}

func (_c *AuthRepo_GetPermission_Call) Return(_a0 domain.Permission, _a1 error) *AuthRepo_GetPermission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTokenByID provides a mock function with given fields: ctx, id
func (_m *AuthRepo) GetTokenByID(ctx context.Context, id uint32) (domain.AccessToken, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.AccessToken
	if rf, ok := ret.Get(0).(func(context.Context, uint32) domain.AccessToken); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.AccessToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_GetTokenByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenByID'
type AuthRepo_GetTokenByID_Call struct {
	*mock.Call
}

// GetTokenByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *AuthRepo_Expecter) GetTokenByID(ctx any, id any) *AuthRepo_GetTokenByID_Call {
	return &AuthRepo_GetTokenByID_Call{Call: _e.mock.On("GetTokenByID", ctx, id)}
}

func (_c *AuthRepo_GetTokenByID_Call) Run(run func(ctx context.Context, id uint32)) *AuthRepo_GetTokenByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *AuthRepo_GetTokenByID_Call) Return(_a0 domain.AccessToken, _a1 error) *AuthRepo_GetTokenByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListAccessToken provides a mock function with given fields: ctx, userID
func (_m *AuthRepo) ListAccessToken(ctx context.Context, userID model.UserID) ([]domain.AccessToken, error) {
	ret := _m.Called(ctx, userID)

	var r0 []domain.AccessToken
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID) []domain.AccessToken); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.AccessToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_ListAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccessToken'
type AuthRepo_ListAccessToken_Call struct {
	*mock.Call
}

// ListAccessToken is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
func (_e *AuthRepo_Expecter) ListAccessToken(ctx any, userID any) *AuthRepo_ListAccessToken_Call {
	return &AuthRepo_ListAccessToken_Call{Call: _e.mock.On("ListAccessToken", ctx, userID)}
}

func (_c *AuthRepo_ListAccessToken_Call) Run(run func(ctx context.Context, userID model.UserID)) *AuthRepo_ListAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID))
	})
	return _c
}

func (_c *AuthRepo_ListAccessToken_Call) Return(_a0 []domain.AccessToken, _a1 error) *AuthRepo_ListAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewAuthRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepo creates a new instance of AuthRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepo(t mockConstructorTestingTNewAuthRepo) *AuthRepo {
	mock := &AuthRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
