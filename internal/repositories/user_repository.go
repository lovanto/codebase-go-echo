package repositories

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/pkg/databases/postgresql"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := postgresql.DB.Find(&users)
	return users, result.Error
}
