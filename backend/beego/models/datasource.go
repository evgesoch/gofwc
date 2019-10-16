package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	database, _ := sql.Open("sqlite3", "./backend/beego/models/maindb.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, text TEXT)")
	statement.Exec()
	database.Close()
}
