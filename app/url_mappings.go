package app

import (
	pc "github.com/leslesnoa/bookstore_users-api/contorollers/ping"
	uc "github.com/leslesnoa/bookstore_users-api/contorollers/users"
)

func mapUrls() {
	router.GET("/ping", pc.Ping)

	router.GET("/users/:user_id", uc.GetUser)
	router.POST("/users", uc.CreateUser)
	router.PUT("/users/:user_id", uc.UpdateUser)
	router.PATCH("/users/:user_id", uc.UpdateUser)
	router.DELETE("/users/:user_id", uc.DeleteUser)
	router.GET("/internal/users/search", uc.Search)
}
