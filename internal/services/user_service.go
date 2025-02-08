package services

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/internal/repositories"
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

func GetUsers(ctx context.Context, limit, offset int) ([]models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	users, err := repositories.GetUsers(ctx, limit, offset)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, errors.New("failed to retrieve users")
	}

	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}

func GetUsersPaginated(ctx context.Context, limit, offset int) ([]models.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	users, totalCount, err := repositories.GetUsersPaginated(ctx, limit, offset)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, 0, errors.New("failed to retrieve users")
	}

	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, totalCount, nil
}

func CreateUser(ctx context.Context, user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}

	existingUser, err := repositories.GetUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error checking existing user: %v", err)
		return errors.New("failed to check existing user")
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	if err := repositories.CreateUser(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return errors.New("failed to create user")
	}

	return nil
}

func UpdateUser(ctx context.Context, id int, user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	user.UpdatedAt = time.Now()

	err := repositories.UpdateUser(ctx, id, user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return errors.New("failed to update user")
	}

	return nil
}

func DeleteUser(ctx context.Context, id int) error {
	err := repositories.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return errors.New("failed to delete user")
	}

	return nil
}
