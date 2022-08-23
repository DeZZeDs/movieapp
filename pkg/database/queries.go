package database

import (
	"movieapp/internal/models"
)

func UserCreate(user *models.User) (err error) {
	_, err = DB.Exec(`INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)`,
		user.FirstName, user.LastName, user.Email, user.Password)
	return
}
