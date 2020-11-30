package service

import (
	"Week02/internal/dao"
	"Week02/internal/model"
)

// UserService struct
type UserService struct {
	dao *dao.UserDao
}

// NewUserService Create user service
func NewUserService() *UserService {
	return &UserService{
		dao: dao.NewUserDao(),
	}
}

// GetUser return user info
func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.dao.Find(id)
}
