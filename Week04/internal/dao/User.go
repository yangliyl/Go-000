package dao

import (
	"database/sql"
	"errors"

	"week04/internal/model"

	pkgerr "github.com/pkg/errors"
)

var DataNotExist = errors.New("model: data not exists")

type UserDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Query(id int64) (*model.User, error) {
	user := model.User{}
	if err := d.db.QueryRow("SELECT id, name FROM user WHERE id = ?", id).Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			err = DataNotExist
		}
		return nil, pkgerr.Wrap(err, "query user fail")
	}
	return &user, nil
}
