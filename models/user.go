package models

import "gorm.io/gorm"

type User struct {
	gorm.Model     `json:"gorm_._model,omitempty"`
	PassportNumber string `gorm:"unique" json:"passport_number,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Name           string `json:"name,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address,omitempty"`
}
