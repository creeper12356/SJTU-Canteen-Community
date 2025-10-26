package v1

import (
	controller_b "SJTU-Canteen-Community/internal/controller/b"
	controller_c "SJTU-Canteen-Community/internal/controller/c"
	"SJTU-Canteen-Community/internal/middleware"
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
var windowCommentRepo *repository.WindowCommentRepository

func initRepos(db *gorm.DB) {
	canteenRepo = repository.NewCanteenRepository(db)
	windowRepo = repository.NewWindowRepository(db)
	dishRepo = repository.NewDishRepository(db)
	windowDishRelationRepo = repository.NewWindowDishRelationRepository(db)
	userRepo = repository.NewUserRepository(db)
	canteenCommentRepo = repository.NewCanteenCommentRepository(db)
	windowCommentRepo = repository.NewWindowCommentRepository(db)

}

func setupBRoutes(r *gin.Engine, db *gorm.DB) {
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

func setupCRoutes(r *gin.Engine, db *gorm.DB) {
	c := r.Group("/c")
	auth := c.Group("/auth")
	{
		AuthService := service.NewAuthService(userRepo)
		AuthController := controller_c.NewAuthController(AuthService)
		auth.POST("/login", AuthController.Login)
		auth.GET("/test_login_status", AuthController.TestLoginStatus)
		auth.POST("/logout", AuthController.Logout)
	}

	CommentService := service.NewCommentService(canteenCommentRepo, canteenRepo, windowRepo, windowDishRelationRepo, windowCommentRepo)
	ContentLikeService := service.NewContentLikeService(db)
	canteen := c.Group("/canteens", middleware.AuthMiddleware())
	{
		CanteenService := service.NewCanteenService(canteenRepo, windowRepo, windowDishRelationRepo, dishRepo)
		CanteenController := controller_c.NewCanteenController(CanteenService)
		canteen.GET("/", CanteenController.ListCanteens)
		canteen.GET("/:canteen_id/windows", CanteenController.ListWindowsOfCanteen)
		canteen.GET("/windows/:window_id/dishes", CanteenController.ListDishesOfWindow)

		CanteenCommentController := controller_c.NewCanteenCommentController(CommentService, ContentLikeService)
		canteen.POST("/:canteen_id/comments", CanteenCommentController.AddCanteenComment)
		canteen.POST("/comments/:comment_id/like", CanteenCommentController.LikeCanteenComment)

		WindowCommentController := controller_c.NewWindowCommentController(CommentService, ContentLikeService)
		canteen.POST("/windows/:window_id/comments", WindowCommentController.AddWindowComment)
		canteen.POST("/windows/comments/:comment_id/like", WindowCommentController.LikeWindowComment)
	}

}

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	initRepos(db)

	setupBRoutes(r, db)
	setupCRoutes(r, db)
}
