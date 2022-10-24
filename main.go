package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nabalmes/enrollment-system/models"
	_ "gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":6969"
)

func main() {
	names := structs.Names(&models.Transactions{})
	fmt.Println(names) // ["Foo", "Bar"]

	fmt.Printf("Go to port System: %v%v/\n", BindIP, Port)
	CreateDB("enrollment_system")

	CreateTable("student_details", "enrollment_system")
	Handlers()
	http.ListenAndServe(Port, nil)
}

func Handlers() {
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./templates"))))
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// http.HandleFunc("/", views.IndexHandler)
}

func CreateDB(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	q := "CREATE DATABASE IF NOT EXISTS {NAME}"
	q = strings.Replace(q, "{NAME}", name, -1)
	// _, err = db.Exec("CREATE DATABASE IF NOT EXIST " + name)
	_, err = db.Exec(q)
	if err != nil {
		panic(err)
	}

	db.Close()

	db, err = sql.Open("mysql", "root:Allen is Great 200%@(127.0.0:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func CreateTable(table string, name string) *sql.DB {
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	o := "USE {DBNAME}"
	o = strings.Replace(o, "{DBNAME}", name, -1)
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	q := "CREATE TABLE IF NOT EXISTS {TABLENAME} (id INT(11) PRIMARY KEY AUTO_INCREMENT, first_name varchar(32), middle_name varchar(32), last_name varchar(32), Age int(155), Gender varchar(32), address varchar(32), birthday datetime, birth_place varchar(32), status varchar(32), nationality varchar(32), religion varchar(32), mobile_number int(11), email varchar(32));"
	// q := "CREATE TABLE IF NOT EXISTS {TABLENAME} (id INT(11) PRIMARY KEY AUTO_INCREMENT, first_name varchar(32), middle_name varchar(32), last_name varchar(32), Age int(155), Gender varchar(32), address varchar(32), birthday datetime, birth_place varchar(32), status varchar(32), nationality varchar(32), religion varchar(32), mobile_number int(11), email varchar(32));"
	q = strings.Replace(q, "{TABLENAME}", table, -1)
	// _, err = db.Exec("CREATE TABLE OF NOT EXIST " + table + "(id INT(11) PRIMARY KEY AUT_INCREMENT, first_name varchar(32), Middle_name varchar(32), last_name varchar(32), Age int(155), Gender varchar(32), address varchar(32), birthday datetime, birth_place varchar(32), status varchar(32), nationality varchar(32), religion varchar(32), mobile_number int(11), email varchar(32);")
	_, err = db.Exec(q)
	if err != nil {
		panic(err)
	}
	db.Close()

	return db
}
