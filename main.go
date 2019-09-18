package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "me"
	password = "password"
	dbname   = "api"
)

func connectDB() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", sqlInfo)

	if err != nil {
		panic(err)
	} else {
		return db
	}
}

func query(db *sql.DB) {
	var id, name, email string
	rows, err := db.Query("select * from users where id=$1", "1")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, name, email)

}

func main() {
	db := connectDB()
	query(db)
}
  
