package controllers

import (
	_ "log"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

var c *gin.Context

// Get all Posts
func GetAllPosts() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// Get a Post
func GetPost() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// Create a new Post
func CreatePost() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// Update a Post
func UpdatePost() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// Delete a Post
func DeletePost() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
