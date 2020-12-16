// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"database/sql"

	"week04/internal/dao"
	"week04/internal/service"

	"github.com/google/wire"
)

func InitializeDemo(db *sql.DB) *service.UserService {
	wire.Build(dao.NewUserDao, service.NewUserService)
	return &service.UserService{}
}
