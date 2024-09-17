package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func show_ui(w http.ResponseWriter, uri string) {
	t, _ := template.ParseFiles(uri)
	t.Execute(w, nil)
}

func index(c *gin.Context) {
	w := c.Writer

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

func Create(c *gin.Context) {
	r := c.Request
	w := c.Writer
	if r.Method == "POST" {
		show_ui(w, "create.gtpl")

		// 条件查询
		err := r.ParseForm()
		if nil == err {
			username := r.Form.Get("username")
			age, _ := strconv.Atoi(r.Form.Get("age"))
			sex := r.Form.Get("sex")

			if sex == "" {
				sex = "unknown" // 默认值
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

func Read(c *gin.Context) {
	r := c.Request
	w := c.Writer
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

func Update(c *gin.Context) {
	r := c.Request
	w := c.Writer

	if r.Method == "POST" {
		show_ui(w, "update.gtpl")

		err := r.ParseForm()
		if nil == err {
			var id = r.Form.Get("id")
			var new_name = r.Form.Get("name")

			// 建立链接
			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			stmt, err := db.Prepare("UPDATE sreetchat.user SET name=? WHERE id=?")
			checkErr(err)
			res, err := stmt.Exec(new_name, id)
			fmt.Fprintln(w, "update success")
			affect, err := res.RowsAffected()
			checkErr(err)
			fmt.Println("操作影响内容", affect, "行")
		}
	}
}

func Delete(c *gin.Context) {
	r := c.Request
	w := c.Writer

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
	router := gin.Default()

	router.GET("/", index)
	router.GET("/create", func(c *gin.Context) { show_ui(c.Writer, "create.gtpl") })
	router.GET("/read", func(c *gin.Context) { show_ui(c.Writer, "read.gtpl") })
	router.GET("/update", func(c *gin.Context) { show_ui(c.Writer, "update.gtpl") })
	router.GET("/delete", func(c *gin.Context) { show_ui(c.Writer, "delete.gtpl") })
	router.POST("/create", Create)
	router.POST("/read", Read)
	router.POST("/update", Update)
	router.POST("/delete", Delete)

	err := router.Run(":12138")
	if nil != err {
		log.Fatal("listen port 12138 failed:", err)
	}
}

func checkErr(err error) {
	if nil != err {
		fmt.Errorf("%s", err)
		panic(err)
	}
}
