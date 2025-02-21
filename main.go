package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDb()
}

func main() {
  r := gin.Default()

  r.GET("/post", controllers.GetAllPosts)
  r.POST("/post", controllers.CreatePost)
  r.GET("/post/:id", controllers.GetPostById)
  r.PATCH("/post/:id", controllers.UpdatePost)
  r.DELETE("/post/:id", controllers.DeletePost)

  r.Run()
}
