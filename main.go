package main

import (
	"GoWitter/model"
	"database/sql"
	"fmt"
	"GoWitter/handler"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tpl     *template.Template
	db      *sql.DB
	dbError error
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		"root", "haruyuki0815", "localhost:3306", "gowitter")
	db, dbError = sql.Open("mysql", dataSourceName)
	if dbError != nil {
		panic(dbError)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	handler.SetUpTemplate(tpl)
	model.SetUpDatabase(db)
	fmt.Println("SetUp Success!")
}

func main() {
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("src/"))))
	http.HandleFunc("/", handler.BaseHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("HELLO WORLD")
}

