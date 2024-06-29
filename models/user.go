package models

import (
	"time"
)

type User struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	PassportNumber string     `json:"passport_number"`
	PassportSerie  string     `json:"passport_serie"`
	Surname        string     `json:"surname"`
	Name           string     `json:"name"`
	Patronymic     string     `json:"patronymic"`
	Address        string     `json:"address"`
}
