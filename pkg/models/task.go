package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	State       bool   `json:"state"`
	UserID      uint   `json:"user_id" gorm:"autoIncrement:false"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
