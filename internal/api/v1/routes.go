package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// user := r.Group("/user")
	// {
	// 	UserService := service.NewUserService(repository.NewUserRepository(db))
	// 	UserController := controller.NewUserController(UserService)
	// }
}
