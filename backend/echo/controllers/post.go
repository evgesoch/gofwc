package controllers

import (
	"log"
	"net/http"
	"strconv"

	echoModels "github.com/evgesoch/gofwc/backend/echo/models"
	"github.com/labstack/echo"
)

// Get all Posts
func GetAllPosts() func(c echo.Context) error {
	return func(c echo.Context) error {
		allPosts, err := echoModels.GetAllPosts()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "PostController(GetAll): Database error, can't fetch posts.")
		}

		return c.JSON(http.StatusOK, allPosts)
	}
}

// Get a Post
func GetPost() func(c echo.Context) error {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("postID"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "PostController(Get): Can't convert URI postID param to int.")
		}

		post, err := echoModels.GetPostByID(postID)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "PostController(Get): Database error, can't fetch the post.")
		}

		if post.ID == 0 {
			return c.String(http.StatusNotFound, "PostController(Get): This Post doesn't exist in the database.")
		}

		return c.JSON(http.StatusOK, post)
	}
}

// Create a Post
func CreatePost() func(c echo.Context) error {
	return func(c echo.Context) error {
		newPost := new(echoModels.Post)

		if err := c.Bind(newPost); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "PostController(Post): Cannot parse JSON request body.")
		}

		newPostID, err := echoModels.CreatePost(newPost.Text)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "PostController(Post): Database error, can't create new post.")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"postID": newPostID,
		})
	}
}

// Update a Post
func UpdatePost() func(c echo.Context) error {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("postID"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "PostController(Put): Can't convert URI postID param to int.")
		}

		// Check if post to update doesn't exist
		allPosts, err := echoModels.GetAllPosts()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "PostController(GetAll): Database error, can't fetch posts.")
		}
		if check := checkIfPostExists(postID, allPosts); !check {
			return c.String(http.StatusNotFound, "PostController(Put): This post doesn't exist.")
		}

		newPost := new(echoModels.Post)

		if err := c.Bind(newPost); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "PostController(Put): Cannot parse JSON request body.")
		}

		if err := echoModels.UpdatePostByID(postID, newPost.Text); err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "PostController(Put): Database error, can't update the post.")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"post_updated": "yes",
		})
	}
}

// Delete a Post
func DeletePost() func(c echo.Context) error {
	return func(c echo.Context) error {
		postID, err := strconv.Atoi(c.Param("postID"))
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "PostController(Delete): Can't convert URI postID param to int.")
		}

		// Check if post to delete doesn't exist
		allPosts, err := echoModels.GetAllPosts()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "PostController(GetAll): Database error, can't fetch posts.")
		}
		if check := checkIfPostExists(postID, allPosts); !check {
			return c.String(http.StatusNotFound, "PostController(Delete): This post doesn't exist.")
		}

		if err := echoModels.DeletePostByID(postID); err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "PostController(Delete): Database error, can't delete the post.")
		}

		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"post_deleted": "yes",
		})
	}
}

// Check if a Post exists in the database by ID
func checkIfPostExists(postID int, postsSlice []*echoModels.Post) bool {
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
