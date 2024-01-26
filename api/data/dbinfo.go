package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func open(path string) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))
	fmt.Println("path = ", path)

	return open(path)
}

var mydb *sql.DB

func init() {
	fmt.Println("Connecting Database...")
	mydb = connectDB()
	if mydb == nil {
		panic("Can't connect Database")
	} else {
		fmt.Println("Finished DB Initilized!")
	}
}

func GetMydb() *sql.DB {
	return mydb
}
