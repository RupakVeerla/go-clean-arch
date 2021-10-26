package service

import (
	"go-demo/entity"
	"go-demo/repository"
)

type UserService interface {
	AddUser(user *entity.User) error
	GetUsers() ([]*entity.User, error)
}

type userService struct{}

var repo repository.Repository

func NewUserService(r repository.Repository) UserService {
	repo = r
	return &userService{}
}

func (*userService) AddUser(user *entity.User) error {
	return repo.Post(user)
}

func (*userService) GetUsers() ([]*entity.User, error) {
	return repo.Get()
}
