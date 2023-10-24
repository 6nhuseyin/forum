package dal

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)
var db *sql.DB

func InitDatabase(){
	var err error
	//Open SQLite database
	db, err = sql.Open("sqlite3","./dal/forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


// Read the schema.sql file
schema, err := os.ReadFile("./dal/schema.sql")
if err != nil {
	log.Fatal(err)
}


//Execute the SQL commands from the schema.sql file
_, err = db.Exec(string(schema))
if err != nil {
	log.Fatal(err)
}

}



