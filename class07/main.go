package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type Todo struct {
	gorm.Model
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID string `json:"user_id"`
}
