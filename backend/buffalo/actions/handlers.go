package actions

import (
	"log"
	"strconv"

	buffaloModels "github.com/evgesoch/gofwc/backend/buffalo/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	packer "github.com/gobuffalo/packr"
)

// Get All Posts
func GetAllPosts(c buffalo.Context) error {
	allPosts, err := buffaloModels.GetAllPosts()
	if err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(GetAll): Database error, can't fetch posts."))
	}

	return c.Render(200, r.JSON(allPosts))
}

// Get a Post
func GetPost(c buffalo.Context) error {
	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		log.Println(err)
		return c.Render(400, r.String("PostController(Get): Can't convert URI postID param to int."))
	}

	post, err := buffaloModels.GetPostByID(postID)
	if err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(Get): Database error, can't fetch the post."))
	}

	if post.ID == 0 {
		return c.Render(404, r.String("PostController(Get): This Post doesn't exist in the database."))
	}

	return c.Render(200, r.JSON(post))
}

// Create a Post
func CreatePost(c buffalo.Context) error {
	newPost := new(buffaloModels.Post)

	if err := c.Bind(newPost); err != nil {
		log.Println(err)
		return c.Render(400, r.String("PostController(Post): Cannot parse JSON request body."))
	}

	newPostID, err := buffaloModels.CreatePost(newPost.Text)
	if err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(Post): Database error, can't create new post."))
	}

	return c.Render(201, r.JSON(map[string]interface{}{
		"postID": newPostID,
	}),
	)
}

// Update a Post
func UpdatePost(c buffalo.Context) error {
	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		log.Println(err)
		return c.Render(400, r.String("PostController(Put): Can't convert URI postID param to int."))
	}

	allPosts, err := buffaloModels.GetAllPosts()
	if err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(GetAll): Database error, can't fetch posts."))
	}
	if check := checkIfPostExists(postID, allPosts); !check {
		return c.Render(404, r.String("PostController(Put): This post doesn't exist."))
	}

	newPost := new(buffaloModels.Post)

	if err := c.Bind(newPost); err != nil {
		log.Println(err)
		return c.Render(400, r.String("PostController(Put): Cannot parse JSON request body."))
	}

	if err := buffaloModels.UpdatePostByID(postID, newPost.Text); err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(Put): Database error, can't update the post."))
	}

	return c.Render(200, r.JSON(map[string]interface{}{
		"post_updated": "yes",
	}),
	)
}

// Delete a Post
func DeletePost(c buffalo.Context) error {
	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		log.Println(err)
		return c.Render(400, r.String("PostController(Delete): Can't convert URI postID param to int."))
	}

	allPosts, err := buffaloModels.GetAllPosts()
	if err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(GetAll): Database error, can't fetch posts."))
	}
	if check := checkIfPostExists(postID, allPosts); !check {
		return c.Render(404, r.String("PostController(Delete): This post doesn't exist."))
	}

	if err := buffaloModels.DeletePostByID(postID); err != nil {
		log.Println(err)
		return c.Render(500, r.String("PostController(Delete): Database error, can't delete the post."))
	}

	return c.Render(204, r.JSON(map[string]interface{}{
		"post_deleted": "yes",
	}),
	)
}

// Serve the index page
func GetIndexPage(c buffalo.Context) error {
	rn := render.New(render.Options{
		DefaultContentType: "text/html",
		HTMLLayout:         "index.html",
		TemplatesBox:       packer.NewBox("../../../frontend"),
	})

	return c.Render(200, rn.HTML("index"))
}

// Check if a Post exists in the database by ID
func checkIfPostExists(postID int, postsSlice []*buffaloModels.Post) bool {
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
