package service

import (
	"log"

	"github.com/shaineminkyaw/grpc_lab/user_info/model"
	"github.com/shaineminkyaw/grpc_lab/user_info/repository"
)

type UserService interface {
	FindByEmail(email string) (*model.User, error)
}

type userService struct {
	UserService repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		UserService: repo,
	}
}

func (us *userService) FindByEmail(email string) (*model.User, error) {
	//
	user, err := us.UserService.FindByEmail(email)
	if err != nil {
		log.Fatalf("error on user service %v", err)
	}
	return user, nil
}
