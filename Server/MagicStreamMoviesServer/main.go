package main

import (
	"fmt"

	"github.com/Furkanberkay/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello")
	})

	router.GET("/movies", controllers.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}
