package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"mnc-stage2/mock"
	"mnc-stage2/src/data"
)

func TestUserService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	mockTransactionRepo := mock.NewMockTransactionRepository(ctrl)
	mockDB := &gorm.DB{}

	userService := NewUserService(mockUserRepo, mockTransactionRepo, mockDB)

	type args struct {
		ctx         context.Context
		firstName   string
		lastName    string
		phoneNumber string
		address     string
		pin         string
	}
	tests := []struct {
		name         string
		args         args
		mockSetup    func()
		expectedUser *data.User
		expectedErr  error
	}{
		{
			name: "success_register_user",
			args: args{
				ctx:         context.Background(),
				firstName:   "John",
				lastName:    "Doe",
				phoneNumber: "1234567890",
				address:     "123 Elm Street",
				pin:         "1234",
			},
			mockSetup: func() {
				mockUserRepo.EXPECT().
					GetUserByPhoneNumber(gomock.Any(), "1234567890").
					Return(nil, nil)
				mockUserRepo.EXPECT().
					UpsertUser(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedUser: &data.User{
				FirstName:   "John",
				LastName:    "Doe",
				PhoneNumber: "1234567890",
				Address:     "123 Elm Street",
			},
			expectedErr: nil,
		},
		{
			name: "error_phone_number_already_registered",
			args: args{
				ctx:         context.Background(),
				firstName:   "Jane",
				lastName:    "Doe",
				phoneNumber: "0987654321",
				address:     "456 Oak Street",
				pin:         "5678",
			},
			mockSetup: func() {
				mockUserRepo.EXPECT().
					GetUserByPhoneNumber(gomock.Any(), "0987654321").
					Return(&data.User{
						ID: uuid.New(),
					}, nil)
			},
			expectedUser: nil,
			expectedErr:  errors.New("phone number already registered"),
		},
		{
			name: "error_upsert_user_failure",
			args: args{
				ctx:         context.Background(),
				firstName:   "Alice",
				lastName:    "Smith",
				phoneNumber: "9876543210",
				address:     "789 Pine Avenue",
				pin:         "7890",
			},
			mockSetup: func() {
				mockUserRepo.EXPECT().
					GetUserByPhoneNumber(gomock.Any(), "9876543210").
					Return(nil, nil)
				mockUserRepo.EXPECT().
					UpsertUser(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedUser: nil,
			expectedErr:  errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			user, err := userService.RegisterUser(tt.args.ctx, tt.args.firstName, tt.args.lastName, tt.args.phoneNumber, tt.args.address, tt.args.pin)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.expectedUser.FirstName, user.FirstName)
				assert.Equal(t, tt.expectedUser.LastName, user.LastName)
				assert.Equal(t, tt.expectedUser.PhoneNumber, user.PhoneNumber)
				assert.Equal(t, tt.expectedUser.Address, user.Address)
			}
		})
	}
}
