package dao

import (
	"database/sql"

	"Week02/internal/model"

	"github.com/pkg/errors"
)

// UserDao struct
type UserDao struct{}

// NewUserDao Create user dao
func NewUserDao() *UserDao {
	return &UserDao{}
}

// Find return user info
func (dao *UserDao) Find(id int) (*model.User, error) {
	_, err := find()
	if err != nil {
		return nil, errors.Wrapf(err, "Not Found User: %d", id)
	}
	return &model.User{}, nil
}

func find() (interface{}, error) {
	return nil, sql.ErrNoRows
}
