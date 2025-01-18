package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"mnc-stage2/src/data"
	"mnc-stage2/src/repository"
)

type UserService interface {
	RegisterUser(ctx context.Context, firstName, lastName, phoneNumber, address, pin string) (*data.User, error)
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

	userID := uuid.New()
	user := &data.User{
		ID:          userID,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		Pin:         pin,
		CreatedAt:   time.Now(),
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
