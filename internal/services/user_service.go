package services

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/internal/repositories"
)

func FetchAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}
