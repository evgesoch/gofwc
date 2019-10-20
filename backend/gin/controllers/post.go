package controllers

import (
	"log"
	"net/http"
	"strconv"

	ginModels "github.com/evgesoch/gofwc/backend/gin/models"
	"github.com/gin-gonic/gin"
)

// Get all Posts
func GetAllPosts() func(c *gin.Context) {
	return func(c *gin.Context) {
		allPosts, err := ginModels.GetAllPosts()
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(GetAll): Database error, can't fetch posts.")
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, allPosts)
	}
}

// Get a Post
func GetPost() func(c *gin.Context) {
	return func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Params.ByName("postID"))
		if err != nil {
			c.String(http.StatusBadRequest, "PostController(Get): Can't convert URI postID param to int.")
			log.Println(err)
			return
		}

		post, err := ginModels.GetPostByID(postID)
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(Get): Database error, can't fetch the post.")
			log.Println(err)
			return
		}

		if post.ID == 0 {
			c.String(http.StatusBadRequest, "PostController(Get): This Post doesn't exist in the database.")
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

// Create a new Post
func CreatePost() func(c *gin.Context) {
	return func(c *gin.Context) {
		newPost := new(ginModels.Post)

		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.String(http.StatusBadRequest, "PostController(Post): Cannot parse JSON request body.")
			log.Println(err)
			return
		}

		newPostID, err := ginModels.CreatePost(newPost.Text)
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(Post): Database error, can't create new post.")
			log.Println(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"postID": newPostID})
	}
}

// Update a Post
func UpdatePost() func(c *gin.Context) {
	return func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Params.ByName("postID"))
		if err != nil {
			c.String(http.StatusBadRequest, "PostController(Put): Can't convert URI postID param to int.")
			log.Println(err)
			return
		}

		// Check if post to update doesn't exist
		allPosts, err := ginModels.GetAllPosts()
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(GetAll): Database error, can't fetch posts.")
			log.Println(err)
			return
		}
		if check := checkIfPostExists(postID, allPosts); !check {
			c.String(http.StatusBadRequest, "PostController(Put): This post doesn't exist.")
			return
		}

		newPost := new(ginModels.Post)

		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.String(http.StatusBadRequest, "PostController(Put): Cannot parse JSON request body.")
			log.Println(err)
			return
		}

		err = ginModels.UpdatePostByID(postID, newPost.Text)
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(Put): Database error, can't update the post.")
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"post_updated": "yes"})
	}
}

// Delete a Post
func DeletePost() func(c *gin.Context) {
	return func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Params.ByName("postID"))
		if err != nil {
			c.String(http.StatusBadRequest, "PostController(Delete): Can't convert URI postID param to int.")
			log.Println(err)
			return
		}

		// Check if post to delete doesn't exist
		allPosts, err := ginModels.GetAllPosts()
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(GetAll): Database error, can't fetch posts.")
			log.Println(err)
			return
		}
		if check := checkIfPostExists(postID, allPosts); !check {
			c.String(http.StatusBadRequest, "PostController(Delete): This post doesn't exist.")
			return
		}

		err = ginModels.DeletePostByID(postID)
		if err != nil {
			c.String(http.StatusInternalServerError, "PostController(Delete): Database error, can't delete the post.")
			log.Println(err)
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"post_deleted": "yes"})
	}
}

// Check if a Post exists in the database by ID
func checkIfPostExists(postID int, postsSlice []*ginModels.Post) bool {
	arr := make([]int, 0)

	for _, v := range postsSlice {
		arr = append(arr, v.ID)
	}

	for _, v := range arr {
		if postID == v {
			return true
		}
	}

	return false
}
