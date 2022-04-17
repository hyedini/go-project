package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/user/post", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")
		fmt.Printf("id: %s; name: %s;\n", id, name)
	})

	router.Run(":8080")
}
