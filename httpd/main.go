package main

import (
	"Go-RestAPI/httpd/handler"
	"Go-RestAPI/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run("localhost:9090")

}
