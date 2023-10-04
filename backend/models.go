package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Events []Event `json:"events"`
}

type User struct {
	gorm.Model
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Events []Event `json:"events"`
}

type Event struct {
	gorm.Model
	ID int `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Organization Organization `json:"organization"`
	Users []User `json:"users"`
}