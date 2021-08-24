package src

import (
	_ "fmt"
	_ "github.com/cesc1802/go_training/utils"
	_ "gorm.io/gorm"
)

type Task struct {
	ID string `json:"id"`
	Content string `json:"content"`
	UserID string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
