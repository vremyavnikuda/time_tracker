package routers

import (
	"github.com/gin-gonic/gin"
	"time_tracker/controllers"
)

// TODO:routers->SetupRouter
func SetupRouter() *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users")
	{
		userGroup.POST("/createUser", controllers.CreateUser)
		userGroup.GET("/getAllUsers", controllers.GetUsers)
		userGroup.GET("/getUsersByID/:id", controllers.GetUserByID)
		userGroup.PUT("/updateUser/:id", controllers.UpdateUser)
		userGroup.DELETE("/deleteUser/:id", controllers.DeleteUser)
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
