package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guregu/null"

	"mnc-stage2/src/data"
	"mnc-stage2/src/service"
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
