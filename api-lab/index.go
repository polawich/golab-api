package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "cleanroom", Completed: false},
	{ID: "2", Item: "cleanroom", Completed: false},
	{ID: "3", Item: "cleanroom", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
	var newTodo todo

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "messagefail"})
	}

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*.html")
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run(":9090")
}