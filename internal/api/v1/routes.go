package v1

import (
	controller_b "SJTU-Canteen-Community/internal/controller/b"
	controller_c "SJTU-Canteen-Community/internal/controller/c"
	"SJTU-Canteen-Community/internal/repository"
	"SJTU-Canteen-Community/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var canteenRepo *repository.CanteenRepository
var windowRepo *repository.WindowRepository
var dishRepo *repository.DishRepository
var windowDishRelationRepo *repository.WindowDishRelationRepository
var userRepo *repository.UserRepository
var canteenCommentRepo *repository.CanteenCommentRepository

func initRepos(db *gorm.DB) {
	canteenRepo = repository.NewCanteenRepository(db)
	windowRepo = repository.NewWindowRepository(db)
	dishRepo = repository.NewDishRepository(db)
	windowDishRelationRepo = repository.NewWindowDishRelationRepository(db)
	userRepo = repository.NewUserRepository(db)
	canteenCommentRepo = repository.NewCanteenCommentRepository(db)

}

func setupBRoutes(r *gin.Engine) {
	b := r.Group("/b")
	canteen := b.Group("/canteens")

	{
		CanteenService := service.NewCanteenService(canteenRepo, windowRepo, windowDishRelationRepo, dishRepo)
		CanteenController := controller_b.NewCanteenController(CanteenService)
		canteen.POST("/", CanteenController.AddCanteen)
		canteen.POST("/:canteen_id/windows", CanteenController.MAddWindowsToCanteen)
		canteen.POST("/windows/:window_id/dishes", CanteenController.MAddDishesToWindow)
	}

	dish := b.Group("/dishes")
	{
		DishService := service.NewDishService(dishRepo)
		DishController := controller_b.NewDishController(DishService)
		dish.POST("/", DishController.MAddDishes)
	}
}

func setupCRoutes(r *gin.Engine) {
	c := r.Group("/c")
	auth := c.Group("/auth")
	{
		AuthService := service.NewAuthService(userRepo)
		AuthController := controller_c.NewAuthController(AuthService)
		auth.POST("/login", AuthController.Login)
		auth.GET("/test_login_status", AuthController.TestLoginStatus)
		auth.POST("/logout", AuthController.Logout)
	}

	CommentService := service.NewCommentService(canteenCommentRepo)
	canteen := c.Group("/canteens")
	{
		CanteenService := service.NewCanteenService(canteenRepo, windowRepo, windowDishRelationRepo, dishRepo)
		CanteenController := controller_c.NewCanteenController(CanteenService)
		canteen.GET("/", CanteenController.ListCanteens)
		canteen.GET("/:canteen_id/windows", CanteenController.ListWindowsOfCanteen)
		canteen.GET("/windows/:window_id/dishes", CanteenController.ListDishesOfWindow)

		CanteenCommentController := controller_c.NewCanteenCommentController(CommentService)
		canteen.GET("/:canteen_id/comments", CanteenCommentController.AddCanteenComment)
	}

}

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	initRepos(db)

	setupBRoutes(r)
	setupCRoutes(r)
}
