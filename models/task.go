package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID uint
	Name string
	Duration int
	Status string
}
