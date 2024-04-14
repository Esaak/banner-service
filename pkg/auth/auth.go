package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthService is an interface for authentication and authorization
type AuthService interface {
	GenerateUserToken(userID int64) (string, error)
	GenerateAdminToken() (string, error)
	AuthenticateUser(token string) (int64, error)
	AuthenticateAdmin(token string) (bool, error)
}

type authService struct {
	userSecret  string
	adminSecret string
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userSecret, adminSecret string) AuthService {
	return &authService{
		userSecret:  userSecret,
		adminSecret: adminSecret,
	}
}

// GenerateUserToken generates a JWT token for a user
func (a *authService) GenerateUserToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.userSecret))
}

// GenerateAdminToken generates a JWT token for an admin
func (a *authService) GenerateAdminToken() (string, error) {
	claims := jwt.MapClaims{
		"isAdmin": true,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.adminSecret))
}

// AuthenticateUser authenticates a user with a JWT token
func (a *authService) AuthenticateUser(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.userSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(float64)
		if !ok {
			return 0, errors.New("invalid user ID in token claims")
		}
		return int64(userID), nil
	}

	return 0, errors.New("invalid token")
}

// AuthenticateAdmin authenticates an admin with a JWT token
func (a *authService) AuthenticateAdmin(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.adminSecret), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok {
			return false, errors.New("invalid admin claim in token")
		}
		return isAdmin, nil
	}

	return false, errors.New("invalid token")
}
