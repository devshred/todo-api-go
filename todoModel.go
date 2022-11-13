package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	todo.ID = uuid.New()
	return
}
