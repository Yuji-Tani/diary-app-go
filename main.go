package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DBMS     string = "mysql"
	USER     string = "*****"
	PASS     string = "*****"
	PROTOCOL string = "tcp(27.133.130.174:3306)"
	DBNAME   string = "diary_app"
	OPTION   string = "?parseTime=true"
	CONNECT         = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + OPTION
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DB初期化
func dbinit() {
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("データベース開けず！（dbInit）")
	}

	db.AutoMigrate(&Todo{})
	//db.CreateTable(&Todo{})
	defer db.Close()

}

func dbGetAll() []Todo {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！（dbGetAll）")
	}
	var todos []Todo
	//db.Order("created_at desc").Find(&todos)
	db.Find(&todos)
	db.Close()

	fmt.Println("Status = " + todos[0].Status)
	fmt.Println("Text = " + todos[0].Text)
	fmt.Println("Status = " + todos[1].Status)
	fmt.Println("Text = " + todos[1].Text)
	fm.Println("Text = " + todos[1].Text)

	return todos

}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbinit()

	// Index
	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()

		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.Run(":9090")

}
