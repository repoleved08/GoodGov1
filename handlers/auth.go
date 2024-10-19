package handlers

import (
	"auth-api/services"
	"encoding/json"
	"net/http"
)

var authService = services.NewAuthService()

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)
	err := authService.Register(data["username"], data["email"], data["password"])
}
