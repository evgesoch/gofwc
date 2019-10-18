package controllers

import (
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
		log.Fatal(err)
	}

	pc.Data["json"] = allPosts
	pc.ServeJSON()
}

// @Title Get
// @Description Get a Post by its ID
// @Param postID path int	true "the postID you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :postID is empty
// @router /:postID [get]
func (pc *PostController) Get() {
	postID, err := strconv.Atoi(pc.Ctx.Input.Param(":postID"))
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(Get): Can't convert URI postID param to int."))
	}

	post, err := models.GetPostByID(postID)
	if err != nil {
		pc.Ctx.Output.SetStatus(500)
		pc.Ctx.Output.Body([]byte("PostController(Get): Database error, can't fetch the post."))
		log.Fatal(err)
	}

	if post.ID == 0 {
		pc.Ctx.Output.SetStatus(404)
		pc.Ctx.Output.Body([]byte("PostController(Get): This Post doesn't exist in the database."))
	}

	pc.Data["json"] = post
	pc.ServeJSON()
}

func (pc *PostController) Post() {

}

func (pc *PostController) Put() {

}

func (pc *PostController) Delete() {

}
