package main

import (
	"log"
	"net/http"
	"strings"

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
	for h, v := range ctx.Request.Header {
		if strings.HasPrefix("ssl", h) {
			log.Printf("%s: %v", h, v)
		}
	}

	ctx.String(200, "pong")
}
