package qthulhu

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func append(c *gin.Context) {
	// p := c.Param("partition")
	// c.String(http.StatusOK, "Hello %s", name)
}

func list(c *gin.Context) {
	p := c.Param("partition")
	c.String(http.StatusOK, "%s", p)
}

func main() {
	router := gin.Default()
	// router.GET("/hello", hello)
	router.GET("/:partition", list)
	router.PUT("/:partition/append", append)

	router.Run(":8080")
}
