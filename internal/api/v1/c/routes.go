package v1

import (
	controller "SJTU-Canteen-Community/internal/controller/c"
	"SJTU-Canteen-Community/internal/repository"
	"SJTU-Canteen-Community/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCRoutes(r *gin.Engine, db *gorm.DB) {
	c := r.Group("/c")
	auth := c.Group("/auth")
	{
		AuthService := service.NewAuthService(repository.NewUserRepository(db))
		AuthController := controller.NewAuthController(AuthService)
		auth.POST("/login", AuthController.Login)
		auth.GET("/test_login_status", AuthController.TestLoginStatus)
		auth.POST("/logout", AuthController.Logout)
	}
}
