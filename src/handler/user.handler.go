package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"mnc-stage2/src/data"
	"mnc-stage2/src/service"
	"mnc-stage2/src/util"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uc *UserHandler) RegisterUser(c *gin.Context) {
	var req = data.RegisterReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response := data.RegisterResponse{
			Status:  "FAILED",
			Message: null.StringFrom("Invalid input").Ptr(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ctx := context.Background()
	resp, err := uc.userService.RegisterUser(ctx, req.FirstName, req.LastName, req.PhoneNumber, req.Address, req.Pin)
	if err != nil {
		if err.Error() == "phone number already registered" {
			response := data.RegisterResponse{
				Status:  "FAILED",
				Message: null.StringFrom("Phone number already registered").Ptr(),
			}
			c.JSON(http.StatusConflict, response)
			return
		}
		response := data.RegisterResponse{
			Status:  "FAILED",
			Message: null.StringFrom("Internal server error").Ptr(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := data.RegisterResponse{
		Status: "SUCCESS",
		Result: &data.UserResponse{
			ID:          resp.ID,
			FirstName:   resp.FirstName,
			LastName:    resp.LastName,
			PhoneNumber: resp.PhoneNumber,
			Address:     resp.Address,
			CreatedAt:   resp.CreatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (uc *UserHandler) Login(c *gin.Context) {
	var req data.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid input"})
		return
	}

	ctx := context.Background()
	resp, err := uc.userService.Login(ctx, req.PhoneNumber, req.Pin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Phone Number and PIN doesn’t match."})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (uc *UserHandler) TopUp(c *gin.Context) {
	var req data.TopUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid input"})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	claims, err := util.GetClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	ctx := context.Background()
	topUpResp, err := uc.userService.TopUp(ctx, claims.UserID, decimal.NewFromInt(req.Amount))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, topUpResp)
}

func (uc *UserHandler) Payment(c *gin.Context) {
	var req data.PaymentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid input"})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	claims, err := util.GetClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	ctx := context.Background()
	payResp, err := uc.userService.Payment(ctx, claims.UserID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, payResp)
}

func (uc *UserHandler) Transfer(c *gin.Context) {
	var req data.TransferReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid input"})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	claims, err := util.GetClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	ctx := context.Background()
	trfResp, err := uc.userService.Transfer(ctx, claims.UserID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, trfResp)
}

func (uc *UserHandler) TransactionReports(c *gin.Context) {
	var req data.TopUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid input"})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	claims, err := util.GetClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILED", "message": "Unauthenticated"})
		return
	}

	ctx := context.Background()
	topUpResp, err := uc.userService.TopUp(ctx, claims.UserID, decimal.NewFromInt(req.Amount))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, topUpResp)
}
