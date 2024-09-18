package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
	checkErr(err)
	// 查看数据
	rows, err := db.Query("SELECT * FROM sreetchat.user")
	checkErr(err)

	for rows.Next() {
		var id int
		var username string
		var age string
		var sex string
		var timeStamp string
		err = rows.Scan(&id, &username, &age, &sex, &timeStamp)
		checkErr(err)

		t, _ := time.Parse("2006-01-02 15:04:05", timeStamp)
		//fmt.Printf("%v, %v, %v, %v, %v\n", id, username, age, sex, t)
		fmt.Fprintf(w, "%v, %v, %v, %v, %v\n", id, username, age, sex, t)
	}
}

func show_ui(w http.ResponseWriter, uri string) {
	t, _ := template.ParseFiles(uri)
	t.Execute(w, nil)
}
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		show_ui(w, "create.gtpl")
	}
	if r.Method == "POST" {
		show_ui(w, "create.gtpl")

		// 条件查询
		err := r.ParseForm()
		if nil == err {
			username := r.Form.Get("username")
			age, _ := strconv.Atoi(r.Form.Get("age"))

			sex := r.Form.Get("sex")
			if sex == "M" {
				sex = "male"
			} else if sex == "F" {
				sex = "female"
			} else if sex == "" {
				sex = "unknown"
			}

			timestamp := time.Now()
			fmt.Printf("POST username:%v, age:%v, sex:%v, timestamp:%v\n", username, age, sex, timestamp)

			//创建链接
			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			//插入资料
			stmt, err := db.Prepare("INSERT user SET username=?,age=?,sex=?,timestamp=?")
			checkErr(err)

			res, err := stmt.Exec(username, age, sex, timestamp)
			checkErr(err)

			id, err := res.LastInsertId()
			fmt.Println("LastInsertId():", id)
			checkErr(err)
		}
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		show_ui(w, "read.gtpl")
	}
	if r.Method == "POST" {
		show_ui(w, "read.gtpl")

		// 条件查询
		err := r.ParseForm()
		if nil == err {
			str_age_low := r.Form.Get("ageLower")
			str_age_upp := r.Form.Get("ageUpper")
			age_low, err := strconv.Atoi(str_age_low)
			age_high, err := strconv.Atoi(str_age_upp)

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
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		show_ui(w, "update.gtpl")
	}

	if r.Method == "POST" {
		show_ui(w, "update.gtpl")

		err := r.ParseForm()
		if nil == err {
			var id = r.Form.Get("id")
			var new_name = r.Form.Get("new_username")
			if new_name == "" {
				fmt.Fprintln(w, id, new_name)
				// fmt.Fprintln(w, "name 为空行")
				return
			}

			// 建立链接
			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			stmt, err := db.Prepare("UPDATE sreetchat.user SET username=? WHERE id=?")
			checkErr(err)
			res, err := stmt.Exec(new_name, id)
			fmt.Fprintln(w, "update success")
			affect, err := res.RowsAffected()
			checkErr(err)
			fmt.Println("操作影响内容", affect, "行")
		}
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		show_ui(w, "delete.gtpl")
	}

	if r.Method == "POST" {
		show_ui(w, "delete.gtpl")

		err := r.ParseForm()
		if nil == err {
			var id = r.Form.Get("id")
			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			stmt, err := db.Prepare("DELETE FROM user WHERE id=" + id)
			checkErr(err)

			result, err := stmt.Exec()
			checkErr(err)
			fmt.Fprintf(w, "Delete result=%s\n", result)

			affect, err := result.RowsAffected()
			checkErr(err)

			fmt.Println(affect)
		}
	}
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
