package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model `json:"gorm_._model"`
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	Duration   int    `json:"duration"`
	Status     string `json:"status"`
}
