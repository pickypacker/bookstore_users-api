package app

import "github.com/pickypacker/bookstore_users-api/controllers"

func mapURLs() {

	router.GET("/", controllers.Root)
	router.GET("/ping", controllers.Ping)
	router.GET("/users/user_id", controllers.GetUser)
	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
}
