package main

import (
	
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Name string `json:"name"`
}


func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root"
    dbName := "demo"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func main() {

	db := dbConn()

	//db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/demo")

	
	fmt.Println("Connected")

	defer db.Close()

	results,err := db.Query("Select name from users")

	if err !=nil{
		panic(err.Error())
	}

	for results.Next(){
		var user User
		err = results.Scan(&user.Name)

		if err !=nil{
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

	

	defer results.Close()

	
	
}