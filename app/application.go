package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	log.Fatalln(router.Run(":8080"))
}
