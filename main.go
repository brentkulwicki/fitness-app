package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//type user struct {
//	id        int
//	name      string
//	email     string
//	createdAt string
//}

//func main() {
//	r := gin.Default()
//	r.POST("/users", postUser)
//
//	r.run("localhost:8085")
//}
//
//func postUser(c *gin.Context) {
//	var newUser user
//
//	if err := c.BindJSON(&newUser); err != nil {
//		return
//	}
//
//	c.IndentedJSON(http.StatusCreated, newUser)
//}

func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {

		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		email := c.PostForm("email")

		fmt.Printf("page: %s; name: %s; message: %s", page, name, email)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
