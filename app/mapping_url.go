package app

import (
	"github.com/pickypacker/bookstore_users-api/controllers/ping"
	"github.com/pickypacker/bookstore_users-api/controllers/users"
)

func mapURLs() {

	router.GET("/", ping.Root)
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
