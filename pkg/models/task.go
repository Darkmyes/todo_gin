package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	State       bool   `json:"state"`
}

type TaskUsers struct {
	Task      Task   `json:"task"`
	TaskUsers []User `json:"users"`
}
