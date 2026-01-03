package service

import (
	"errors"
	"fmt"
	"healtech-backend/server/internal/config"
	"healtech-backend/server/internal/models"
	"healtech-backend/server/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Config *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{Config: cfg}
}
func (s *AuthService) LoginMiddleware(username, password string) (string, error) {
	// 1. Get User
	// /fmt.Println("-->", username)
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	fmt.Println("-->", user)
	// 2. Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		// fmt.Println("thiserror")
		return "", errors.New("invalid credentials")
	}

	// 3. Generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.RegisterID,
		"role_id": user.RoleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.Config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) RegisterMiddleware(username, password string, registerId, roleId int) (string, error) {
	// 1. Get User
	// /fmt.Println("-->", username)
	status, err := repository.RegisterUser(username, password, registerId, roleId)
	if err != nil {
		// loggermdl.LogError(time.Now().Format("2006-01-02 15:04:00") + "::" + loggermdl.GetCallers(1) + "::" + err.Error())
		return "", errors.New("Error Occurred")
	}
	return status, nil
}

func (s *AuthService) ValidateToken(tokenStr string) (*models.RegisterDetails, error) {
	// Simple validation returning a stub user object with ID derived from token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("invalid user_id in token")
	}

	return &models.RegisterDetails{ID: int(userIDFloat)}, nil
}
