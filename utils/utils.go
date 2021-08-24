package utils

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"
	_ "net/http"
)

func Respond(writer http.ResponseWriter, data map[string]interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}
