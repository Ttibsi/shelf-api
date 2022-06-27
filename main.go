package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db_name string

func Set_db_name(name string) {
    db_name = name
}

func execute(cmd string) {
    db, err := sql.Open("mysql", "root:@/"+db_name)
    if err != nil { panic(err) }
    defer db.Close()

    fmt.Println("Executing: " + cmd)

    _, err = db.Exec(cmd)
    if err != nil { panic(err) }
    db.Close()
}

func CreateDB() {
    //root:@tcp(127.0.0.1:3306)/
    // default username/password @ tcp conn to default mysql port
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
    if err != nil { panic(err) }
    defer db.Close()

    _, err = db.Exec("CREATE DATABASE " + db_name)
    if err != nil { panic(err) }
    db.Close()
}

func CreateTable(tbl_name string) {
    execute("CREATE TABLE " + tbl_name)
}
