package app

import (
	"github.com/StarCloud2/StarCloud2-api/controllers/ping"
	"github.com/StarCloud2/StarCloud2-api/controllers/prometheus"
	"github.com/StarCloud2/StarCloud2-api/controllers/users"
)

func mapUrls() {
	router.GET("/metrics", prometheus.Handler())

	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
