package main

import (
	"github.com/LintaoAmons/go-crud/initializers"
	"github.com/LintaoAmons/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
