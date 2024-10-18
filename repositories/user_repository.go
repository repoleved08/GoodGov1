package repositories

import (
	"auth-api/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	
}
