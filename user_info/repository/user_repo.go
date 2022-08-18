package repository

import (
	"log"

	"github.com/shaineminkyaw/grpc_lab/user_info/ds"
	"github.com/shaineminkyaw/grpc_lab/user_info/model"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
}

type userRepo struct {
	Source ds.DataSource
}

func NewUserRepo(dd *ds.DataSource) UserRepository {
	return &userRepo{
		Source: *dd,
	}
}

func (ur *userRepo) FindByEmail(email string) (*model.User, error) {
	//
	user := &model.User{}
	err := ur.Source.Data.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("error on filter email %v", err)
	}
	return user, nil
}
