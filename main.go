package main

import (
	"database/sql"
	"net/http"

	"github.com/nabalmes/enrollment-system/views"
)

const (
	BindIP = "0.0.0.0"
	Port   = "6969"
)

func main() {
	CreateDB("student_details")

	CreateTable("students", "student_details")
	Handlers()
	http.ListenAndServe(Port, nil)
}

func Handlers() {
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./templates"))))
	httt.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", views.IndexHandler)
}

func CreateDB(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		pa
	}
}
