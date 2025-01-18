package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"mnc-stage2/src/data"
	"mnc-stage2/src/repository"
	"mnc-stage2/src/util"
)

type UserService interface {
	RegisterUser(ctx context.Context, firstName, lastName, phoneNumber, address, pin string) (*data.User, error)
	Login(ctx context.Context, phoneNumber, pin string) (*data.LoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, firstName, lastName, phoneNumber, address, pin string) (*data.User, error) {
	existingUser, err := s.userRepo.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("phone number already registered")
	}

	hashedPin, err := util.HashPin(pin)
	if err != nil {
		return nil, fmt.Errorf("failed to hash PIN: %v", err)
	}

	userID := uuid.New()
	user := &data.User{
		ID:          userID,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		Pin:         hashedPin,
		CreatedAt:   time.Now(),
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, phoneNumber, pin string) (*data.LoginResponse, error) {
	user, err := s.userRepo.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil || user == nil {
		msg := "invalid credentials"
		return &data.LoginResponse{
			Status:  "error",
			Message: &msg,
		}, nil
	}

	// Verify PIN
	if !util.VerifyPin(user.Pin, pin) {
		msg := "invalid credentials"
		return &data.LoginResponse{
			Status:  "error",
			Message: &msg,
		}, nil
	}

	// Generate tokens
	accessToken, err := util.GenerateAccessToken(user.ID.String())
	if err != nil {
		msg := "failed to generate access token"
		return &data.LoginResponse{
			Status:  "error",
			Message: &msg,
		}, nil
	}

	refreshToken, err := util.GenerateRefreshToken(user.ID.String())
	if err != nil {
		msg := "failed to generate refresh token"
		return &data.LoginResponse{
			Status:  "error",
			Message: &msg,
		}, nil
	}

	// Create response
	return &data.LoginResponse{
		Status: "success",
		Result: &data.LoginRsp{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
