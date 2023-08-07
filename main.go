package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "user"
	dbname   = "db"
)

func main() {
	fmt.Println("Docker Postgres Example")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	performSelect(db)
}

func performSelect(db *sql.DB) {
	fmt.Println("Fetching All Users")
	rows, err := db.Query(`SELECT * FROM users`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var user_id int
		var username string
		var full_name string
		var password string
		var email string
		var created_on time.Time
		var last_login time.Time

		err = rows.Scan(&user_id, &username, &full_name, &password, &email, &created_on, &last_login)
		CheckError(err)

		fmt.Println(user_id, username, full_name, password, email, created_on, last_login)
	}
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
