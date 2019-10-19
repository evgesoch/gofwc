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
	datasourceName = "C:/Users/Evgenios/go/src/github.com/evgesoch/gofwc/backend/gin/models/maindb.db"
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
	db, _ = sql.Open("sqlite3", "./backend/gin/models/maindb.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT)")
	statement.Exec()
	db.Close()
}

// Get all Posts from the database
func GetAllPosts() ([]*Post, error) {
	OpenDB()

	allPosts, err := db.Query("SELECT * FROM posts ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer CloseDB()
	defer allPosts.Close()

	posts := make([]*Post, 0)

	for allPosts.Next() {
		post := new(Post)
		err := allPosts.Scan(&post.ID, &post.Text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := allPosts.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return posts, nil
}

// Get a Post from the database
func GetPostByID(id int) (*Post, error) {
	OpenDB()

	fetchedPost, err := db.Query("SELECT id, text FROM posts WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer CloseDB()
	defer fetchedPost.Close()

	post := new(Post)

	for fetchedPost.Next() {
		err = fetchedPost.Scan(&post.ID, &post.Text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	if err := fetchedPost.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return post, nil
}

// Create a new Post in the database
func CreatePost(text string) (int64, error) {
	OpenDB()

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

	defer CloseDB()

	return lastID, nil
}

// Update a Post in the database
func UpdatePostByID(id int, text string) error {
	OpenDB()

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

	defer CloseDB()

	return nil
}

// Delete a Post from the database
func DeletePostByID(id int) error {
	OpenDB()

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

	defer CloseDB()

	return nil
}
