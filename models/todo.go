package models

import (
	"errors"
	"time"
)

type TodoStatus string

const (
	Progress  TodoStatus = "progress"
	Completed TodoStatus = "completed"
)

func (s TodoStatus) IsValid() bool {
	return s == Progress || s == Completed
}

func (s TodoStatus) Validate() error {
	if !s.IsValid() {
		return errors.New("invalid status: must be either 'progress' or 'completed'")
	}
	return nil
}

type Todo struct {
	TodoID    uint       `gorm:"primaryKey" json:"todo_id"`
	Content   string     `json:"content"`
	Status    TodoStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
