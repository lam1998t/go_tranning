package src

import (
	_ "fmt"
	_ "github.com/cesc1802/go_training/utils"
	_ "gorm.io/gorm"
)

type User struct {
	ID string `json:"id"`
	Password string `json:"password"`
	MaxTodo string `json:"max_todo"`
}
