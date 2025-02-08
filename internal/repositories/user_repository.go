package repositories

import (
	"codebase-go-echo/internal/models"
	"codebase-go-echo/pkg/databases/postgresql"
	"context"
	"errors"

	"gorm.io/gorm"
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

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := postgresql.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
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

func CreateUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	db := postgresql.DB.WithContext(ctx)
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx context.Context, id int, user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	db := postgresql.DB.WithContext(ctx)
	result := db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func DeleteUser(ctx context.Context, id int) error {
	db := postgresql.DB.WithContext(ctx)
	result := db.Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
