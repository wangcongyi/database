package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
		fmt.Println("link succeeded")
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

func inser(db *sql.DB) {
	stmt, err := db.Prepare("insert into users(name, email) values ($1,$2)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("golang", "golangtest@qq.com")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert into users success")
	}
}

func delete(db *sql.DB) {
	stms, err := db.Prepare("DELETE FROM users WHERE id=$1")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stms.Exec(2)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("delete from users success")
	}
}

func main() {
	db := connectDB()
	query(db)
	inser(db)
	delete(db)
}
  
