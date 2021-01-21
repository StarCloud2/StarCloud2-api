package app

import (
	"github.com/StarCloud2/StarCloud2-api/controllers/ping"
	"github.com/StarCloud2/StarCloud2-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)

	router.POST("/users", users.CreateUser)
}
