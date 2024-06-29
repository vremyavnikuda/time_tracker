package controllers

import (
	"net/http"
	"strconv"
	"task/database"
	"task/models"
	"time"

	"github.com/gin-gonic/gin"
)

// StartTimeEntry
// @Summary Start a new time entry
// @Description Start a new time entry for a user
// @Tags time
// @Accept json
// @Produce json
// @Param user_id body uint true "User ID"
// @Param task body string true "Task description"
// @Success 200 {object} models.TimeEntry
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /time/start [post]
func StartTimeEntry(c *gin.Context) {
	var input struct {
		UserID uint   `json:"user_id" binding:"required"`
		Task   string `json:"task" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	timeEntry := models.TimeEntry{
		UserID:    input.UserID,
		Task:      input.Task,
		StartTime: time.Now(),
	}

	database.DB.Create(&timeEntry)
	c.JSON(http.StatusOK, timeEntry)
}

// EndTimeEntry
// @Summary End an existing time entry
// @Description End an existing time entry for a user
// @Tags time
// @Accept json
// @Produce json
// @Param id path uint true "Time Entry ID"
// @Success 200 {object} models.TimeEntry
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "Time entry not found"
// @Router /time/end/{id} [post]
func EndTimeEntry(c *gin.Context) {
	id := c.Param("id")
	entryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	var timeEntry models.TimeEntry
	if err := database.DB.First(&timeEntry, entryID).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "Time entry not found"})
		return
	}

	timeEntry.EndTime = time.Now()
	timeEntry.Duration = int64(timeEntry.EndTime.Sub(timeEntry.StartTime).Seconds())

	database.DB.Save(&timeEntry)
	c.JSON(http.StatusOK, timeEntry)
}

// GetTimeEntries
// @Summary Get time entries
// @Description Get time entries for a user within a specified period
// @Tags time
// @Accept json
// @Produce json
// @Param user_id query uint true "User ID"
// @Param start_date query string true "Start Date" format(date)
// @Param end_date query string true "End Date" format(date)
// @Success 200 {array} models.TimeEntry
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /time [get]
func GetTimeEntries(c *gin.Context) {
	var input struct {
		UserID    uint   `form:"user_id" binding:"required"`
		StartDate string `form:"start_date" binding:"required"`
		EndDate   string `form:"end_date" binding:"required"`
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	startTime, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start date"})
		return
	}

	endTime, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end date"})
		return
	}

	var timeEntries []models.TimeEntry
	database.DB.Where("user_id = ? AND start_time BETWEEN ? AND ?", input.UserID, startTime, endTime).Order("duration DESC").Find(&timeEntries)
	c.JSON(http.StatusOK, timeEntries)
}
