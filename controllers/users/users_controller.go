package users

import (
	"encoding/json"
	"github.com/StarCloud2/StarCloud2-api/domain/users"
	"github.com/StarCloud2/StarCloud2-api/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		log.Fatalln(err)
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle json error
		log.Fatalln(err)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		//TODO: Handle error
		log.Fatalln(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
