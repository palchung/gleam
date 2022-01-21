package user

import (
	"github.com/gin-gonic/gin"
	"thefreepress/controller/auth"
)

type Todo struct {
	UserID 	uint64 	`json:"user_id"`
	Title 	string	`json:"title"`	
}


func CreateTodo(c *gin.Context) {
	var td *Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "unauthorized")
		return
	}
	userId, err = auth.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	td.UserID = userId

	c.JSON(http.StatusCreated, td)
}