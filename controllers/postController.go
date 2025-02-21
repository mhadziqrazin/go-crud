package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
  var body models.Post

  c.Bind(&body)

  result := initializers.DB.Create(&body)

  if result.Error != nil {
    c.Status(400)
    return
  }

  c.JSON(201, body)
}

func GetAllPosts(c *gin.Context) {
  var posts *[]models.Post

  initializers.DB.Find(&posts)

  c.JSON(200, posts)
}

func GetPostById(c *gin.Context) {
  var post models.Post

  initializers.DB.First(&post, c.Param("id"))

  c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
  var body models.Post
  var post models.Post

  c.Bind(&body)

  initializers.DB.First(&post, c.Param("id"))

  initializers.DB.Model(&post).Updates(&body)

  c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
  initializers.DB.Delete(&models.Post{}, c.Param("id"))

  c.Status(204)
}
