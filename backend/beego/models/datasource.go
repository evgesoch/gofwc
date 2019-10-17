package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// The Post struct
type Post struct {
	ID   int
	Text string
}

var (
	db             *sql.DB
	driver         = "sqlite3"
	datasourceName = "./backend/beego/models/maindb.db"
)

// Open the database
func OpenDB() {
	var err error

	db, err = sql.Open(driver, datasourceName)
	if err != nil {
		log.Panic(err)
	}
}

// Close the database
func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Panic(err)
	}
}

// Create the database
func CreateDB() {
	db, _ = sql.Open("sqlite3", "./backend/beego/models/maindb.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, text TEXT)")
	statement.Exec()
	db.Close()
}

// Get all Posts from the database
func GetAllPosts() ([]*Post, error) {
	allPosts, err := db.Query("SELECT * FROM posts ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer allPosts.Close()

	posts := make([]*Post, 0)

	for allPosts.Next() {
		post := new(Post)
		err := allPosts.Scan(&post.ID, &post.Text)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := allPosts.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// Get a Post from the database
func GetPostByID(postID int) (*Post, error) {
	fetchedPost, err := db.Query("SELECT id, text FROM posts WHERE id=" + string(postID))
	if err != nil {
		return nil, err
	}
	defer fetchedPost.Close()

	post := new(Post)
	err = fetchedPost.Scan(&post.ID, &post.Text)
	if err != nil {
		return nil, err
	}

	if err := fetchedPost.Err(); err != nil {
		return nil, err
	}

	return post, nil
}

// Create a new Post in the database
func CreatePost(id int, text string) (int, error) {
	newPost, err := db.Query("INSERT INTO posts (id, text) VALUES (?, ?)", id, text)
	if err != nil {
		return 0, err
	}
	defer newPost.Close()

	post := new(Post)
	err = newPost.Scan(&post.ID)
	if err != nil {
		return 0, err
	}

	if err := newPost.Err(); err != nil {
		return 0, err
	}

	return post.ID, nil
}

// Update a Post in the database
func UpdatePostByID(id int, text string) error {
	updatedPost, err := db.Query("UPDATE posts SET text=" + text + " WHERE id=" + string(id))
	if err != nil {
		return err
	}
	defer updatedPost.Close()

	if err := updatedPost.Err(); err != nil {
		return err
	}

	return nil
}

// Delete a Post from the database
func DeletePostByID(id int) error {
	deletedPost, err := db.Query("DELETE FROM posts WHERE id=" + string(id))
	if err != nil {
		return err
	}
	defer deletedPost.Close()

	if err := deletedPost.Err(); err != nil {
		return err
	}

	return nil
}
