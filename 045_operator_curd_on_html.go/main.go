package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	// db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
	// checkErr(err)
	// // 查看数据
	// rows, err := db.Query("SELECT * FROM sreetchat.user")
	// checkErr(err)

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var age int
	// 	err = rows.Scan(&id, &name, &age)
	// 	checkErr(err)
	// 	fmt.Println(id, name, age)
	// }

}

func show_ui(w http.ResponseWriter, tmpl string) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, nil)
}
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		show_ui(w, "create.gtpl")
	}
	if r.Method == "POSt" {
		show_ui(w, "read.gtpl")

		// 条件查询
		err := r.ParseForm()
		if nil == err {
			// 检查年龄
			str_age_low := r.FormValue("ageLower")
			str_age_high := r.FormValue("ageUpper")
			age_low, err := strconv.Atoi(str_age_low)
			age_high, err := strconv.Atoi(str_age_high)

			if nil != err {
				fmt.Println(err)
				fmt.Fprintln(w, err)
			} else {
				// 建立链接
				db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
				checkErr(err)

				// 查询数据
				rows, err := db.Query("SELECT * FROM sreetchat.user WHERE age BETWEEN ? AND ?", age_low, age_high)
				checkErr(err)

				for rows.Next() {
					var id int
					var name string
					var age string
					var sex string
					var time_stamp string
					err = rows.Scan(&id, &name, &age, &sex, &time_stamp)
					checkErr(err)

					t, _ := time.Parse("2006-01-02 15:04:05", time_stamp)
					fmt.Fprintf(w, "%v,%v, %v,%v %v </br>", id, name, age, sex, t)
				}
			}
		}
	}
}

func Read(w http.ResponseWriter, r *http.Request) {

}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/read", Read)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	err := http.ListenAndServe(":12138", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}
}

func checkErr(err error) {
	if nil != err {
		fmt.Errorf("%s", err)
		panic(err)
	}
}
