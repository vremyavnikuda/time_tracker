package models

import "gorm.io/gorm"

// TODO:models->User
type User struct {
	gorm.Model     `json:"gorm_._model,omitempty"`
	PassportNumber string `gorm:"unique" json:"passport_number"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}
