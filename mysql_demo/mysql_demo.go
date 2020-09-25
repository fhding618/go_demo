package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int32
	age  int32
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:12345@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	tx, _ := db.Begin()
	row := db.QueryRow("SELECT * FROM users")
	user := new(User)
	err = row.Scan(&user.id, &user.age, &user.name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", user)
	tx.Commit()
	defer db.Close()
}
