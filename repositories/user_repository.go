package repositories

import (
	"time_tracker/config"
	"time_tracker/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func GetUsers(surname, name, address string, page, pageSize int) ([]models.User, error) {
	var users []models.User
	query := config.DB

	if surname != "" {
		query = query.Where("surname LIKE ?", "%"+surname+"%")
	}
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if address != "" {
		query = query.Where("address LIKE ?", "%"+address+"%")
	}

	offset := (page - 1) * pageSize
	result := query.Offset(offset).Limit(pageSize).Find(&users)
	return users, result.Error
}
