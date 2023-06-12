package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "fitness_app"
)

func main() {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {

		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		email := c.PostForm("email")
		sqlStatement := `INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, now())`

		_, err := db.Exec(sqlStatement, name, email)
		if err != nil {
			panic(err)
		}

		fmt.Printf("page: %s; name: %s; message: %s", page, name, email)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
