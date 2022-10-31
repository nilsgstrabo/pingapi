package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := router()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}

func router() http.Handler {
	r := gin.Default()
	r.RemoveExtraSlash = true
	r.GET("/ping", func(ctx *gin.Context) { ctx.String(200, "pong") })
	return r
}
