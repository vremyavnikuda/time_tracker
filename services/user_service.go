package services

import (
	"time_tracker/models"
	"time_tracker/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUserByID(id uint) (models.User, error) {
	return repositories.GetUserByID(id)
}

func UpdateUser(user *models.User) error {
	existingUser, err := repositories.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	existingUser.Name = user.Name
	existingUser.Surname = user.Surname
	existingUser.Patronymic = user.Patronymic
	existingUser.Address = user.Address

	return repositories.UpdateUser(&existingUser)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}

func GetUsers(surname, name, address string, page, pageSize int) ([]models.User, error) {
	return repositories.GetUsers(surname, name, address, page, pageSize)
}
