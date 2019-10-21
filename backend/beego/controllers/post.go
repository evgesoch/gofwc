package controllers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/evgesoch/gofwc/backend/beego/models"
)

// Main controller to hadle requests for Post
type PostController struct {
	beego.Controller
}

// @Title GetAll
// @Description Get all Posts
// @Success 200 {post} models.Post
// @Failure 403 :postID is empty
// @router / [get]
func (pc *PostController) GetAll() {
	allPosts, err := models.GetAllPosts()
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(GetAll): Database error, can't fetch posts."))
		log.Println(err)
		return
	}

	pc.Data["json"] = allPosts
	pc.ServeJSON()
}

// @Title Get
// @Description Get a Post by its ID
// @Param postID path int true "The Post's ID you want to get"
// @Success 200 {post} models.Post
// @Failure 403 :postID is empty
// @router /:postID [get]
func (pc *PostController) Get() {
	postID, err := strconv.Atoi(pc.Ctx.Input.Param(":postID"))
	if err != nil {
		pc.Ctx.Output.SetStatus(400)
		pc.Ctx.Output.Body([]byte("PostController(Get): Can't convert URI postID param to int."))
		log.Println(err)
		return
	}

	post, err := models.GetPostByID(postID)
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(Get): Database error, can't fetch the post."))
		log.Println(err)
		return
	}

	if post.ID == 0 {
		pc.Ctx.Output.SetStatus(404)
		pc.Ctx.Output.Body([]byte("PostController(Get): This Post doesn't exist in the database."))
	}

	pc.Data["json"] = post
	pc.ServeJSON()
}

// @Title Post
// @Description Create a new Post
// @Param text body string true	"The new Post's text"
// @Success 200 {string} models.Post.Id
// @Failure 403 body is empty
// @router / [post]
func (pc *PostController) Post() {
	params := make(map[string]interface{})

	json.Unmarshal(pc.Ctx.Input.RequestBody, &params)

	postText, ok := params["text"].(string)
	if !ok {
		pc.Ctx.Output.SetStatus(400)
		pc.Ctx.Output.Body([]byte("PostController(Post): The post's text isn't a string."))
		return
	}

	newPostID, err := models.CreatePost(postText)
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(Post): Database error, can't create new post."))
		log.Println(err)
		return
	}

	pc.Ctx.Output.SetStatus(201)
	pc.Data["json"] = map[string]interface{}{
		"postID": newPostID,
	}
	pc.ServeJSON()
}

// @Title Put
// @Description Update a Post by its ID
// @Param postID path int true "The Post's ID you want to update"
// @Param text body string true	"The updated Post's text"
// @Success 201 {post} models.Post
// @Failure 403 :postID is empty
// @router /:postID [put]
func (pc *PostController) Put() {
	postID, err := strconv.Atoi(pc.Ctx.Input.Param(":postID"))
	if err != nil {
		pc.Ctx.Output.SetStatus(400)
		pc.Ctx.Output.Body([]byte("PostController(Put): Can't convert URI postID param to int."))
		log.Println(err)
		return
	}

	// Check if post to update doesn't exist
	allPosts, err := models.GetAllPosts()
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(GetAll): Database error, can't fetch posts."))
		log.Println(err)
		return
	}
	if check := checkIfPostExists(postID, allPosts); !check {
		pc.Ctx.Output.SetStatus(404)
		pc.Ctx.Output.Body([]byte("PostController(Put): This post doesn't exist."))
		return
	}

	params := make(map[string]interface{})

	json.Unmarshal(pc.Ctx.Input.RequestBody, &params)

	postText, ok := params["text"].(string)
	if !ok {
		pc.Ctx.Output.SetStatus(400)
		pc.Ctx.Output.Body([]byte("PostController(Put): The updated post's text isn't a string."))
		return
	}

	if err := models.UpdatePostByID(postID, postText); err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(Put): Database error, can't update the post."))
		log.Println(err)
		return
	}

	pc.Data["json"] = map[string]interface{}{
		"post_updated": "yes",
	}
	pc.ServeJSON()
}

// @Title Delete
// @Description Delete a Post by its ID
// @Param postID path int true "The Post's ID you want to delete"
// @Success 204 {post} models.Post
// @Failure 403 :postID is empty
// @router /:postID [delete]
func (pc *PostController) Delete() {
	postID, err := strconv.Atoi(pc.Ctx.Input.Param(":postID"))
	if err != nil {
		pc.Ctx.Output.SetStatus(400)
		pc.Ctx.Output.Body([]byte("PostController(Delete): Can't convert URI postID param to int."))
		log.Println(err)
		return
	}

	// Check if post to delete doesn't exist
	allPosts, err := models.GetAllPosts()
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(GetAll): Database error, can't fetch posts."))
		log.Println(err)
		return
	}
	if check := checkIfPostExists(postID, allPosts); !check {
		pc.Ctx.Output.SetStatus(404)
		pc.Ctx.Output.Body([]byte("PostController(Delete): This post doesn't exist."))
		return
	}

	if err := models.DeletePostByID(postID); err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(Delete): Database error, can't delete the post."))
		log.Println(err)
		return
	}

	pc.Ctx.Output.SetStatus(204)
}

// Check if a Post exists in the database by ID
func checkIfPostExists(postID int, postsSlice []*models.Post) bool {
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
