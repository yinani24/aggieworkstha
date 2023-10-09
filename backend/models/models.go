package models

import "gorm.io/gorm"

type EventType struct{
	gorm.Model

	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Organization string `json:"organization"`
}

type Organization struct {
	gorm.Model

	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Events []EventType `json:"eventtype"`
}

type User struct {
	gorm.Model
	
	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Events []EventType `json:"eventtype"`
}

type Event struct {
	gorm.Model

	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Organization string `json:"organization"`
	Users []User `json:"users"`
}