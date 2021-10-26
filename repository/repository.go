package repository

import "go-demo/entity"

type Repository interface {
	Post(user *entity.User) error
	Get() ([]*entity.User, error)
}
