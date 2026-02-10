package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Хранилище данных (временно в памяти)
var organizations = make(map[string]Organization)
var schedules = make(map[string]Schedule)

// getOrganizations получает все организации
func getOrganizations(c *gin.Context) {
	var result []Organization
	for _, org := range organizations {
		result = append(result, org)
	}
	c.JSON(http.StatusOK, result)
}

// getOrganization получает организацию по ID
func getOrganization(c *gin.Context) {
	id := c.Param("id")
	org, exists := organizations[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	c.JSON(http.StatusOK, org)
}

// createOrganization создает новую организацию
func createOrganization(c *gin.Context) {
	var req CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерируем UUID
	id := uuid.New().String()
	now := time.Now()

	org := Organization{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	organizations[id] = org
	c.JSON(http.StatusCreated, org)
}

// updateOrganization обновляет организацию
func updateOrganization(c *gin.Context) {
	id := c.Param("id")
	
	_, exists := organizations[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	var req UpdateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	org := organizations[id]
	org.Name = req.Name
	org.Description = req.Description
	org.UpdatedAt = time.Now()

	organizations[id] = org
	c.JSON(http.StatusOK, org)
}

// deleteOrganization удаляет организацию
func deleteOrganization(c *gin.Context) {
	id := c.Param("id")
	
	_, exists := organizations[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	delete(organizations, id)
	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}

// getSchedule получает расписание по организации
func getSchedule(c *gin.Context) {
	orgID := c.Query("org_id")
	if orgID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "org_id is required"})
		return
	}
	
	var result []Schedule
	for _, s := range schedules {
		if s.OrgID == orgID {
			result = append(result, s)
		}
	}
	c.JSON(http.StatusOK, result)
}

// getScheduleItem получает конкретное расписание по ID
func getScheduleItem(c *gin.Context) {
	id := c.Param("id")
	schedule, exists := schedules[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule item not found"})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

// createSchedule создает новое событие в расписании
func createSchedule(c *gin.Context) {
	var req CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем существование организации
	_, orgExists := organizations[req.OrgID]
	if !orgExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization not found"})
		return
	}

	// Генерируем UUID
	id := uuid.New().String()
	now := time.Now()

	schedule := Schedule{
		ID:          id,
		OrgID:       req.OrgID,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	schedules[id] = schedule
	c.JSON(http.StatusCreated, schedule)
}

// updateSchedule обновляет событие в расписании
func updateSchedule(c *gin.Context) {
	id := c.Param("id")
	
	schedule, exists := schedules[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule item not found"})
		return
	}

	var req UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule.Title = req.Title
	schedule.Description = req.Description
	schedule.StartTime = req.StartTime
	schedule.EndTime = req.EndTime
	schedule.UpdatedAt = time.Now()

	schedules[id] = schedule
	c.JSON(http.StatusOK, schedule)
}

// deleteSchedule удаляет событие из расписания
func deleteSchedule(c *gin.Context) {
	id := c.Param("id")
	
	_, exists := schedules[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule item not found"})
		return
	}

	delete(schedules, id)
	c.JSON(http.StatusOK, gin.H{"message": "Schedule item deleted successfully"})
}