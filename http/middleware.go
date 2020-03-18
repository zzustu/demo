package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

func tMid(c *gin.Context) {
	log.Printf("%s", "Hello Mid")
	log.Printf("%s", "Hello Mid Over")
}

func t2Mid(c *gin.Context) {
	log.Printf("%s", "T2 Mid")
	c.Next()
	log.Printf("%s", "T2 Over.....")
}