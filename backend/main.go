package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Organization представляет организацию
type Organization struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// Schedule представляет расписание
type Schedule struct {
	ID          string `json:"id" binding:"required"`
	OrgID       string `json:"org_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
}

// Хранилище данных (временно в памяти)
var organizations = []Organization{}
var schedules = []Schedule{}

func main() {
	r := gin.Default()

	// Настройка CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API маршруты
	api := r.Group("/api")
	{
		// Организации
		api.GET("/organizations", getOrganizations)
		api.POST("/organizations", createOrganization)
		
		// Расписание
		api.GET("/schedule", getSchedule)
		api.POST("/schedule", createSchedule)
	}

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	r.Run(":8080")
}

// Получить все организации
func getOrganizations(c *gin.Context) {
	c.JSON(http.StatusOK, organizations)
}

// Создать организацию
func createOrganization(c *gin.Context) {
	var org Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organizations = append(organizations, org)
	c.JSON(http.StatusCreated, org)
}

// Получить расписание по организации
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

// Создать событие в расписании
func createSchedule(c *gin.Context) {
	var schedule Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	schedules = append(schedules, schedule)
	c.JSON(http.StatusCreated, schedule)
}