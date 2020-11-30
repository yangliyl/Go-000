package service

import (
	"database/sql"
	"errors"

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
	user, err := s.dao.Find(id)
	if errors.Is(err, sql.ErrNoRows) {
		return user, nil
	}
	return nil, err
}
