package routes

import (
	swaggerFiles "github.com/swaggo/files"
	"task/controllers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.AddUser)
		userRoutes.POST("/full", controllers.AddUserWithFullInfo)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
	}

	r.POST("/info", controllers.GetUserInfo)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	timeRoutes := r.Group("/time")
	{
		timeRoutes.POST("/start", controllers.StartTimeEntry)
		timeRoutes.POST("/end/:id", controllers.EndTimeEntry)
		timeRoutes.GET("/", controllers.GetTimeEntries)
	}

	return r
}
