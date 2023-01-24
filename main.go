package main

import (
	"fmt"
	"github.com/LintaoAmons/go-crud/controller"
	"github.com/LintaoAmons/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("sdfHello 22")

	r := gin.Default()

	r.POST("/posts", controller.PostCreate)
	r.GET("/posts", controller.PostsIndex)
	r.GET("/posts/:id", controller.GetPost)
	r.PUT("/posts/:id", controller.UpdatePost)
	r.DELETE("/posts/:id", controller.DeletePost)

	r.Run()
}
