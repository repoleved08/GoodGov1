package repositories

import (
	"auth-api/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", user.Username, user.Email, user.Password)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	row := r.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}
