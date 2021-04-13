package models

import "gorm.io/gorm"

type UserTask struct {
	gorm.Model
	UserID uint
	TaskID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Task   Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
