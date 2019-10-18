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
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT)")
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
func GetPostByID(id int) (*Post, error) {
	fetchedPost, err := db.Query("SELECT id, text FROM posts WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	defer fetchedPost.Close()

	post := new(Post)

	for fetchedPost.Next() {
		err = fetchedPost.Scan(&post.ID, &post.Text)
		if err != nil {
			return nil, err
		}
	}

	if err := fetchedPost.Err(); err != nil {
		return nil, err
	}

	return post, nil
}

// Create a new Post in the database
func CreatePost(text string) (int64, error) {
	prstmt, err := db.Prepare("INSERT INTO posts (text) VALUES (?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	result, err := prstmt.Exec(text)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return lastID, nil
}

// Update a Post in the database
func UpdatePostByID(id int, text string) error {
	prstmt, err := db.Prepare("UPDATE posts SET text=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = prstmt.Exec(text, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Delete a Post from the database
func DeletePostByID(id int) error {
	prstmt, err := db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = prstmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
