package main

import (
	"time"
)

// Organization представляет организацию
type Organization struct {
	ID          string    `json:"id" binding:"required,uuid4"`
	Name        string    `json:"name" binding:"required,min=2,max=100"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Schedule представляет расписание
type Schedule struct {
	ID          string    `json:"id" binding:"required,uuid4"`
	OrgID       string    `json:"org_id" binding:"required,uuid4"`
	Title       string    `json:"title" binding:"required,min=2,max=200"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required,gtfield=StartTime"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateOrganizationRequest запрос на создание организации
type CreateOrganizationRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description,omitempty"`
}

// UpdateOrganizationRequest запрос на обновление организации
type UpdateOrganizationRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description,omitempty"`
}

// CreateScheduleRequest запрос на создание расписания
type CreateScheduleRequest struct {
	OrgID       string    `json:"org_id" binding:"required,uuid4"`
	Title       string    `json:"title" binding:"required,min=2,max=200"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required,gtfield=StartTime"`
}

// UpdateScheduleRequest запрос на обновление расписания
type UpdateScheduleRequest struct {
	Title       string    `json:"title" binding:"required,min=2,max=200"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required,gtfield=StartTime"`
}