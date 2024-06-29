package routers

import (
	"github.com/gin-gonic/gin"
	"time_tracker/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users")
	{
		userGroup.POST("", controllers.CreateUser)
		userGroup.GET("", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("", controllers.CreateTask)
		taskGroup.GET("", controllers.GetTasks)
		taskGroup.POST("/start", controllers.StartTask)
		taskGroup.POST("/stop", controllers.StopTask)
	}

	return router
}
