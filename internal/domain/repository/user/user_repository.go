package repository

import (
	model "king/internal/domain/model/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	// UpdateUser(user *model.User) error
	// DeleteUser(user *model.User) error
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormOrderRepository{db: db};
}

func(r *GormOrderRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user);
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func(r *GormOrderRepository) GetUserByUsername(username string) (*model.User, error) {
	
	var user model.User

    err := r.db.Where("username = ?", username).Find(&user).Error;
	if err != nil {
		return nil, err
	}

	return &user, nil
}