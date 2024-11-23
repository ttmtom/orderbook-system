package repository

import (
	"gorm.io/gorm"
	"log/slog"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) IsUserExist(email string) bool {
	var exists bool

	results := ur.DB.Model(&model.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists)

	if results.Error != nil {
		slog.Info("user exist", results.Error.Error())
		return false
	}
	slog.Info("ess", exists)

	return exists
}

func (ur *UserRepository) updatedIDHashOnCreateUserDone(user *model.User) (tx *gorm.DB) {
	result := ur.DB.Model(&user).Updates(map[string]interface{}{
		"id_hash": security.HashUserId(user.ID),
	})

	return result
}

func (ur *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	result := ur.DB.Create(&user)

	if result.Error != nil {
		slog.Info("Error on creating user", result.Error)
		return nil, result.Error
	}

	result = ur.updatedIDHashOnCreateUserDone(user)

	return user, result.Error
}

func (ur *UserRepository) getUserById(id uint) (*model.User, error) {
	var user *model.User
	result := ur.DB.Model(&model.User{}).
		Select("*").
		Where("id = ?", id).
		First(&user)

	if result.Error != nil {
		slog.Info("Error on getting user by id", result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) GetUserByIdHash(id string) (*model.User, error) {
	var user *model.User
	result := ur.DB.Model(&model.User{}).
		Select("*").
		Where("id_hash = ?", id).
		First(&user)

	if result.Error != nil {
		slog.Info("Error on getting user by id", result.Error)
		return nil, result.Error
	}

	return user, nil
}
