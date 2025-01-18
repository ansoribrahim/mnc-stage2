package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"mnc-stage2/src/constant"
	"mnc-stage2/src/data"
	"mnc-stage2/src/repository"
	"mnc-stage2/src/util"
)

type UserService interface {
	RegisterUser(ctx context.Context, firstName, lastName, phoneNumber, address, pin string) (*data.User, error)
	Login(ctx context.Context, phoneNumber, pin string) (*data.LoginResponse, error)
	TopUp(ctx context.Context, userID string, amount decimal.Decimal) (*data.TopUpResponse, error)
	Payment(ctx context.Context, userID string, req data.PaymentReq) (*data.PaymentResponse, error)
	Transfer(ctx context.Context, userID string, req data.TransferReq) (*data.TransferResponse, error)
}

type userService struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
	db              *gorm.DB
}

func NewUserService(userRepo repository.UserRepository, transactionRepo repository.TransactionRepository, db *gorm.DB) UserService {
	return &userService{userRepo: userRepo, transactionRepo: transactionRepo, db: db}
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

	err = s.userRepo.UpsertUser(ctx, user)
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

	return &data.LoginResponse{
		Status: "success",
		Result: &data.LoginRsp{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *userService) TopUp(ctx context.Context, userID string, amount decimal.Decimal) (*data.TopUpResponse, error) {

	var err error

	if amount.Compare(decimal.Zero) <= 0 {
		return nil, errors.New("invalid amount")
	}

	tx := s.db.WithContext(ctx).Begin()
	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	user, err := s.userRepo.GetUserByID(ctx, userID, true)
	if err != nil {
		return nil, err
	}

	balanceBefore := user.Balance
	user.Balance = user.Balance.Add(amount)

	if err := s.userRepo.UpsertUser(ctx, user); err != nil {
		return nil, err
	}

	transactionID := uuid.New()
	transaction := &data.Transaction{
		ID:            transactionID,
		UserID:        user.ID,
		Type:          constant.TRANSACTION_TYPE_TOPUP,
		Amount:        amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Balance,
		Remarks:       "",
		Status:        "SUCCESS",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
	}

	if err := s.transactionRepo.InsertTransaction(ctx, transaction); err != nil {
		return nil, err
	}

	return &data.TopUpResponse{
		Status: "SUCCESS",
		Result: &data.TopUpRsp{
			TopUpID:       transaction.ID.String(),
			AmountTopUp:   transaction.Amount,
			BalanceBefore: transaction.BalanceBefore,
			BalanceAfter:  transaction.BalanceAfter,
			CreatedDate:   transaction.CreatedAt.String(),
		},
		Message: nil,
	}, nil
}

func (s *userService) Payment(ctx context.Context, userID string, req data.PaymentReq) (*data.PaymentResponse, error) {

	var err error

	if req.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	tx := s.db.WithContext(ctx).Begin()
	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	user, err := s.userRepo.GetUserByID(ctx, userID, true)
	if err != nil {
		return nil, err
	}

	balanceBefore := user.Balance
	user.Balance = user.Balance.Sub(decimal.NewFromInt(req.Amount))

	if err := s.userRepo.UpsertUser(ctx, user); err != nil {
		return nil, err
	}

	transactionID := uuid.New()
	transaction := &data.Transaction{
		ID:            transactionID,
		UserID:        user.ID,
		Type:          constant.TRANSACTION_TYPE_PAYMENT,
		Amount:        decimal.NewFromInt(req.Amount),
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Balance,
		Remarks:       req.Remarks,
		Status:        "SUCCESS",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
	}

	if err := s.transactionRepo.InsertTransaction(ctx, transaction); err != nil {
		return nil, err
	}

	return &data.PaymentResponse{
		Status: "SUCCESS",
		Result: &data.PaymentRsp{
			PaymentID:     transaction.ID.String(),
			Amount:        decimal.NewFromInt(req.Amount),
			Remarks:       req.Remarks,
			BalanceBefore: transaction.BalanceBefore,
			BalanceAfter:  transaction.BalanceAfter,
			CreatedDate:   transaction.CreatedAt.String(),
		},
		Message: nil,
	}, nil
}

func (s *userService) Transfer(ctx context.Context, userID string, req data.TransferReq) (*data.TransferResponse, error) {

	var err error

	if req.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	tx := s.db.WithContext(ctx).Begin()
	ctx = util.NewTxContext(ctx, tx)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	user, err := s.userRepo.GetUserByID(ctx, userID, true)
	if err != nil {
		return nil, err
	}

	balanceBefore := user.Balance
	user.Balance = user.Balance.Sub(decimal.NewFromInt(req.Amount))

	if err := s.userRepo.UpsertUser(ctx, user); err != nil {
		return nil, err
	}

	transactionID := uuid.New()
	transaction := &data.Transaction{
		ID:            transactionID,
		UserID:        user.ID,
		Type:          constant.TRANSACTION_TYPE_TRANSFER,
		Amount:        decimal.NewFromInt(req.Amount),
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Balance,
		Remarks:       req.Remarks,
		Status:        "SUCCESS",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
	}

	if err := s.transactionRepo.InsertTransaction(ctx, transaction); err != nil {
		return nil, err
	}

	return &data.TransferResponse{
		Status: "SUCCESS",
		Result: &data.TransferRsp{
			TransferID:    transaction.ID.String(),
			Amount:        decimal.NewFromInt(req.Amount),
			Remarks:       req.Remarks,
			BalanceBefore: transaction.BalanceBefore,
			BalanceAfter:  transaction.BalanceAfter,
			CreatedDate:   transaction.CreatedAt.String(),
		},
		Message: nil,
	}, nil
}
