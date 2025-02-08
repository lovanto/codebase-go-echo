package services

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/internal/repositories"
	"context"
	"errors"
	"log"
	"time"
)

func GetUsers(ctx context.Context, limit, offset int) ([]models.UserResponse, error) {
	// Set a timeout to avoid long DB queries
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	users, err := repositories.GetUsers(ctx, limit, offset)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, errors.New("failed to retrieve users")
	}

	// Convert to response format to avoid exposing sensitive fields
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
	// Set a timeout to avoid long DB queries
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	users, totalCount, err := repositories.GetUsersPaginated(ctx, limit, offset)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, 0, errors.New("failed to retrieve users")
	}

	// Convert to response format
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
