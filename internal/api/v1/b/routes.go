package v1

import (
	controller "SJTU-Canteen-Community/internal/controller/b"
	"SJTU-Canteen-Community/internal/repository"
	"SJTU-Canteen-Community/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBRoutes(r *gin.Engine, db *gorm.DB) {
	b := r.Group("/b")
	canteen := b.Group("/canteen")
	CanteenRepo := repository.NewCanteenRepository(db)
	WindowRepo := repository.NewWindowRepository(db)
	DishRepo := repository.NewDishRepository(db)
	WindowDishRelationRepo := repository.NewWindowDishRelationRepository(db)

	{
		CanteenService := service.NewCanteenService(CanteenRepo, WindowRepo, WindowDishRelationRepo, DishRepo)
		CanteenController := controller.NewCanteenController(CanteenService)
		canteen.POST("/", CanteenController.AddCanteen)
		canteen.POST("/windows", CanteenController.MAddWindowsToCanteen)
		canteen.POST("/windows/dishes", CanteenController.MAddDishesToWindow)
	}

	dish := b.Group("/dish")
	{
		DishService := service.NewDishService(DishRepo)
		DishController := controller.NewDishController(DishService)
		dish.POST("/", DishController.MAddDishes)
	}
}
