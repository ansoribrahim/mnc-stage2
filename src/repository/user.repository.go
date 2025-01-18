package repository

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"mnc-stage2/src/data"
	"mnc-stage2/src/util"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *data.User) error
	GetUserByPhoneNumber(ctx context.Context, userID string) (*data.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *data.User) error {
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.WithContext(ctx).
		Create(&user).
		Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user": util.Dump(user),
		}).Error(err)
		return err
	}

	return nil
}

func (r *userRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*data.User, error) {
	var user data.User
	if err := r.db.WithContext(ctx).First(&user, "phone_number = ?", phoneNumber).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
