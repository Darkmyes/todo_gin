package models

import "time"

type UserTask struct {
	CreatedAt time.Time
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`
	TaskID    uint `gorm:"primaryKey;autoIncrement:false"`
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Task      Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
