package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"task/database"
	"task/models"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users from the database
// @Summary Get all users
// @Description Get all users from the database
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// AddUser adds a new user to the database
// @Summary Add a new user
// @Description Add a new user to the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Add user"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to fetch user info"
// @Router /users [post]
func AddUser(c *gin.Context) {
	var input struct {
		PassportNumber string `json:"passportNumber" binding:"required"`
		PassportSerie  string `json:"passportSerie" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Fetch user info from external API
	url := fmt.Sprintf("%s?passportSerie=%s&passportNumber=%s", os.Getenv("PEOPLE_API_URL"), input.PassportSerie, input.PassportNumber)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Surname    string `json:"surname"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Address    string `json:"address"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode user info"})
		return
	}

	user := models.User{
		PassportNumber: input.PassportNumber,
		Surname:        userInfo.Surname,
		Name:           userInfo.Name,
		Patronymic:     userInfo.Patronymic,
		Address:        userInfo.Address,
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// GetUserInfo fetches user info based on passport number and returns it
// @Summary Get user info
// @Description Get user info based on passport number
// @Tags users
// @Accept  json
// @Produce  json
// @Param passportNumber body string true "Passport number"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to fetch user info"
// @Router /info [post]
func GetUserInfo(c *gin.Context) {
	var input struct {
		PassportNumber string `json:"passportNumber" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Assuming the PassportSerie is the first part of PassportNumber
	passportParts := strings.Fields(input.PassportNumber)
	if len(passportParts) != 2 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid passport number format"})
		return
	}

	passportSerie := passportParts[0]
	passportNumber := passportParts[1]

	// Fetch user info from external API
	url := fmt.Sprintf("%s?passportSerie=%s&passportNumber=%s", os.Getenv("PEOPLE_API_URL"), passportSerie, passportNumber)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Surname    string `json:"surname"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Address    string `json:"address"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode user info"})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

// DeleteUser removes a user from the database by ID
// @Summary Delete a user
// @Description Delete a user from the database by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string "message": "User deleted"
// @Failure 404 {object} map[string]string "error": "User not found"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}

// UpdateUser updates user information in the database
// @Summary Update a user
// @Description Update user information in the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "User not found"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	var input struct {
		Surname    string `json:"surname"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Address    string `json:"address"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user)
}

// AddUserWithFullInfo adds a new user with the given passport number and series
// @Summary Add a new user with full info
// @Description Add a new user with the given passport number and series
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Add user with full info"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to fetch user info"
// @Router /users/full [post]
func AddUserWithFullInfo(c *gin.Context) {
	var input struct {
		PassportNumber string `json:"passportNumber" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Assuming the PassportSerie is the first part of PassportNumber
	passportParts := strings.Fields(input.PassportNumber)
	if len(passportParts) != 2 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid passport number format"})
		return
	}

	passportSerie := passportParts[0]
	passportNumber := passportParts[1]

	// Fetch user info from external API
	url := fmt.Sprintf("%s?passportSerie=%s&passportNumber=%s", os.Getenv("PEOPLE_API_URL"), passportSerie, passportNumber)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Surname    string `json:"surname"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Address    string `json:"address"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode user info"})
		return
	}

	user := models.User{
		PassportNumber: input.PassportNumber,
		Surname:        userInfo.Surname,
		Name:           userInfo.Name,
		Patronymic:     userInfo.Patronymic,
		Address:        userInfo.Address,
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
