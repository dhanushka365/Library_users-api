package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//create some endpoint
func GetUser(c *gin.Context) {
 c.String(http.StatusNotImplemented,"implemet me!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implemet me!")
}

func SearchUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implemet me!")
}
