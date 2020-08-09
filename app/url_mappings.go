package app

import (
	"github.com/judesantos/go-bookstore_users_api/controllers/ping"
	"github.com/judesantos/go-bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)

	router.POST("/users", users.CreateUser)

	//router.GET("/users/search", users.SearchUser)
}
