package repository

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"mnc-stage2/src/data"
	"mnc-stage2/src/util"
)

type UserRepository interface {
	UpsertUser(ctx context.Context, user *data.User) error
	GetUserByPhoneNumber(ctx context.Context, userID string) (*data.User, error)
	GetUserByID(ctx context.Context, userID string, lock bool) (*data.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UpsertUser(ctx context.Context, user *data.User) error {
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.WithContext(ctx).
		Save(&user).
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
		logrus.WithFields(logrus.Fields{
			"phone_number": phoneNumber,
		}).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID string, lock bool) (*data.User, error) {
	var user data.User
	query := r.db.WithContext(ctx)

	// Apply pessimistic lock only if lock is true
	if lock {
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	if err := query.First(&user, "id = ?", userID).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": userID,
		}).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
