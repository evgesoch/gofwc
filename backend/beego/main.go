package main

import (
	"github.com/evgesoch/gofwc/backend/beego/models"

	_ "github.com/evgesoch/gofwc/backend/beego/routers"
)

func main() {

	//models.CreateDB()

	//Open the main database
	models.OpenDB()

	// Create a Post
	/*a, err := models.CreatePost("hi 25th text")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The new Post has ID: ", a)*/

	// Update a Post
	/*err := models.UpdatePostByID(3, "hi 3rd text")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("updated Post")*/

	// Delete a Post
	/*err := models.DeletePostByID(24)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deleted Post")*/

	// Get all Posts
	/*b, err := models.GetAllPosts()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("All Posts: ", b)*/

	// Get one Post
	/*b, err := models.GetPostByID(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("A Posts: ", b)*/

	/*c, _ := models.GetPostByID(1)
	fmt.Println("A Post: ", c)*/

	models.CloseDB()

	/*if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()*/
}
