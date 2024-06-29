package services

import (
	"fmt"
	"time"
	"time_tracker/models"
	"time_tracker/repositories"
)

type StartTaskRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type StopTaskRequest struct {
	TaskID uint `json:"task_id" binding:"required"`
}

func StartTask(req StartTaskRequest) error {
	task := models.Task{
		UserID:   req.UserID,
		Name:     req.Name,
		Status:   "in_progress",
		Duration: 0,
	}
	return repositories.CreateTask(&task)
}

func StopTask(req StopTaskRequest) error {
	task, err := repositories.GetTaskByID(req.TaskID)
	if err != nil {
		return err
	}

	if task.Status != "in_progress" {
		return fmt.Errorf("task is not in progress")
	}

	task.Status = "completed"
	task.Duration = int(time.Since(task.CreatedAt).Minutes())

	return repositories.UpdateTask(&task)
}
