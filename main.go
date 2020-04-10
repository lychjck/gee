package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello nice</h1>")
	})
	r.RUN(":8080")
}
