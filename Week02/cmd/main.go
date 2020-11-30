package main

import (
	"fmt"
	"net/http"

	"Week02/service"
)

func main() {
	user, err := service.NewUserService().GetUser(1)
	if err != nil {
		fmt.Printf("%+v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, user)
}
