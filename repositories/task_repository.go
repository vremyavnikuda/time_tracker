package repositories

import (
	"time_tracker/config"
	"time_tracker/models"
)

// TODO:repositories->CreateTask
func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

// TODO:repositories->GetTasksByUserID
func GetTasksByUserID(userID int) ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Where("user_id = ?", userID).Find(&tasks)
	return tasks, result.Error
}

// TODO:repositories->GetTaskByID
func GetTaskByID(taskID uint) (models.Task, error) {
	var task models.Task
	result := config.DB.First(&task, taskID)
	return task, result.Error
}

// TODO:repositories->UpdateTask
func UpdateTask(task *models.Task) error {
	return config.DB.Save(task).Error
}

// TODO:repositories->DeleteTask
func DeleteTask(taskID uint) error {
	return config.DB.Delete(&models.Task{}, taskID).Error
}
