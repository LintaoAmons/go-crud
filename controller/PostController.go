package controller

import (
	"github.com/LintaoAmons/go-crud/initializers"
	"github.com/LintaoAmons/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// get POST data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create iTitle
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{"post": post})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
