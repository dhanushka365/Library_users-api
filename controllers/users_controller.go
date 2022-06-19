package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/library_users-api/domain/users"
	"github.com/library_users-api/services"
)

//create some endpoint
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemet me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	//Invalid json body
	if err := c.ShouldBindJSON(&user);err !=nil{
		fmt.Println(err)
		//TODO:return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO:handle user creation error
		return
	}
	fmt.Println(user)
	//fmt.Println(err)
	//fmt.Println(string(bytes))
	c.JSON(http.StatusCreated, result)
}

func SearchUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemet me!")
}
