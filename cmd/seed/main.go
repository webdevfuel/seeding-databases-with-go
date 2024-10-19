package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name  string
	Email string
}

var users []User = []User{
	{
		Name:  "Alice Johnson",
		Email: "alice@example.com",
	},
	{
		Name:  "Bob Smith",
		Email: "bob@example.com",
	},
	{
		Name:  "Charlie Davis",
		Email: "charlie@example.com",
	},
	{
		Name:  "Diana White",
		Email: "diana@example.com",
	},
}

func main() {
	db, err := sql.Open("sqlite3", "../../example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("delete from users;")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into users(name, email) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, u := range users {
		_, err = stmt.Exec(u.Name, u.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
