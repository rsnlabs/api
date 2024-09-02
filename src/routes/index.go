package routes

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.String(200, "Welcome to the RsnAI Public API!")
}

func RegisterHomeRoute(r *gin.Engine) {
	r.GET("/", HomeHandler)
}