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
	r.GET("/ping", pong)
	return r
}

func pong(ctx *gin.Context) {
	log.Printf("received ping with %v headers\n", len(ctx.Request.Header))
	for h, v := range ctx.Request.Header {
		log.Printf("%s: %v\n", h, v)
	}

	ctx.String(200, "pong")
}
