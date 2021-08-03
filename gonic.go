package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{})
}

func init() {
	gin.SetMode("release")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", home)

	// we'll change this to unix socket later
	router.Run(":9872")
}