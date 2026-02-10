package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Настройка CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
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
		api.GET("/organizations/:id", getOrganization)
		api.POST("/organizations", createOrganization)
		api.PUT("/organizations/:id", updateOrganization)
		api.DELETE("/organizations/:id", deleteOrganization)
		
		// Расписание
		api.GET("/schedule", getSchedule)
		api.GET("/schedule/:id", getScheduleItem)
		api.POST("/schedule", createSchedule)
		api.PUT("/schedule/:id", updateSchedule)
		api.DELETE("/schedule/:id", deleteSchedule)
	}

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	r.Run(":8080")
}
