package auth_mocks

import (
	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) GenerateUserToken(userID int64) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) GenerateAdminToken() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) AuthenticateUser(token string) (int64, error) {
	args := m.Called(token)
	return args.Get(0).(int64), args.Error(1)
}

func (m *AuthServiceMock) AuthenticateAdmin(token string) (bool, error) {
	args := m.Called(token)
	return args.Bool(0), args.Error(1)
}
