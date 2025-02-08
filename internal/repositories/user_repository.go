package repositories

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/pkg/databases/postgresql"
	"context"
)

func GetUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	var users []models.User

	result := postgresql.DB.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetUsersPaginated(ctx context.Context, limit, offset int) ([]models.User, int, error) {
	var users []models.User
	var totalCount int64

	result := postgresql.DB.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&users)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	postgresql.DB.Model(&models.User{}).Count(&totalCount)

	return users, int(totalCount), nil
}
