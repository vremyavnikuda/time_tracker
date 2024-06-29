package controllers

import (
	"net/http"
	"strconv"
	"time_tracker/models"
	"time_tracker/repositories"
	"time_tracker/services"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	tasks, err := repositories.GetTasksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func StartTask(c *gin.Context) {
	var req services.StartTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.StartTask(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task started successfully"})
}

func StopTask(c *gin.Context) {
	var req services.StopTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.StopTask(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task stopped successfully"})
}
