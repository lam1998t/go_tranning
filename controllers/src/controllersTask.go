package src

import (
	_ "encoding/json"
	_ "github.com/cesc1802/go_training/models/src"
	_ "net/http"
)

type Task struct {
	ID 			string `json:"id"`
	Content	 	string `json:"content"`
	UserID 		string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
