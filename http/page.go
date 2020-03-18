package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type pageBody struct {
	Msg string `json:"msg"`
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, pageBody{Msg: "Welcome"})
}

func notfound(c *gin.Context) {
	c.JSON(http.StatusNotFound, pageBody{Msg: "404 NotFound"})
}

