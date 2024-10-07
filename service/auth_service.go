package service

import (
	"api-service-go/dto"
	"api-service-go/entity"
	errorhandler "api-service-go/errorHandler"
	"api-service-go/helper"
	"api-service-go/repository"
)


type AuthService interface {
	Register(*dto.RegisterRequest)error
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthRepository(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {

	// check email exists
	if emailExists, _ := s.repository.EmailExists(req.Email); emailExists {
		return &errorhandler.BadRequestError{
			Message: "Email already exists",
		}
	}

	// check password and confirm password
	if req.Password != req.ConfirmPassword {
		return &errorhandler.BadRequestError{
			Message: "Password and confirm password does not match",
		}
	}
	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	user := entity.Users{
		Name: req.Name,
		Email: req.Email,
		Gender: req.Gender,
		Password: passwordHash,
	}
	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	return nil
}