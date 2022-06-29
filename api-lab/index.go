package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func Post(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"result": u})
}

func main() {
	r := gin.Default()
	r.GET("/", GetHelloWorld)
	r.POST("/", Post)
	r.Run()
}
