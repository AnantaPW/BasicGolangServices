package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	ID              string `json:"id"`
	ResponseMessage string `json:"message"`
}

var responses = []response{
	{ID: "1", ResponseMessage: "Hello, this is response 1"},
	{ID: "2", ResponseMessage: "Hello, this is response 2"},
	{ID: "3", ResponseMessage: "Hello, this is response 3"},
}

func getResponseByID(ginContext *gin.Context) {
	id := ginContext.Param("id")

	for _, response := range responses {
		if response.ID == id {
			ginContext.IndentedJSON(http.StatusOK, response)
			return
		}
	}
	ginContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "response not found"})
}

func getResponses(ginContext *gin.Context) {
	ginContext.IndentedJSON(http.StatusOK, responses)
}

func main() {
	router := gin.Default()

	//route path
	router.GET("/responses", getResponses)
	router.GET("/responses/:id", getResponseByID)

	//run router
	router.Run("localhost:8080")
}
