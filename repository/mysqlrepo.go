package repository

import (
	"errors"
	"fmt"
	"go-demo/entity"
)

type sqlRepo struct{}

func NewSqlRepo() Repository {
	return &sqlRepo{}
}

var users []*entity.User

func (*sqlRepo) Post(user *entity.User) error {
	for _, u := range users {
		if user.ID == u.ID {
			return fmt.Errorf("user already exists with ID %d", u.ID)
		}
	}
	users = append(users, user)
	return nil
}

func (*sqlRepo) Get() ([]*entity.User, error) {
	if len(users) == 0 {
		return nil, errors.New("no user data")
	}
	return users, nil
}
