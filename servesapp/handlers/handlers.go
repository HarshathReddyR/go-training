package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func API() *gin.Engine {

	router := gin.New()
	router.GET("/check", Msg)
	return router
}
func Msg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "This is Project"})
}
