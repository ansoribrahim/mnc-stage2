package src

import (
	"github.com/gin-gonic/gin"

	"mnc-stage2/src/handler"
)

func RegisterRoutes(r *gin.Engine, userController *handler.UserHandler) {
	r.POST("/register", userController.RegisterUser)
}
