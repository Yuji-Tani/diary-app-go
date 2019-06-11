package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	appengine.Main()

	router.Run(":9090")

}
