package src

import (
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"mnc-stage2/src/handler"
)

func RegisterRoutes(r *gin.Engine, userController *handler.UserHandler) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	r.Use(gin.Logger())

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.Login)
	r.POST("/topup", userController.TopUp)
	r.POST("/pay", userController.Payment)
	r.POST("/transfer", userController.Payment)
	r.POST("/transaction", userController.Payment)
}
