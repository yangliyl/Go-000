package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"Week02/service"
)

func main() {
	user, err := service.NewUserService().GetUser(1)
	if err != nil {
		fmt.Printf("%+v\n", err)
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, user)
}
