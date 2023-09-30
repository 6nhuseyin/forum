package main

import (
    "net/http"
)

type User struct {
    ID       int
    Email    string
    Username string
    Password string
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.ListenAndServe(":8080", nil)
}






// below is the old code:
/*
package main

import (
	"database/sql"
	"fmt"
	"log"

	//database driver for SQLite ->
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, err := sql.Open("sqlite3", "forumdatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, username TEXT, email TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("INSERT INTO users(username, email, password) VALUES(?, ? , ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("bambi", "bambi@gmail.com", "IloveBonesandSometimesMike")

	rows, _ := database.Query("SELECT id, username, email, password FROM users")
	var id int
	var username string
	var password string
	var email string

	for rows.Next() {
		rows.Scan(&id, &username, &password, &email)

		fmt.Println("id is: ", id)
		fmt.Println("username is: ", username)
		fmt.Println("email is: ", email)
		fmt.Println("password is: ", password)

	}
}
*/