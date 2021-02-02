package repository

import (
	"labelproject-back/model"

	"github.com/jinzhu/gorm"
)

// UserRepository defines functions for model.User
type UserRepository interface {
	//
	GetUser(userName string) ([]*model.User, error)

	//
	GetAuthoritiesByID(userID int64) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

// This statement verifies interface compliance.
var userRepositoryInstance = &userRepository{}

// UserRepositoryInstance returns the instance of userRepository
func UserRepositoryInstance(db *gorm.DB) UserRepository {
	userRepositoryInstance.db = db
	return userRepositoryInstance
}

//
func (r *userRepository) GetUser(userName string) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Where("user_name = ?", userName).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

//
func (r *userRepository) GetAuthoritiesByID(userID int64) (string, error) {
	var user model.User
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Authorities, nil
}
