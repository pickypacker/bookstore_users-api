package app

import "github.com/pickypacker/bookstore_users-api/controllers"

func mapURLs() {

	router.GET("/", controllers.Root)
	router.GET("/ping", controllers.Ping)
}
