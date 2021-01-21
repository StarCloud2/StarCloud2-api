package services

import "github.com/StarCloud2/StarCloud2-api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
