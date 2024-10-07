package repository

import (
	"api-service-go/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExists(email string) (bool, error)
	Register(user *entity.Users) error
}


type authRespository struct {
	db *gorm.DB
}


func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRespository{
		db: db,
	}
}

func (r *authRespository) EmailExists(email string) (bool, error) {
	var user entity.Users

	err := r.db.First(&user, "email = ?", email).Error
	return err == nil, err

}

func (r *authRespository) Register(user *entity.Users) error {
	err := r.db.Create(&user).Error
	return err
}