package controllers

import (
	"log"
	"strconv"

	irisModels "github.com/evgesoch/gofwc/backend/iris/models"
	"github.com/kataras/iris"
)

// Get All Posts
func GetAllPosts() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		allPosts, err := irisModels.GetAllPosts()
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.Text("PostController(GetAll): Database error, can't fetch posts.")
			return
		}

		ctx.StatusCode(200)
		ctx.JSON(allPosts)
	}
}

// Get a Post
func GetPost() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		postID, err := strconv.Atoi(ctx.Params().Get("postID"))
		if err != nil {
			log.Println(err)
			ctx.StatusCode(400)
			ctx.Text("PostController(Get): Can't convert URI postID param to int.")
			return
		}

		post, err := irisModels.GetPostByID(postID)
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.Text("PostController(Get): Database error, can't fetch the post.")
			return
		}

		if post.ID == 0 {
			ctx.StatusCode(404)
			ctx.Text("PostController(Get): This Post doesn't exist in the database.")
			return
		}

		ctx.StatusCode(200)
		ctx.JSON(post)
	}
}

// Create a Post
func CreatePost() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		newPost := new(irisModels.Post)

		if err := ctx.ReadJSON(newPost); err != nil {
			log.Println(err)
			ctx.StatusCode(400)
			ctx.Text("PostController(Post): Cannot parse JSON request body.")
			return
		}

		newPostID, err := irisModels.CreatePost(newPost.Text)
		if err != nil {
			ctx.StatusCode(500)
			ctx.Text("PostController(Post): Database error, can't create new post.")
			return
		}

		ctx.StatusCode(201)
		ctx.JSON(map[string]interface{}{
			"postID": newPostID,
		})
	}
}

// Update a Post
func UpdatePost() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		postID, err := strconv.Atoi(ctx.Params().Get("postID"))
		if err != nil {
			log.Println(err)
			ctx.StatusCode(400)
			ctx.Text("PostController(Put): Can't convert URI postID param to int.")
			return
		}

		// Check if post to update doesn't exist
		allPosts, err := irisModels.GetAllPosts()
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.Text("PostController(GetAll): Database error, can't fetch posts.")
			return
		}
		if check := checkIfPostExists(postID, allPosts); !check {
			ctx.StatusCode(404)
			ctx.Text("PostController(Put): This post doesn't exist.")
			return
		}

		newPost := new(irisModels.Post)

		if err := ctx.ReadJSON(newPost); err != nil {
			log.Println(err)
			ctx.StatusCode(400)
			ctx.Text("PostController(Put): Cannot parse JSON request body.")
			return
		}

		if err := irisModels.UpdatePostByID(postID, newPost.Text); err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.Text("PostController(Put): Database error, can't update the post.")
			return
		}

		ctx.StatusCode(200)
		ctx.JSON(map[string]interface{}{
			"post_updated": "yes",
		})
	}
}

// Delete a Post
func DeletePost() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		postID, err := strconv.Atoi(ctx.Params().Get("postID"))
		if err != nil {
			log.Println(err)
			ctx.StatusCode(400)
			ctx.Text("PostController(Delete): Can't convert URI postID param to int.")
			return
		}

		// Check if post to delete doesn't exist
		allPosts, err := irisModels.GetAllPosts()
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.Text("PostController(GetAll): Database error, can't fetch posts.")
			return
		}
		if check := checkIfPostExists(postID, allPosts); !check {
			ctx.StatusCode(404)
			ctx.Text("PostController(Delete): This post doesn't exist.")
			return
		}

		if err := irisModels.DeletePostByID(postID); err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.Text("PostController(Delete): Database error, can't delete the post.")
			return
		}

		ctx.StatusCode(204)
		ctx.JSON(map[string]interface{}{
			"post_deleted": "yes",
		})
	}
}

// Serve the index page
func GetIndexPage() func(ctx iris.Context) {
	return func(ctx iris.Context) {

		ctx.ServeFile("../../frontend/index.html", true)
	}
}

// Check if a Post exists in the database by ID
func checkIfPostExists(postID int, postsSlice []*irisModels.Post) bool {
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
