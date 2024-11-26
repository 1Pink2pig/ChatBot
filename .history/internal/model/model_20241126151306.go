package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `gorm:"column:id;" json:"id"`
	Name     string `gorm:"column:name;" json:"name"`
	Username string `gorm:"column:username; primary_key;" json:"username"`
	Passwd   string `gorm:"column:passwd;" json:"password"`

	Balance  float64 `gorm:"column:balance;" json:""`
}