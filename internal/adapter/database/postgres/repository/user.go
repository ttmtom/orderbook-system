package repository

import (
	"gorm.io/gorm"
	"log/slog"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) IsUserExist(email string) bool {
	var exists bool

	results := ur.db.Model(&model.User{}).
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
	result := ur.db.Model(&user).Updates(map[string]interface{}{
		"id_hash": security.HashUserId(user.ID),
	})

	return result
}

func (ur *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		result := ur.db.Create(&user)
		if result.Error != nil {
			return result.Error
		}

		result = ur.updatedIDHashOnCreateUserDone(user)
		return result.Error
	})
	if err != nil {
		slog.Info("Error on creating user", err.Error)
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) getUserById(id uint) (*model.User, error) {
	var user *model.User
	result := ur.db.Model(&model.User{}).
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
	result := ur.db.Model(&model.User{}).
		Select("*").
		Where("id_hash = ?", id).
		First(&user)

	if result.Error != nil {
		slog.Info("Error on getting user by id", result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) GetUserById(id uint) (*model.User, error) {
	var user *model.User
	result := ur.db.Model(&model.User{}).
		Select("*").
		Where("id = ?", id).
		First(&user)

	if result.Error != nil {
		slog.Info("Error on getting user by id", result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User
	result := ur.db.Model(&model.User{}).
		Select("*").
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		slog.Info("Error on getting user by email", result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) UpdateUserLoginAt(user *model.User) {
	ur.db.Model(&user).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
	})
}

func (ur *UserRepository) UpdateUserLastAccessAt(userId string) {
	ur.db.Model(&model.User{}).
		Where("id_hash = ?", userId).
		Updates(
			map[string]interface{}{
				"last_access_at": time.Now(),
			})
}
