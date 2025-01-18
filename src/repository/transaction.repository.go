package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"mnc-stage2/src/data"
	"mnc-stage2/src/util"
)

type TransactionRepository interface {
	InsertTransaction(ctx context.Context, transaction *data.Transaction) error
	GetTransactionsByUserID(ctx context.Context, userID string) ([]data.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) InsertTransaction(ctx context.Context, transaction *data.Transaction) error {
	tx := util.GetTxFromContext(ctx, r.db)

	err := tx.WithContext(ctx).
		Create(&transaction).
		Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"transaction": util.Dump(transaction),
		}).Error(err)
		return err
	}

	return nil
}

func (r *transactionRepository) GetTransactionsByUserID(ctx context.Context, userID string) ([]data.Transaction, error) {
	var transactions []data.Transaction
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": userID,
		}).Error(err)
		return nil, err
	}
	return transactions, nil
}
