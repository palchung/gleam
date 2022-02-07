package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	UserID int64  `json:"user_id"`
}

func (p *profileHandler) CreateTodo(c *gin.Context) {
	var td Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	userId := GetUserIdByToken(p, c)
	td.UserID = userId

	// save data into database
	testdata := p.db.AllUsers()

	c.JSON(http.StatusCreated, testdata)
}
