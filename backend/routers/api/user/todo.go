package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	UserID 	string 	`json:"user_id"`
	Title 	string	`json:"title"`	
}


func (h *profileHandler) CreateTodo(c *gin.Context) {
	var td Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invlid json")
		return
	}
	metadata, err := h.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userId, err := h.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	td.UserID = userId

	// save data into database

	c.JSON(http.StatusCreated, td)
}