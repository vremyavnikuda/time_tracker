package repositories

import (
	"time_tracker/config"
	"time_tracker/models"
)

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func GetTasksByUserID(userID int) ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Where("user_id = ?", userID).Find(&tasks)
	return tasks, result.Error
}

func GetTaskByID(taskID uint) (models.Task, error) {
	var task models.Task
	result := config.DB.First(&task, taskID)
	return task, result.Error
}

func UpdateTask(task *models.Task) error {
	return config.DB.Save(task).Error
}

func DeleteTask(taskID uint) error {
	return config.DB.Delete(&models.Task{}, taskID).Error
}
