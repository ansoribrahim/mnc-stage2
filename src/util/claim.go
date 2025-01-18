package util

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims structure for access and refresh tokens
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Helper to get environment variable or return a default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GenerateAccessToken creates a JWT access token
func GenerateAccessToken(userID string) (string, error) {
	accessTokenSecret := getEnv("ACCESS_TOKEN_SECRET", "")
	if accessTokenSecret == "" {
		return "", errors.New("access token secret is not set in environment variables")
	}

	expirationTime := time.Now().Add(15 * time.Minute) // Access token valid for 15 minutes
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessTokenSecret))
}

// GenerateRefreshToken creates a JWT refresh token
func GenerateRefreshToken(userID string) (string, error) {
	refreshTokenSecret := getEnv("REFRESH_TOKEN_SECRET", "")
	if refreshTokenSecret == "" {
		return "", errors.New("refresh token secret is not set in environment variables")
	}

	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Refresh token valid for 7 days
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshTokenSecret))
}

func GetClaims(tokenString string) (*Claims, error) {
	// Extract the token part after "Bearer " if it's present
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return nil, errors.New("invalid token format")
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Define a function to parse and validate the JWT token
	accessTokenSecret := getEnv("ACCESS_TOKEN_SECRET", "")
	if accessTokenSecret == "" {
		return nil, errors.New("access token secret is not set in environment variables")
	}

	// Parse the token and validate claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(accessTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the extracted claims
	return claims, nil
}
