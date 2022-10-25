package views

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	//Database connection

	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/enrollment_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to Connect to the Database ", err)
	}

	if r.Method == "POST" {

		firstname := r.FormValue("firstname")
		middlename := r.FormValue("middlename")
		lastname := r.FormValue("lastname")
		age := r.FormValue("age")
		gender := r.FormValue("gender")
		address := r.FormValue("address")
		birthday := r.FormValue("birthday")
		birthplace := r.FormValue("birthplace")
		status := r.FormValue("status")
		nationality := r.FormValue("nationality")
		religion := r.FormValue("religion")
		mobile_number := r.FormValue("mobile_number")
		email := r.FormValue("email")

		_ = db.Exec("USE enrollment_system;")
		_ = db.Exec("INSERT INTO student_details(first_name, middle_name, last_name, Age, Gender, address, birthday, birth_place, status, nationality, religion, mobile_number, email) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)", firstname, middlename, lastname, age, gender, address, birthday, birthplace, status, nationality, religion, mobile_number, email)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data["Title"] = "index | enrollment-system"
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, data)
}
